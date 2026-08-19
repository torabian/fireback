import { useLocale } from "./useLocale";
import { useEffect } from "react";

export function useRtlClass() {
  const { locale } = useLocale();

  useEffect(() => {
    document
      .querySelector("html")
      ?.setAttribute("dir", ["fa", "ar"].includes(locale) ? "rtl" : "ltr");
  }, [locale]);
}
