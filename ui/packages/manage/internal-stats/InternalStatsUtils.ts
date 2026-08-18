import { InternalStatsSnapshotActionRes } from "@fireback/manage/sdk/internalstats/InternalStatsSnapshotAction";

export type InternalStatItem = InstanceType<
  typeof InternalStatsSnapshotActionRes.Items
>;

/**
 * Groups items into an array of [key, items] tuples, preserving the first-seen
 * order of `keyOf(item)` rather than sorting - InternalStatsSnapshot already
 * hands back every stat "in a stable display order (grouped by category)" (see
 * InternalStats.emi.yml), so re-sorting here would just fight the server's own
 * grouping.
 */
function groupPreservingOrder<T>(
  items: T[],
  keyOf: (item: T) => string,
): Array<[string, T[]]> {
  const order: string[] = [];
  const byKey = new Map<string, T[]>();
  for (const item of items) {
    const key = keyOf(item);
    if (!byKey.has(key)) {
      byKey.set(key, []);
      order.push(key);
    }
    byKey.get(key)!.push(item);
  }
  return order.map((key) => [key, byKey.get(key)!]);
}

export function groupByCategory(
  items: InternalStatItem[],
): Array<[string, InternalStatItem[]]> {
  return groupPreservingOrder(items, (item) => item.category);
}

// Preferred facet ordering for the chart - unknown units (anything Collector.go
// starts reporting later) fold in alphabetically after these, rather than being
// dropped.
const UNIT_ORDER = ["percent", "bytes", "seconds", "mhz", "load", "count"];

export function groupByUnit(
  items: InternalStatItem[],
): Array<[string, InternalStatItem[]]> {
  // Only items with a unit are numeric metrics - unit is "empty for non-numeric
  // stats" (see InternalStats.emi.yml), so a plain string field like hostname
  // or platform never ends up mistaken for a chartable value.
  const numeric = items.filter((item) => item.unit);
  const grouped = groupPreservingOrder(numeric, (item) => item.unit);
  return grouped.sort(([a], [b]) => {
    const ai = UNIT_ORDER.indexOf(a);
    const bi = UNIT_ORDER.indexOf(b);
    if (ai === -1 && bi === -1) return a.localeCompare(b);
    if (ai === -1) return 1;
    if (bi === -1) return -1;
    return ai - bi;
  });
}

const KNOWN_SEVERITIES = ["ok", "warn", "critical", "info"];

/** Anything InternalStats.emi.yml doesn't document (yet) reads as "info". */
export function normalizeSeverity(severity: string): string {
  return KNOWN_SEVERITIES.includes(severity) ? severity : "info";
}
