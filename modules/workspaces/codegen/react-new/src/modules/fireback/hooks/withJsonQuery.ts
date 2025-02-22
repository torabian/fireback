import {
  Context,
  FilterOperation,
  JsonQuery,
} from "../definitions/definitions";
import { flatMapDeep, get } from "lodash";

export function withJsonQuery(items: Array<any>, ctx: Context): Array<any> {
  const searchParams = new URLSearchParams(ctx.url);
  let jq0 = searchParams.get("jsonQuery");
  let jq: JsonQuery = null;
  if (jq0) {
    jq = JSON.parse(jq0);
  }

  return jsonQueryFilter(items, jq);
}

export function jsonQueryFilter(items: Array<any>, jq: JsonQuery) {
  console.log(25, items, jq);
  const jq1: Array<{
    name: string;
    filter: { value: any; operation: FilterOperation };
  }> = flatMapDeep(jq, (value, key, collection) => {
    let result: any = [];
    let name = key;

    // Check if it's a nested object. Do not take it if it has value, means it's last child
    if (value && typeof value === "object" && !value.value) {
      const keys = Object.keys(value);
      if (keys.length) {
        for (let key of keys) {
          result.push({
            name: `${name}.${key}`,
            filter: value[key],
          });
        }
      }
    } else {
      result.push({
        name,
        filter: value,
      });
    }

    return result;
  });

  return items.filter((item: any, index: number) => {
    for (let property of jq1) {
      const fieldValue = get(item, property.name);

      if (!fieldValue) {
        continue;
      }

      switch (property.filter.operation) {
        case "equal":
          if (fieldValue !== property.filter.value) {
            return false;
          }
          break;
        case "contains":
          if (!fieldValue.includes(property.filter.value)) {
            return false;
          }
          break;
        case "notContains":
          if (fieldValue.includes(property.filter.value)) {
            return false;
          }
          break;
        case "endsWith":
          if (!fieldValue.endsWith(property.filter.value)) {
            return false;
          }
          break;
        case "startsWith":
          if (!fieldValue.startsWith(property.filter.value)) {
            return false;
          }
          break;
        case "greaterThan":
          if (fieldValue < property.filter.value) {
            return false;
          }
          break;
        case "greaterThanOrEqual":
          if (fieldValue <= property.filter.value) {
            return false;
          }
          break;
        case "lessThan":
          if (fieldValue > property.filter.value) {
            return false;
          }
          break;
        case "lessThanOrEqual":
          if (fieldValue >= property.filter.value) {
            return false;
          }
          break;

        case "notEqual":
          if (fieldValue === property.filter.value) {
            return false;
          }
          break;
      }
    }

    return true;
  });
}
