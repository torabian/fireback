export function localeFromPath(path: string) {
  let locale = "en";

  const match = path.match(/^\/(fa|en|ar|pl|de|ua|ru)\//);
  if (match && match[1]) {
    locale = match[1];
  }

  if (!["fa", "en", "ar", "pl", "de", "ru", "ua"].includes(locale)) {
    return "en";
  }

  return locale;
}
