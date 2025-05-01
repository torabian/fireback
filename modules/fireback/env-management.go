package fireback

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

/*
Manages the code related to switching and creating environment (files)
*/

func EnvManagement(xapp *FirebackApp) cli.Command {
	return cli.Command{
		Name:  "env",
		Usage: "Manages the environments and .env files",
		Subcommands: []cli.Command{
			{
				Name:  "switch",
				Usage: "Switches between environments, if you specify the --to it would look for that file, if not it would ask",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "to",
						Required: false,
						Usage:    "The environment file, for example dev, prod, or any other, skip .env prefix",
					},
				},
				Action: func(c *cli.Context) error {
					envs, err := findEnvFiles()
					if err != nil {
						log.Fatalln("Error searching .env files: ", err)
						return err
					}

					if len(envs) == 0 || (len(envs) == 1 && Contains(envs, ".env")) {
						log.Fatalln("There are no environments to switch to, you need more than .env file to switch between")
						return nil
					}

					envName := c.String("to")

					if !c.IsSet("to") {
						envName = AskForSelect("Which environment file you want to switch to?", envs)
					}

					if envName == ".env" {
						log.Fatalln("You cannot switch to .env, because thats the only file which fireback will read.")
						return nil
					}

					switchEnvironment(envName)
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "Shows a list of .env files, basically any file starts with .env",
				Action: func(c *cli.Context) error {

					envs, err := findEnvFiles()
					if err != nil {
						log.Fatalln("Error searching .env files: ", err)
						return err
					}

					for _, item := range envs {
						fmt.Println(item)
					}

					if len(envs) == 0 {
						fmt.Println("There are not environments at all, you can use 'init' command to create default one")
					}

					return nil
				},
			},
			{
				Name:  "new",
				Usage: "Creates a new environment file, and initializes it",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "name",
						Required: false,
						Usage:    "file name of the new environment, which will be created",
					},
				},
				Action: func(c *cli.Context) error {

					envName := c.String("name")

					if !c.IsSet("name") {
						envName = AskForInput("Define the file name of env", "")
					}

					if !strings.HasPrefix(envName, ".env.") {
						envName = ".env." + envName
					}

					if Exists(envName) {
						log.Fatalln("Environment ", envName, " already exists. You can see a list of envs by 'list' command")
						return nil
					}

					if validation := validateEnvFileName(envName); validation != "" {
						log.Fatalln(validation)
						return nil
					}

					// InitEnvironment(xapp, envName)

					return nil
				},
			},
		},
	}
}

func validateEnvFileName(name string) string {
	// Check if the name is empty
	if name == "" {
		return "File name is empty"
	}

	// Check if the name starts with ".env."
	if !strings.HasPrefix(name, ".env.") {
		return "File name must start with '.env.'"
	}

	// Check for invalid path characters
	if name != filepath.Base(name) {
		return "File name contains invalid path characters"
	}

	// Check if the name is too short
	if len(name) <= 5 { // Minimum ".env." is 5 characters
		return "File name is too short"
	}

	// Validate allowed characters (letters, numbers, dashes, underscores, and periods)
	validNamePattern := `^[a-zA-Z0-9._-]+$`
	if !regexp.MustCompile(validNamePattern).MatchString(name) {
		return "File name contains invalid characters. Only letters, numbers, '.', '-', and '_' are allowed."
	}

	// Check for trailing or leading whitespace
	if strings.TrimSpace(name) != name {
		return "File name has leading or trailing whitespace"
	}

	// Check for length limits (optional, based on file system constraints)
	if len(name) > 255 {
		return "File name exceeds the maximum length of 255 characters"
	}

	return "" // No issues
}

func findEnvFiles() ([]string, error) {
	var envFiles []string

	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.Type().IsRegular() && len(entry.Name()) >= 4 && entry.Name()[:4] == ".env" {
			envFiles = append(envFiles, entry.Name())
		}
	}

	return envFiles, nil
}

func switchEnvironment(to string) error {

	fmt.Println("Switching to environment:", to)
	return nil
}
