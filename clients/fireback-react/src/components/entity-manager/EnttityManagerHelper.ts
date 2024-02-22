import { Filter } from "@/definitions/definitions";
import { set } from "lodash";

/**
 * Filters is this specific datatable way of filtering things.
 * Our server, expects query dsl (json) instead of older string sql alike one.
 * This function converts it in that style
 * @param filters
 */
export function filtersToJsonQuery(
  filters: Array<Filter | undefined | null> | undefined | null
) {
  const jq = {};

  for (let filter of filters || []) {
    if (!filter) {
      continue;
    }

    if (filter.columnName) {
      set(jq, filter.columnName, {
        operation: filter.operation,
        value: filter.value,
      });
    }
  }

  return jq;
}
