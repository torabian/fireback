import { useRouter } from "@/Router";

export function localeFromPath(path: string) {
  let locale = "en";

  const match = path.match(/^\/(fa|en|ar|pl|de)\//);
  if (match && match[1]) {
    locale = match[1];
  }

  return locale;
}

export function useLocale() {
  const router = useRouter();

  let locale = "en";
  let region = "us";
  let dir = "ltr";

  if (process.env.REACT_APP_FORCED_LOCALE) {
    locale = process.env.REACT_APP_FORCED_LOCALE;
  } else if (router.query.locale) {
    locale = `${router.query.locale}`;
  } else {
    locale = localeFromPath(router.asPath);
  }

  if (locale === "fa") {
    region = "ir";
    dir = "rtl";
  }

  return { locale, asPath: router.asPath, region, dir };
}
