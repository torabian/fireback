// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
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
import { RemoteQueryContext } from "../../core/react-tools";

export function usePostPriceTag({
  queryClient,
  query,
  execFnOverride,
}: {
  queryClient: QueryClient;
  query?: any;
  execFnOverride?: any;
}) {
  query = query || {};

  const { options, execFn } = useContext(RemoteQueryContext);

  const fnx = execFnOverride
    ? PriceTagActions.fnExec(execFnOverride(options))
    : execFn
    ? PriceTagActions.fnExec(execFn(options))
    : PriceTagActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postPriceTag(entity);

  const mutation = useMutation<
    IResponse<currency.PriceTagEntity>,
    IResponse<currency.PriceTagEntity>,
    Partial<currency.PriceTagEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<currency.PriceTagEntity> | undefined,
    item: IResponse<currency.PriceTagEntity>
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    // To me it seems this is not a good or any correct strategy to update the store.
    // When we are posting, we want to add it there, that's it. Not updating it.
    // We have patch, but also posting with ID is possible.

    // if (data?.data?.items && item.data) {
    //   data.data.items = data.data.items.map((t) => {
    //     if (
    //       item.data !== undefined &&
    //       PriceTagActions.isPriceTagEntityEqual(t, item.data)
    //     ) {
    //       return item.data;
    //     }

    //     return t;
    //   });
    // } else if (data?.data && item.data) {
    //   data.data.items = [item.data, ...(data?.data?.items || [])];
    // }

    data.data.items = [item.data, ...(data?.data?.items || [])];

    return data;
  };

  const submit = (
    values: Partial<currency.PriceTagEntity>,
    formikProps?: FormikHelpers<Partial<currency.PriceTagEntity>>
  ): Promise<IResponse<currency.PriceTagEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<currency.PriceTagEntity>) {
          queryClient.setQueryData<IResponseList<currency.PriceTagEntity>>(
            "*currency.PriceTagEntity",
            (data) => fnUpdater(data, response)
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

  return { mutation, submit, fnUpdater };
}
