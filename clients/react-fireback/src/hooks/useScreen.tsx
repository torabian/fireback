import { useRouter } from "@/Router";
import { useT } from "./useT";
import { useLocale } from "./useLocale";
import { useQueryClient } from "react-query";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { enTranslations } from "@/translations/en";

export function useScreen(fn: (t: typeof enTranslations) => string) {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();

  const queryClient = useQueryClient();

  usePageTitle(fn(t));

  return { t, router, locale, queryClient };
}
