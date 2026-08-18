// A minimal JsonLogic (jsonlogic.com) evaluator - only used by this demo's fake
// data source (useFakeCursorListQuery.ts) to actually apply the filter tree
// DataGridList's column filters build (ui-core/components/data-grid-list/
// jsonLogic.ts) against in-memory rows, so the demo genuinely filters instead of
// just showing the built query. A real backend-backed queryHook never needs this
// - it just serializes the tree into a request, the same way the Go side
// (modules/fireback/jsonsql/JsonLogicToSql.go) already parses it server-side.
//
// Covers the operators DataGridList's own filter builder can produce (var,
// contains, and) plus enough of the rest of the spec (or, ==, !=, <, <=, >, >=)
// to make a hand-typed raw-JSON power filter (see jsonLogic.ts's "looks like
// JSON" escape hatch) work too.
export function evaluateJsonLogic(node: unknown, data: Record<string, unknown>): boolean {
  if (node === null || typeof node !== "object") {
    return Boolean(node);
  }
  if (Array.isArray(node)) {
    return node.every((n) => evaluateJsonLogic(n, data));
  }

  const entries = Object.entries(node as Record<string, unknown>);
  if (entries.length === 0) {
    return true;
  }
  const [op, rawArgs] = entries[0];
  const args = Array.isArray(rawArgs) ? rawArgs : [rawArgs];

  const resolve = (n: unknown): unknown => {
    if (n !== null && typeof n === "object" && !Array.isArray(n)) {
      const keys = Object.keys(n as object);
      if (keys.length === 1 && keys[0] === "var") {
        const path = (n as any).var as string;
        return path.split(".").reduce<any>((acc, key) => acc?.[key], data);
      }
      return evaluateJsonLogic(n, data);
    }
    return n;
  };

  switch (op) {
    case "and":
      return args.every((n) => evaluateJsonLogic(n, data));
    case "or":
      return args.some((n) => evaluateJsonLogic(n, data));
    case "var":
      return Boolean(resolve({ var: args[0] }));
    case "contains": {
      const haystack = String(resolve(args[0]) ?? "").toLowerCase();
      const needle = String(resolve(args[1]) ?? "").toLowerCase();
      return haystack.includes(needle);
    }
    case "==":
      // eslint-disable-next-line eqeqeq
      return resolve(args[0]) == resolve(args[1]);
    case "!=":
      // eslint-disable-next-line eqeqeq
      return resolve(args[0]) != resolve(args[1]);
    case "<":
      return (resolve(args[0]) as any) < (resolve(args[1]) as any);
    case "<=":
      return (resolve(args[0]) as any) <= (resolve(args[1]) as any);
    case ">":
      return (resolve(args[0]) as any) > (resolve(args[1]) as any);
    case ">=":
      return (resolve(args[0]) as any) >= (resolve(args[1]) as any);
    default:
      return true;
  }
}
