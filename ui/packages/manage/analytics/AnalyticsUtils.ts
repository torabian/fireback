import { AnalyticsOverviewActionRes } from "@fireback/manage/sdk/abac/AnalyticsOverviewAction";

export type AnalyticsItem = InstanceType<typeof AnalyticsOverviewActionRes.Items>;
export type AnalyticsSeries = InstanceType<typeof AnalyticsOverviewActionRes.Series>;

/**
 * Groups items into an array of [category, items] tuples, preserving the first-seen
 * order of each item's category rather than sorting - AnalyticsOverviewAction already
 * hands back every stat "in a stable display order (grouped by category)" (see
 * Abac.emi.yml's AnalyticsOverview action), so re-sorting here would just fight the
 * server's own grouping. Mirrors internal-stats/InternalStatsUtils.ts's own
 * groupByCategory - same shape of response, same reason to group the same way.
 */
export function groupByCategory<T extends { category: string }>(
  items: T[],
): Array<[string, T[]]> {
  const order: string[] = [];
  const byCategory = new Map<string, T[]>();
  for (const item of items) {
    if (!byCategory.has(item.category)) {
      byCategory.set(item.category, []);
      order.push(item.category);
    }
    byCategory.get(item.category)!.push(item);
  }
  return order.map((category) => [category, byCategory.get(category)!]);
}

const KNOWN_SEVERITIES = ["ok", "warn", "critical", "info"];

/** Anything the backend doesn't document (yet) reads as "info". */
export function normalizeSeverity(severity: string): string {
  return KNOWN_SEVERITIES.includes(severity) ? severity : "info";
}
