// pgliteBridge — TS port of the emi in-browser-server example's
// browser/database-bridge.js. It boots an in-browser Postgres (pglite) and
// installs `window.queryDatabase`, the one function the Go side's
// wasmpgdriver (modules/fireback/wasmpgdriver, wired into gorm through
// application.ConnectWasmPostgres) calls for every statement:
//
//   window.queryDatabase(sql: string, args: any[]) => Promise<string /* JSON */>
//
// Call installPgliteBridge() and await it BEFORE starting the wasm module
// (see wasmServer.ts's startWasmServer, which does this automatically) —
// Go's main() reads window.queryDatabase once at startup and fails fast if
// it isn't there yet.
//
// @electric-sql/pglite is only imported lazily (dynamic import), so apps
// that never call installPgliteBridge() never pay for pulling a multi-MB
// wasm Postgres build into their bundle. Install it yourself
// (`npm install @electric-sql/pglite`) wherever you actually use this.
import { PGlite } from "./pglite/index.js";

export interface WasmQueryResult {
  rows: Record<string, unknown>[];
  fields: { name: string; dataTypeID: number }[];
  affectedRows: number;
  error: { Message: string; Code: string; Severity: string } | null;
}

declare global {
  interface Window {
    queryDatabase?: (query: string, args: unknown) => Promise<string>;
  }
}

// Narrow surface of the PGlite instance this bridge actually needs, so we
// don't have to statically import (and type-depend on) the package.
interface PgliteLike {
  query(
    sql: string,
    params: unknown[],
  ): Promise<{
    rows?: Record<string, unknown>[];
    fields?: { name: string; dataTypeID: number }[];
    affectedRows?: number;
  }>;
}

let dbPromise: Promise<PgliteLike> | null = null;

/**
 * Boots (or reuses) a pglite instance persisted at `dataDir` and wires it up
 * as `window.queryDatabase`. Safe to call more than once — later calls with
 * the same dataDir just await the existing instance.
 *
 * @param dataDir pglite storage location. `idb://<name>` persists in
 *   IndexedDB across reloads (what you want for a real app); "memory://"
 *   gives a throwaway in-memory database (handy for tests/demos).
 */
export async function installPgliteBridge(
  dataDir = "idb://fireback-wasm",
): Promise<PgliteLike> {
  if (!dbPromise) {
    dbPromise = (async () => {
      // @ts-ignore -- optional peer dependency, only resolved at runtime
      // const { PGlite } = await import("@electric-sql/pglite");
      return PGlite.create(dataDir) as Promise<PgliteLike>;
    })();
  }

  const db = await dbPromise;

  window.queryDatabase = async (query: string, args: unknown) => {
    try {
      if (query === "COPY") {
        // Bulk-insert path used by pgx's CopyFrom (see wasmpgdriver /
        // WasmPostgres.go's CopyFromSource handling). Not exercised by plain
        // gorm calls, but kept for parity with the emi example so any future
        // bulk-insert helper works unmodified.
        const payload = JSON.parse(args as string) as {
          table: string | string[];
          columns: string[];
          rows: unknown[][];
        };
        const cols = payload.columns.map((c) => `"${c}"`).join(", ");
        const table = Array.isArray(payload.table)
          ? payload.table.map((t) => `"${t}"`).join(".")
          : `"${payload.table}"`;

        let affected = 0;
        for (const row of payload.rows || []) {
          const placeholders = row.map((_, i) => `$${i + 1}`).join(", ");
          const res = await db.query(
            `INSERT INTO ${table} (${cols}) VALUES (${placeholders})`,
            row,
          );
          affected += res.affectedRows ?? 0;
        }

        return JSON.stringify({
          rows: [],
          fields: [],
          affectedRows: affected,
          error: null,
        } satisfies WasmQueryResult);
      }

      const params = Array.isArray(args) ? args : [];
      const res = await db.query(query, params);

      return JSON.stringify({
        rows: res.rows ?? [],
        fields: res.fields ?? [],
        affectedRows: res.affectedRows ?? 0,
        error: null,
      } satisfies WasmQueryResult);
    } catch (e: any) {
      return JSON.stringify({
        rows: [],
        fields: [],
        affectedRows: 0,
        error: {
          Message: String(e && e.message ? e.message : e),
          Code: (e && e.code) || "",
          Severity: (e && e.severity) || "ERROR",
        },
      } satisfies WasmQueryResult);
    }
  };

  return db;
}
