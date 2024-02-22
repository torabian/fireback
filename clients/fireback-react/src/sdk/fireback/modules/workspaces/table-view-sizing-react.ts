// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: workspaces
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { TableViewSizingActions } from "./table-view-sizing-actions";
import * as workspaces from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  core,
  IResponse,
  ExecApi,
  mutationErrorsToFormik,
  IResponseList,
} from "../../core/http-tools";
/**
 * Gives you formik forms, all mutations, submit actions, and error handling,
 * and provides internal store for all changes happens through this
 * for modules
 */
export function useWorkspaces(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? TableViewSizingActions.fnExec(execFn(options))
    : TableViewSizingActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const tableViewSizingsQuery = useQuery(
    ["*[]workspaces.TableViewSizingEntity", options],
    () => Q().getTableViewSizings(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const tableViewSizingsExportQuery = useQuery(
    ["*[]workspaces.TableViewSizingEntity", options],
    () => Q().getTableViewSizingsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const tableViewSizingByUniqueIdQuery = useQuery(
    ["*workspaces.TableViewSizingEntity", options],
    (uniqueId: string) => Q().getTableViewSizingByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post tableViewSizing

  const mutationPostTableViewSizing = useMutation<
    IResponse<workspaces.TableViewSizingEntity>,
    IResponse<workspaces.TableViewSizingEntity>,
    workspaces.TableViewSizingEntity
  >((entity) => {
    return Q().postTableViewSizing(entity);
  });

  // Only entities are having a store in front-end

  const fnPostTableViewSizingUpdater = (
    data: IResponseList<workspaces.TableViewSizingEntity> | undefined,
    item: IResponse<workspaces.TableViewSizingEntity>
  ) => {
    return [];

    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items && item.data) {
      data.data.items = data.data.items.map((t) => {
        if (
          item.data !== undefined &&
          TableViewSizingActions.isTableViewSizingEntityEqual(t, item.data)
        ) {
          return item.data;
        }

        return t;
      });
    } else if (data?.data && item.data) {
      data.data.items = [item.data, ...(data?.data?.items || [])];
    }

    return data;
  };

  const submitPostTableViewSizing = (
    values: workspaces.TableViewSizingEntity,
    formikProps?: FormikHelpers<workspaces.TableViewSizingEntity>
  ): Promise<IResponse<workspaces.TableViewSizingEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostTableViewSizing.mutate(values, {
        onSuccess(response: IResponse<workspaces.TableViewSizingEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.TableViewSizingEntity>
          >("*[]workspaces.TableViewSizingEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.TableViewSizingEntity) => {
                if (item.uniqueId === response.data?.uniqueId) {
                  return response.data;
                }

                return item;
              }
            );

            return data;
          });

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  // patch tableViewSizing

  const mutationPatchTableViewSizing = useMutation<
    IResponse<workspaces.TableViewSizingEntity>,
    IResponse<workspaces.TableViewSizingEntity>,
    workspaces.TableViewSizingEntity
  >((entity) => {
    return Q().patchTableViewSizing(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchTableViewSizingUpdater = (
    data: IResponseList<workspaces.TableViewSizingEntity> | undefined,
    item: IResponse<workspaces.TableViewSizingEntity>
  ) => {
    return [];

    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items && item.data) {
      data.data.items = data.data.items.map((t) => {
        if (
          item.data !== undefined &&
          TableViewSizingActions.isTableViewSizingEntityEqual(t, item.data)
        ) {
          return item.data;
        }

        return t;
      });
    } else if (data?.data && item.data) {
      data.data.items = [item.data, ...(data?.data?.items || [])];
    }

    return data;
  };

  const submitPatchTableViewSizing = (
    values: workspaces.TableViewSizingEntity,
    formikProps?: FormikHelpers<workspaces.TableViewSizingEntity>
  ): Promise<IResponse<workspaces.TableViewSizingEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchTableViewSizing.mutate(values, {
        onSuccess(response: IResponse<workspaces.TableViewSizingEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.TableViewSizingEntity>
          >("*[]workspaces.TableViewSizingEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.TableViewSizingEntity) => {
                if (item.uniqueId === response.data?.uniqueId) {
                  return response.data;
                }

                return item;
              }
            );

            return data;
          });

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  // patch tableViewSizings

  const mutationPatchTableViewSizings = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.TableViewSizingEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.TableViewSizingEntity>>,
    core.BulkRecordRequest<workspaces.TableViewSizingEntity>
  >((entity) => {
    return Q().patchTableViewSizings(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchTableViewSizings = (
    values: core.BulkRecordRequest<workspaces.TableViewSizingEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.TableViewSizingEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.TableViewSizingEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchTableViewSizings.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.TableViewSizingEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.TableViewSizingEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.TableViewSizingEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.TableViewSizingEntity>
                ) => {
                  if (item.uniqueId === response.data?.uniqueId) {
                    return response.data;
                  }

                  return item;
                }
              );

              return data;
            }
          );

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  // Deleting an entity
  const mutationDeleteTableViewSizing = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteTableViewSizing();
  });

  const fnDeleteTableViewSizingUpdater = (
    data: IResponseList<workspaces.TableViewSizingEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key =
          TableViewSizingActions.getTableViewSizingEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteTableViewSizing = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.TableViewSizingEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteTableViewSizing.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.TableViewSizingEntity>
          >("*[]workspaces.TableViewSizingEntity", (data) =>
            fnDeleteTableViewSizingUpdater(data, values)
          );

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  return {
    queryClient,
    tableViewSizingsQuery,
    tableViewSizingsExportQuery,
    tableViewSizingByUniqueIdQuery,
    mutationPostTableViewSizing,
    submitPostTableViewSizing,
    mutationPatchTableViewSizing,
    submitPatchTableViewSizing,
    mutationPatchTableViewSizings,
    submitPatchTableViewSizings,
    mutationDeleteTableViewSizing,
    submitDeleteTableViewSizing,
  };
}
