package fireback

import (
	"log"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/envm"
	"github.com/urfave/cli/v3"
)

// Function which needs to be called to generate the server.

func CommonHeadlessAppStart(x *application.Application, onDatabaseCompleted func()) {

	// Load the application configuration
	envm.LoadFirebackAppConfiguration(&config)

	// Use the logger
	initLogger()

	if !excludeDatabaseConnection() {
		db, dbErr := CreateDatabasePool()
		if db == nil && dbErr != nil {
			log.Fatalln("Database error on initialize connection:", dbErr)
		}

		if onDatabaseCompleted != nil {
			onDatabaseCompleted()
		}

		// Give every registered module a chance to start whatever background work it
		// needs (e.g. eventbus.ModuleSetup starts the event bus goroutine here) - only
		// modules a project actually registered in its own Modules list run anything,
		// instead of fireback starting things like the event bus unconditionally for
		// every project.
		for _, module := range x.Modules {
			if module.OnAppStart == nil {
				continue
			}
			if err := module.OnAppStart(x); err != nil {
				log.Fatalln("Module", module.Name, "failed to start:", err)
			}
		}
	}

	x.CliActions = func() []*cli.Command {
		return GetCommonWebServerCliActions(x)
	}

	RunApp(x)
}
