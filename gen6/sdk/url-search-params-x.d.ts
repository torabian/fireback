/**
 * Extended URLSearchParams that stores data in a nested object
 * and keeps compatibility with URLSearchParams methods.
 */
export declare class URLSearchParamsX extends URLSearchParams {
    /** Internal data store */
    private data;
    constructor(init?: string[][] | Record<string, string> | string | URLSearchParams);
    /** Remove a key from the store */
    delete(name: string): void;
    /** Append a value to an array or create a new array */
    append(name: string, value: string): void;
    /** Get an iterator of top-level keys */
    keys(): Generator<string, void, unknown>;
    /** Number of top-level keys */
    get size(): number;
    /** Sort top-level keys */
    sort(): void;
    /** Get an iterator of top-level values */
    values(): Generator<any, void, unknown>;
    /** Get a single value by key */
    get(name: string): string | null;
    /** Check if key exists */
    has(name: string): boolean;
    /** Iterate over top-level keys and values */
    forEach(callbackfn: (value: any, key: string, parent: any) => void): void;
    /** Get all values for a key as array */
    getAll(name: string): string[];
    /** Get an iterator of key/value pairs (flattened) */
    entries(): URLSearchParamsIterator<[string, string]>;
    /** Convert to query string */
    toString(): string;
    /** Set a key to a value */
    set(name: string, value: any): this;
    /** Convert entries to plain object */
    toObject(): Record<string, any>;
}
