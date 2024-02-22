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
import { CurrencyActions } from "./currency-actions";
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
    ? CurrencyActions.fnExec(execFn(options))
    : CurrencyActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const currencysQuery = useQuery(
    ["*[]currency.CurrencyEntity", options],
    () => Q().getCurrencys(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const currencysExportQuery = useQuery(
    ["*[]currency.CurrencyEntity", options],
    () => Q().getCurrencysExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const currencyByUniqueIdQuery = useQuery(
    ["*currency.CurrencyEntity", options],
    (uniqueId: string) => Q().getCurrencyByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post currency

  const mutationPostCurrency = useMutation<
    IResponse<currency.CurrencyEntity>,
    IResponse<currency.CurrencyEntity>,
    currency.CurrencyEntity
  >((entity) => {
    return Q().postCurrency(entity);
  });

  // Only entities are having a store in front-end

  const fnPostCurrencyUpdater = (
    data: IResponseList<currency.CurrencyEntity> | undefined,
    item: IResponse<currency.CurrencyEntity>
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
          CurrencyActions.isCurrencyEntityEqual(t, item.data)
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

  const submitPostCurrency = (
    values: currency.CurrencyEntity,
    formikProps?: FormikHelpers<currency.CurrencyEntity>
  ): Promise<IResponse<currency.CurrencyEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostCurrency.mutate(values, {
        onSuccess(response: IResponse<currency.CurrencyEntity>) {
          queryClient.setQueriesData<IResponseList<currency.CurrencyEntity>>(
            "*[]currency.CurrencyEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: currency.CurrencyEntity) => {
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

  // patch currency

  const mutationPatchCurrency = useMutation<
    IResponse<currency.CurrencyEntity>,
    IResponse<currency.CurrencyEntity>,
    currency.CurrencyEntity
  >((entity) => {
    return Q().patchCurrency(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchCurrencyUpdater = (
    data: IResponseList<currency.CurrencyEntity> | undefined,
    item: IResponse<currency.CurrencyEntity>
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
          CurrencyActions.isCurrencyEntityEqual(t, item.data)
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

  const submitPatchCurrency = (
    values: currency.CurrencyEntity,
    formikProps?: FormikHelpers<currency.CurrencyEntity>
  ): Promise<IResponse<currency.CurrencyEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchCurrency.mutate(values, {
        onSuccess(response: IResponse<currency.CurrencyEntity>) {
          queryClient.setQueriesData<IResponseList<currency.CurrencyEntity>>(
            "*[]currency.CurrencyEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: currency.CurrencyEntity) => {
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

  // patch currencys

  const mutationPatchCurrencys = useMutation<
    IResponse<core.BulkRecordRequest<currency.CurrencyEntity>>,
    IResponse<core.BulkRecordRequest<currency.CurrencyEntity>>,
    core.BulkRecordRequest<currency.CurrencyEntity>
  >((entity) => {
    return Q().patchCurrencys(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchCurrencys = (
    values: core.BulkRecordRequest<currency.CurrencyEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<currency.CurrencyEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<currency.CurrencyEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchCurrencys.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<currency.CurrencyEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<currency.CurrencyEntity>>
          >(
            "*[]core.BulkRecordRequest[currency.CurrencyEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<currency.CurrencyEntity>) => {
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
  const mutationDeleteCurrency = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteCurrency();
  });

  const fnDeleteCurrencyUpdater = (
    data: IResponseList<currency.CurrencyEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = CurrencyActions.getCurrencyEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteCurrency = (
    values: string[],
    formikProps?: FormikHelpers<currency.CurrencyEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteCurrency.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<currency.CurrencyEntity>>(
            "*[]currency.CurrencyEntity",
            (data) => fnDeleteCurrencyUpdater(data, values)
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
    currencysQuery,
    currencysExportQuery,
    currencyByUniqueIdQuery,
    mutationPostCurrency,
    submitPostCurrency,
    mutationPatchCurrency,
    submitPatchCurrency,
    mutationPatchCurrencys,
    submitPatchCurrencys,
    mutationDeleteCurrency,
    submitDeleteCurrency,
  };
}
