import { useEffect, useState } from "react";

export function localeFromPath(path: string) {
  let locale = "en";

  const match = path.match(/\/(fa|en|ar|pl|de)\//);
  if (match && match[1]) {
    locale = match[1];
  }

  return locale;
}

export function useWindowHash() {
  const [hash, setHash] = useState(window.location.toString());
  useEffect(() => {
    const handleLocationChange = () => {
      setHash(window.location.hash);
    };

    window.addEventListener("popstate", handleLocationChange);
    window.addEventListener("pushState", handleLocationChange); // Custom event for pushState
    window.addEventListener("replaceState", handleLocationChange); // Custom event for replaceState

    return () => {
      window.removeEventListener("popstate", handleLocationChange);
      window.removeEventListener("pushState", handleLocationChange);
      window.removeEventListener("replaceState", handleLocationChange);
    };
  }, []);

  return { hash };
}

export function usePureLocale() {
  const { hash } = useWindowHash();
  let locale = "en";
  let region = "us";
  let dir = "ltr";

  if (process.env.REACT_APP_FORCED_LOCALE) {
    locale = process.env.REACT_APP_FORCED_LOCALE;
  } else {
    locale = localeFromPath(hash);
  }

  if (locale === "fa") {
    region = "ir";
    dir = "rtl";
  }

  return { locale, region, dir };
}
