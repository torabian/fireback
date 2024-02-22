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
import { PublicJoinKeyActions } from "./public-join-key-actions";
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
    ? PublicJoinKeyActions.fnExec(execFn(options))
    : PublicJoinKeyActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const publicJoinKeysQuery = useQuery(
    ["*[]workspaces.PublicJoinKeyEntity", options],
    () => Q().getPublicJoinKeys(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const publicJoinKeysExportQuery = useQuery(
    ["*[]workspaces.PublicJoinKeyEntity", options],
    () => Q().getPublicJoinKeysExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const publicJoinKeyByUniqueIdQuery = useQuery(
    ["*workspaces.PublicJoinKeyEntity", options],
    (uniqueId: string) => Q().getPublicJoinKeyByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post publicJoinKey

  const mutationPostPublicJoinKey = useMutation<
    IResponse<workspaces.PublicJoinKeyEntity>,
    IResponse<workspaces.PublicJoinKeyEntity>,
    workspaces.PublicJoinKeyEntity
  >((entity) => {
    return Q().postPublicJoinKey(entity);
  });

  // Only entities are having a store in front-end

  const fnPostPublicJoinKeyUpdater = (
    data: IResponseList<workspaces.PublicJoinKeyEntity> | undefined,
    item: IResponse<workspaces.PublicJoinKeyEntity>
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
          PublicJoinKeyActions.isPublicJoinKeyEntityEqual(t, item.data)
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

  const submitPostPublicJoinKey = (
    values: workspaces.PublicJoinKeyEntity,
    formikProps?: FormikHelpers<workspaces.PublicJoinKeyEntity>
  ): Promise<IResponse<workspaces.PublicJoinKeyEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostPublicJoinKey.mutate(values, {
        onSuccess(response: IResponse<workspaces.PublicJoinKeyEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.PublicJoinKeyEntity>
          >("*[]workspaces.PublicJoinKeyEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.PublicJoinKeyEntity) => {
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

  // patch publicJoinKey

  const mutationPatchPublicJoinKey = useMutation<
    IResponse<workspaces.PublicJoinKeyEntity>,
    IResponse<workspaces.PublicJoinKeyEntity>,
    workspaces.PublicJoinKeyEntity
  >((entity) => {
    return Q().patchPublicJoinKey(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchPublicJoinKeyUpdater = (
    data: IResponseList<workspaces.PublicJoinKeyEntity> | undefined,
    item: IResponse<workspaces.PublicJoinKeyEntity>
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
          PublicJoinKeyActions.isPublicJoinKeyEntityEqual(t, item.data)
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

  const submitPatchPublicJoinKey = (
    values: workspaces.PublicJoinKeyEntity,
    formikProps?: FormikHelpers<workspaces.PublicJoinKeyEntity>
  ): Promise<IResponse<workspaces.PublicJoinKeyEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchPublicJoinKey.mutate(values, {
        onSuccess(response: IResponse<workspaces.PublicJoinKeyEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.PublicJoinKeyEntity>
          >("*[]workspaces.PublicJoinKeyEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.PublicJoinKeyEntity) => {
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

  // patch publicJoinKeys

  const mutationPatchPublicJoinKeys = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>>,
    core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>
  >((entity) => {
    return Q().patchPublicJoinKeys(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchPublicJoinKeys = (
    values: core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchPublicJoinKeys.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.PublicJoinKeyEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>
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
  const mutationDeletePublicJoinKey = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deletePublicJoinKey();
  });

  const fnDeletePublicJoinKeyUpdater = (
    data: IResponseList<workspaces.PublicJoinKeyEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = PublicJoinKeyActions.getPublicJoinKeyEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeletePublicJoinKey = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.PublicJoinKeyEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeletePublicJoinKey.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.PublicJoinKeyEntity>
          >("*[]workspaces.PublicJoinKeyEntity", (data) =>
            fnDeletePublicJoinKeyUpdater(data, values)
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
    publicJoinKeysQuery,
    publicJoinKeysExportQuery,
    publicJoinKeyByUniqueIdQuery,
    mutationPostPublicJoinKey,
    submitPostPublicJoinKey,
    mutationPatchPublicJoinKey,
    submitPatchPublicJoinKey,
    mutationPatchPublicJoinKeys,
    submitPatchPublicJoinKeys,
    mutationDeletePublicJoinKey,
    submitDeletePublicJoinKey,
  };
}
