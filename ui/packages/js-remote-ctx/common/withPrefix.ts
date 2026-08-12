export function withPrefix<T extends Record<string, any>>(
  prefix: string,
  fields: T
): T {
  const out: Record<string, any> = {};
  for (const [k, v] of Object.entries(fields)) {
    if (typeof v === "string") {
      out[k] = `${prefix}.${v}`;
    } else if (typeof v === "object" && v !== null) {
      out[k] = v;
    }
  }
  return out as T;
}

export function at(source: string, ...args: number[]): string {
  args.forEach((item) => {
    source = source.replace("[:i]", `[${item}]`);
  });

  return source;
}
