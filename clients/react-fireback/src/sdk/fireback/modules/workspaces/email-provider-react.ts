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
import { EmailProviderActions } from "./email-provider-actions";
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
    ? EmailProviderActions.fnExec(execFn(options))
    : EmailProviderActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const emailProvidersQuery = useQuery(
    ["*[]workspaces.EmailProviderEntity", options],
    () => Q().getEmailProviders(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const emailProvidersExportQuery = useQuery(
    ["*[]workspaces.EmailProviderEntity", options],
    () => Q().getEmailProvidersExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const emailProviderByUniqueIdQuery = useQuery(
    ["*workspaces.EmailProviderEntity", options],
    (uniqueId: string) => Q().getEmailProviderByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post emailProvider

  const mutationPostEmailProvider = useMutation<
    IResponse<workspaces.EmailProviderEntity>,
    IResponse<workspaces.EmailProviderEntity>,
    workspaces.EmailProviderEntity
  >((entity) => {
    return Q().postEmailProvider(entity);
  });

  // Only entities are having a store in front-end

  const fnPostEmailProviderUpdater = (
    data: IResponseList<workspaces.EmailProviderEntity> | undefined,
    item: IResponse<workspaces.EmailProviderEntity>
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
          EmailProviderActions.isEmailProviderEntityEqual(t, item.data)
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

  const submitPostEmailProvider = (
    values: workspaces.EmailProviderEntity,
    formikProps?: FormikHelpers<workspaces.EmailProviderEntity>
  ): Promise<IResponse<workspaces.EmailProviderEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostEmailProvider.mutate(values, {
        onSuccess(response: IResponse<workspaces.EmailProviderEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailProviderEntity>
          >("*[]workspaces.EmailProviderEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailProviderEntity) => {
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

  // patch emailProvider

  const mutationPatchEmailProvider = useMutation<
    IResponse<workspaces.EmailProviderEntity>,
    IResponse<workspaces.EmailProviderEntity>,
    workspaces.EmailProviderEntity
  >((entity) => {
    return Q().patchEmailProvider(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchEmailProviderUpdater = (
    data: IResponseList<workspaces.EmailProviderEntity> | undefined,
    item: IResponse<workspaces.EmailProviderEntity>
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
          EmailProviderActions.isEmailProviderEntityEqual(t, item.data)
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

  const submitPatchEmailProvider = (
    values: workspaces.EmailProviderEntity,
    formikProps?: FormikHelpers<workspaces.EmailProviderEntity>
  ): Promise<IResponse<workspaces.EmailProviderEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchEmailProvider.mutate(values, {
        onSuccess(response: IResponse<workspaces.EmailProviderEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailProviderEntity>
          >("*[]workspaces.EmailProviderEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailProviderEntity) => {
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

  // patch emailProviders

  const mutationPatchEmailProviders = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.EmailProviderEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.EmailProviderEntity>>,
    core.BulkRecordRequest<workspaces.EmailProviderEntity>
  >((entity) => {
    return Q().patchEmailProviders(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchEmailProviders = (
    values: core.BulkRecordRequest<workspaces.EmailProviderEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.EmailProviderEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.EmailProviderEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchEmailProviders.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.EmailProviderEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.EmailProviderEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.EmailProviderEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.EmailProviderEntity>
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
  const mutationDeleteEmailProvider = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteEmailProvider();
  });

  const fnDeleteEmailProviderUpdater = (
    data: IResponseList<workspaces.EmailProviderEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = EmailProviderActions.getEmailProviderEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteEmailProvider = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.EmailProviderEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteEmailProvider.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.EmailProviderEntity>
          >("*[]workspaces.EmailProviderEntity", (data) =>
            fnDeleteEmailProviderUpdater(data, values)
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
    emailProvidersQuery,
    emailProvidersExportQuery,
    emailProviderByUniqueIdQuery,
    mutationPostEmailProvider,
    submitPostEmailProvider,
    mutationPatchEmailProvider,
    submitPatchEmailProvider,
    mutationPatchEmailProviders,
    submitPatchEmailProviders,
    mutationDeleteEmailProvider,
    submitDeleteEmailProvider,
  };
}
