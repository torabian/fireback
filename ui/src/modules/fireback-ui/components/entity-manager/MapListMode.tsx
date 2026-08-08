import { type QueryArchiveColumn } from "../../types/QueryArchiveColumn";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import { useQueryClient } from "@tanstack/react-query";
import { AutoCardDrawer } from "./AutoCardDrawer";

export const MapListMode = ({
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  udf,
  q,
}: {
  udf: any;
  q: any;
  deleteHook?: any;
  columns: QueryArchiveColumn[];
  uniqueIdHrefHandler?: (id: string) => void;
}) => {
  const s = useS(strings);

  const queryClient = useQueryClient();

  const delHook =
    deleteHook &&
    deleteHook({
      queryClient,
    });

  const rows: any = q.query.data?.data?.items || [];

  const items = q.query?.data?.data?.items || [];
  const total = q.query?.data?.data?.totalItems || 0;

  return (
    <>
      {total === 0 && <p>{s.table.noRecords}</p>}
      {(items || []).map((item: any) => {
        return (
          <AutoCardDrawer
            key={item.uniqueId}
            style={{}}
            uniqueIdHrefHandler={uniqueIdHrefHandler}
            columns={columns}
            content={item}
          />
        );
      })}
    </>
  );
};
