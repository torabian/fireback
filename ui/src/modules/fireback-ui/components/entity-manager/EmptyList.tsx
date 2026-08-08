import { source } from "../../hooks/source";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export const EmptyList = () => {
  const s = useS(strings);
  return (
    <div className="empty-list-indicator">
      <img src={source("/common/empty.png")} />
      <div>{s.table.noRecords}</div>
    </div>
  );
};
