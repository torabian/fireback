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
	reflect "reflect"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
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
		log.Default().Println("Yaml file is broken", path)
		return err
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
		log.Default().Println("Yaml file is broken", path)
		return err
	}

	return nil
}

type EnvironmentUris struct {
	CurrentDirectory       string `json:"currentDirectory" yaml:"currentDirectory"`
	CurrentDirectoryConfig string `json:"currentDirectoryConfig" yaml:"currentDirectoryConfig"`
	BinaryDirectory        string `json:"binaryDirectory" yaml:"binaryDirectory"`
	ConfigFileName         string `json:"configFileName" yaml:"configFileName"`
	AppDataDirectory       string `json:"appDataDirectory" yaml:"appDataDirectory"`
	AppLogDirectory        string `json:"appLogDirectory" yaml:"appLogDirectory"`
	OsAppDataDirectory     string `json:"osAppDataDirectory" yaml:"osAppDataDirectory"`
	ProductUniqueDirectory string `json:"productUniqueDirectory" yaml:"productUniqueDirectory"`
}

func (x *EnvironmentUris) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func (x *EnvironmentUris) Yaml() string {
	if x != nil {
		str, _ := yaml.Marshal(x)
		return (string(str))
	}
	return ""
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

	configFileName := ".env"

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

/*
In many places in an app, we might need a set of essential variables depending on the environment,
paths, tmp folders, etc. This function computes all and returns them as a map.
Put data which won't change on the app life time, such as os, current cwd, etc here
*/
func EssentialVariablesMap() map[string]string {
	uris := GetEnvironmentUris()

	return map[string]string{
		"appDataDirectory":       uris.AppDataDirectory,
		"currentDirectory":       uris.CurrentDirectory,
		"currentDirectoryConfig": uris.CurrentDirectoryConfig,
		"productUniqueDirectory": uris.ProductUniqueDirectory,
		"binaryDirectory":        uris.BinaryDirectory,
		"configFileName":         uris.ConfigFileName,
		"osAppDataDirectory":     uris.OsAppDataDirectory,
		"appLogDirectory":        uris.AppLogDirectory,
	}
}

func LoadXappConfiguration() {

	uri, err3 := ResolveConfigurationUri()
	if err3 != nil {
		// log.Default().Println("there are no configuration files found, using environment variables")
	} else {
		// log.Default().Printf("looking for config file: %v \r\n", uri)
		err := godotenv.Load(uri)
		if err != nil {
			log.Printf("environment variable file expected: %s was not loaded. Error: %v", uri, err)
		}
	}

	envconfig.MustProcess("", &config)
}

func HandleEnvVars(spec interface{}) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	filename := ".env." + env
	err := godotenv.Load(filename)
	if err != nil {
		log.Printf("environment variable file expected: %s was not loaded. Error: %v", filename, err)
	}

	envconfig.MustProcess("", spec)
}

func init() {
	LoadXappConfiguration()
}

func structToEnvMap(config interface{}) (map[string]string, error) {
	envMap := make(map[string]string)
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		envKey := field.Tag.Get("envconfig")

		if envKey != "" {
			value := val.Field(i).Interface()

			switch v := value.(type) {
			case string:
				envMap[envKey] = v
			case bool:
				envMap[envKey] = strconv.FormatBool(v)
			case int, int64:
				envMap[envKey] = strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
			case float64:
				envMap[envKey] = strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
			default:
				return nil, fmt.Errorf("unsupported type: %s", reflect.TypeOf(v))
			}
		}
	}

	return envMap, nil
}

func SaveEnvFile(config interface{}, filename string) error {
	envMap, err := structToEnvMap(config)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range envMap {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			return err
		}
	}

	return nil
}
