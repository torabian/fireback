import { useEffect, useState } from "react";
import { fetchx } from "@fireback/js-remote-ctx/common/fetchx";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";

// Proves the whole chain actually round-trips, not just that the wasm module
// booted: WithWasmServer -> WasmFetchOverrideContext -> WithFireback's
// FetchxContext -> fetchx() -> wasmFetch() -> window.handleWasmRequest ->
// cmd/fireback-wasm's mux -> back out again, same path any generated SDK
// action (RemoteAwareQuery, ...) would take instead of a network call.
//
// GET /whoami is the one route cmd/fireback-wasm/main.go actually implements
// today (hardcoded, no gorm/entities wired in yet - see that file's own
// comments) - a stand-in for "point real ApplicationRoutes/generated actions
// at this once the wasm binary implements more than a stub".
export function WasmWhoamiDemo() {
  const ctx = useFetchxContext();
  const [result, setResult] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchx("/whoami", {}, ctx)
      .then((res) => res.text())
      .then(setResult)
      .catch((err) => setError(err instanceof Error ? err.message : String(err)));
  }, [ctx]);

  return (
    <div style={{ padding: "2rem", fontFamily: "monospace", maxWidth: 720 }}>
      <h1>Fireback — wasm entry point</h1>
      <p>
        This app (<code>src/apps/wasm-demo</code>) never talks to a network
        backend: <code>WithWasmServer</code> downloaded and booted a fireback
        server compiled to wasm (<code>cmd/fireback-wasm</code>) right here in
        this tab, backed by an in-browser Postgres (
        <a href="https://pglite.dev/" target="_blank" rel="noreferrer">
          pglite
        </a>
        ). See <code>App.tsx</code> for the wiring, and{" "}
        <code>@fireback/wasm-server</code>'s own files for how it's kept out
        of every other app's bundle.
      </p>

      {error && (
        <pre style={{ color: "crimson", whiteSpace: "pre-wrap" }}>{error}</pre>
      )}
      {!error && !result && <p>Calling GET /whoami on the in-browser server…</p>}
      {result && (
        <>
          <p>GET /whoami, answered entirely inside this tab:</p>
          <pre
            style={{
              padding: "1rem",
              background: "rgba(128, 128, 128, 0.12)",
              borderRadius: 6,
              overflowX: "auto",
            }}
          >
            {result}
          </pre>
        </>
      )}
    </div>
  );
}

export default WasmWhoamiDemo;
