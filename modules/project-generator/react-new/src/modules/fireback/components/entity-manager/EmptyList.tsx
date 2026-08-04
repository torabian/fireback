import { source } from "../../hooks/source";
import { useT } from "../../hooks/useT";

export const EmptyList = () => {
  const t = useT();
  return (
    <div className="empty-list-indicator">
      <img src={source("/common/empty.png")} />
      <div>{t.table.noRecords}</div>
    </div>
  );
};
