// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: licenses
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { ProductPlanActions } from "./product-plan-actions";
import * as licenses from "./index";
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
export function useLicenses(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? ProductPlanActions.fnExec(execFn(options))
    : ProductPlanActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const productPlansQuery = useQuery(
    ["*[]licenses.ProductPlanEntity", options],
    () => Q().getProductPlans(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const productPlansExportQuery = useQuery(
    ["*[]licenses.ProductPlanEntity", options],
    () => Q().getProductPlansExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const productPlanByUniqueIdQuery = useQuery(
    ["*licenses.ProductPlanEntity", options],
    (uniqueId: string) => Q().getProductPlanByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post productPlan

  const mutationPostProductPlan = useMutation<
    IResponse<licenses.ProductPlanEntity>,
    IResponse<licenses.ProductPlanEntity>,
    licenses.ProductPlanEntity
  >((entity) => {
    return Q().postProductPlan(entity);
  });

  // Only entities are having a store in front-end

  const fnPostProductPlanUpdater = (
    data: IResponseList<licenses.ProductPlanEntity> | undefined,
    item: IResponse<licenses.ProductPlanEntity>
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
          ProductPlanActions.isProductPlanEntityEqual(t, item.data)
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

  const submitPostProductPlan = (
    values: licenses.ProductPlanEntity,
    formikProps?: FormikHelpers<licenses.ProductPlanEntity>
  ): Promise<IResponse<licenses.ProductPlanEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostProductPlan.mutate(values, {
        onSuccess(response: IResponse<licenses.ProductPlanEntity>) {
          queryClient.setQueriesData<IResponseList<licenses.ProductPlanEntity>>(
            "*[]licenses.ProductPlanEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: licenses.ProductPlanEntity) => {
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

  // patch productPlan

  const mutationPatchProductPlan = useMutation<
    IResponse<licenses.ProductPlanEntity>,
    IResponse<licenses.ProductPlanEntity>,
    licenses.ProductPlanEntity
  >((entity) => {
    return Q().patchProductPlan(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchProductPlanUpdater = (
    data: IResponseList<licenses.ProductPlanEntity> | undefined,
    item: IResponse<licenses.ProductPlanEntity>
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
          ProductPlanActions.isProductPlanEntityEqual(t, item.data)
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

  const submitPatchProductPlan = (
    values: licenses.ProductPlanEntity,
    formikProps?: FormikHelpers<licenses.ProductPlanEntity>
  ): Promise<IResponse<licenses.ProductPlanEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchProductPlan.mutate(values, {
        onSuccess(response: IResponse<licenses.ProductPlanEntity>) {
          queryClient.setQueriesData<IResponseList<licenses.ProductPlanEntity>>(
            "*[]licenses.ProductPlanEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: licenses.ProductPlanEntity) => {
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

  // patch productPlans

  const mutationPatchProductPlans = useMutation<
    IResponse<core.BulkRecordRequest<licenses.ProductPlanEntity>>,
    IResponse<core.BulkRecordRequest<licenses.ProductPlanEntity>>,
    core.BulkRecordRequest<licenses.ProductPlanEntity>
  >((entity) => {
    return Q().patchProductPlans(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchProductPlans = (
    values: core.BulkRecordRequest<licenses.ProductPlanEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<licenses.ProductPlanEntity>
    >
  ): Promise<IResponse<core.BulkRecordRequest<licenses.ProductPlanEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchProductPlans.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<licenses.ProductPlanEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<licenses.ProductPlanEntity>>
          >(
            "*[]core.BulkRecordRequest[licenses.ProductPlanEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<licenses.ProductPlanEntity>) => {
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
  const mutationDeleteProductPlan = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteProductPlan();
  });

  const fnDeleteProductPlanUpdater = (
    data: IResponseList<licenses.ProductPlanEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = ProductPlanActions.getProductPlanEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteProductPlan = (
    values: string[],
    formikProps?: FormikHelpers<licenses.ProductPlanEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteProductPlan.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<licenses.ProductPlanEntity>>(
            "*[]licenses.ProductPlanEntity",
            (data) => fnDeleteProductPlanUpdater(data, values)
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
    productPlansQuery,
    productPlansExportQuery,
    productPlanByUniqueIdQuery,
    mutationPostProductPlan,
    submitPostProductPlan,
    mutationPatchProductPlan,
    submitPatchProductPlan,
    mutationPatchProductPlans,
    submitPatchProductPlans,
    mutationDeleteProductPlan,
    submitDeleteProductPlan,
  };
}
