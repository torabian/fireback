import { BUILD_VARIABLES } from "./build-variables";

/**
 * Every locale code any part of the app has ever recognized (union of the old,
 * now-removed per-file `localeFromPath` regexes that used to pull this out of
 * the URL - see useLocale.ts/usePureLocale.tsx/localeFromPath.ts's own history).
 * A stored value outside this set (a stray value from an old build, a
 * hand-edited localStorage, ...) is treated as absent rather than trusted.
 */
export const KNOWN_LOCALES = ["en", "fa", "ar", "pl", "de", "ua", "ru"];

export const LOCALE_STORAGE_KEY = "app_locale";

function supportedLocales(): string[] {
  const configured = (BUILD_VARIABLES.SUPPORTED_LANGUAGES || "")
    .split(",")
    .map((s) => s.trim())
    .filter(Boolean);

  // A build that never set VITE_SUPPORTED_LANGUAGES falls back to every known
  // locale rather than none, so this never silently rejects everything.
  return configured.length > 0
    ? KNOWN_LOCALES.filter((l) => configured.includes(l))
    : KNOWN_LOCALES;
}

function isSupportedLocale(value: unknown): value is string {
  return typeof value === "string" && supportedLocales().includes(value);
}

function readStoredLocale(): string | null {
  if (typeof window === "undefined") {
    return null;
  }
  try {
    return window.localStorage.getItem(LOCALE_STORAGE_KEY);
  } catch {
    // Storage disabled/unavailable (private browsing, etc.) - fall back below.
    return null;
  }
}

/**
 * Resolution order: a build forced to a single locale always wins (existing
 * VITE_FORCED_LOCALE behavior, unchanged) - otherwise whatever the user
 * previously chose (persisted by setLocale, read back here), validated
 * against what this build actually supports - otherwise the build's own
 * default.
 */
function resolveLocale(): string {
  if (BUILD_VARIABLES.FORCED_LOCALE) {
    return BUILD_VARIABLES.FORCED_LOCALE;
  }

  const stored = readStoredLocale();
  if (isSupportedLocale(stored)) {
    return stored;
  }

  return BUILD_VARIABLES.DEFAULT_LOCALE;
}

let currentLocale = resolveLocale();
const listeners = new Set<() => void>();

/** Snapshot getter for useSyncExternalStore - must return the same reference
 * (here, a primitive string, so equality is enough) until something actually
 * changes, which is exactly what the module-level `currentLocale` gives us. */
export function getLocale(): string {
  return currentLocale;
}

/**
 * Changes the app's locale everywhere at once: persists it (so it survives a
 * reload/new tab) and notifies every mounted useLocale()/usePureLocale()
 * instance to re-render, regardless of which component called this - there's
 * no Provider to wire up, any component can just import and call it (see
 * personal-settings/InterfaceSettings.tsx for the actual switcher UI).
 * No-ops under a forced locale, the same as passing an unsupported value.
 */
export function setLocale(locale: string): void {
  if (BUILD_VARIABLES.FORCED_LOCALE || !isSupportedLocale(locale)) {
    return;
  }
  if (locale === currentLocale) {
    return;
  }
  currentLocale = locale;
  try {
    window.localStorage.setItem(LOCALE_STORAGE_KEY, locale);
  } catch {
    // Storage disabled/unavailable - the in-memory value above still applies
    // for the rest of this session, it just won't survive a reload.
  }
  listeners.forEach((listener) => listener());
}

export function subscribeToLocale(listener: () => void): () => void {
  listeners.add(listener);
  return () => listeners.delete(listener);
}
