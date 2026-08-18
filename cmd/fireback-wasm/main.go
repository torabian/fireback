//go:build wasm

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall/js"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	fireback "github.com/torabian/fireback/modules/fireback"
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

	// Every module's generated Config struct (fireback's own core one, abac's
	// otpLockoutSeconds/selfServiceBaseUrl, ...) is populated the same way on
	// both build targets: LoadConfiguration() -> emigo.HandleEnvVars() ->
	// envconfig reading os.Getenv per field (see emigo/Config.go's !wasm
	// implementation vs emigo/ConfigWasm.go's wasm one). There's no .env file
	// to read inside a browser sandbox though, so applyEnvFromJs seeds Go's
	// (real, if in-memory) os.Setenv store from window.firebackEnv - a plain
	// object the host page sets before calling go.run() (see
	// ui/packages/wasm-server/wasmServer.ts's WasmServerOptions.env) - which
	// has to happen before any LoadConfiguration() call below reads it.
	applyEnvFromJs()

	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration. Only applied
	// as a fallback - window.firebackEnv above is allowed to override it.
	if os.Getenv("PRODUCT_UNIQUE_NAME") == "" {
		os.Setenv("PRODUCT_UNIQUE_NAME", "fireback")
	}

	// Primes both configs from the env vars just applied - same reasoning as
	// modules/fireback/Entrypoint.go's CommonHeadlessAppStart (the non-wasm
	// startup path), just without that function's own !wasm-only pieces
	// (envm.LoadFirebackAppConfiguration's interactive .env prompting, the
	// full xapp.Modules ConfigProvider loop). Return values are discarded -
	// LoadConfiguration also stores into its package-level config var, which
	// is what every other call site (e.g. abacdefs.LoadConfiguration() inside
	// ClassicPassportRequestOtpActionImplementation.go) actually reads.
	fireback.LoadConfiguration()
	abacdefs.LoadConfiguration()

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

	con, err := application.ConnectWasmPostgres(queryFunc, nil)
	if err != nil {
		fmt.Println("failed to connect wasm database:", err)
		return
	} else {
		fmt.Println("Postgres connection enabled via gorm.")
	}

	fireback.SetDbRef(con)

	if err := abac.AbacMigration(fireback.GetDbRef()); err == nil {
		fmt.Println(1, "Migration done without an issue!")
	} else {
		fmt.Println(2, "Migration has some errors", err.Error())
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

	// Hardcoded, same stand-in status as /whoami above - the real one is
	// abac.CheckPassportMethodsAction (modules/abac/defs/
	// CheckPassportMethodsAction.go, GResponse<CheckPassportMethodsActionRes>
	// - hence the data.item wrapper below, not a flat object like /whoami's).
	// ui/packages/selfservice/Welcome.screen.tsx (the wasm-demo app's own
	// unauthenticated landing page - see ui/src/apps/wasm-demo/App.tsx) calls
	// this to decide which signin buttons to show; without it 404ing forever
	// left that screen stuck on its loading spinner. Only "email" is turned
	// on here - clicking it still won't get you signed in for real (there's
	// no ClassicSigninAction route here yet), but the screen itself now
	// renders instead of spinning.
	mux.HandleFunc("GET /passports/available-methods", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
	"data": {
		"item": {
			"email": true,
			"phone": false,
			"google": false,
			"facebook": false,
			"googleOAuthClientKey": "",
			"facebookAppId": "",
			"enabledRecaptcha2": false,
			"recaptcha2ClientKey": ""
		}
	}
}`)
	})

	{
		method, url, handler := abacdefs.CheckClassicPassportActionHttpHandler(abac.CheckClassicPassportAction)
		pattern := fmt.Sprintf("%v %v", strings.ToUpper(method), url)
		mux.HandleFunc(pattern, handler)
	}

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

// applyEnvFromJs copies every own-enumerable key of window.firebackEnv (a
// plain string->string object; see WasmServerOptions.env in
// ui/packages/wasm-server/wasmServer.ts) into Go's env via os.Setenv, so
// every module's LoadConfiguration() sees them the same way it would see
// real process env vars on a non-wasm build. window.firebackEnv itself is
// optional - a host page that doesn't set it just gets every config field's
// hardcoded default (the same as running the non-wasm binary with no .env
// file and no env vars set).
func applyEnvFromJs() {
	env := js.Global().Get("firebackEnv")
	if env.IsUndefined() || env.IsNull() {
		return
	}

	keys := js.Global().Get("Object").Call("keys", env)
	for i := 0; i < keys.Length(); i++ {
		key := keys.Index(i).String()
		os.Setenv(key, env.Get(key).String())
	}
}
