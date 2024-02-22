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
import { GsmProviderActions } from "./gsm-provider-actions";
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
    ? GsmProviderActions.fnExec(execFn(options))
    : GsmProviderActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const gsmProvidersQuery = useQuery(
    ["*[]workspaces.GsmProviderEntity", options],
    () => Q().getGsmProviders(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const gsmProvidersExportQuery = useQuery(
    ["*[]workspaces.GsmProviderEntity", options],
    () => Q().getGsmProvidersExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const gsmProviderByUniqueIdQuery = useQuery(
    ["*workspaces.GsmProviderEntity", options],
    (uniqueId: string) => Q().getGsmProviderByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post gsmProvider

  const mutationPostGsmProvider = useMutation<
    IResponse<workspaces.GsmProviderEntity>,
    IResponse<workspaces.GsmProviderEntity>,
    workspaces.GsmProviderEntity
  >((entity) => {
    return Q().postGsmProvider(entity);
  });

  // Only entities are having a store in front-end

  const fnPostGsmProviderUpdater = (
    data: IResponseList<workspaces.GsmProviderEntity> | undefined,
    item: IResponse<workspaces.GsmProviderEntity>
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
          GsmProviderActions.isGsmProviderEntityEqual(t, item.data)
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

  const submitPostGsmProvider = (
    values: workspaces.GsmProviderEntity,
    formikProps?: FormikHelpers<workspaces.GsmProviderEntity>
  ): Promise<IResponse<workspaces.GsmProviderEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostGsmProvider.mutate(values, {
        onSuccess(response: IResponse<workspaces.GsmProviderEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.GsmProviderEntity>
          >("*[]workspaces.GsmProviderEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.GsmProviderEntity) => {
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

  // patch gsmProvider

  const mutationPatchGsmProvider = useMutation<
    IResponse<workspaces.GsmProviderEntity>,
    IResponse<workspaces.GsmProviderEntity>,
    workspaces.GsmProviderEntity
  >((entity) => {
    return Q().patchGsmProvider(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchGsmProviderUpdater = (
    data: IResponseList<workspaces.GsmProviderEntity> | undefined,
    item: IResponse<workspaces.GsmProviderEntity>
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
          GsmProviderActions.isGsmProviderEntityEqual(t, item.data)
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

  const submitPatchGsmProvider = (
    values: workspaces.GsmProviderEntity,
    formikProps?: FormikHelpers<workspaces.GsmProviderEntity>
  ): Promise<IResponse<workspaces.GsmProviderEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchGsmProvider.mutate(values, {
        onSuccess(response: IResponse<workspaces.GsmProviderEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.GsmProviderEntity>
          >("*[]workspaces.GsmProviderEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.GsmProviderEntity) => {
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

  // patch gsmProviders

  const mutationPatchGsmProviders = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.GsmProviderEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.GsmProviderEntity>>,
    core.BulkRecordRequest<workspaces.GsmProviderEntity>
  >((entity) => {
    return Q().patchGsmProviders(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchGsmProviders = (
    values: core.BulkRecordRequest<workspaces.GsmProviderEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.GsmProviderEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.GsmProviderEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchGsmProviders.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.GsmProviderEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.GsmProviderEntity>>
          >(
            "*[]core.BulkRecordRequest[workspaces.GsmProviderEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.GsmProviderEntity>
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
  const mutationDeleteGsmProvider = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteGsmProvider();
  });

  const fnDeleteGsmProviderUpdater = (
    data: IResponseList<workspaces.GsmProviderEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = GsmProviderActions.getGsmProviderEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteGsmProvider = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.GsmProviderEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteGsmProvider.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.GsmProviderEntity>>(
            "*[]workspaces.GsmProviderEntity",
            (data) => fnDeleteGsmProviderUpdater(data, values)
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
    gsmProvidersQuery,
    gsmProvidersExportQuery,
    gsmProviderByUniqueIdQuery,
    mutationPostGsmProvider,
    submitPostGsmProvider,
    mutationPatchGsmProvider,
    submitPatchGsmProvider,
    mutationPatchGsmProviders,
    submitPatchGsmProviders,
    mutationDeleteGsmProvider,
    submitDeleteGsmProvider,
  };
}
