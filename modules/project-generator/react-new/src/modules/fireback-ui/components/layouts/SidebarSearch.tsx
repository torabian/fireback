import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export default function SidebarSearch() {
  const s = useS(strings);
  return (
    <form action="/" method="GET" className="search-form">
      <input
        type="search"
        placeholder={s.components.findASetting}
        className="search-field"
      />
      <button type="submit" className="search-button">
        <img src="/adp/icons/mglass.svg" />
      </button>
    </form>
  );
}
