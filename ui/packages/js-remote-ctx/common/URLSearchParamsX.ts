import { stringify, parse } from "qs";

/**
 * Extended URLSearchParams that stores data in a nested object
 * and keeps compatibility with URLSearchParams methods.
 */
export class URLSearchParamsX extends URLSearchParams {
  /** Internal data store */
  private data: Record<string, any> = {};

  constructor(
    init?: string[][] | Record<string, string> | string | URLSearchParams
  ) {
    super(init);
    if (init) {
      if (typeof init === "string") {
        Object.assign(this.data, parse(init));
      } else if (init instanceof URLSearchParams) {
        Object.assign(this.data, parse(init.toString()));
      } else if (Array.isArray(init)) {
        init.forEach(([k, v]) => (this.data[k] = v));
      } else {
        Object.assign(this.data, init);
      }
    }
  }

  /** Remove a key from the store */
  override delete(name: string): void {
    delete this.data[name];
  }

  /** Append a value to an array or create a new array */
  override append(name: string, value: string): void {
    if (this.data[name] === undefined) this.data[name] = value;
    else if (Array.isArray(this.data[name])) this.data[name].push(value);
    else this.data[name] = [this.data[name], value];
  }

  /** Get an iterator of top-level keys */
  override keys(): URLSearchParamsIterator<string> {
    const obj = this.data;
    return (function* (): Generator<string, undefined, unknown> {
      for (const key of Object.keys(obj)) {
        yield key;
      }
      return undefined;
    })();
  }

  /** Number of top-level keys */
  override get size(): number {
    return Object.keys(this.data).length;
  }

  /** Sort top-level keys */
  override sort(): void {
    const sorted: Record<string, any> = {};
    Object.keys(this.data)
      .sort()
      .forEach((key) => {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
        sorted[key] = this.data[key];
      });
    this.data = sorted;
  }

  /** Get an iterator of top-level values */
  override values(): URLSearchParamsIterator<string> {
    const obj = this.data;
    return (function* (): Generator<string, undefined, unknown> {
      for (const key of Object.keys(obj)) {
        const val = obj[key];
        // Make sure val is string
        yield String(val);
      }
      return undefined;
    })();
  }

  /** Get a single value by key */
  override get(name: string): string | null {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    const val = this.data[name];
    if (val == null) return null;
    return Array.isArray(val) ? String(val[0]) : String(val);
  }

  /** Check if key exists */
  override has(name: string): boolean {
    return this.data[name] !== undefined;
  }

  /** Iterate over top-level keys and values */
  override forEach(
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    callbackfn: (value: any, key: string, parent: any) => void
  ): void {
    for (const key of Object.keys(this.data)) {
      callbackfn(this.data[key], key, this);
    }
  }

  /** Get all values for a key as array */
  override getAll(name: string): string[] {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    const val = this.data[name];
    if (val === undefined) return [];
    return Array.isArray(val) ? val.map(String) : [String(val)];
  }

  /** Get an iterator of key/value pairs (flattened) */
  override entries() {
    const params = new URLSearchParams(stringify(this.data));
    return params.entries();
  }

  /** Convert to query string */
  override toString(): string {
    return stringify(this.data);
  }

  /** Set a key to a value */
  override set(name: string, value: any): this {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    this.data[name] = value;
    return this;
  }

  /** Convert entries to plain object */
  toObject(): Record<string, any> {
    return Object.fromEntries(this.entries());
  }

  // eslint-disable-next-line no-unused-private-class-members
  protected getTyped(key: string, type: string) {
    const val = this.get(key);
    if (val == null) return null;
    const t = type.toLowerCase();
    if (t.includes("number")) return Number(val);
    if (t.includes("bool")) return val === "true";
    return val;
  }
}
