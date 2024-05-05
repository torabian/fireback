import { localizeNumber } from "@/fireback/hooks/fonts";
import { useLocale } from "@/fireback/hooks/useLocale";
import { useT } from "@/fireback/hooks/useT";

export function NotFound404() {
  const t = useT();
  const { locale } = useLocale();
  return (
    <>
      <div className="not-found-pagex">
        <div className="content">
          <p>{t.not_found_404}</p>
        </div>
      </div>
    </>
  );
}
