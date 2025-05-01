import {random} from 'lodash';

export namespace inlineFiltering {
  export function NUMBER<T>(criteria: T, item: T, field: keyof T): boolean {
    if (
      criteria[field] !== undefined &&
      criteria[field] !== null &&
      criteria[field] !== item[field]
    ) {
      return true;
    }

    return false;
  }
  export function STRING<T>(criteria: T, item: T, field: keyof T): boolean {
    if (
      criteria[field] &&
      !((item as any)[field] || '')
        .toLowerCase()
        .includes((criteria as any)[field].toLowerCase())
    ) {
      return true;
    }

    return false;
  }

  export function BOOLEAN<T>(criteria: T, item: T, field: keyof T): boolean {
    if (
      ((criteria as any)[field] === true ||
        (criteria as any)[field] === false) &&
      (criteria as any)[field] !== (item as any).underAgeAccess
    ) {
      return true;
    }

    return false;
  }
}
export async function delay(min: number, max?: number) {
  return new Promise(r => {
    setTimeout(() => {
      r(true);
    }, random(min, max || min + 1));
  });
}
