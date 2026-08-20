import { type Filter, type Sorting } from "@devexpress/dx-react-grid";
import { parse, stringify } from "qs";
import { useEffect, useRef, useState } from "react";
import { useLocation } from "react-router-dom";
import {
  type IMenuActionItem,
  useMenuTools,
} from "../components/action-menu/ActionMenu";
import { commonDialogs } from "../components/overlay/CommonOverlays";
import { strings } from "../components/strings/translations";
import { httpErrorHanlder } from "./api";
import { type Filters } from "./datatabletools";
import { osResources } from "./resources";
import { useDebouncedEffect } from "./useDebouncedEffect";
import { KeyboardAction } from "./useExportTools";
import { useKeyCombination } from "./useKeyPress";
import { useRouter } from "./useRouter";
import { useS } from "./useS";

export function useDatatableFiltering({
  urlMask,
  submitDelete,
  onRecordsDeleted,
  initialFilters,
}: {
  urlMask: string;
  onRecordsDeleted?: (items: string[]) => void;
  submitDelete?: any;
  initialFilters?: Partial<Filters>;
}) {
  const s = useS(strings);
  const router = useRouter();

  const { confirmModal } = commonDialogs();
  const { withDebounce } = useDebouncedEffect();
  const init = {
    itemsPerPage: 100,
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
    } catch (error) {}

    setFilters({ ...init, ...filters });
    setDebouncedFilters({ ...init, ...filters });
  }, [search]);

  const [selection, setSelection$] = useState<Array<string>>([]);
  // const [queryHash, setQueryHash] = useState("{}");

  const computeQueryKey = (filters) => {
    const queryHashItems = { ...filters };
    delete queryHashItems.itemsPerPage;
    // cursor advances page-to-page within the same logical query (same as
    // startIndex used to), so it must not be part of the hash either -
    // otherwise every "load next page" would look like a brand new query.
    delete queryHashItems.cursor;
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
      newFilters.cursor = "";
    }

    setFilters(newFilters);

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

  // Advances a cursor-paginated query to its next page: `cursor` should be
  // the `next.cursor` value from the most recent GResponse. `reset: false`
  // so it doesn't clobber startIndex or restart the query.
  const setCursor = (cursor: string) => {
    setFilter({ cursor }, false);
  };

  const onFiltersChange = (filters: Filter[] | undefined) => {
    let newFilters = { cursor: "" };
    setFilter(newFilters);
  };

  const deleteItems = async () => {
    confirmModal({
      title: s.confirm,
      confirmLabel: s.common.yes,
      cancelLabel: s.common.no,
      description: s.deleteConfirmMessage,
    })
      .promise.then(({ type }) => {
        if (type !== "resolved") {
          return;
        }

        return submitDelete(
          JSON.stringify({ uniqueIds: selection }),
          null as any,
        ).then((response: any) => {
          // Bug fix: our fetch layer resolves (doesn't reject) a backend-rejected
          // delete - e.g. WorkspaceAwareDeleteAction refusing to delete the root
          // workspace - so `response` here can itself be an error envelope
          // (`{error: {...}}`) even though this .then() ran. Same
          // response.error?.toJSON?.() ?? response.error unwrapping
          // CommonEntityManager.tsx's onSubmit already uses. Without this check, a
          // rejected delete silently did nothing - no toast, no error, and (before
          // this fix) still called onRecordsDeleted as if it had succeeded.
          const errorInfo = response?.error?.toJSON?.() ?? response?.error;
          if (
            errorInfo?.message ||
            errorInfo?.messageTranslated ||
            errorInfo?.errors?.length
          ) {
            httpErrorHanlder({ error: errorInfo }, s);
            return;
          }

          onRecordsDeleted && onRecordsDeleted(selection);
        });
      })
      // A genuinely rejected promise (network failure, etc.) - distinct from the
      // resolved-but-failed case handled above.
      .catch((err) => httpErrorHanlder(err, s));
  };

  const deleteAction = (): IMenuActionItem => ({
    label: s.deleteAction,
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
    setCursor,
    selection,
    setSelection,
    onFiltersChange,
    queryHash,
    setPageSize,
    debouncedFilters,
  };
}

export type Udf = ReturnType<typeof useDatatableFiltering>;
