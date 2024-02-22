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
import { WorkspaceActions } from "./workspace-actions";
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
    ? WorkspaceActions.fnExec(execFn(options))
    : WorkspaceActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const cteWorkspacesQuery = useQuery(
    ["*[]workspaces.WorkspaceEntity", options],
    () => Q().getCteWorkspaces(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspacesQuery = useQuery(
    ["*[]workspaces.WorkspaceEntity", options],
    () => Q().getWorkspaces(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspacesExportQuery = useQuery(
    ["*[]workspaces.WorkspaceEntity", options],
    () => Q().getWorkspacesExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspaceByUniqueIdQuery = useQuery(
    ["*workspaces.WorkspaceEntity", options],
    (uniqueId: string) => Q().getWorkspaceByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post workspace

  const mutationPostWorkspace = useMutation<
    IResponse<workspaces.WorkspaceEntity>,
    IResponse<workspaces.WorkspaceEntity>,
    workspaces.WorkspaceEntity
  >((entity) => {
    return Q().postWorkspace(entity);
  });

  // Only entities are having a store in front-end

  const fnPostWorkspaceUpdater = (
    data: IResponseList<workspaces.WorkspaceEntity> | undefined,
    item: IResponse<workspaces.WorkspaceEntity>
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
          WorkspaceActions.isWorkspaceEntityEqual(t, item.data)
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

  const submitPostWorkspace = (
    values: workspaces.WorkspaceEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceEntity>
  ): Promise<IResponse<workspaces.WorkspaceEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostWorkspace.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.WorkspaceEntity>>(
            "*[]workspaces.WorkspaceEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.WorkspaceEntity) => {
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

  // patch workspace

  const mutationPatchWorkspace = useMutation<
    IResponse<workspaces.WorkspaceEntity>,
    IResponse<workspaces.WorkspaceEntity>,
    workspaces.WorkspaceEntity
  >((entity) => {
    return Q().patchWorkspace(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchWorkspaceUpdater = (
    data: IResponseList<workspaces.WorkspaceEntity> | undefined,
    item: IResponse<workspaces.WorkspaceEntity>
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
          WorkspaceActions.isWorkspaceEntityEqual(t, item.data)
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

  const submitPatchWorkspace = (
    values: workspaces.WorkspaceEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceEntity>
  ): Promise<IResponse<workspaces.WorkspaceEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspace.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.WorkspaceEntity>>(
            "*[]workspaces.WorkspaceEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.WorkspaceEntity) => {
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

  // patch workspaces

  const mutationPatchWorkspaces = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceEntity>>,
    core.BulkRecordRequest<workspaces.WorkspaceEntity>
  >((entity) => {
    return Q().patchWorkspaces(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchWorkspaces = (
    values: core.BulkRecordRequest<workspaces.WorkspaceEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.WorkspaceEntity>
    >
  ): Promise<IResponse<core.BulkRecordRequest<workspaces.WorkspaceEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspaces.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.WorkspaceEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.WorkspaceEntity>>
          >(
            "*[]core.BulkRecordRequest[workspaces.WorkspaceEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<workspaces.WorkspaceEntity>) => {
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
  const mutationDeleteWorkspace = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteWorkspace();
  });

  const fnDeleteWorkspaceUpdater = (
    data: IResponseList<workspaces.WorkspaceEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = WorkspaceActions.getWorkspaceEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteWorkspace = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.WorkspaceEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteWorkspace.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.WorkspaceEntity>>(
            "*[]workspaces.WorkspaceEntity",
            (data) => fnDeleteWorkspaceUpdater(data, values)
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
    cteWorkspacesQuery,
    workspacesQuery,
    workspacesExportQuery,
    workspaceByUniqueIdQuery,
    mutationPostWorkspace,
    submitPostWorkspace,
    mutationPatchWorkspace,
    submitPatchWorkspace,
    mutationPatchWorkspaces,
    submitPatchWorkspaces,
    mutationDeleteWorkspace,
    submitDeleteWorkspace,
  };
}
