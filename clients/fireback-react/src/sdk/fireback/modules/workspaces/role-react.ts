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
import { RoleActions } from "./role-actions";
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
    ? RoleActions.fnExec(execFn(options))
    : RoleActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const rolesQuery = useQuery(
    ["*[]workspaces.RoleEntity", options],
    () => Q().getRoles(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const rolesExportQuery = useQuery(
    ["*[]workspaces.RoleEntity", options],
    () => Q().getRolesExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const roleByUniqueIdQuery = useQuery(
    ["*workspaces.RoleEntity", options],
    (uniqueId: string) => Q().getRoleByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post role

  const mutationPostRole = useMutation<
    IResponse<workspaces.RoleEntity>,
    IResponse<workspaces.RoleEntity>,
    workspaces.RoleEntity
  >((entity) => {
    return Q().postRole(entity);
  });

  // Only entities are having a store in front-end

  const fnPostRoleUpdater = (
    data: IResponseList<workspaces.RoleEntity> | undefined,
    item: IResponse<workspaces.RoleEntity>
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
          RoleActions.isRoleEntityEqual(t, item.data)
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

  const submitPostRole = (
    values: workspaces.RoleEntity,
    formikProps?: FormikHelpers<workspaces.RoleEntity>
  ): Promise<IResponse<workspaces.RoleEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostRole.mutate(values, {
        onSuccess(response: IResponse<workspaces.RoleEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.RoleEntity>>(
            "*[]workspaces.RoleEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.RoleEntity) => {
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

  // patch role

  const mutationPatchRole = useMutation<
    IResponse<workspaces.RoleEntity>,
    IResponse<workspaces.RoleEntity>,
    workspaces.RoleEntity
  >((entity) => {
    return Q().patchRole(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchRoleUpdater = (
    data: IResponseList<workspaces.RoleEntity> | undefined,
    item: IResponse<workspaces.RoleEntity>
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
          RoleActions.isRoleEntityEqual(t, item.data)
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

  const submitPatchRole = (
    values: workspaces.RoleEntity,
    formikProps?: FormikHelpers<workspaces.RoleEntity>
  ): Promise<IResponse<workspaces.RoleEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchRole.mutate(values, {
        onSuccess(response: IResponse<workspaces.RoleEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.RoleEntity>>(
            "*[]workspaces.RoleEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.RoleEntity) => {
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

  // patch roles

  const mutationPatchRoles = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.RoleEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.RoleEntity>>,
    core.BulkRecordRequest<workspaces.RoleEntity>
  >((entity) => {
    return Q().patchRoles(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchRoles = (
    values: core.BulkRecordRequest<workspaces.RoleEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<workspaces.RoleEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<workspaces.RoleEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchRoles.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<workspaces.RoleEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.RoleEntity>>
          >("*[]core.BulkRecordRequest[workspaces.RoleEntity]", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: core.BulkRecordRequest<workspaces.RoleEntity>) => {
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
  const mutationDeleteRole = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteRole();
  });

  const fnDeleteRoleUpdater = (
    data: IResponseList<workspaces.RoleEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = RoleActions.getRoleEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteRole = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.RoleEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteRole.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.RoleEntity>>(
            "*[]workspaces.RoleEntity",
            (data) => fnDeleteRoleUpdater(data, values)
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
    rolesQuery,
    rolesExportQuery,
    roleByUniqueIdQuery,
    mutationPostRole,
    submitPostRole,
    mutationPatchRole,
    submitPatchRole,
    mutationPatchRoles,
    submitPatchRoles,
    mutationDeleteRole,
    submitDeleteRole,
  };
}
