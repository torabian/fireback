import {
  Grid,
  Table,
  TableColumnResizing,
  TableEditColumn,
  TableEditRow,
  TableFilterRow,
  TableHeaderRow,
  TableRowDetail,
  TableSelection,
} from "@devexpress/dx-react-grid-bootstrap4";

import {
  ChangeSet,
  Column,
  EditingState,
  FilteringState,
  IntegratedSelection,
  Row,
  RowDetailState,
  SelectionState,
  SortingState,
  TableColumnWidthInfo,
} from "@devexpress/dx-react-grid";
import { UseQueryResult, useQueryClient } from "react-query";
import { LineLoader } from "../line-loader/LineLoader";

import { DatatableColumn } from "@/definitions/definitions";
import { httpErrorHanlder } from "@/helpers/api";
import { useT } from "@/hooks/useT";
import classNames from "classnames";
import React, { useRef, useState } from "react";
import { QueryErrorView } from "../error-view/QueryError";
import { WithPermissions } from "../layouts/WithPermissions";
import { CustomPageSize } from "./CustomPageSize";
import { CustomPagination } from "./CustomPagination";

const TableActivityIndicator = ({
  query,
}: {
  query?: UseQueryResult<any, unknown>;
}) => {
  if (query?.isLoading) {
    return <LineLoader className="table-activity-indicator" />;
  }

  return null;
};

