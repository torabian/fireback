import React, { type ReactNode, useMemo } from "react";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
import { useWasmServer } from "./useWasmServer";
import { wasmFetchOverride } from "./wasmServer";
import type { WasmDownloadProgress, WasmServerOptions } from "./wasmServer";
import { WasmFetchOverrideContext } from "./WasmFetchContext";

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
  if (BUILD_VARIABLES.USE_WASM_SERVER !== "true") {
    return <>{children}</>;
  }

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
  const { ready, error, progress } = useWasmServer(options);
  // Stable across renders (useWasmServer's own boot is memoized too - see
  // startWasmServer) - provided once ready so WithFireback (via
  // WasmFetchContext.ts) can pick it up without ever importing this package
  // directly. See that file's own doc comment for why that indirection exists.
  const fetchOverride = useMemo(() => wasmFetchOverride(), []);

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
    return <>{fallback?.booting ?? <DefaultBootingScreen progress={progress} />}</>;
  }

  return (
    <WasmFetchOverrideContext.Provider value={fetchOverride}>
      {children}
    </WasmFetchOverrideContext.Provider>
  );
}

const BYTES_PER_MB = 1024 * 1024;

function DefaultBootingScreen({
  progress,
}: {
  progress: WasmDownloadProgress | null;
}) {
  const percent =
    progress?.total != null
      ? Math.min(100, Math.round((progress.loaded / progress.total) * 100))
      : null;

  return (
    <div style={{ padding: "2rem", fontFamily: "monospace" }}>
      <div>Starting in-browser server…</div>

      {progress && progress.loaded > 0 && (
        <div style={{ marginTop: "0.75rem" }}>
          <div
            style={{
              width: 240,
              height: 8,
              borderRadius: 4,
              overflow: "hidden",
              background: "rgba(128, 128, 128, 0.25)",
            }}
          >
            <div
              style={{
                // No content-length to compute a real percentage from? Show
                // a fixed-width bar instead of a 0%-forever one, so it still
                // reads as "something is happening" rather than stuck.
                width: percent != null ? `${percent}%` : "40%",
                height: "100%",
                background: "currentColor",
                transition: "width 0.15s ease",
              }}
            />
          </div>
          <div style={{ marginTop: "0.25rem", fontSize: "0.85em" }}>
            {percent != null
              ? `${mb(progress.loaded)} / ${mb(progress.total!)} MB (${percent}%)`
              : `${mb(progress.loaded)} MB downloaded`}
          </div>
        </div>
      )}
    </div>
  );
}

function mb(bytes: number): string {
  return (bytes / BYTES_PER_MB).toFixed(1);
}

export default WithWasmServer;
