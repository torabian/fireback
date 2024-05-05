import { useLocale } from "./useLocale";

import { enTranslations } from "../../translations/en";
import { faTranslations } from "../../translations/fa";

const locales: any = {
  en: enTranslations,
  fa: faTranslations,
};

export function useT(): typeof enTranslations {
  const { locale } = useLocale();

  if (!locale) {
    return enTranslations;
  }

  return locales[locale];
}
