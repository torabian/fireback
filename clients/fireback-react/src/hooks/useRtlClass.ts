import { useLocale } from "@/hooks/useLocale";
import { useEffect } from "react";

export function useRtlClass() {
  const { locale, asPath } = useLocale();

  useEffect(() => {
    document
      .querySelector("html")
      ?.setAttribute("dir", ["fa", "ar"].includes(locale) ? "rtl" : "ltr");
  }, [asPath]);
}
