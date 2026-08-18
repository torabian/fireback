// CommonHeadlessAppStart bootstraps CLI dispatch (RunApp), which the wasm
// binary never does - it wires its own minimal HTTP mux + DB connection
// directly in cmd/fireback-wasm/main.go instead (see that file's own
// comments). Excluding this file from wasm builds isn't just tidiness: it
// transitively imports envm.LoadFirebackAppConfiguration, which pulls in
// manifoldco/promptui -> chzyer/readline for interactive .env prompts -
// readline has no wasm platform support at all (no _js.go files, so it
// doesn't even compile, let alone run), so this file has to be excluded for
// any wasm build to succeed, whether or not CommonHeadlessAppStart is ever
// called.
//go:build !wasm

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

	// Same idea, per module: a module's own `config:` block (Abac.emi.yml's
	// otpLockoutSeconds, for instance) only actually gets read from the environment
	// when something calls that module's generated LoadConfiguration() - ConfigProvider
	// does exactly that (see application.ModuleProvider.ConfigProvider's doc comment),
	// but until now nothing ever called it unless GetCombinedConfigInfo happened to run
	// first (i.e. "config list"). Without this, every other consumer of that module's
	// config - its own generated per-field "get"/"set" CLI commands included - would
	// see only the struct's hardcoded default, never an env or .env override, exactly
	// the gap LoadFirebackAppConfiguration above already closes for fireback's own core
	// config. Priming here, unconditionally and before RunApp dispatches to any
	// subcommand, gives every module's config the same guarantee.
	for _, module := range x.Modules {
		if module.ConfigProvider != nil {
			module.ConfigProvider()
		}
	}

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
