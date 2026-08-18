import { createContext, useContext } from "react";
import type { TypedRequestInit } from "@fireback/js-remote-ctx/common/fetchx";

export type FetchOverrideFn = (
  input: RequestInfo | URL,
  init?: TypedRequestInit,
) => Promise<Response>;

// Bridges the wasm-backed fetch override from WithWasmServer down to
// @fireback/enterprise-shell's WithFireback - shared by *every* app
// (projectname, self-service, ...) - WITHOUT WithFireback ever having to
// import anything else from @fireback/wasm-server directly.
//
// That matters because @fireback/wasm-server transitively imports pglite
// (an in-browser Postgres compiled to wasm - see pgliteBridge.ts): a real
// static import of it anywhere in a file every app shares ships pglite's
// .wasm/.data assets (tens of MB) in *every* app's build output, even ones
// that never touch wasm mode - confirmed empirically while wiring this up:
// building projectname/manage with a plain
// `BUILD_VARIABLES.USE_WASM_SERVER === "true" ? wasmFetchOverride() : undefined`
// in WithFireback added a ~16.5MB pglite-*.wasm/.data pair to the output
// even though that build's flag is "false" - a runtime-only guard doesn't
// stop Rollup from emitting an asset a static import graph references.
//
// This file is the one thing safe for WithFireback to import from
// @fireback/wasm-server: it has no other imports, and defines nothing beyond
// a context + a tiny hook, so it can never pull pglite in with it. Only an
// app that actually mounts WithWasmServer (see WithWasmServer.tsx) - which
// *does* import the real wasm/pglite code, and is meant to be imported only
// from that app's own entry point, never from shared infrastructure - ever
// provides a value here.
//
// undefined (the default, no Provider mounted) means "not running in wasm
// mode" - WithFireback falls back to its normal network fetch behavior.
export const WasmFetchOverrideContext = createContext<
  FetchOverrideFn | undefined
>(undefined);

export function useWasmFetchOverride(): FetchOverrideFn | undefined {
  return useContext(WasmFetchOverrideContext);
}
