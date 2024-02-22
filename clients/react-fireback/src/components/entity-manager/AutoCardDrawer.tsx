import { QueryArchiveColumn } from "@/definitions/common";
import Link from "../link/Link";

export function AutoCardDrawer({
  content,
  columns,
  uniqueIdHrefHandler,
  style,
}: {
  style: any;
  content: any;
  columns: QueryArchiveColumn[];
  uniqueIdHrefHandler?: (id: string) => void;
}) {
  return (
    <Link
      className="auto-card-list-item card mb-2 p-3"
      style={style}
      href={uniqueIdHrefHandler && uniqueIdHrefHandler(content.uniqueId)}
    >
      {columns.map((col) => {
        let v = col.getCellValue ? col.getCellValue(content) : "";
        if (!v) {
          v = col.name ? content[col.name] : "";
        }
        if (!v) {
          v = "-";
        }

        if (col.name === "uniqueId") {
          return null;
        }
        return (
          <div className="row" key={col.title}>
            <div className="col-5">{col.title}:</div>
            <div className="col-7">{v}</div>
          </div>
        );
      })}
    </Link>
  );
}
