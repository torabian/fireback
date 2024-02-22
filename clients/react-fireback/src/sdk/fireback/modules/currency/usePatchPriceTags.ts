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

export function usePatchPriceTags({
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

  const fn = (entity: any) => Q().patchPriceTags(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>,
    IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>,
    Partial<core.BulkRecordRequest<currency.PriceTagEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<core.BulkRecordRequest<currency.PriceTagEntity>>,
    response: IResponse<
      core.BulkRecordRequest<core.BulkRecordRequest<currency.PriceTagEntity>>
    >
  ) => {
    if (!data || !data.data) {
      return data;
    }

    const records = response?.data?.records || [];

    if (data.data.items && records.length > 0) {
      data.data.items = data.data.items.map((m) => {
        const editedVersion = records.find((l) => l.uniqueId === m.uniqueId);
        if (editedVersion) {
          return {
            ...m,
            ...editedVersion,
          };
        }
        return m;
      });
    }

    return data;
  };

  const submit = (
    values: Partial<core.BulkRecordRequest<currency.PriceTagEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<currency.PriceTagEntity>>
    >
  ): Promise<IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<currency.PriceTagEntity>>
        ) {
          queryClient.setQueriesData("*currency.PriceTagEntity", (data) =>
            fnUpdater(data, response)
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
