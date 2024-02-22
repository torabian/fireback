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
import { AppMenuActions } from "./app-menu-actions";
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
    ? AppMenuActions.fnExec(execFn(options))
    : AppMenuActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const cteAppMenusQuery = useQuery(
    ["*[]workspaces.AppMenuEntity", options],
    () => Q().getCteAppMenus(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const appMenusQuery = useQuery(
    ["*[]workspaces.AppMenuEntity", options],
    () => Q().getAppMenus(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const appMenusExportQuery = useQuery(
    ["*[]workspaces.AppMenuEntity", options],
    () => Q().getAppMenusExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const appMenuByUniqueIdQuery = useQuery(
    ["*workspaces.AppMenuEntity", options],
    (uniqueId: string) => Q().getAppMenuByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post appMenu

  const mutationPostAppMenu = useMutation<
    IResponse<workspaces.AppMenuEntity>,
    IResponse<workspaces.AppMenuEntity>,
    workspaces.AppMenuEntity
  >((entity) => {
    return Q().postAppMenu(entity);
  });

  // Only entities are having a store in front-end

  const fnPostAppMenuUpdater = (
    data: IResponseList<workspaces.AppMenuEntity> | undefined,
    item: IResponse<workspaces.AppMenuEntity>
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
          AppMenuActions.isAppMenuEntityEqual(t, item.data)
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

  const submitPostAppMenu = (
    values: workspaces.AppMenuEntity,
    formikProps?: FormikHelpers<workspaces.AppMenuEntity>
  ): Promise<IResponse<workspaces.AppMenuEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostAppMenu.mutate(values, {
        onSuccess(response: IResponse<workspaces.AppMenuEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.AppMenuEntity>>(
            "*[]workspaces.AppMenuEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.AppMenuEntity) => {
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

  // patch appMenu

  const mutationPatchAppMenu = useMutation<
    IResponse<workspaces.AppMenuEntity>,
    IResponse<workspaces.AppMenuEntity>,
    workspaces.AppMenuEntity
  >((entity) => {
    return Q().patchAppMenu(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchAppMenuUpdater = (
    data: IResponseList<workspaces.AppMenuEntity> | undefined,
    item: IResponse<workspaces.AppMenuEntity>
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
          AppMenuActions.isAppMenuEntityEqual(t, item.data)
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

  const submitPatchAppMenu = (
    values: workspaces.AppMenuEntity,
    formikProps?: FormikHelpers<workspaces.AppMenuEntity>
  ): Promise<IResponse<workspaces.AppMenuEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchAppMenu.mutate(values, {
        onSuccess(response: IResponse<workspaces.AppMenuEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.AppMenuEntity>>(
            "*[]workspaces.AppMenuEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.AppMenuEntity) => {
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

  // patch appMenus

  const mutationPatchAppMenus = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.AppMenuEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.AppMenuEntity>>,
    core.BulkRecordRequest<workspaces.AppMenuEntity>
  >((entity) => {
    return Q().patchAppMenus(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchAppMenus = (
    values: core.BulkRecordRequest<workspaces.AppMenuEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.AppMenuEntity>
    >
  ): Promise<IResponse<core.BulkRecordRequest<workspaces.AppMenuEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchAppMenus.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<workspaces.AppMenuEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.AppMenuEntity>>
          >(
            "*[]core.BulkRecordRequest[workspaces.AppMenuEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<workspaces.AppMenuEntity>) => {
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
  const mutationDeleteAppMenu = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteAppMenu();
  });

  const fnDeleteAppMenuUpdater = (
    data: IResponseList<workspaces.AppMenuEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = AppMenuActions.getAppMenuEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteAppMenu = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.AppMenuEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteAppMenu.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.AppMenuEntity>>(
            "*[]workspaces.AppMenuEntity",
            (data) => fnDeleteAppMenuUpdater(data, values)
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
    cteAppMenusQuery,
    appMenusQuery,
    appMenusExportQuery,
    appMenuByUniqueIdQuery,
    mutationPostAppMenu,
    submitPostAppMenu,
    mutationPatchAppMenu,
    submitPatchAppMenu,
    mutationPatchAppMenus,
    submitPatchAppMenus,
    mutationDeleteAppMenu,
    submitDeleteAppMenu,
  };
}
