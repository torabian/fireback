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
import { WorkspaceInviteActions } from "./workspace-invite-actions";
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
    ? WorkspaceInviteActions.fnExec(execFn(options))
    : WorkspaceInviteActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const workspaceInvitesQuery = useQuery(
    ["*[]workspaces.WorkspaceInviteEntity", options],
    () => Q().getWorkspaceInvites(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspaceInvitesExportQuery = useQuery(
    ["*[]workspaces.WorkspaceInviteEntity", options],
    () => Q().getWorkspaceInvitesExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspaceInviteByUniqueIdQuery = useQuery(
    ["*workspaces.WorkspaceInviteEntity", options],
    (uniqueId: string) => Q().getWorkspaceInviteByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post workspaceInvite

  const mutationPostWorkspaceInvite = useMutation<
    IResponse<workspaces.WorkspaceInviteEntity>,
    IResponse<workspaces.WorkspaceInviteEntity>,
    workspaces.WorkspaceInviteEntity
  >((entity) => {
    return Q().postWorkspaceInvite(entity);
  });

  // Only entities are having a store in front-end

  const fnPostWorkspaceInviteUpdater = (
    data: IResponseList<workspaces.WorkspaceInviteEntity> | undefined,
    item: IResponse<workspaces.WorkspaceInviteEntity>
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
          WorkspaceInviteActions.isWorkspaceInviteEntityEqual(t, item.data)
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

  const submitPostWorkspaceInvite = (
    values: workspaces.WorkspaceInviteEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceInviteEntity>
  ): Promise<IResponse<workspaces.WorkspaceInviteEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostWorkspaceInvite.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceInviteEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.WorkspaceInviteEntity>
          >("*[]workspaces.WorkspaceInviteEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.WorkspaceInviteEntity) => {
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

  // patch workspaceInvite

  const mutationPatchWorkspaceInvite = useMutation<
    IResponse<workspaces.WorkspaceInviteEntity>,
    IResponse<workspaces.WorkspaceInviteEntity>,
    workspaces.WorkspaceInviteEntity
  >((entity) => {
    return Q().patchWorkspaceInvite(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchWorkspaceInviteUpdater = (
    data: IResponseList<workspaces.WorkspaceInviteEntity> | undefined,
    item: IResponse<workspaces.WorkspaceInviteEntity>
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
          WorkspaceInviteActions.isWorkspaceInviteEntityEqual(t, item.data)
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

  const submitPatchWorkspaceInvite = (
    values: workspaces.WorkspaceInviteEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceInviteEntity>
  ): Promise<IResponse<workspaces.WorkspaceInviteEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspaceInvite.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceInviteEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.WorkspaceInviteEntity>
          >("*[]workspaces.WorkspaceInviteEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.WorkspaceInviteEntity) => {
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

  // patch workspaceInvites

  const mutationPatchWorkspaceInvites = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>,
    core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
  >((entity) => {
    return Q().patchWorkspaceInvites(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchWorkspaceInvites = (
    values: core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspaceInvites.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.WorkspaceInviteEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
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
  const mutationDeleteWorkspaceInvite = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteWorkspaceInvite();
  });

  const fnDeleteWorkspaceInviteUpdater = (
    data: IResponseList<workspaces.WorkspaceInviteEntity> | undefined,
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
          WorkspaceInviteActions.getWorkspaceInviteEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteWorkspaceInvite = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.WorkspaceInviteEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteWorkspaceInvite.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.WorkspaceInviteEntity>
          >("*[]workspaces.WorkspaceInviteEntity", (data) =>
            fnDeleteWorkspaceInviteUpdater(data, values)
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
    workspaceInvitesQuery,
    workspaceInvitesExportQuery,
    workspaceInviteByUniqueIdQuery,
    mutationPostWorkspaceInvite,
    submitPostWorkspaceInvite,
    mutationPatchWorkspaceInvite,
    submitPatchWorkspaceInvite,
    mutationPatchWorkspaceInvites,
    submitPatchWorkspaceInvites,
    mutationDeleteWorkspaceInvite,
    submitDeleteWorkspaceInvite,
  };
}
