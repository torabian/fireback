import { BUILD_VARIABLES } from "@/modules/fireback/hooks/build-variables";
import { type KeyValue } from "../../../definitions/definitions";
import { enTranslations } from "../../../translations/en";

export const interfaceLanguages = (t: typeof enTranslations): KeyValue[] => [
  BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("en") ?
  {
    label: t.locale.englishWorldwide,
    value: "en",
  } : undefined,
  /// #endif
  BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("fa") ?
  {
    label: t.locale.persianIran,
    value: "fa",
  } : undefined,
  /// #endif
  BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("ru") ?
  {
    label: "Russian (Русский)",
    value: "ru",
  } : undefined,
  /// #endif
  BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("pl") ?
  {
    label: t.locale.polishPoland,
    value: "pl",
  } : undefined,
  /// #endif
  BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("ua") ?
  {
    label: "Ukrainain (українська)",
    value: "ua",
  } : undefined,
  /// #endif
].filter(Boolean);
