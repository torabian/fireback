package workspaces

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

func detectProjectNameFromFolder() string {
	// Detect if it's a nodejs project. Node projects are having a package.json file,
	// we can read the package name from there

	if Exists("package.json") {
		body := &NpmPackageJson{}
		ReadJsonFile("package.json", body)

		if body.Name != "" {
			return body.Name
		}
	}

	return ""
}

func askProjectName() (string, error) {
	validate := func(input string) error {
		re := regexp.MustCompile(`^[a-z0-9-]*$`)

		if strings.Trim(input, " ") == "" || input == "" || !re.MatchString(input) {
			return errors.New("project name can only be lowercase and dash, and not empty")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "Project name",
		Validate: validate,
		Default:  detectProjectNameFromFolder(),
	}

	variable, err := promptVariable.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return variable, nil
}

func askSQLiteDatabaseLocation(projectName string) (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("enter the database path on file system, eg. /tmp/database1.db")
		}
		return nil
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	promptVariable := promptui.Prompt{
		Label:    "Database file location (.db)",
		Validate: validate,
		Default:  filepath.Join(workingDirectory, projectName+"-database.db"),
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}

func askMysqlDsn() (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("you need to enter dsn (eg: username:password@protocol(address)/dbname?param=value)")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "DSN Connection (eg: username:password@protocol(address)/dbname?param=value)",
		Validate: validate,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}
