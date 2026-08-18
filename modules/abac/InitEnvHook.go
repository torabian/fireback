//go:build !wasm

package abac

import (
	"fmt"
	"log"
	"os"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

func OnInitEnvHook(c *cli.Command) error {

	appConfig := fireback.GetConfig()

	// 5. Ask for the storage folder as well

	appConfig.Save(".env")

	fmt.Println("Creating storage directory, where all files will be uploaded to:", appConfig.Storage)
	if err := os.Mkdir(appConfig.Storage, os.ModePerm); err != nil {
		fmt.Println("Folder for storage exists or inaccessible.")
	}

	fmt.Println("Your new project has been created successfully.")
	fmt.Println("\nIf you want to start the project with HTTP Server, run:")
	fmt.Println("$ " + fireback.GetExePath() + " start \n ")
	fmt.Println("You can also run the project on daemon, as a system server to presist the connection: (good for production)")
	fmt.Println("$ " + fireback.GetExePath() + " service load \n ")

	if r := fireback.AskForSelect("Do you want repair the workspaces (adds necessary content to tables)?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		RepairTheWorkspaces()
	} else {
		return nil
	}

	if r := fireback.AskForSelect("Do you want to authorize the cli, by creating possible root account?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		IntegrateAuthFlow(c)

	}

	return nil
}
