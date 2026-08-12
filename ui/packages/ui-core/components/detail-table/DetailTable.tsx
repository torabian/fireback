import Link from "../link/Link";

// `columns` is a zero-argument function that returns the already-resolved column
// definitions - callers close over their own module-local translation strings (each
// module has its own shape now, there's no shared central translation object anymore)
// and pass e.g. `columns={() => columns(s, uiS)}` rather than a raw, still-needs-t
// function reference.
export const CommonRowDetail = ({ row, uniqueIdHrefHandler, columns }: any) => {
  return (
    <>
      {(row.children || []).map((item: any) => {
        return (
          <tr>
            <td></td>
            <td></td>
            {columns().map((col: any) => {
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
