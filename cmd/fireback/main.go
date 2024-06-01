package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/torabian/fireback/modules/workspaces"
)

func inlineTest() {

	data, err := ioutil.ReadFile("/tmp/t44/sample.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	uri := workspaces.GetLineContext(data, 4, 5)

	log.Println("uri: ", uri)
	fmt.Println(1, uri)
}

func main() {

	file, err := os.OpenFile("/tmp/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)

	// inlineTest()
	// // This is an important setting for some kind of app which will be installed
	// // it makes it easier for fireback to find the configuration.
	// os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	// // This AppStart function is a wrapper for few things commonly can handle entire backend project
	// // startup. For mobile or desktop might other functionality be used.
	xapp.CommonHeadlessAppStart(func() {
		// // 	// If anything needs to be done after database initialized
		// // 	// workspaces.RegionalContentSyncSeeders()
		// // 	// workspaces.AppMenuSyncSeeders()
	})

}
