import { localizeNumber } from "../../hooks/fonts";
import { source } from "../../hooks/source";
import { useLocale } from "../../hooks/useLocale";
import { useT } from "../../hooks/useT";

export function NotFound404() {
  const t = useT();
  const { locale } = useLocale();
  return (
    <>
      <div className="not-found-pagex">
        <img src={source("/common/error.svg")} />
        <div className="content">
          <p>{t.not_found_404}</p>
        </div>
      </div>
    </>
  );
}
