import { Filter, Sorting } from "@devexpress/dx-react-grid";

export interface Filters {
  itemsPerPage: number;
  startIndex: number;
  query: string;
  rawFilters?: Filter[];
  sorting?: Sorting[];
}

export const oprationCast = (operation?: string) => {
  if (operation === "contains") {
    return "%";
  }
  if (operation === "equal") {
    return "=";
  }

  return "";
};

const camelToSnakeCase = (str: string) =>
  str.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`);

export function dxFilterToSqlAlike(
  filters: Array<Filter | undefined> | undefined,
  overrides?: { [key: string]: string }
): string {
  if (!filters) {
    return "";
  }

  let query = [];

  for (const item of filters) {
    if (!item) {
      continue;
    }
    query.push(
      `${camelToSnakeCase(item.columnName)} ${oprationCast(
        item.operation
      )} "${item.value?.replaceAll('"', `\"`)}"`
    );
  }

  return query.join(" and ");
}

export function urlStringToFilters(): Partial<Filters> {
  const params: Partial<Filters> = Object.fromEntries(
    new URLSearchParams(location.search)
  );

  if (params.itemsPerPage) {
    params.itemsPerPage = +params.itemsPerPage;
  }

  if (params.startIndex) {
    params.startIndex = +params.startIndex;
  }

  if (params.rawFilters) {
    try {
      params.rawFilters = JSON.parse(params.rawFilters as any) || [];
    } catch (err) {
      params.rawFilters = [];
    }
  }

  if (params.sorting) {
    try {
      params.sorting = JSON.parse(params.sorting as any) || [];
    } catch (err) {
      params.sorting = [];
    }
  }

  return params;
}
