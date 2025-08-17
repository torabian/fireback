export class ListResponse<T> {
  items: T[] = [];

  constructor(private ItemClass: { new (): T }) {}

  static unserialize<T>(raw: any, ItemClass: { new (): T }): ListResponse<T> {
    const list = new ListResponse(ItemClass);
    if (Array.isArray(raw)) {
      list.items = raw.map((x) => Object.assign(new ItemClass(), x));
    } else if (raw.items && Array.isArray(raw.items)) {
      list.items = raw.items.map((x) => Object.assign(new ItemClass(), x));
    }
    return list;
  }

  static By<T>(x: { new (): T }) {
    return class {
      static unserialize(raw: unknown) {
        return ListResponse.unserialize(raw, x);
      }
    };
  }
}
