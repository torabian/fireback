import { useRef, useState } from "react";

export function useReindexedContent(udf: any) {
  const previousQuery = useRef<any>();
  const keyref = useRef<any>({});
  let [indexedData, setIndexedData] = useState<Array<any>>([]);

  const reindex = (rows: Array<any>, jsonQueryKey: string) => {
    const toAdd = rows.filter((row) => {
      if (!keyref.current[row.uniqueId]) {
        keyref.current[row.uniqueId] = true;
        return true;
      }

      return false;
    });

    setIndexedData([...indexedData, ...toAdd].filter(Boolean));
    previousQuery.current = jsonQueryKey;
  };

  return { reindex, indexedData };
}
