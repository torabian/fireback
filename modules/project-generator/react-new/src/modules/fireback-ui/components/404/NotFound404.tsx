import { source } from "../../hooks/source";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export function NotFound404() {
  const s = useS(strings);
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
