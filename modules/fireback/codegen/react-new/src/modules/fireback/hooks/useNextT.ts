import { enTranslations } from "../translations/en";
import { faTranslations } from "../translations/fa";
import { useRouter } from "next/router";


const locales: any = {
  en: enTranslations,
  fa: faTranslations,
};

export function useT(): typeof enTranslations {
  let locale = "en";
  const router = useRouter();

  if (router.route.includes("/fa")) {
    locale = "fa";
  }

  return locales[locale];
}
