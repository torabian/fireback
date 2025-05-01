import { useRef, useState } from "react";
import { Udf } from "../../hooks/useDatatableFiltering";

export function useReindexedContent(udf: Udf) {
  const previousQuery = useRef<any>();
  let [indexedData, setIndexedData] = useState<Array<any>>([]);

  // Keep a uniqueid reference to make sure no duplicates will ever be shown in case
  // of another mistake somewhere else.
  const keyref = useRef<any>({});
  const reindex = (
    rows: Array<any>,
    jsonQueryKey: string,
    onKeyChange?: () => void
  ) => {
    if (jsonQueryKey === previousQuery.current) {
      const toAdd = rows.filter((row) => {
        if (!keyref.current[row.uniqueId]) {
          keyref.current[row.uniqueId] = true;
          return true;
        }

        return false;
      });
      setIndexedData([...indexedData, ...toAdd].filter(Boolean));
    } else {
      setIndexedData([...rows].filter(Boolean));
      onKeyChange?.();
    }

    previousQuery.current = jsonQueryKey;
  };

  return { reindex, indexedData };
}
