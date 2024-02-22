// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: licenses
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { ActivationKeyActions } from "./activation-key-actions";
import * as licenses from "./index";
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
export function useLicenses(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? ActivationKeyActions.fnExec(execFn(options))
    : ActivationKeyActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const activationKeysQuery = useQuery(
    ["*[]licenses.ActivationKeyEntity", options],
    () => Q().getActivationKeys(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const activationKeysExportQuery = useQuery(
    ["*[]licenses.ActivationKeyEntity", options],
    () => Q().getActivationKeysExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const activationKeyByUniqueIdQuery = useQuery(
    ["*licenses.ActivationKeyEntity", options],
    (uniqueId: string) => Q().getActivationKeyByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post activationKey

  const mutationPostActivationKey = useMutation<
    IResponse<licenses.ActivationKeyEntity>,
    IResponse<licenses.ActivationKeyEntity>,
    licenses.ActivationKeyEntity
  >((entity) => {
    return Q().postActivationKey(entity);
  });

  // Only entities are having a store in front-end

  const fnPostActivationKeyUpdater = (
    data: IResponseList<licenses.ActivationKeyEntity> | undefined,
    item: IResponse<licenses.ActivationKeyEntity>
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
          ActivationKeyActions.isActivationKeyEntityEqual(t, item.data)
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

  const submitPostActivationKey = (
    values: licenses.ActivationKeyEntity,
    formikProps?: FormikHelpers<licenses.ActivationKeyEntity>
  ): Promise<IResponse<licenses.ActivationKeyEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostActivationKey.mutate(values, {
        onSuccess(response: IResponse<licenses.ActivationKeyEntity>) {
          queryClient.setQueriesData<
            IResponseList<licenses.ActivationKeyEntity>
          >("*[]licenses.ActivationKeyEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: licenses.ActivationKeyEntity) => {
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

  // patch activationKey

  const mutationPatchActivationKey = useMutation<
    IResponse<licenses.ActivationKeyEntity>,
    IResponse<licenses.ActivationKeyEntity>,
    licenses.ActivationKeyEntity
  >((entity) => {
    return Q().patchActivationKey(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchActivationKeyUpdater = (
    data: IResponseList<licenses.ActivationKeyEntity> | undefined,
    item: IResponse<licenses.ActivationKeyEntity>
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
          ActivationKeyActions.isActivationKeyEntityEqual(t, item.data)
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

  const submitPatchActivationKey = (
    values: licenses.ActivationKeyEntity,
    formikProps?: FormikHelpers<licenses.ActivationKeyEntity>
  ): Promise<IResponse<licenses.ActivationKeyEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchActivationKey.mutate(values, {
        onSuccess(response: IResponse<licenses.ActivationKeyEntity>) {
          queryClient.setQueriesData<
            IResponseList<licenses.ActivationKeyEntity>
          >("*[]licenses.ActivationKeyEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: licenses.ActivationKeyEntity) => {
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

  // patch activationKeys

  const mutationPatchActivationKeys = useMutation<
    IResponse<core.BulkRecordRequest<licenses.ActivationKeyEntity>>,
    IResponse<core.BulkRecordRequest<licenses.ActivationKeyEntity>>,
    core.BulkRecordRequest<licenses.ActivationKeyEntity>
  >((entity) => {
    return Q().patchActivationKeys(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchActivationKeys = (
    values: core.BulkRecordRequest<licenses.ActivationKeyEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<licenses.ActivationKeyEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<licenses.ActivationKeyEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchActivationKeys.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<licenses.ActivationKeyEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<licenses.ActivationKeyEntity>>
          >(
            "*[]core.BulkRecordRequest[licenses.ActivationKeyEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<licenses.ActivationKeyEntity>
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
  const mutationDeleteActivationKey = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteActivationKey();
  });

  const fnDeleteActivationKeyUpdater = (
    data: IResponseList<licenses.ActivationKeyEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = ActivationKeyActions.getActivationKeyEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteActivationKey = (
    values: string[],
    formikProps?: FormikHelpers<licenses.ActivationKeyEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteActivationKey.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<licenses.ActivationKeyEntity>>(
            "*[]licenses.ActivationKeyEntity",
            (data) => fnDeleteActivationKeyUpdater(data, values)
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
    activationKeysQuery,
    activationKeysExportQuery,
    activationKeyByUniqueIdQuery,
    mutationPostActivationKey,
    submitPostActivationKey,
    mutationPatchActivationKey,
    submitPatchActivationKey,
    mutationPatchActivationKeys,
    submitPatchActivationKeys,
    mutationDeleteActivationKey,
    submitDeleteActivationKey,
  };
}
