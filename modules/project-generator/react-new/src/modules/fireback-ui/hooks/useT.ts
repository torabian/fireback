import { enTranslations } from "../../fireback/translations/en";
import { useLocale } from "./useLocale";
import { faTranslations } from "../../fireback/translations/fa";

const locales: any = {
  en: enTranslations,
  fa: faTranslations,
};

export function useT(): typeof enTranslations {
  const { locale } = useLocale();

  if (!locale || !locales[locale]) {
    return enTranslations;
  }

  return locales[locale];
}
