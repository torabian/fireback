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
import { WorkspaceTypeActions } from "./workspace-type-actions";
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
    ? WorkspaceTypeActions.fnExec(execFn(options))
    : WorkspaceTypeActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const workspaceTypesQuery = useQuery(
    ["*[]workspaces.WorkspaceTypeEntity", options],
    () => Q().getWorkspaceTypes(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspaceTypesExportQuery = useQuery(
    ["*[]workspaces.WorkspaceTypeEntity", options],
    () => Q().getWorkspaceTypesExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const workspaceTypeByUniqueIdQuery = useQuery(
    ["*workspaces.WorkspaceTypeEntity", options],
    (uniqueId: string) => Q().getWorkspaceTypeByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post workspaceType

  const mutationPostWorkspaceType = useMutation<
    IResponse<workspaces.WorkspaceTypeEntity>,
    IResponse<workspaces.WorkspaceTypeEntity>,
    workspaces.WorkspaceTypeEntity
  >((entity) => {
    return Q().postWorkspaceType(entity);
  });

  // Only entities are having a store in front-end

  const fnPostWorkspaceTypeUpdater = (
    data: IResponseList<workspaces.WorkspaceTypeEntity> | undefined,
    item: IResponse<workspaces.WorkspaceTypeEntity>
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
          WorkspaceTypeActions.isWorkspaceTypeEntityEqual(t, item.data)
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

  const submitPostWorkspaceType = (
    values: workspaces.WorkspaceTypeEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceTypeEntity>
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostWorkspaceType.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceTypeEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.WorkspaceTypeEntity>
          >("*[]workspaces.WorkspaceTypeEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.WorkspaceTypeEntity) => {
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

  // patch workspaceType

  const mutationPatchWorkspaceType = useMutation<
    IResponse<workspaces.WorkspaceTypeEntity>,
    IResponse<workspaces.WorkspaceTypeEntity>,
    workspaces.WorkspaceTypeEntity
  >((entity) => {
    return Q().patchWorkspaceType(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchWorkspaceTypeUpdater = (
    data: IResponseList<workspaces.WorkspaceTypeEntity> | undefined,
    item: IResponse<workspaces.WorkspaceTypeEntity>
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
          WorkspaceTypeActions.isWorkspaceTypeEntityEqual(t, item.data)
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

  const submitPatchWorkspaceType = (
    values: workspaces.WorkspaceTypeEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceTypeEntity>
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspaceType.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceTypeEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.WorkspaceTypeEntity>
          >("*[]workspaces.WorkspaceTypeEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.WorkspaceTypeEntity) => {
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

  // patch workspaceTypeDistinct

  const mutationPatchWorkspaceTypeDistinct = useMutation<
    IResponse<workspaces.WorkspaceTypeEntity>,
    IResponse<workspaces.WorkspaceTypeEntity>,
    workspaces.WorkspaceTypeEntity
  >((entity) => {
    return Q().patchWorkspaceTypeDistinct(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchWorkspaceTypeDistinctUpdater = (
    data: IResponseList<workspaces.WorkspaceTypeEntity> | undefined,
    item: IResponse<workspaces.WorkspaceTypeEntity>
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
          WorkspaceTypeActions.isWorkspaceTypeEntityEqual(t, item.data)
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

  const submitPatchWorkspaceTypeDistinct = (
    values: workspaces.WorkspaceTypeEntity,
    formikProps?: FormikHelpers<workspaces.WorkspaceTypeEntity>
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspaceTypeDistinct.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceTypeEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.WorkspaceTypeEntity>
          >("*[]workspaces.WorkspaceTypeEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.WorkspaceTypeEntity) => {
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

  const workspaceTypeDistinctQuery = useQuery(
    ["*workspaces.WorkspaceTypeEntity", options],
    () => Q().getWorkspaceTypeDistinct(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // patch workspaceTypes

  const mutationPatchWorkspaceTypes = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>>,
    core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>
  >((entity) => {
    return Q().patchWorkspaceTypes(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchWorkspaceTypes = (
    values: core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchWorkspaceTypes.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.WorkspaceTypeEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>
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
  const mutationDeleteWorkspaceType = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteWorkspaceType();
  });

  const fnDeleteWorkspaceTypeUpdater = (
    data: IResponseList<workspaces.WorkspaceTypeEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = WorkspaceTypeActions.getWorkspaceTypeEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteWorkspaceType = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.WorkspaceTypeEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteWorkspaceType.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.WorkspaceTypeEntity>
          >("*[]workspaces.WorkspaceTypeEntity", (data) =>
            fnDeleteWorkspaceTypeUpdater(data, values)
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

  const publicWorkspaceTypesQuery = useQuery(
    ["*[]workspaces.WorkspaceTypeEntity", options],
    () => Q().getPublicWorkspaceTypes(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  return {
    queryClient,
    workspaceTypesQuery,
    workspaceTypesExportQuery,
    workspaceTypeByUniqueIdQuery,
    mutationPostWorkspaceType,
    submitPostWorkspaceType,
    mutationPatchWorkspaceType,
    submitPatchWorkspaceType,
    mutationPatchWorkspaceTypeDistinct,
    submitPatchWorkspaceTypeDistinct,
    workspaceTypeDistinctQuery,
    mutationPatchWorkspaceTypes,
    submitPatchWorkspaceTypes,
    mutationDeleteWorkspaceType,
    submitDeleteWorkspaceType,
    publicWorkspaceTypesQuery,
  };
}
