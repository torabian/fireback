import { createElement } from "react";
import { QueryArchiveColumn } from "../../definitions/common";
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
  const Component: any = uniqueIdHrefHandler ? Link : "span";
  return (
    <Component
      className="auto-card-list-item card mb-2 p-3"
      style={style}
      href={uniqueIdHrefHandler(content.uniqueId)}
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
          <div className="row auto-card-drawer" key={col.title}>
            <div className="col-6">{col.title}:</div>
            <div className="col-6">{v}</div>
          </div>
        );
      })}
    </Component>
  );
}
