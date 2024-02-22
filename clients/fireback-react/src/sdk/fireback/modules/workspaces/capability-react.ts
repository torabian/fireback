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
import { CapabilityActions } from "./capability-actions";
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
    ? CapabilityActions.fnExec(execFn(options))
    : CapabilityActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const capabilitysQuery = useQuery(
    ["*[]workspaces.CapabilityEntity", options],
    () => Q().getCapabilitys(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const capabilitysExportQuery = useQuery(
    ["*[]workspaces.CapabilityEntity", options],
    () => Q().getCapabilitysExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const capabilityByUniqueIdQuery = useQuery(
    ["*workspaces.CapabilityEntity", options],
    (uniqueId: string) => Q().getCapabilityByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post capability

  const mutationPostCapability = useMutation<
    IResponse<workspaces.CapabilityEntity>,
    IResponse<workspaces.CapabilityEntity>,
    workspaces.CapabilityEntity
  >((entity) => {
    return Q().postCapability(entity);
  });

  // Only entities are having a store in front-end

  const fnPostCapabilityUpdater = (
    data: IResponseList<workspaces.CapabilityEntity> | undefined,
    item: IResponse<workspaces.CapabilityEntity>
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
          CapabilityActions.isCapabilityEntityEqual(t, item.data)
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

  const submitPostCapability = (
    values: workspaces.CapabilityEntity,
    formikProps?: FormikHelpers<workspaces.CapabilityEntity>
  ): Promise<IResponse<workspaces.CapabilityEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostCapability.mutate(values, {
        onSuccess(response: IResponse<workspaces.CapabilityEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.CapabilityEntity>
          >("*[]workspaces.CapabilityEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.CapabilityEntity) => {
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

  // patch capability

  const mutationPatchCapability = useMutation<
    IResponse<workspaces.CapabilityEntity>,
    IResponse<workspaces.CapabilityEntity>,
    workspaces.CapabilityEntity
  >((entity) => {
    return Q().patchCapability(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchCapabilityUpdater = (
    data: IResponseList<workspaces.CapabilityEntity> | undefined,
    item: IResponse<workspaces.CapabilityEntity>
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
          CapabilityActions.isCapabilityEntityEqual(t, item.data)
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

  const submitPatchCapability = (
    values: workspaces.CapabilityEntity,
    formikProps?: FormikHelpers<workspaces.CapabilityEntity>
  ): Promise<IResponse<workspaces.CapabilityEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchCapability.mutate(values, {
        onSuccess(response: IResponse<workspaces.CapabilityEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.CapabilityEntity>
          >("*[]workspaces.CapabilityEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.CapabilityEntity) => {
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

  // patch capabilitys

  const mutationPatchCapabilitys = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.CapabilityEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.CapabilityEntity>>,
    core.BulkRecordRequest<workspaces.CapabilityEntity>
  >((entity) => {
    return Q().patchCapabilitys(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchCapabilitys = (
    values: core.BulkRecordRequest<workspaces.CapabilityEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.CapabilityEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.CapabilityEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchCapabilitys.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.CapabilityEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.CapabilityEntity>>
          >(
            "*[]core.BulkRecordRequest[workspaces.CapabilityEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<workspaces.CapabilityEntity>) => {
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
  const mutationDeleteCapability = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteCapability();
  });

  const fnDeleteCapabilityUpdater = (
    data: IResponseList<workspaces.CapabilityEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = CapabilityActions.getCapabilityEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteCapability = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.CapabilityEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteCapability.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.CapabilityEntity>>(
            "*[]workspaces.CapabilityEntity",
            (data) => fnDeleteCapabilityUpdater(data, values)
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

  const capabilitiesTreeQuery = useQuery(
    ["*workspaces.CapabilitiesResult", options],
    () => Q().getCapabilitiesTree(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  return {
    queryClient,
    capabilitysQuery,
    capabilitysExportQuery,
    capabilityByUniqueIdQuery,
    mutationPostCapability,
    submitPostCapability,
    mutationPatchCapability,
    submitPatchCapability,
    mutationPatchCapabilitys,
    submitPatchCapabilitys,
    mutationDeleteCapability,
    submitDeleteCapability,
    capabilitiesTreeQuery,
  };
}
