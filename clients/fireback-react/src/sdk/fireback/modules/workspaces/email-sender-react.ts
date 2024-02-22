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
import { EmailSenderActions } from "./email-sender-actions";
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
    ? EmailSenderActions.fnExec(execFn(options))
    : EmailSenderActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const emailSendersQuery = useQuery(
    ["*[]workspaces.EmailSenderEntity", options],
    () => Q().getEmailSenders(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const emailSendersExportQuery = useQuery(
    ["*[]workspaces.EmailSenderEntity", options],
    () => Q().getEmailSendersExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const emailSenderByUniqueIdQuery = useQuery(
    ["*workspaces.EmailSenderEntity", options],
    (uniqueId: string) => Q().getEmailSenderByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post emailSender

  const mutationPostEmailSender = useMutation<
    IResponse<workspaces.EmailSenderEntity>,
    IResponse<workspaces.EmailSenderEntity>,
    workspaces.EmailSenderEntity
  >((entity) => {
    return Q().postEmailSender(entity);
  });

  // Only entities are having a store in front-end

  const fnPostEmailSenderUpdater = (
    data: IResponseList<workspaces.EmailSenderEntity> | undefined,
    item: IResponse<workspaces.EmailSenderEntity>
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
          EmailSenderActions.isEmailSenderEntityEqual(t, item.data)
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

  const submitPostEmailSender = (
    values: workspaces.EmailSenderEntity,
    formikProps?: FormikHelpers<workspaces.EmailSenderEntity>
  ): Promise<IResponse<workspaces.EmailSenderEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostEmailSender.mutate(values, {
        onSuccess(response: IResponse<workspaces.EmailSenderEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailSenderEntity>
          >("*[]workspaces.EmailSenderEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailSenderEntity) => {
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

  // patch emailSender

  const mutationPatchEmailSender = useMutation<
    IResponse<workspaces.EmailSenderEntity>,
    IResponse<workspaces.EmailSenderEntity>,
    workspaces.EmailSenderEntity
  >((entity) => {
    return Q().patchEmailSender(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchEmailSenderUpdater = (
    data: IResponseList<workspaces.EmailSenderEntity> | undefined,
    item: IResponse<workspaces.EmailSenderEntity>
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
          EmailSenderActions.isEmailSenderEntityEqual(t, item.data)
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

  const submitPatchEmailSender = (
    values: workspaces.EmailSenderEntity,
    formikProps?: FormikHelpers<workspaces.EmailSenderEntity>
  ): Promise<IResponse<workspaces.EmailSenderEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchEmailSender.mutate(values, {
        onSuccess(response: IResponse<workspaces.EmailSenderEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailSenderEntity>
          >("*[]workspaces.EmailSenderEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailSenderEntity) => {
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

  // patch emailSenders

  const mutationPatchEmailSenders = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.EmailSenderEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.EmailSenderEntity>>,
    core.BulkRecordRequest<workspaces.EmailSenderEntity>
  >((entity) => {
    return Q().patchEmailSenders(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchEmailSenders = (
    values: core.BulkRecordRequest<workspaces.EmailSenderEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.EmailSenderEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.EmailSenderEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchEmailSenders.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.EmailSenderEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.EmailSenderEntity>>
          >(
            "*[]core.BulkRecordRequest[workspaces.EmailSenderEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.EmailSenderEntity>
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
  const mutationDeleteEmailSender = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteEmailSender();
  });

  const fnDeleteEmailSenderUpdater = (
    data: IResponseList<workspaces.EmailSenderEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = EmailSenderActions.getEmailSenderEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteEmailSender = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.EmailSenderEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteEmailSender.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.EmailSenderEntity>>(
            "*[]workspaces.EmailSenderEntity",
            (data) => fnDeleteEmailSenderUpdater(data, values)
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
    emailSendersQuery,
    emailSendersExportQuery,
    emailSenderByUniqueIdQuery,
    mutationPostEmailSender,
    submitPostEmailSender,
    mutationPatchEmailSender,
    submitPatchEmailSender,
    mutationPatchEmailSenders,
    submitPatchEmailSenders,
    mutationDeleteEmailSender,
    submitDeleteEmailSender,
  };
}
