import { BUILD_VARIABLES } from "../../fireback-ui/hooks/build-variables";
import { type KeyValue } from "../../fireback-ui/types/KeyValue";
import { strings } from "./strings/translations";

export const interfaceLanguages = (s: typeof strings): KeyValue[] =>
  [
    BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("en")
      ? {
          label: s.locale.englishWorldwide,
          value: "en",
        }
      : undefined,
    /// #endif
    BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("fa")
      ? {
          label: s.locale.persianIran,
          value: "fa",
        }
      : undefined,
    /// #endif
    BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("ru")
      ? {
          label: s.languages.russian,
          value: "ru",
        }
      : undefined,
    /// #endif
    BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("pl")
      ? {
          label: s.locale.polishPoland,
          value: "pl",
        }
      : undefined,
    /// #endif
    BUILD_VARIABLES.SUPPORTED_LANGUAGES.includes("ua")
      ? {
          label: s.languages.ukrainian,
          value: "ua",
        }
      : undefined,
    /// #endif
  ].filter(Boolean);
