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
import { NotificationConfigActions } from "./notification-config-actions";
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
    ? NotificationConfigActions.fnExec(execFn(options))
    : NotificationConfigActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const notificationConfigsQuery = useQuery(
    ["*[]workspaces.NotificationConfigEntity", options],
    () => Q().getNotificationConfigs(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const notificationConfigsExportQuery = useQuery(
    ["*[]workspaces.NotificationConfigEntity", options],
    () => Q().getNotificationConfigsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const notificationConfigByUniqueIdQuery = useQuery(
    ["*workspaces.NotificationConfigEntity", options],
    (uniqueId: string) => Q().getNotificationConfigByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post notificationConfig

  const mutationPostNotificationConfig = useMutation<
    IResponse<workspaces.NotificationConfigEntity>,
    IResponse<workspaces.NotificationConfigEntity>,
    workspaces.NotificationConfigEntity
  >((entity) => {
    return Q().postNotificationConfig(entity);
  });

  // Only entities are having a store in front-end

  const fnPostNotificationConfigUpdater = (
    data: IResponseList<workspaces.NotificationConfigEntity> | undefined,
    item: IResponse<workspaces.NotificationConfigEntity>
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
          NotificationConfigActions.isNotificationConfigEntityEqual(
            t,
            item.data
          )
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

  const submitPostNotificationConfig = (
    values: workspaces.NotificationConfigEntity,
    formikProps?: FormikHelpers<workspaces.NotificationConfigEntity>
  ): Promise<IResponse<workspaces.NotificationConfigEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostNotificationConfig.mutate(values, {
        onSuccess(response: IResponse<workspaces.NotificationConfigEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.NotificationConfigEntity>
          >("*[]workspaces.NotificationConfigEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.NotificationConfigEntity) => {
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

  // patch notificationConfig

  const mutationPatchNotificationConfig = useMutation<
    IResponse<workspaces.NotificationConfigEntity>,
    IResponse<workspaces.NotificationConfigEntity>,
    workspaces.NotificationConfigEntity
  >((entity) => {
    return Q().patchNotificationConfig(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchNotificationConfigUpdater = (
    data: IResponseList<workspaces.NotificationConfigEntity> | undefined,
    item: IResponse<workspaces.NotificationConfigEntity>
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
          NotificationConfigActions.isNotificationConfigEntityEqual(
            t,
            item.data
          )
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

  const submitPatchNotificationConfig = (
    values: workspaces.NotificationConfigEntity,
    formikProps?: FormikHelpers<workspaces.NotificationConfigEntity>
  ): Promise<IResponse<workspaces.NotificationConfigEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchNotificationConfig.mutate(values, {
        onSuccess(response: IResponse<workspaces.NotificationConfigEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.NotificationConfigEntity>
          >("*[]workspaces.NotificationConfigEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.NotificationConfigEntity) => {
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

  // patch notificationConfigs

  const mutationPatchNotificationConfigs = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.NotificationConfigEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.NotificationConfigEntity>>,
    core.BulkRecordRequest<workspaces.NotificationConfigEntity>
  >((entity) => {
    return Q().patchNotificationConfigs(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchNotificationConfigs = (
    values: core.BulkRecordRequest<workspaces.NotificationConfigEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.NotificationConfigEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.NotificationConfigEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchNotificationConfigs.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.NotificationConfigEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.NotificationConfigEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.NotificationConfigEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.NotificationConfigEntity>
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
  const mutationDeleteNotificationConfig = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteNotificationConfig();
  });

  const fnDeleteNotificationConfigUpdater = (
    data: IResponseList<workspaces.NotificationConfigEntity> | undefined,
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
          NotificationConfigActions.getNotificationConfigEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteNotificationConfig = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.NotificationConfigEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteNotificationConfig.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.NotificationConfigEntity>
          >("*[]workspaces.NotificationConfigEntity", (data) =>
            fnDeleteNotificationConfigUpdater(data, values)
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

  // post notification/testmail

  const mutationPostNotificationTestmail = useMutation<
    IResponse<core.OkayResponse>,
    IResponse<core.OkayResponse>,
    workspaces.TestMailDto
  >((entity) => {
    return Q().postNotificationTestmail(entity);
  });

  // Only entities are having a store in front-end

  const submitPostNotificationTestmail = (
    values: workspaces.TestMailDto,
    formikProps?: FormikHelpers<workspaces.TestMailDto>
  ): Promise<IResponse<core.OkayResponse>> => {
    return new Promise((resolve, reject) => {
      mutationPostNotificationTestmail.mutate(values, {
        onSuccess(response: IResponse<core.OkayResponse>) {
          queryClient.setQueriesData<IResponseList<workspaces.TestMailDto>>(
            "*[]workspaces.TestMailDto",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.TestMailDto) => {
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

  const notificationWorkspaceConfigQuery = useQuery(
    ["*workspaces.NotificationConfigEntity", options],
    () => Q().getNotificationWorkspaceConfig(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // patch notification/workspace/config

  const mutationPatchNotificationWorkspaceConfig = useMutation<
    IResponse<workspaces.NotificationConfigEntity>,
    IResponse<workspaces.NotificationConfigEntity>,
    workspaces.NotificationConfigEntity
  >((entity) => {
    return Q().patchNotificationWorkspaceConfig(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchNotificationWorkspaceConfigUpdater = (
    data: IResponseList<workspaces.NotificationConfigEntity> | undefined,
    item: IResponse<workspaces.NotificationConfigEntity>
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
          NotificationConfigActions.isNotificationConfigEntityEqual(
            t,
            item.data
          )
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

  const submitPatchNotificationWorkspaceConfig = (
    values: workspaces.NotificationConfigEntity,
    formikProps?: FormikHelpers<workspaces.NotificationConfigEntity>
  ): Promise<IResponse<workspaces.NotificationConfigEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchNotificationWorkspaceConfig.mutate(values, {
        onSuccess(response: IResponse<workspaces.NotificationConfigEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.NotificationConfigEntity>
          >("*[]workspaces.NotificationConfigEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.NotificationConfigEntity) => {
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

  return {
    queryClient,
    notificationConfigsQuery,
    notificationConfigsExportQuery,
    notificationConfigByUniqueIdQuery,
    mutationPostNotificationConfig,
    submitPostNotificationConfig,
    mutationPatchNotificationConfig,
    submitPatchNotificationConfig,
    mutationPatchNotificationConfigs,
    submitPatchNotificationConfigs,
    mutationDeleteNotificationConfig,
    submitDeleteNotificationConfig,
    mutationPostNotificationTestmail,
    submitPostNotificationTestmail,
    notificationWorkspaceConfigQuery,
    mutationPatchNotificationWorkspaceConfig,
    submitPatchNotificationWorkspaceConfig,
  };
}
