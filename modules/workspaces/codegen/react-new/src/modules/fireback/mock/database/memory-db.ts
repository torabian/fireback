import { get } from "lodash";
import { uuidv4 } from "../../hooks/api";
import { Context } from "../../hooks/mock-tools";
import { BaseEntity } from "../../sdk/core/definitions";

type Criteria = Record<string, any>;

function flattenObject(
  obj: Criteria,
  parentKey = "",
  result: Criteria = {}
): Criteria {
  for (const key in obj) {
    if (obj.hasOwnProperty(key)) {
      const newKey = parentKey ? `${parentKey}.${key}` : key;

      if (
        typeof obj[key] === "object" &&
        !Array.isArray(obj[key]) &&
        !obj[key].operation
      ) {
        flattenObject(obj[key], newKey, result);
      } else {
        result[newKey] = obj[key];
      }
    }
  }

  return result;
}

function removeVowels(input: string): string {
  const diacriticsMap: Record<string, string> = {
    ł: "l",
    Ł: "L",
    // Add more mappings here if needed
  };

  return input
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/[aeiouAEIOU]/g, "")
    .replace(/[łŁ]/g, (match) => diacriticsMap[match] || match)
    .toLowerCase();
}

function filterData(data, criteria) {
  const flattenCriteria = flattenObject(criteria);

  return data.filter((item) => {
    return Object.keys(flattenCriteria).every((field) => {
      let { operation, value } = flattenCriteria[field];
      value = removeVowels(value || "");
      const fieldValue = removeVowels(get(item, field) || "");
      if (!fieldValue) return false;

      switch (operation) {
        case "contains":
          return fieldValue.includes(value);
        case "equals":
          return fieldValue === value;
        case "startsWith":
          return fieldValue.startsWith(value);
        case "endsWith":
          return fieldValue.endsWith(value);
        // Add more operations as needed
        default:
          return false;
      }
    });
  });
}

export class MemoryEntity<T extends BaseEntity> {
  constructor(private content: T[]) {}

  items(ctx?: Context): T[] {
    const jsonQuery = JSON.parse((ctx as any).jsonQuery);
    let filtered = filterData(this.content, jsonQuery);
    return filtered.filter((item, index) => {
      if (index < ctx.startIndex - 1) {
        return false;
      }

      if (ctx.itemsPerPage && index > ctx.startIndex + ctx.itemsPerPage - 1) {
        return false;
      }

      return true;
    });
  }

  total(): number {
    return this.content.length;
  }

  create(entity: Partial<T>): T {
    const newT = {
      ...entity,
      uniqueId: uuidv4().substr(0, 12),
    } as T;

    this.content.push(newT);
    return newT;
  }

  getOne(uniqueId: string): T | null {
    return this.content.find((item) => item.uniqueId === uniqueId);
  }

  deletes(uniqueId: string[]): boolean {
    this.content = this.content.filter(
      (item) => !uniqueId.includes(item.uniqueId)
    );

    return true;
  }

  patchOne(entity: Partial<T>): T | null {
    this.content = this.content.map((item) => {
      if (item.uniqueId === entity.uniqueId) {
        return {
          ...item,
          ...entity,
        };
      }

      return item;
    });

    return entity as T;
  }
}

export const QueryToId = (name: string) => {
  return name.split(" or ").map((item) => {
    return item.split(" = ")[1].trim();
  });
};
