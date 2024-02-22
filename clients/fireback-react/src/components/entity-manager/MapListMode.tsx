import { QueryArchiveColumn } from "@/definitions/common";
import { useT } from "@/hooks/useT";
import { useQueryClient } from "react-query";
import { AutoCardDrawer } from "./AutoCardDrawer";
import ReactPullToRefresh from "react-pull-to-refresh";

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
  const t = useT();

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
      {total === 0 && <p>{t.table.noRecords}</p>}
      {items.map((item: any) => {
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
