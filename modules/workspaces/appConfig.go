package workspaces

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

var UNIQUE_MACHINE_ID string = ""
var UNIQUE_APP_PUBLIC_KEY string = ""

func DashedString(input string) string {
	out := ""
	for index, char := range input {
		out += string(char)
		if index != 0 && (index+1)%4 == 0 && len(input) != index+1 {
			out += "-"
		}
	}

	return out
}

// Some products do not need the configuration file, they include it in the binary instead
var BundledConfig *AppConfig

func BuiltInConfig() AppConfig {
	if BundledConfig != nil {
		return *BundledConfig
	}

	config := AppConfig{}

	config.Name = "fireback"
	config.Log.StdErr = "/tmp/fireback-err.log"
	config.Log.StdOut = "/tmp/fireback-out.log"
	config.Service.MacIdentifier = "com.pixelplux.fireback"
	config.Service.DebianIdentifier = "fireback"
	config.Service.WindowsIdentifier = "com.pixelplux.fireback"
	config.PublicServer.Enabled = true
	config.PublicServer.Host = "localhost"
	config.PublicServer.Port = "4500"
	config.Drive.Port = "4502"
	config.Drive.Enabled = true
	config.Mqtt.MqttVersion = "3.1"
	config.Mqtt.ConnectTimeout = 10
	config.Mqtt.KeepAlive = 60
	config.Mqtt.CleanSession = true

	config.PublicServer.GrpcPort = "4510"

	config.BackOfficeServer.Enabled = false
	config.BackOfficeServer.Host = "localhost"
	config.BackOfficeServer.Port = "4501"

	config.Database.Vendor = "sqlite"
	config.Database.Database = OsGetDefaultDatabase()
	config.Headers.AccessControlAllowOrigin = "*"
	config.Headers.AccessControlAllowHeaders = "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With, Workspace, Workspace-Id, Role-Id, Deep, query"
	config.Drive.Storage = "fireback-file-storage"

	config.Gin.Mode = "production"
	config.SmartUI.Enabled = true
	config.SmartUI.RedirectOnSuccess = "/auth?exchangeKey=%exchangeKey%"

	return config
}

var initConfig = false
var cfg AppConfig

func NecessaryConfigAudit(config AppConfig) {

	if config.MailServer.Provider == "" {
		fmt.Println("Warning: mailserver.vendor is not set. You need to setup the mail server, otherwise users will not receive the confirm account server, forget password.")
	}

	if config.BackOfficeServer.Enabled {
		fmt.Println("Warning: backOfficeServer.enabled is true. Make sure you protected the port from public traffic, anyone can access your data using port " + config.BackOfficeServer.Port)
	}

}

type NpmPackageJson struct {
	Name string `json:"name"`
}

func ReadJsonFile[T any](path string, data *T) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &data)

	return nil
}

func ReadYamlFile[T any](path string, data *T) error {

	f, err := os.Open(path)
	if err != nil {
		return errors.New("cannot read file")
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Config yaml file is broken.")
		fmt.Println(err)
	}

	return nil
}

