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

func askEnvironmentName(originalName string) (string, error) {
	validate := func(input string) error {
		re := regexp.MustCompile(`^[a-z0-9-]*$`)

		if strings.Trim(input, " ") == "" || input == "" || !re.MatchString(input) {
			return errors.New("environment name can only be lowercase and dash, and not empty")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "Environment name (used for system service files and launchctl, and switching env)",
		Validate: validate,
		Default:  originalName,
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

func askPostgresDsn() (string, error) {

	validate := func(input string) error {
		if input == "" {
			return errors.New("you need to enter dsn (eg: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai)")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    "DSN Connection (eg: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai)",
		Validate: validate,
	}

	value, err := promptVariable.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return value, nil
}
