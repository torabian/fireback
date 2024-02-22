import { localizeNumber } from "@/hooks/fonts";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";

export function NotFound404() {
  const t = useT();
  const { locale } = useLocale();
  return (
    <>
      <div className="not-found-page">
        <div className="content">
          <p>{t.not_found_404}</p>
          <div className="font-404">
            <h1>
              {localizeNumber("4", locale)}
              <span>{localizeNumber("0", locale)}</span>
              {localizeNumber("4", locale)}
            </h1>
          </div>
        </div>
      </div>
    </>
  );
}
