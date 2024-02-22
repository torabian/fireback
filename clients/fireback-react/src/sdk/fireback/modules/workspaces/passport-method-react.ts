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
import { PassportMethodActions } from "./passport-method-actions";
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
    ? PassportMethodActions.fnExec(execFn(options))
    : PassportMethodActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const passportMethodsQuery = useQuery(
    ["*[]workspaces.PassportMethodEntity", options],
    () => Q().getPassportMethods(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const passportMethodsExportQuery = useQuery(
    ["*[]workspaces.PassportMethodEntity", options],
    () => Q().getPassportMethodsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const passportMethodByUniqueIdQuery = useQuery(
    ["*workspaces.PassportMethodEntity", options],
    (uniqueId: string) => Q().getPassportMethodByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post passportMethod

  const mutationPostPassportMethod = useMutation<
    IResponse<workspaces.PassportMethodEntity>,
    IResponse<workspaces.PassportMethodEntity>,
    workspaces.PassportMethodEntity
  >((entity) => {
    return Q().postPassportMethod(entity);
  });

  // Only entities are having a store in front-end

  const fnPostPassportMethodUpdater = (
    data: IResponseList<workspaces.PassportMethodEntity> | undefined,
    item: IResponse<workspaces.PassportMethodEntity>
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
          PassportMethodActions.isPassportMethodEntityEqual(t, item.data)
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

  const submitPostPassportMethod = (
    values: workspaces.PassportMethodEntity,
    formikProps?: FormikHelpers<workspaces.PassportMethodEntity>
  ): Promise<IResponse<workspaces.PassportMethodEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostPassportMethod.mutate(values, {
        onSuccess(response: IResponse<workspaces.PassportMethodEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.PassportMethodEntity>
          >("*[]workspaces.PassportMethodEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.PassportMethodEntity) => {
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

  // patch passportMethod

  const mutationPatchPassportMethod = useMutation<
    IResponse<workspaces.PassportMethodEntity>,
    IResponse<workspaces.PassportMethodEntity>,
    workspaces.PassportMethodEntity
  >((entity) => {
    return Q().patchPassportMethod(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchPassportMethodUpdater = (
    data: IResponseList<workspaces.PassportMethodEntity> | undefined,
    item: IResponse<workspaces.PassportMethodEntity>
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
          PassportMethodActions.isPassportMethodEntityEqual(t, item.data)
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

  const submitPatchPassportMethod = (
    values: workspaces.PassportMethodEntity,
    formikProps?: FormikHelpers<workspaces.PassportMethodEntity>
  ): Promise<IResponse<workspaces.PassportMethodEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchPassportMethod.mutate(values, {
        onSuccess(response: IResponse<workspaces.PassportMethodEntity>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.PassportMethodEntity>
          >("*[]workspaces.PassportMethodEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.PassportMethodEntity) => {
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

  // patch passportMethods

  const mutationPatchPassportMethods = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.PassportMethodEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.PassportMethodEntity>>,
    core.BulkRecordRequest<workspaces.PassportMethodEntity>
  >((entity) => {
    return Q().patchPassportMethods(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchPassportMethods = (
    values: core.BulkRecordRequest<workspaces.PassportMethodEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.PassportMethodEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.PassportMethodEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchPassportMethods.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.PassportMethodEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<workspaces.PassportMethodEntity>
            >
          >(
            "*[]core.BulkRecordRequest[workspaces.PassportMethodEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<workspaces.PassportMethodEntity>
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
  const mutationDeletePassportMethod = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deletePassportMethod();
  });

  const fnDeletePassportMethodUpdater = (
    data: IResponseList<workspaces.PassportMethodEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = PassportMethodActions.getPassportMethodEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeletePassportMethod = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.PassportMethodEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeletePassportMethod.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<workspaces.PassportMethodEntity>
          >("*[]workspaces.PassportMethodEntity", (data) =>
            fnDeletePassportMethodUpdater(data, values)
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
    passportMethodsQuery,
    passportMethodsExportQuery,
    passportMethodByUniqueIdQuery,
    mutationPostPassportMethod,
    submitPostPassportMethod,
    mutationPatchPassportMethod,
    submitPatchPassportMethod,
    mutationPatchPassportMethods,
    submitPatchPassportMethods,
    mutationDeletePassportMethod,
    submitDeletePassportMethod,
  };
}
