import { QueryArchiveColumn } from "@/definitions/common";
import { Filter } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { use, useCallback, useEffect, useRef, useState } from "react";
import {
  PullDownContent,
  PullToRefresh,
  RefreshContent,
  ReleaseContent,
} from "react-js-pull-to-refresh";
import { useQueryClient } from "react-query";
import AutoSizer from "react-virtualized-auto-sizer";
import InfiniteLoader from "react-window-infinite-loader";
import { AutoCardDrawer } from "./AutoCardDrawer";
import { QueryErrorView } from "../error-view/QueryError";
const { FixedSizeList } = require("react-window");

export const FlatListMode = ({
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  udf,
  jsonQuery,
  q,
}: {
  udf: any;
  q: any;

  deleteHook?: any;
  columns: QueryArchiveColumn[];
  uniqueIdHrefHandler?: (id: string) => void;
  jsonQuery?: any;
}) => {
  const t = useT();
  // Used for cashing the query
  // const indexedData = useRef<Array<any>>([]);
  const previousQuery = useRef<any>();
  let [indexedData, setIndexedData] = useState<Array<any>>([]);
  const [pullToRefreshEnabled, setPTREnabled] = useState<boolean>(true);

  const queryClient = useQueryClient();

  const delHook =
    deleteHook &&
    deleteHook({
      queryClient,
    });

  const reindex = (rows: Array<any>, jsonQueryKey: string) => {
    const index = udf.debouncedFilters.startIndex || 0;
    if (previousQuery.current !== jsonQueryKey) {
      indexedData = [];
      previousQuery.current = jsonQueryKey;
    }

    for (
      let i = index;
      i < (udf.debouncedFilters.itemsPerPage || 0) + index;
      i++
    ) {
      let m = i;
      if (index > 0) {
        m -= index;
      }
      if (rows[m]) {
        indexedData[i] = rows[m];
      }
    }
    setIndexedData([...indexedData]);
    previousQuery.current = jsonQueryKey;
  };

  useEffect(() => {
    const rows: any = q.query.data?.data?.items || [];

    reindex(rows, q.query.data?.jsonQuery);
  }, [q.query.data?.data?.items]);

  const Item = ({ index, style }: any) => {
    const data = indexedData[index];

    if (!data) {
      return null;
    }

    return (
      <AutoCardDrawer
        key={indexedData[index]?.uniqueId}
        style={{
          ...style,
          top: style.top + 10,
          height: style.height - 10,
          width: style.width,
        }}
        uniqueIdHrefHandler={uniqueIdHrefHandler}
        columns={columns}
        content={indexedData[index]}
      />
    );
  };

  const onScroll = ({ scrollOffset }: any) => {
    if (scrollOffset === 0 && !pullToRefreshEnabled) {
      setPTREnabled(true);
    } else if (scrollOffset > 0 && pullToRefreshEnabled) {
      setPTREnabled(false);
    }
  };

  const onRefresh = useCallback(() => {
    q.query.refetch();

    return Promise.resolve(true);
  }, []);

  const total = q.query?.data?.data?.totalItems || 0;

  return (
    <>
      <PullToRefresh
        pullDownContent={<PullDownContent label="" />}
        // pullDownContent={<span />}
        releaseContent={<ReleaseContent label="" />}
        refreshContent={<RefreshContent label="" />}
        pullDownThreshold={200}
        onRefresh={onRefresh}
        // triggerHeight={200}
        triggerHeight={pullToRefreshEnabled ? 500 : 0}
        startInvisible={true}
      >
        <div style={{ height: "calc(100vh - 140px)" }}>
          <QueryErrorView query={q.query} />

          {indexedData.length === 0 && !q.query?.isError && (
            <span>{t.table.noRecords}</span>
          )}
          <InfiniteLoader
            isItemLoaded={(index) => {
              return !!indexedData[index];
            }}
            itemCount={total}
            loadMoreItems={async (startIndex, stopIndex) => {
              udf.setFilter({
                startIndex,
                itemsPerPage: stopIndex - startIndex,
              });
            }}
          >
            {({ onItemsRendered, ref }) => (
              <AutoSizer>
                {({ height, width }: any) => (
                  <FixedSizeList
                    height={height}
                    itemCount={indexedData.length}
                    itemSize={columns.length * 25 + 10}
                    width={width}
                    onScroll={onScroll}
                    onItemsRendered={onItemsRendered}
                    ref={ref}
                  >
                    {Item}
                  </FixedSizeList>
                )}
              </AutoSizer>
            )}
          </InfiniteLoader>
        </div>
      </PullToRefresh>
    </>
  );
};
