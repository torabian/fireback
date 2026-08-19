import EssentialApp from "@fireback/enterprise-shell/EssentialApp";
import { WithWasmServer } from "@fireback/wasm-server/WithWasmServer";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";

// wasm-demo is its own app entry point (VITE_TARGET_APP == 'wasm-demo', see
// src/index.tsx) specifically so @fireback/wasm-server - which transitively
// pulls in pglite, an in-browser Postgres compiled to wasm - is only ever
// imported here, never from projectname's or self-service's own App.tsx.
// Each app is its own Vite/Rollup entry with its own module graph (that's
// the whole point of VITE_TARGET_APP - see src/index.tsx's own comment), so
// as long as WithWasmServer/wasmFetchOverride are imported *only* by files
// under this folder, pglite's multi-MB .wasm/.data assets never end up in
// projectname's or self-service's build output. Confirmed empirically while
// wiring this up (see WasmFetchContext.ts's doc comment for the numbers).
//
// WithWasmServer downloads and boots cmd/fireback-wasm, then gates mounting
// its children until window.handleWasmRequest is live - see that
// component's own doc comment for the boot/progress/error screens it shows
// meanwhile. EssentialApp (the same shell projectname uses - sidebar, panel
// layout, WithSelfServiceRoutes' auth gate) then picks up the resulting
// fetch override two ways, neither of which requires it (or WithFireback,
// or WithSelfServiceRoutes - all shared with every other app) to import
// wasm-server directly:
//   - WithFireback's FetchxContext, for every generated SDK action
//   - checkSessionViaWhoami (via WithSelfServiceRoutes), for the whoami
//     check SessionGate runs before rendering the sidebar
// both via WasmFetchOverrideContext - see that file's own doc comment.
//
// Rendering the sidebar at all still needs a session already in localStorage
// (nima_ui_session, see @fireback/auth-client/authenticationUtils) *before*
// this app boots - AuthenticationProvider reads it synchronously and trusts
// it outright (no checkValidity configured here, same as projectname/
// self-service), same as a real signed-in visitor. There's no signin flow
// wired up for this app to produce that session itself (cmd/fireback-wasm
// only implements GET /whoami so far), so seeding it is the caller's job:
// see e2e/cypress/e2e/wasm-demo.cy.ts, which seeds a fake one before
// visiting and asserts .sidebar shows up once whoami answers 200.
function App() {
  return (
    <WithWasmServer
      options={{
        wasmUrl: BUILD_VARIABLES.PUBLIC_URL + "wasm_exec.js",
        wasmExecUrl: BUILD_VARIABLES.PUBLIC_URL + "fireback.wasm",
      }}
    >
      <EssentialApp ApplicationRoutes={ApplicationRoutes} />
    </WithWasmServer>
  );
}

export default App;