func ReadEmbedFileContent(fsRef *embed.FS, path string) (string, error) {

	data, err := fsRef.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ReadYamlFileEmbed[T any](fsRef *embed.FS, path string, data *T) error {

	f, err := fsRef.Open(path)
	if err != nil {
		return errors.New("cannot read file")
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Config yaml file is broken.")
		fmt.Println(err)
	}

	return nil
}

type EnvironmentUris struct {
	CurrentDirectory       string `json:"currentDirectory"`
	CurrentDirectoryConfig string `json:"currentDirectoryConfig"`
	BinaryDirectory        string `json:"binaryDirectory"`
	ConfigFileName         string `json:"configFileName"`
	AppDataDirectory       string `json:"appDataDirectory"`
	AppLogDirectory        string `json:"appLogDirectory"`
	OsAppDataDirectory     string `json:"osAppDataDirectory"`
	ProductUniqueDirectory string `json:"productUniqueDirectory"`
}

func GetEnvironmentUris() *EnvironmentUris {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Default project is fireback.
	PRODUCT_UNIQUE_DIR := os.Getenv("PRODUCT_UNIQUE_NAME")
	if PRODUCT_UNIQUE_DIR == "" {
		PRODUCT_UNIQUE_DIR = "fireback"
	}

	configFileName := PRODUCT_UNIQUE_DIR + "-configuration.yml"

	binaryDirectory := filepath.Dir(ex) +
		fmt.Sprintf("%c", os.PathSeparator)

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	currentDirectory := workingDirectory + fmt.Sprintf("%c", os.PathSeparator)

	osAppDataDirectory := ""
	switch cos := runtime.GOOS; cos {

	case "darwin":

		dirname, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		osAppDataDirectory = filepath.Join(dirname, "Library/Application Support") + fmt.Sprintf("%c", os.PathSeparator)
		// For root use this. It's weird to install a software as root on mac anyways.
		// osAppDataDirectory = filepath.Join("/Library/Application Support") + fmt.Sprintf("%c", os.PathSeparator)

	case "windows":
		osAppDataDirectory = "C:\\ProgramData" + fmt.Sprintf("%c", os.PathSeparator)

	case "linux":
		osAppDataDirectory = "/var/lib" + fmt.Sprintf("%c", os.PathSeparator)
	}

	appDataDirectory := filepath.Join(osAppDataDirectory, PRODUCT_UNIQUE_DIR) + fmt.Sprintf("%c", os.PathSeparator)
	appLogDirectory := filepath.Join(osAppDataDirectory, PRODUCT_UNIQUE_DIR) + fmt.Sprintf("%c", os.PathSeparator) + "logs" + fmt.Sprintf("%c", os.PathSeparator)

	currentDirectoryConfig := filepath.Join(currentDirectory, configFileName)

	os.MkdirAll(appDataDirectory, os.ModePerm)
	os.MkdirAll(appLogDirectory, os.ModePerm)

	return &EnvironmentUris{
		CurrentDirectory:       currentDirectory,
		CurrentDirectoryConfig: currentDirectoryConfig,
		ProductUniqueDirectory: PRODUCT_UNIQUE_DIR,
		BinaryDirectory:        binaryDirectory,
		ConfigFileName:         configFileName,
		OsAppDataDirectory:     osAppDataDirectory,
		AppDataDirectory:       appDataDirectory,
		AppLogDirectory:        appLogDirectory,
	}
}

func ResolveConfigurationUri() (string, error) {
	uris := GetEnvironmentUris()

	/**
	*	This is how we check for the configuration by priority:
	*	0-CONFIG_PATH ENV, it overrides configuration, if not found, project fails.
	*	1-CWD/PRODUCT_UNIQUE_NAME-configuration.yml
	*	2- OS Specific App Data folder, per installation (for example on fireback makes no sense, )
	*	3-BinaryLocation/PRODUCT_UNIQUE_NAME-configuration.yml
	*	But on other products makes sense, which there will be a copy only
	**/

	uri1 := os.Getenv("CONFIG_PATH")
	if uri1 != "" && Exists(uri1) {
		return uri1, nil
	}

	uri2 := uris.CurrentDirectory + uris.ConfigFileName
	if Exists(uri2) {
		return uri2, nil
	}

	uri3 := uris.AppDataDirectory + uris.ConfigFileName
	if Exists(uri3) {
		return uri3, nil
	}

	uri4 := uris.BinaryDirectory + uris.ConfigFileName
	if Exists(uri4) {
		return uri4, nil
	}

	return "", errors.New("No configuration file has been found, checked 4 directories: " + uri1 + "\n" + uri2 + "\n" + uri3 + "\n" + uri4)
}

func GetAppConfig() AppConfig {
	uris := GetEnvironmentUris()

	if initConfig {
		return cfg
	}

	cfg = BuiltInConfig()
	cfg.Drive.Storage = strings.ReplaceAll(cfg.Drive.Storage, "{appDataDirectory}", uris.AppDataDirectory)
	cfg.Database.Database = strings.ReplaceAll(cfg.Database.Database, "{appDataDirectory}", uris.AppDataDirectory)

	initConfig = true

	uri, err := ResolveConfigurationUri()
	f, err := os.Open(uri)
	if err != nil {

		fmt.Println("Fireback cannot start without a configuration file. Either:\n ")
		fmt.Println("  * create a new project with `fireback init` in this directory,")
		fmt.Println("  * Run fireback with CONFIG_PATH environment variable, eg. \n\n 'CONFIG_PATH=/tmp/fireback-project.yml fireback start'")
		// os.Exit(100)
		return cfg

	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Config yaml file is broken.", uri)
		fmt.Println(err)
	}
	cfg.Drive.Storage = strings.ReplaceAll(cfg.Drive.Storage, "{appDataDirectory}", uris.AppDataDirectory)

	return cfg
}

func WriteAppConfig(cfgx AppConfig) {
	uri, err := ResolveConfigurationUri()
	cfg = cfgx

	if err != nil {
		fmt.Println("Configuration did not found on disk", uri)
	}

	body, err := yaml.Marshal(cfgx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote config", uri)
	os.WriteFile(uri, body, 0644)
}
