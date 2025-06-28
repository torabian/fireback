import { QueryArchiveColumn } from "../../definitions/common";
import { useT } from "../../hooks/useT";
import { FC, ReactNode, useCallback, useEffect, useRef, useState } from "react";
import {
  PullDownContent,
  PullToRefresh,
  RefreshContent,
  ReleaseContent,
} from "../../thirdparty/react-pull-to-refresh";
import { useQueryClient } from "react-query";
import AutoSizer from "react-virtualized-auto-sizer";
import InfiniteLoader from "react-window-infinite-loader";
import { QueryErrorView } from "../error-view/QueryError";
import { AutoCardDrawer } from "./AutoCardDrawer";
import { EmptyList } from "./EmptyList";
const { FixedSizeList } = require("react-window");

// Define the props
interface CardProps<T> {
  content: T;
}

// Extend FC with static methods
export interface CardComponentType<T> extends FC<CardProps<T>> {
  getHeight: () => number;
}

export const FlatListMode = ({
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  udf,
  jsonQuery,
  q,
  CardComponent,
}: {
  udf: any;
  q: any;

  deleteHook?: any;
  columns: QueryArchiveColumn[];
  uniqueIdHrefHandler?: (id: string) => void;
  jsonQuery?: any;
  CardComponent?: CardComponentType<unknown>;
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
    const newData = [...indexedData]; // shallow copy

    if (previousQuery.current !== jsonQueryKey) {
      newData.length = 0; // reset immutably
      previousQuery.current = jsonQueryKey;
    }

    for (
      let i = index;
      i < (udf.debouncedFilters.itemsPerPage || 0) + index;
      i++
    ) {
      const m = i - index;
      if (rows[m]) {
        newData[i] = rows[m];
      }
    }

    setIndexedData(newData);
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

    if (CardComponent) {
      return (
        <CardComponent
          key={indexedData[index]?.uniqueId}
          content={indexedData[index]}
        />
      );
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
        releaseContent={<ReleaseContent />}
        refreshContent={<RefreshContent />}
        pullDownThreshold={200}
        onRefresh={onRefresh}
        // triggerHeight={200}
        triggerHeight={pullToRefreshEnabled ? 500 : 0}
        startInvisible={true}
      >
        {indexedData.length === 0 && !q.query?.isError ? (
          <div style={{ height: "calc(100vh - 130px)" }}>
            <EmptyList />
          </div>
        ) : (
          <div style={{ height: "calc(100vh - 130px)" }}>
            <QueryErrorView query={q.query} />

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
                      itemSize={
                        CardComponent?.getHeight
                          ? CardComponent.getHeight()
                          : columns.length * 24 + 10
                      }
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
        )}
      </PullToRefresh>
    </>
  );
};
