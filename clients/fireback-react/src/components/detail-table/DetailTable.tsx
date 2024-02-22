import { useT } from "@/hooks/useT";
import Link from "../link/Link";

export const CommonRowDetail = ({ row, uniqueIdHrefHandler, columns }: any) => {
  const t = useT();
  return (
    <>
      {(row.children || []).map((item: any) => {
        return (
          <tr>
            <td></td>
            <td></td>
            {columns(t).map((col: any) => {
              let v = item.getCellValue
                ? item.getCellValue(row)
                : item[col.name];
              if (col.name === "uniqueId") {
                return (
                  <td>
                    <Link href={uniqueIdHrefHandler && uniqueIdHrefHandler(v)}>
                      {v}
                    </Link>
                  </td>
                );
              }
              return <td>{v}</td>;
            })}
          </tr>
        );
      })}
    </>
  );
};
