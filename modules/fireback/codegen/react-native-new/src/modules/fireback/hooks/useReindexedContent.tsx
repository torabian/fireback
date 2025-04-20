import {useRef, useState} from 'react';

import {DataFilteringResult} from '@/modules/fireback/hooks/useDatatableFiltering';

export function useReindexedContent(udf: DataFilteringResult) {
  const previousQuery = useRef<any>();
  let [indexedData, setIndexedData] = useState<Array<any>>([]);

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
    setIndexedData([...indexedData].filter(Boolean));
    previousQuery.current = jsonQueryKey;
  };

  return {reindex, indexedData};
}