export function CommonDataTable({
  rows,
  columns,
  withFilters,
  booleanColumns,
  query,
  columnSizes,
  defaultColumnWidths,
  onColumnWidthsChange,
  udf,
  RowDetail,
  permissions,
  withPagination,
  children,
  bulkEditHook,
  inlineInsertHook,
  tableClass,
}: {
  rows: Row[];
  columns: Column[] | DatatableColumn[];
  booleanColumns?: string[];
  permissions?: string[];
  withFilters?: boolean;
  withPagination?: boolean;
  RowDetail?: any;
  tableClass?: string;
  query?: UseQueryResult<any, any>;
  defaultColumnWidths: TableColumnWidthInfo[];
  udf: any;
  children?: React.ReactNode;
  columnSizes?: TableColumnWidthInfo[];
  onColumnWidthsChange?:
    | ((nextColumnWidths: TableColumnWidthInfo[]) => void)
    | undefined;
  bulkEditHook?: any;
  inlineInsertHook?: any;
}) {
  const t = useT();

  const {
    filters,
    setSorting,
    setStartIndex,
    selection,
    setSelection,
    setPageSize,
    onFiltersChange,
  } = udf;

  const sel = selection.map((uniqueId: string) => {
    return rows.findIndex((d) => d.uniqueId === uniqueId);
  });

  const totalCount = query?.data?.data?.totalItems || rows.length || 0;

  let currentPage = 0;

  if (totalCount > 0) {
    currentPage = Math.ceil((filters?.startIndex || 0) / filters.itemsPerPage);
  }

  const commitChanges = ({ added, changed, deleted }: ChangeSet) => {};

  const gridRows = rows.map((row, index) => {
    return {
      ...row,
      // id: row.id || row.uniqueId,
    };
  });

  const rowsChanges = useRef<{
    [key: string]: any;
  }>();

  const rowsAdded = useRef<{
    [key: string]: any;
  }>();

  const queryClient = useQueryClient();
  // const { submit: patchAll, mutation: patchAllMutation } = bulkEditHook({
  //   queryClient,
  // });

  const bulkPatchHook = bulkEditHook
    ? bulkEditHook({
        queryClient,
      })
    : undefined;

  const inlineInsertActions = inlineInsertHook
    ? inlineInsertHook({
        queryClient,
      })
    : undefined;

  const createItems = (onExecute: () => void) => {
    const v = rowsAdded.current;
    if (!v) {
      onExecute();
      return;
    }

    for (let item of Object.values(v)) {
      Promise.resolve(inlineInsertActions.submit(item));
    }
  };

  const patchItems = (onExecute: () => void) => {
    const v = rowsChanges.current;
    if (!v) {
      onExecute();
      return;
    }

    let items = Object.keys(v).map((index) => {
      return {
        ...(v[index] ? v[index] : {}),
        uniqueId: gridRows[index as any].uniqueId,
      };
    });

    bulkPatchHook
      ?.submit({ records: items })
      .then((res: any) => {
        onExecute();
      })
      .catch((e: any) => httpErrorHanlder(e, t));
  };

  const onSubmitChanges = async (onExecute: () => void) => {
    patchItems(onExecute);
    // return;

    // Creation is broken

    // createItems(onExecute);
  };

  return (
    <WithPermissions permissions={permissions}>
      <div className={tableClass || "table-container"}>
        <TableActivityIndicator query={query} />
        <Grid rows={gridRows} columns={columns as any}>
          <SortingState
            sorting={filters.sorting}
            onSortingChange={setSorting}
          />
          <EditingState
            onAddedRowsChange={(change) => {
              rowsAdded.current = change;
            }}
            onRowChangesChange={(change) => {
              rowsChanges.current = change;
            }}
            onCommitChanges={commitChanges}
          />
          <SelectionState
            selection={sel}
            onSelectionChange={(selIndex) =>
              setSelection(
                gridRows
                  .filter((x, i) => selIndex.includes(i))
                  .map((d) => d.uniqueId)
              )
            }
          />
          {children}

          <RowDetailState />
          <Table
            noDataCellComponent={() => {
              if (!query) {
                return null;
              }
              return (
                <th colSpan={10}>
                  {query?.isLoading || query?.isError ? null : (
                    <span className="datatable-no-data">
                      {t.table.noRecords}
                    </span>
                  )}
                  <QueryErrorView query={query} />
                </th>
              );
            }}
          />

          {onColumnWidthsChange && (
            <TableColumnResizing
              columnWidths={
                columnSizes
                // columns.map((m) => ({ columnName: m.name, width: 100 })) as any
              }
              onColumnWidthsChange={onColumnWidthsChange}
            />
          )}
          <TableHeaderRow showSortingControls />
          <TableEditRow
            cellComponent={(props) => {
              const currentFieldErr =
                bulkPatchHook?.mutation?.error?.error?.errors?.find(
                  (f: any) => f.location === props.column.name
                );
              return (
                <td colSpan={props.colSpan}>
                  {(props.column as any).inline ? (
                    <input
                      value={props.value}
                      onChange={(e) => props.onValueChange(e.target.value)}
                      type="text"
                      className={classNames("form-control", {
                        "is-invalid": !!currentFieldErr,
                      })}
                    />
                  ) : (
                    <span>{props.value}</span>
                  )}
                </td>
              );
            }}
          />
          {(!!bulkEditHook || !!inlineInsertHook) && (
            <TableEditColumn
              showEditCommand={!!bulkEditHook}
              showAddCommand={!!inlineInsertHook}
              commandComponent={(props, mm) => {
                if (props.id === "commit") {
                  return (
                    <button
                      className="table-row-action"
                      onClick={() => {
                        onSubmitChanges(props.onExecute);
                      }}
                    >
                      {props.text}
                    </button>
                  );
                }

                if (props.id === "cancel") {
                  return (
                    <button
                      className="table-row-action"
                      onClick={() => {
                        bulkPatchHook?.mutation?.reset();
                        props.onExecute();
                      }}
                    >
                      {props.text}
                    </button>
                  );
                }

                return (
                  <button
                    className="table-row-action"
                    onClick={props.onExecute}
                  >
                    {props.text}
                  </button>
                );
              }}
            />
          )}
          <IntegratedSelection />
          <TableSelection highlightRow selectByRowClick showSelectAll />
          {RowDetail && (
            <TableRowDetail
              cellComponent={(data) => {
                return <>{data.children}</>;
              }}
              rowComponent={(data) => {
                return <>{data.children}</>;
              }}
              contentComponent={RowDetail}
            />
          )}

          {withFilters !== false && (
            <FilteringState
              filters={filters.rawFilters || []}
              onFiltersChange={onFiltersChange}
            />
          )}
          {withFilters !== false && (
            <TableFilterRow messages={t.table.filter} />
          )}
        </Grid>
        {withPagination !== false && (
          <div className="table-footer-actions">
            <CustomPageSize
              currentPageSize={filters.itemsPerPage || 10}
              onPageSizeChange={setPageSize}
              pageSizes={[2, 10, 15, 20, 40, 50]}
            />
            <span style={{ padding: "0 10px" }}></span>

            <CustomPagination
              totalCount={totalCount}
              currentPage={currentPage}
              onCurrentPageChange={(page) => {
                setStartIndex(page * filters.itemsPerPage);
              }}
              pageSize={filters.itemsPerPage || 10}
              onPageSizeChange={setPageSize}
            />
          </div>
        )}
      </div>
    </WithPermissions>
  );
}
