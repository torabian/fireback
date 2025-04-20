import {useState} from 'react';
import {useDebouncedEffect} from './useDebouncedEffect';

export declare type FilterOperation = string;

export interface Filter {
  /** Specifies the name of a column whose value is used for filtering. */
  columnName: string;
  /** Specifies the operation name. The value is 'contains' if the operation name is not set. */
  operation?: FilterOperation;
  /** Specifies the filter value. */
  value?: any;
}

export declare type SortingDirection = 'asc' | 'desc';

/** Describes the sorting applied to a column */
export interface Sorting {
  /** Specifies a column's name to which the sorting is applied. */
  columnName: string;
  /** Specifies a column's sorting order. */
  direction: SortingDirection;
}

export interface Filters {
  itemsPerPage: number;
  startIndex: number;
  sorting?: Sorting[];
}

export interface DataFiltering {
  urlMask?: string;
  onRecordsDeleted?: () => void;
  submitDelete?: any;
  initialFilters?: Partial<Filters>;
}

export interface DataFilteringResult {
  filters: Partial<Filters>;
  setFilters: React.Dispatch<React.SetStateAction<Partial<Filters>>>;
  setFilter: (
    newFiltersObj: Partial<Filters>,
    withAddressBarChang?: boolean,
  ) => void;
  setSorting: (sorting: Sorting[] | undefined) => void;
  selection: string[];
  setStartIndex: (index: number) => void;
  setSelection: (selection: string[]) => void;
  onFiltersChange: (filters: Filter[] | undefined) => void;
  setPageSize: (page: number) => void;
  debouncedFilters: Partial<Filters>;
  increaseIndex: (amount: number) => void;
}

export function useDatatableFiltering({initialFilters}: DataFiltering) {
  const {withDebounce} = useDebouncedEffect();
  const init = {
    itemsPerPage: 15,
    startIndex: 0,
    sorting: [],
    ...(initialFilters || {}),
  };

  const [filters, setFilters] = useState<Partial<Filters>>(init);
  const [debouncedFilters, setDebouncedFilters] =
    useState<Partial<Filters>>(init);

  const [selection, setSelection$] = useState<Array<string>>([]);

  const setSelection = (selection: string[]) => {
    setSelection$(selection);
  };

  const setFilter = (
    newFiltersObj: Partial<Filters>,
    withAddressBarChang = true,
  ) => {
    const newFilters = {
      ...filters,
      ...newFiltersObj,
    };

    const reflectToAddressBar: Object = {
      ...newFilters,
      sorting: newFilters.sorting && JSON.stringify(newFilters.sorting),
    };

    if (withAddressBarChang) {
      withDebounce(() => setLocationWithFilters(reflectToAddressBar), 250);
    }
    setFilters(newFilters);
    withDebounce(() => setDebouncedFilters(newFilters), 500);
  };

  const setLocationWithFilters = (filters: Object) => {
    const searchParams = new URLSearchParams();
    const params = filters as any;
    Object.keys(params).forEach(
      key => params[key] !== undefined && searchParams.append(key, params[key]),
    );

    // const q = searchParams.toString();
    // router.push(
    //   `/${locale}/${urlMask}`.replace("//", "/"),
    //   `/${locale}/${urlMask}?${q}`.replace("//", "/"),
    //   {
    //     shallow: true,
    //   }
    // );
  };

  const setPageSize = (page: number) => {
    setFilter({itemsPerPage: page});
  };

  const setSorting = (sorting: Sorting[] | undefined) => {
    setFilter({sorting});
  };

  const setStartIndex = (index: number) => {
    setFilter({startIndex: index});
  };

  const increaseIndex = (amount: number) => {
    setFilter({startIndex: (filters.startIndex || 0) + amount});
  };

  const onFiltersChange = (filters: Filter[] | undefined) => {
    let newFilters = {rawFilters: filters, startIndex: 0};
    setFilter(newFilters);
  };

  return {
    filters,
    setFilters,
    setFilter,
    setSorting,
    setStartIndex,
    selection,
    setSelection,
    onFiltersChange,
    increaseIndex,
    setPageSize,
    debouncedFilters,
  };
}
