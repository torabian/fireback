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
import { BackupTableMetaActions } from "./backup-table-meta-actions";
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
    ? BackupTableMetaActions.fnExec(execFn(options))
    : BackupTableMetaActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const backupTableMetasQuery = useQuery(
    ["*[]workspaces.BackupTableMetaEntity", options],
    () => Q().getBackupTableMetas(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const backupTableMetasExportQuery = useQuery(
    ["*[]workspaces.BackupTableMetaEntity", options],
    () => Q().getBackupTableMetasExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const backupTableMetaByUniqueIdQuery = useQuery(
    ["*workspaces.BackupTableMetaEntity", options],
    (uniqueId: string) => Q().getBackupTableMetaByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post backupTableMeta

  const mutationPostBackupTableMeta = useMutation<
    IResponse<workspaces.BackupTableMetaEntity>,
    IResponse<workspaces.BackupTableMetaEntity>,
    workspaces.BackupTableMetaEntity
  >((entity) => {
    return Q().postBackupTableMeta(entity);
  });

  // Only entities are having a store in front-end

  const fnPostBackupTableMetaUpdater = (
    data: IResponseList<workspaces.BackupTableMetaEntity> | undefined,
    item: IResponse<workspaces.BackupTableMetaEntity>
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
          BackupTableMetaActions.isBackupTableMetaEntityEqual(t, item.data)
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

  const submitPostBackupTableMeta = (
    values: workspaces.BackupTableMetaEntity,
    formikProps?: FormikHelpers<workspaces.BackupTableMetaEntity>
  ): Promise<IResponse<workspaces.BackupTableMetaEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostBackupTableMeta.mutate(values, {
        onSuccess(response: IResponse<workspaces.BackupTableMetaEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.BackupTableMetaEntity>
          >("*[]workspaces.BackupTableMetaEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.BackupTableMetaEntity) => {
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

  // patch backupTableMeta

  const mutationPatchBackupTableMeta = useMutation<
    IResponse<workspaces.BackupTableMetaEntity>,
    IResponse<workspaces.BackupTableMetaEntity>,
    workspaces.BackupTableMetaEntity
  >((entity) => {
    return Q().patchBackupTableMeta(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchBackupTableMetaUpdater = (
    data: IResponseList<workspaces.BackupTableMetaEntity> | undefined,
    item: IResponse<workspaces.BackupTableMetaEntity>
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
          BackupTableMetaActions.isBackupTableMetaEntityEqual(t, item.data)
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

  const submitPatchBackupTableMeta = (
    values: workspaces.BackupTableMetaEntity,
    formikProps?: FormikHelpers<workspaces.BackupTableMetaEntity>
  ): Promise<IResponse<workspaces.BackupTableMetaEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchBackupTableMeta.mutate(values, {
        onSuccess(response: IResponse<workspaces.BackupTableMetaEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.BackupTableMetaEntity>
          >("*[]workspaces.BackupTableMetaEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.BackupTableMetaEntity) => {
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

  // patch backupTableMetas

  const mutationPatchBackupTableMetas = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>,
    core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
  >((entity) => {
    return Q().patchBackupTableMetas(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchBackupTableMetas = (
    values: core.BulkRecordRequest<workspaces.BackupTableMetaEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchBackupTableMetas.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.BackupTableMetaEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
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
  const mutationDeleteBackupTableMeta = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteBackupTableMeta();
  });

  const fnDeleteBackupTableMetaUpdater = (
    data: IResponseList<workspaces.BackupTableMetaEntity> | undefined,
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
          BackupTableMetaActions.getBackupTableMetaEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteBackupTableMeta = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.BackupTableMetaEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteBackupTableMeta.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.BackupTableMetaEntity>
          >("*[]workspaces.BackupTableMetaEntity", (data) =>
            fnDeleteBackupTableMetaUpdater(data, values)
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
    backupTableMetasQuery,
    backupTableMetasExportQuery,
    backupTableMetaByUniqueIdQuery,
    mutationPostBackupTableMeta,
    submitPostBackupTableMeta,
    mutationPatchBackupTableMeta,
    submitPatchBackupTableMeta,
    mutationPatchBackupTableMetas,
    submitPatchBackupTableMetas,
    mutationDeleteBackupTableMeta,
    submitDeleteBackupTableMeta,
  };
}
