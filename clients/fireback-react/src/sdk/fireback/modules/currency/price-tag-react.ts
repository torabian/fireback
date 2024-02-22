// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: currency
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { PriceTagActions } from "./price-tag-actions";
import * as currency from "./index";
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
export function useCurrency(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? PriceTagActions.fnExec(execFn(options))
    : PriceTagActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const priceTagsQuery = useQuery(
    ["*[]currency.PriceTagEntity", options],
    () => Q().getPriceTags(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const priceTagsExportQuery = useQuery(
    ["*[]currency.PriceTagEntity", options],
    () => Q().getPriceTagsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const priceTagByUniqueIdQuery = useQuery(
    ["*currency.PriceTagEntity", options],
    (uniqueId: string) => Q().getPriceTagByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post priceTag

  const mutationPostPriceTag = useMutation<
    IResponse<currency.PriceTagEntity>,
    IResponse<currency.PriceTagEntity>,
    currency.PriceTagEntity
  >((entity) => {
    return Q().postPriceTag(entity);
  });

  // Only entities are having a store in front-end

  const fnPostPriceTagUpdater = (
    data: IResponseList<currency.PriceTagEntity> | undefined,
    item: IResponse<currency.PriceTagEntity>
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
          PriceTagActions.isPriceTagEntityEqual(t, item.data)
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

  const submitPostPriceTag = (
    values: currency.PriceTagEntity,
    formikProps?: FormikHelpers<currency.PriceTagEntity>
  ): Promise<IResponse<currency.PriceTagEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostPriceTag.mutate(values, {
        onSuccess(response: IResponse<currency.PriceTagEntity>) {
          queryClient.setQueriesData<IResponseList<currency.PriceTagEntity>>(
            "*[]currency.PriceTagEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: currency.PriceTagEntity) => {
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

  // patch priceTag

  const mutationPatchPriceTag = useMutation<
    IResponse<currency.PriceTagEntity>,
    IResponse<currency.PriceTagEntity>,
    currency.PriceTagEntity
  >((entity) => {
    return Q().patchPriceTag(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchPriceTagUpdater = (
    data: IResponseList<currency.PriceTagEntity> | undefined,
    item: IResponse<currency.PriceTagEntity>
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
          PriceTagActions.isPriceTagEntityEqual(t, item.data)
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

  const submitPatchPriceTag = (
    values: currency.PriceTagEntity,
    formikProps?: FormikHelpers<currency.PriceTagEntity>
  ): Promise<IResponse<currency.PriceTagEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchPriceTag.mutate(values, {
        onSuccess(response: IResponse<currency.PriceTagEntity>) {
          queryClient.setQueriesData<IResponseList<currency.PriceTagEntity>>(
            "*[]currency.PriceTagEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: currency.PriceTagEntity) => {
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

  // patch priceTags

  const mutationPatchPriceTags = useMutation<
    IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>,
    IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>,
    core.BulkRecordRequest<currency.PriceTagEntity>
  >((entity) => {
    return Q().patchPriceTags(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchPriceTags = (
    values: core.BulkRecordRequest<currency.PriceTagEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<currency.PriceTagEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchPriceTags.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<currency.PriceTagEntity>>
          >(
            "*[]core.BulkRecordRequest[currency.PriceTagEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<currency.PriceTagEntity>) => {
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
  const mutationDeletePriceTag = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deletePriceTag();
  });

  const fnDeletePriceTagUpdater = (
    data: IResponseList<currency.PriceTagEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = PriceTagActions.getPriceTagEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeletePriceTag = (
    values: string[],
    formikProps?: FormikHelpers<currency.PriceTagEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeletePriceTag.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<currency.PriceTagEntity>>(
            "*[]currency.PriceTagEntity",
            (data) => fnDeletePriceTagUpdater(data, values)
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
    priceTagsQuery,
    priceTagsExportQuery,
    priceTagByUniqueIdQuery,
    mutationPostPriceTag,
    submitPostPriceTag,
    mutationPatchPriceTag,
    submitPatchPriceTag,
    mutationPatchPriceTags,
    submitPatchPriceTags,
    mutationDeletePriceTag,
    submitDeletePriceTag,
  };
}
