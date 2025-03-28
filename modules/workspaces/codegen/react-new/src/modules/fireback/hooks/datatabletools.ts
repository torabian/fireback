import { Filter, Sorting } from "@devexpress/dx-react-grid";

export interface Filters {
  itemsPerPage: number;
  startIndex: number;
  sort?: string;
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

  let query: string[] = [];

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
