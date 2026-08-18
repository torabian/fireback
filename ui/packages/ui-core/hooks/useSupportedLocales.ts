import { BUILD_VARIABLES } from "./build-variables";

/**
 * The locales this build ships translations for, e.g. ["en", "fa"] - parsed from
 * BUILD_VARIABLES.SUPPORTED_LANGUAGES (VITE_SUPPORTED_LANGUAGES), the same env var
 * the rest of the app's own i18n is built from. Shared by every "one field per
 * language" editor of a TString value (TStringFilterDrawer's column filter,
 * FormTString's edit modal, ...) so they all offer the same locale list instead of
 * each re-parsing the env var themselves.
 */
export function useSupportedLocales(): string[] {
  return (BUILD_VARIABLES.SUPPORTED_LANGUAGES || "en")
    .split(",")
    .map((l) => l.trim())
    .filter(Boolean);
}
