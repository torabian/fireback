import type { FileValidationRule } from "./types";

export function buildAcceptString(rules?: FileValidationRule[]): string | undefined {
  if (!rules?.length) return undefined;

  const parts = rules.map((rule) => {
    if (rule.extension) return rule.extension;
    if (rule.mimeStartsWith?.endsWith("/")) return rule.mimeStartsWith + "*";
    return rule.mimeStartsWith || "";
  });

  return parts.filter(Boolean).join(",") || undefined;
}

export function formatMB(bytes: number): string {
  return `${Math.round(bytes / 1024 / 1024)}MB`;
}

/** Largest maxSize configured across all rules, for a single "max file size" hint shown up front. */
export function getOverallMaxSizeHint(rules?: FileValidationRule[]): string | null {
  if (!rules?.length) return null;
  const sizes = rules.map((rule) => rule.maxSize).filter((size): size is number => !!size);
  if (!sizes.length) return null;
  return formatMB(Math.max(...sizes));
}

/** The maxSize of the rule that would apply to this specific file, regardless of pass/fail. */
export function getApplicableMaxSize(
  file: File,
  rules?: FileValidationRule[],
): number | null {
  if (!rules?.length) return null;
  for (const rule of rules) {
    const matchesType = !rule.mimeStartsWith || file.type.startsWith(rule.mimeStartsWith);
    if (matchesType) return rule.maxSize ?? null;
  }
  return null;
}

export type FileValidationFailure =
  | { code: "type" }
  | { code: "size"; maxSize: number };

export function validateFileAgainstRules(
  file: File,
  rules?: FileValidationRule[],
): FileValidationFailure | null {
  if (!rules?.length) return null;

  for (const rule of rules) {
    const matchesType = !rule.mimeStartsWith || file.type.startsWith(rule.mimeStartsWith);
    if (matchesType) {
      if (rule.maxSize && file.size > rule.maxSize) {
        return { code: "size", maxSize: rule.maxSize };
      }
      return null;
    }
  }

  return { code: "type" };
}
