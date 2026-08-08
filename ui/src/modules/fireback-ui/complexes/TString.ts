/**
 * TypeScript counterpart of modules/fireback/complexes/TString.go.
 *
 * A "translatable string": a locale -> value map, e.g.
 * { en: "Hello", fa: "سلام" }, serialized to/from the API as either that
 * object or, if the caller doesn't care about translations, a bare string -
 * which is always read as DEFAULT_LOCALE's value. This mirrors exactly how
 * the Go side accepts a plain string or a {locale: value} object for any
 * TString field.
 */

export const DEFAULT_LOCALE = "en";

export type TStringInput =
  | string
  | Record<string, string>
  | TString
  | null
  | undefined;

export class TString {
  private values: Record<string, string>;

  constructor(input?: TStringInput) {
    this.values = TString.toRecord(input);
  }

  /** Normalizes any accepted input shape into a plain locale->value record. */
  private static toRecord(input: TStringInput): Record<string, string> {
    if (input == null) {
      return {};
    }
    if (typeof input === "string") {
      return { [DEFAULT_LOCALE]: input };
    }
    if (input instanceof TString) {
      return { ...input.values };
    }
    return { ...input };
  }

  /** Same as `new TString(input)`, useful in functional/pipe-style code. */
  static from(input: TStringInput): TString {
    return new TString(input);
  }

  /**
   * Returns the value for `locale`, falling back to DEFAULT_LOCALE, then to
   * any single value present, then "".
   */
  get(locale: string = DEFAULT_LOCALE): string {
    if (locale in this.values) {
      return this.values[locale];
    }
    if (DEFAULT_LOCALE in this.values) {
      return this.values[DEFAULT_LOCALE];
    }
    const first = Object.values(this.values)[0];
    return first ?? "";
  }

  /** Sets locale's value and returns `this`, so calls can be chained. */
  set(locale: string, value: string): this {
    this.values[locale] = value;
    return this;
  }

  has(locale: string): boolean {
    return locale in this.values;
  }

  /** Every locale code that currently has a value set. */
  locales(): string[] {
    return Object.keys(this.values);
  }

  /** Same shape TString.Go's Value()/MarshalJSON produce - a plain object. */
  toRecord(): Record<string, string> {
    return { ...this.values };
  }

  /** So JSON.stringify(entity) sends the same {locale: value} shape as Go. */
  toJSON(): Record<string, string> {
    return this.toRecord();
  }

  /** So template strings / string concatenation use get(DEFAULT_LOCALE). */
  toString(): string {
    return this.get(DEFAULT_LOCALE);
  }

  isEmpty(): boolean {
    return Object.keys(this.values).length === 0;
  }
}
