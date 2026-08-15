//go:build wasm

package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall/js"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/application"
)

// Fireback doesn't come with default ui in the cmd anymore.
// Fireback itself has 2 uis: Manage and SelfService.
// Developer needs to build them if necessary and put the static files in workspaces
// Folder. Fireback serves them on /manage and /selfservice, similarly child projects
// Can serve those react projects if they wanted to.
// //go:embed all:ui
// var ui embed.FS

var xapp = &application.Application{

	Modules: []*application.ModuleProvider{},
}

// SampleRecord is a throwaway model proving gorm's ordinary API (migration,
// create, query) works unmodified against the wasm bridge, reachable from the
// browser exactly like a real endpoint. Swap this out for real fireback
// entities once a module's EntityProvider is wired here.
type SampleRecord struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Label string `json:"label"`
}

func main() {

	fmt.Println(1, xapp)
	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", "fireback")

	// queryDatabase is injected by js-remote-ctx's installPgliteBridge()
	// (ui/packages/js-remote-ctx/common/pgliteBridge.ts, a TS port of the emi
	// in-browser-server example's browser/database-bridge.js) and is backed
	// by an in-browser pglite instance. application.ConnectWasmPostgres
	// routes gorm's normal Postgres dialector through it instead of a real
	// TCP dial — see modules/fireback/wasmpgdriver and
	// application/DatabaseConnectionWasm.go for the plumbing. Everything
	// below is stock gorm; it has no idea it isn't talking to a real server.
	queryFunc := js.Global().Get("queryDatabase")
	if queryFunc.IsUndefined() {
		fmt.Println("queryDatabase is not defined — call installPgliteBridge() (js-remote-ctx) before starting the wasm module")
		return
	}

	_, err := application.ConnectWasmPostgres(queryFunc, nil)
	if err != nil {
		fmt.Println("failed to connect wasm database:", err)
		return
	} else {
		fmt.Println("Postgres connection enabled via gorm.")
	}

	// A normal net/http router. gorm.DB is captured by closure the same way
	// any real fireback handler would grab it off application context —
	// nothing here is wasm-specific.
	mux := http.NewServeMux()

	// Handle Who am I api call here. Hardcoded, nothing imported beyond what
	// main.go already pulls in — a stand-in for the real abac WhoamiAction
	// until that's wired in for real.
	mux.HandleFunc("GET /whoami", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
	"userId": "mock-user-id",
	"workspaces": [
		{
			"name": "Mock Workspace",
			"uniqueId": "mock-workspace-id",
			"capabilities": [],
			"roles": []
		}
	]
}`)
	})

	// emigo.LiftWasmServer always dispatches through mux.ServeHTTP (it needs
	// a concrete *http.ServeMux, not just an http.Handler), so a single
	// catch-all "/" on a wrapping mux is enough to see every request before
	// it reaches the real routes — same trick you'd use for any net/http
	// logging middleware, just routed through one extra mux since that's
	// what LiftWasmServer's signature requires.
	logged := http.NewServeMux()
	logged.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[wasm] %s %s\n", r.Method, r.URL.String())
		mux.ServeHTTP(w, r)
	}))

	// Exposes window.handleWasmRequest(method, url, body, headersJSON) ->
	// Promise<JSON {status, headers, body}>. js-remote-ctx's
	// wasmFetchOverride() (common/wasmServer.ts) is the browser-side half:
	// plug it into a FetchxContext as fetchOverrideFn and every generated
	// SDK call in the app transparently lands on this mux instead of the
	// network.
	emigo.LiftWasmServer(logged, nil)

	// Keep the Go runtime alive so the exposed callback stays callable.
	select {}
}
