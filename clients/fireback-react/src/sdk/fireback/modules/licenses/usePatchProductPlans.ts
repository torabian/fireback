// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
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
import { RemoteQueryContext } from "../../core/react-tools";

export function usePatchProductPlans({
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
    ? ProductPlanActions.fnExec(execFnOverride(options))
    : execFn
    ? ProductPlanActions.fnExec(execFn(options))
    : ProductPlanActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchProductPlans(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<licenses.ProductPlanEntity>>,
    IResponse<core.BulkRecordRequest<licenses.ProductPlanEntity>>,
    Partial<core.BulkRecordRequest<licenses.ProductPlanEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<core.BulkRecordRequest<licenses.ProductPlanEntity>>,
    response: IResponse<
      core.BulkRecordRequest<core.BulkRecordRequest<licenses.ProductPlanEntity>>
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
    values: Partial<core.BulkRecordRequest<licenses.ProductPlanEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<licenses.ProductPlanEntity>>
    >
  ): Promise<IResponse<core.BulkRecordRequest<licenses.ProductPlanEntity>>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<licenses.ProductPlanEntity>
          >
        ) {
          queryClient.setQueriesData("*licenses.ProductPlanEntity", (data) =>
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
