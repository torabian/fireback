import React, { type ReactNode } from "react";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
import { useWasmServer } from "./useWasmServer";
import type { WasmServerOptions } from "./wasmServer";

// WithWasmServer — wraps the app's essential router (see App.tsx) and, when
// BUILD_VARIABLES.USE_WASM_SERVER ("VITE_USE_WASM_SERVER" in
// src/apps/*/build-variables/*.json) is on, holds off mounting `children`
// until an entire fireback backend, compiled to wasm (cmd/fireback-wasm),
// has downloaded and booted inside the tab, backed by an in-browser Postgres
// (pglite). enterprise-shell's WithFireback picks up the resulting
// wasm-backed FetchxContext automatically (it checks the same build
// variable), so nothing downstream — the router, the generated SDK actions —
// has to know the backend it's calling is running in the same tab instead of
// over the network.
//
// When the flag is off this is a pure passthrough: children mount
// immediately and wasmServer.ts/pgliteBridge.ts are never even touched.
//
//   <WithWasmServer>
//     <EssentialApp ApplicationRoutes={ApplicationRoutes} />
//   </WithWasmServer>
export function WithWasmServer({
  children,
  options,
  fallback,
}: {
  children: ReactNode;
  /** Passed straight through to startWasmServer/useWasmServer. */
  options?: WasmServerOptions;
  /** Replaces the default boot/error screens below. */
  fallback?: { booting?: ReactNode; error?: (error: Error) => ReactNode };
}) {
  // if (BUILD_VARIABLES.USE_WASM_SERVER !== "true") {
  //   return <>{children}</>;
  // }

  return (
    <WasmServerGate options={options} fallback={fallback}>
      {children}
    </WasmServerGate>
  );
}

// Split into its own component so useWasmServer (which starts the boot as a
// side effect) is only ever invoked while the feature is enabled — a plain
// `if` around a hook call inside one component would break the rules of
// hooks the moment the flag changed between renders.
function WasmServerGate({
  children,
  options,
  fallback,
}: {
  children: ReactNode;
  options?: WasmServerOptions;
  fallback?: { booting?: ReactNode; error?: (error: Error) => ReactNode };
}) {
  const { ready, error } = useWasmServer(options);

  if (error) {
    return (
      <>
        {fallback?.error?.(error) ?? (
          <div style={{ padding: "2rem", fontFamily: "monospace" }}>
            Failed to start the in-browser server: {error.message}
          </div>
        )}
      </>
    );
  }

  if (!ready) {
    return (
      <>
        {fallback?.booting ?? (
          <div style={{ padding: "2rem", fontFamily: "monospace" }}>
            Starting in-browser server…
          </div>
        )}
      </>
    );
  }

  return <>{children}</>;
}

export default WithWasmServer;
