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
import { UserActions } from "./user-actions";
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
    ? UserActions.fnExec(execFn(options))
    : UserActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const usersQuery = useQuery(
    ["*[]workspaces.UserEntity", options],
    () => Q().getUsers(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const usersExportQuery = useQuery(
    ["*[]workspaces.UserEntity", options],
    () => Q().getUsersExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const userByUniqueIdQuery = useQuery(
    ["*workspaces.UserEntity", options],
    (uniqueId: string) => Q().getUserByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post user

  const mutationPostUser = useMutation<
    IResponse<workspaces.UserEntity>,
    IResponse<workspaces.UserEntity>,
    workspaces.UserEntity
  >((entity) => {
    return Q().postUser(entity);
  });

  // Only entities are having a store in front-end

  const fnPostUserUpdater = (
    data: IResponseList<workspaces.UserEntity> | undefined,
    item: IResponse<workspaces.UserEntity>
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
          UserActions.isUserEntityEqual(t, item.data)
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

  const submitPostUser = (
    values: workspaces.UserEntity,
    formikProps?: FormikHelpers<workspaces.UserEntity>
  ): Promise<IResponse<workspaces.UserEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostUser.mutate(values, {
        onSuccess(response: IResponse<workspaces.UserEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.UserEntity>>(
            "*[]workspaces.UserEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.UserEntity) => {
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

  // patch user

  const mutationPatchUser = useMutation<
    IResponse<workspaces.UserEntity>,
    IResponse<workspaces.UserEntity>,
    workspaces.UserEntity
  >((entity) => {
    return Q().patchUser(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchUserUpdater = (
    data: IResponseList<workspaces.UserEntity> | undefined,
    item: IResponse<workspaces.UserEntity>
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
          UserActions.isUserEntityEqual(t, item.data)
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

  const submitPatchUser = (
    values: workspaces.UserEntity,
    formikProps?: FormikHelpers<workspaces.UserEntity>
  ): Promise<IResponse<workspaces.UserEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchUser.mutate(values, {
        onSuccess(response: IResponse<workspaces.UserEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.UserEntity>>(
            "*[]workspaces.UserEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.UserEntity) => {
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

  // patch users

  const mutationPatchUsers = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.UserEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.UserEntity>>,
    core.BulkRecordRequest<workspaces.UserEntity>
  >((entity) => {
    return Q().patchUsers(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchUsers = (
    values: core.BulkRecordRequest<workspaces.UserEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<workspaces.UserEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<workspaces.UserEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchUsers.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<workspaces.UserEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.UserEntity>>
          >("*[]core.BulkRecordRequest[workspaces.UserEntity]", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: core.BulkRecordRequest<workspaces.UserEntity>) => {
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

  // Deleting an entity
  const mutationDeleteUser = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteUser();
  });

  const fnDeleteUserUpdater = (
    data: IResponseList<workspaces.UserEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = UserActions.getUserEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteUser = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.UserEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteUser.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.UserEntity>>(
            "*[]workspaces.UserEntity",
            (data) => fnDeleteUserUpdater(data, values)
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
    usersQuery,
    usersExportQuery,
    userByUniqueIdQuery,
    mutationPostUser,
    submitPostUser,
    mutationPatchUser,
    submitPatchUser,
    mutationPatchUsers,
    submitPatchUsers,
    mutationDeleteUser,
    submitDeleteUser,
  };
}
