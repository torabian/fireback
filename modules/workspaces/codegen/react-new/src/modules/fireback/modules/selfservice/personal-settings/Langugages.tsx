import { KeyValue } from "../../../definitions/definitions";
import { enTranslations } from "../../../translations/en";

export const interfaceLanguages = (t: typeof enTranslations): KeyValue[] => [
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("en")
  {
    label: t.locale.englishWorldwide,
    value: "en",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("fa")
  {
    label: t.locale.persianIran,
    value: "fa",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("ru")
  {
    label: "Russian (Русский)",
    value: "ru",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("pl")
  {
    label: t.locale.polishPoland,
    value: "pl",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("ua")
  {
    label: "Ukrainain (українська)",
    value: "ua",
  },
  /// #endif
];
