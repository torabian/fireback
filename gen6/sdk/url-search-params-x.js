import { stringify, parse } from "qs";
/**
 * Extended URLSearchParams that stores data in a nested object
 * and keeps compatibility with URLSearchParams methods.
 */
export class URLSearchParamsX extends URLSearchParams {
    constructor(init) {
        super(init);
        /** Internal data store */
        this.data = {};
        if (init) {
            if (typeof init === "string") {
                Object.assign(this.data, parse(init));
            }
            else if (init instanceof URLSearchParams) {
                Object.assign(this.data, parse(init.toString()));
            }
            else if (Array.isArray(init)) {
                init.forEach(([k, v]) => (this.data[k] = v));
            }
            else {
                Object.assign(this.data, init);
            }
        }
    }
    /** Remove a key from the store */
    delete(name) {
        delete this.data[name];
    }
    /** Append a value to an array or create a new array */
    append(name, value) {
        if (this.data[name] === undefined)
            this.data[name] = value;
        else if (Array.isArray(this.data[name]))
            this.data[name].push(value);
        else
            this.data[name] = [this.data[name], value];
    }
    /** Get an iterator of top-level keys */
    keys() {
        return (function* (obj) {
            for (const key of Object.keys(obj))
                yield key;
        })(this.data);
    }
    /** Number of top-level keys */
    get size() {
        return Object.keys(this.data).length;
    }
    /** Sort top-level keys */
    sort() {
        const sorted = {};
        Object.keys(this.data)
            .sort()
            .forEach((key) => {
            sorted[key] = this.data[key];
        });
        this.data = sorted;
    }
    /** Get an iterator of top-level values */
    values() {
        return (function* (obj) {
            for (const key of Object.keys(obj))
                yield obj[key];
        })(this.data);
    }
    /** Get a single value by key */
    get(name) {
        const val = this.data[name];
        if (val == null)
            return null;
        return Array.isArray(val) ? String(val[0]) : String(val);
    }
    /** Check if key exists */
    has(name) {
        return this.data[name] !== undefined;
    }
    /** Iterate over top-level keys and values */
    forEach(callbackfn) {
        for (const key of Object.keys(this.data)) {
            callbackfn(this.data[key], key, this);
        }
    }
    /** Get all values for a key as array */
    getAll(name) {
        const val = this.data[name];
        if (val === undefined)
            return [];
        return Array.isArray(val) ? val.map(String) : [String(val)];
    }
    /** Get an iterator of key/value pairs (flattened) */
    entries() {
        const params = new URLSearchParams(stringify(this.data));
        return params.entries();
    }
    /** Convert to query string */
    toString() {
        return stringify(this.data);
    }
    /** Set a key to a value */
    set(name, value) {
        this.data[name] = value;
        return this;
    }
    /** Convert entries to plain object */
    toObject() {
        return Object.fromEntries(this.entries());
    }
}
