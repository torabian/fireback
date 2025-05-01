package abac

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/torabian/fireback/modules/fireback"
)

func OnInitEnvHook() error {

	appConfig := fireback.GetConfig()
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Remember to use always absolute path for database, and storage.
	appConfig.Storage = fireback.AskFolderName("Storage folder (all upload files from users will go here)", filepath.Join(workingDirectory, "storage"))
	appConfig.TusPort = fireback.AskPortName("TUS File upload port", "4506")

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

	if r := fireback.AskForSelect("Do you want to add the seed data, menu items, etc?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}
		AppMenuSyncSeeders()
	} else {
		return nil
	}

	if r := fireback.AskForSelect("Do you want to create a root admin for project?", []string{"yes", "no"}); r == "yes" {
		db, dbErr := fireback.CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		if err := InteractiveUserAdmin(fireback.QueryDSL{
			WorkspaceHas: []string{ROOT_ALL_ACCESS},
			WorkspaceId:  "system",
			ItemsPerPage: 10,
		}); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
