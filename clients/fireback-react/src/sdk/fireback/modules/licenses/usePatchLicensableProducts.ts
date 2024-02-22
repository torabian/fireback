// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
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
import { RemoteQueryContext } from "../../core/react-tools";

export function usePatchLicensableProducts({
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
    ? LicensableProductActions.fnExec(execFnOverride(options))
    : execFn
    ? LicensableProductActions.fnExec(execFn(options))
    : LicensableProductActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchLicensableProducts(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<licenses.LicensableProductEntity>>,
    IResponse<core.BulkRecordRequest<licenses.LicensableProductEntity>>,
    Partial<core.BulkRecordRequest<licenses.LicensableProductEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<licenses.LicensableProductEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<licenses.LicensableProductEntity>
      >
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
    values: Partial<core.BulkRecordRequest<licenses.LicensableProductEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<licenses.LicensableProductEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<licenses.LicensableProductEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<licenses.LicensableProductEntity>
          >
        ) {
          queryClient.setQueriesData(
            "*licenses.LicensableProductEntity",
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
