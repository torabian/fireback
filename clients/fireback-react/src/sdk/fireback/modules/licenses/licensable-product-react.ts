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
import { LicensableProductActions } from "./licensable-product-actions";
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
    ? LicensableProductActions.fnExec(execFn(options))
    : LicensableProductActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const licensableProductsQuery = useQuery(
    ["*[]licenses.LicensableProductEntity", options],
    () => Q().getLicensableProducts(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const licensableProductsExportQuery = useQuery(
    ["*[]licenses.LicensableProductEntity", options],
    () => Q().getLicensableProductsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const licensableProductByUniqueIdQuery = useQuery(
    ["*licenses.LicensableProductEntity", options],
    (uniqueId: string) => Q().getLicensableProductByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post licensableProduct

  const mutationPostLicensableProduct = useMutation<
    IResponse<licenses.LicensableProductEntity>,
    IResponse<licenses.LicensableProductEntity>,
    licenses.LicensableProductEntity
  >((entity) => {
    return Q().postLicensableProduct(entity);
  });

  // Only entities are having a store in front-end

  const fnPostLicensableProductUpdater = (
    data: IResponseList<licenses.LicensableProductEntity> | undefined,
    item: IResponse<licenses.LicensableProductEntity>
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
          LicensableProductActions.isLicensableProductEntityEqual(t, item.data)
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

  const submitPostLicensableProduct = (
    values: licenses.LicensableProductEntity,
    formikProps?: FormikHelpers<licenses.LicensableProductEntity>
  ): Promise<IResponse<licenses.LicensableProductEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostLicensableProduct.mutate(values, {
        onSuccess(response: IResponse<licenses.LicensableProductEntity>) {
          queryClient.setQueriesData<
            IResponseList<licenses.LicensableProductEntity>
          >("*[]licenses.LicensableProductEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: licenses.LicensableProductEntity) => {
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

  // patch licensableProduct

  const mutationPatchLicensableProduct = useMutation<
    IResponse<licenses.LicensableProductEntity>,
    IResponse<licenses.LicensableProductEntity>,
    licenses.LicensableProductEntity
  >((entity) => {
    return Q().patchLicensableProduct(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchLicensableProductUpdater = (
    data: IResponseList<licenses.LicensableProductEntity> | undefined,
    item: IResponse<licenses.LicensableProductEntity>
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
          LicensableProductActions.isLicensableProductEntityEqual(t, item.data)
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

  const submitPatchLicensableProduct = (
    values: licenses.LicensableProductEntity,
    formikProps?: FormikHelpers<licenses.LicensableProductEntity>
  ): Promise<IResponse<licenses.LicensableProductEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchLicensableProduct.mutate(values, {
        onSuccess(response: IResponse<licenses.LicensableProductEntity>) {
          queryClient.setQueriesData<
            IResponseList<licenses.LicensableProductEntity>
          >("*[]licenses.LicensableProductEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: licenses.LicensableProductEntity) => {
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

  // patch licensableProducts

  const mutationPatchLicensableProducts = useMutation<
    IResponse<core.BulkRecordRequest<licenses.LicensableProductEntity>>,
    IResponse<core.BulkRecordRequest<licenses.LicensableProductEntity>>,
    core.BulkRecordRequest<licenses.LicensableProductEntity>
  >((entity) => {
    return Q().patchLicensableProducts(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchLicensableProducts = (
    values: core.BulkRecordRequest<licenses.LicensableProductEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<licenses.LicensableProductEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<licenses.LicensableProductEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchLicensableProducts.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<licenses.LicensableProductEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<licenses.LicensableProductEntity>
            >
          >(
            "*[]core.BulkRecordRequest[licenses.LicensableProductEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<licenses.LicensableProductEntity>
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
  const mutationDeleteLicensableProduct = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteLicensableProduct();
  });

  const fnDeleteLicensableProductUpdater = (
    data: IResponseList<licenses.LicensableProductEntity> | undefined,
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
          LicensableProductActions.getLicensableProductEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteLicensableProduct = (
    values: string[],
    formikProps?: FormikHelpers<licenses.LicensableProductEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteLicensableProduct.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<licenses.LicensableProductEntity>
          >("*[]licenses.LicensableProductEntity", (data) =>
            fnDeleteLicensableProductUpdater(data, values)
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

  // post licensableProducts/generate

  const mutationPostLicensableProductsGenerate = useMutation<
    IResponse<licenses.LicensableProductEntity>,
    IResponse<licenses.LicensableProductEntity>,
    licenses.LicensableProductEntity
  >((entity) => {
    return Q().postLicensableProductsGenerate(entity);
  });

  // Only entities are having a store in front-end

  const fnPostLicensableProductsGenerateUpdater = (
    data: IResponseList<licenses.LicensableProductEntity> | undefined,
    item: IResponse<licenses.LicensableProductEntity>
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
          LicensableProductActions.isLicensableProductEntityEqual(t, item.data)
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

  const submitPostLicensableProductsGenerate = (
    values: licenses.LicensableProductEntity,
    formikProps?: FormikHelpers<licenses.LicensableProductEntity>
  ): Promise<IResponse<licenses.LicensableProductEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostLicensableProductsGenerate.mutate(values, {
        onSuccess(response: IResponse<licenses.LicensableProductEntity>) {
          queryClient.setQueriesData<
            IResponseList<licenses.LicensableProductEntity>
          >("*[]licenses.LicensableProductEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: licenses.LicensableProductEntity) => {
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
    licensableProductsQuery,
    licensableProductsExportQuery,
    licensableProductByUniqueIdQuery,
    mutationPostLicensableProduct,
    submitPostLicensableProduct,
    mutationPatchLicensableProduct,
    submitPatchLicensableProduct,
    mutationPatchLicensableProducts,
    submitPatchLicensableProducts,
    mutationDeleteLicensableProduct,
    submitDeleteLicensableProduct,
    mutationPostLicensableProductsGenerate,
    submitPostLicensableProductsGenerate,
  };
}
