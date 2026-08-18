import { FirebackEssentialRouterManager } from "@fireback/enterprise-shell/EssentialRouter";
import { Route } from "react-router-dom";
import { WasmWhoamiDemo } from "./WasmWhoamiDemo";

// Real routes now that checkSessionViaWhoami (ui-core/components/session-gate)
// is wasm-aware (see its own doc comment + WithSelfServiceRoutes.tsx) -
// FirebackEssentialRouterManager's auth-gated routing has somewhere valid to
// land once SessionGate's whoami check resolves against the in-browser
// server instead of hanging on a real network call that was never coming.
//
// A single demo route (the same WasmWhoamiDemo App.tsx used before switching
// to the full EssentialApp shell) - proves content renders inside the real
// sidebar layout, not just the sidebar chrome itself. Add real
// routes/entities here once cmd/fireback-wasm implements more than
// GET /whoami.

// ~ auto:useRouteImport

export function ApplicationRoutes({ routerId }: { routerId?: string }) {
  // ~ auto:useRouteDefs

  return (
    <FirebackEssentialRouterManager routerId={routerId}>
      {/* ~ auto:useRouteJsx */}
      <Route path="whoami" element={<WasmWhoamiDemo />}></Route>
    </FirebackEssentialRouterManager>
  );
}
