import { Filter, Sorting } from "@devexpress/dx-react-grid";
import { parse, stringify } from "qs";
import { useContext, useEffect, useRef, useState } from "react";
import { useLocation } from "react-router-dom";
import {
  IMenuActionItem,
  useMenuTools,
} from "../components/action-menu/ActionMenu";
import { ModalContext } from "../components/modal/Modal";
import { commonDialogs } from "../components/overlay/CommonOverlays";
import { KeyboardAction } from "../definitions/definitions";
import { Filters } from "../hooks/datatabletools";
import { useRouter } from "../hooks/useRouter";
import { osResources } from "../resources/resources";
import { DeleteRequest } from "../sdk/core/http-tools";
import { useDebouncedEffect } from "./useDebouncedEffect";
import { useKeyCombination } from "./useKeyPress";
import { useT } from "./useT";

export function useDatatableFiltering({
  urlMask,
  submitDelete,
  onRecordsDeleted,
  initialFilters,
}: {
  urlMask: string;
  onRecordsDeleted?: () => void;
  submitDelete?: any;
  initialFilters?: Partial<Filters>;
}) {
  const t = useT();
  const router = useRouter();

  const { confirmModal } = commonDialogs();
  const { withDebounce } = useDebouncedEffect();
  const init = {
    itemsPerPage: 100,
    startIndex: 0,
    sorting: [],
    ...(initialFilters || {}),
  };

  const [filters, setFilters] = useState<Partial<Filters>>(init);
  const [debouncedFilters, setDebouncedFilters] =
    useState<Partial<Filters>>(init);

  const { search } = useLocation();

  const locked = useRef(false);
  useEffect(() => {
    if (locked.current) {
      return;
    }
    locked.current = true;

    let filters: Partial<Filters> = {};

    try {
      filters = parse(search.substring(1));

      // startIndex is being removed, and I am not sure if that's a good idea or not.
      // When user scrolls down a lot, and then tries to share the link with someone or refresh,
      // He doesn't have the content which was there before hand, therefor sees a empty screen
      delete filters.startIndex;
    } catch (error) {}

    setFilters({ ...init, ...filters });
    setDebouncedFilters({ ...init, ...filters });
  }, [search]);

  const [selection, setSelection$] = useState<Array<string>>([]);
  // const [queryHash, setQueryHash] = useState("{}");

  const computeQueryKey = (filters) => {
    const queryHashItems = { ...filters };
    delete queryHashItems.startIndex;
    delete queryHashItems.itemsPerPage;
    if (queryHashItems?.sorting?.length === 0) {
      delete queryHashItems.sorting;
    }
    return JSON.stringify(queryHashItems);
  };

  const queryHash = computeQueryKey(filters);

  const setSelection = (selection: string[]) => {
    setSelection$(selection);
  };

  const setFilter = (newFiltersObj: Partial<Filters>, reset = true) => {
    const newFilters = {
      ...filters,
      ...newFiltersObj,
    };

    if (reset) {
      newFilters.startIndex = 0;
    }

    setFilters(newFilters);
    // setQueryHash(computeQueryKey(newFilters));

    router.push("?" + stringify(newFilters), undefined, {}, true);
    withDebounce(() => {
      setDebouncedFilters(newFilters);
    }, 500);
  };

  const setPageSize = (page: number) => {
    setFilter({ itemsPerPage: page }, false);
  };

  const toSortString = (sorting: Sorting[]) => {
    return sorting
      .map((sort) => `${sort.columnName} ${sort.direction}`)
      .join(", ");
  };

  const setSorting = (sorting: Sorting[] | undefined) => {
    setFilter({ sorting, sort: toSortString(sorting) }, false);
  };

  const setStartIndex = (index: number) => {
    setFilter({ startIndex: index }, false);
  };

  const onFiltersChange = (filters: Filter[] | undefined) => {
    let newFilters = { startIndex: 0 };
    setFilter(newFilters);
  };

  const useModal = useContext(ModalContext);

  const idsToQuery = (items: string[]): DeleteRequest => {
    return {
      query: items.map((t) => `unique_id = ${t}`).join(" or "),
      uniqueId: "",
    };
  };

  const deleteItems = async () => {
    confirmModal({
      title: t.confirm,
      confirmLabel: t.common.yes,
      cancelLabel: t.common.no,
      description: t.deleteConfirmMessage,
    })
      .promise.then(({ type }) => {
        if (type === "resolved") {
          return submitDelete(idsToQuery(selection), null as any);
        }
      })
      .then(() => {
        onRecordsDeleted && onRecordsDeleted();
      });
  };

  const deleteAction = (): IMenuActionItem => ({
    label: t.deleteAction,
    onSelect() {
      deleteItems();
    },
    icon: osResources.delete,
    uniqueActionKey: "GENERAL_DELETE_ACTION",
  });

  const { addActions, removeActionMenu } = useMenuTools();

  useEffect(() => {
    if (selection.length > 0 && typeof submitDelete !== "undefined") {
      return addActions("table-selection", [deleteAction()]);
    } else {
      removeActionMenu("table-selection");
    }
  }, [selection]);

  useKeyCombination(KeyboardAction.Delete, () => {
    if (selection.length > 0 && typeof submitDelete !== "undefined") {
      deleteItems();
    }
  });

  return {
    filters,
    setFilters,
    setFilter,
    setSorting,
    setStartIndex,
    selection,
    setSelection,
    onFiltersChange,
    queryHash,
    setPageSize,
    debouncedFilters,
  };
}

export type Udf = ReturnType<typeof useDatatableFiltering>;
