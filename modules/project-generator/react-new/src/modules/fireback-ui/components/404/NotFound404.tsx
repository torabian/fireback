import { localizeNumber } from "../../hooks/fonts";
import { source } from "../../hooks/source";
import { useLocale } from "../../hooks/useLocale";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export function NotFound404() {
  const s = useS(strings);
  const { locale } = useLocale();
  return (
    <>
      <div className="not-found-pagex">
        <img src={source("/common/error.svg")} />
        <div className="content">
          <p>{s.not_found_404}</p>
        </div>
      </div>
    </>
  );
}
