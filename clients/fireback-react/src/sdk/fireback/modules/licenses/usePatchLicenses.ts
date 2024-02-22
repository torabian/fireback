// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { LicenseActions } from "./license-actions";
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

export function usePatchLicenses({
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
    ? LicenseActions.fnExec(execFnOverride(options))
    : execFn
    ? LicenseActions.fnExec(execFn(options))
    : LicenseActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchLicenses(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>,
    IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>,
    Partial<core.BulkRecordRequest<licenses.LicenseEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<core.BulkRecordRequest<licenses.LicenseEntity>>,
    response: IResponse<
      core.BulkRecordRequest<core.BulkRecordRequest<licenses.LicenseEntity>>
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
    values: Partial<core.BulkRecordRequest<licenses.LicenseEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<licenses.LicenseEntity>>
    >
  ): Promise<IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>
        ) {
          queryClient.setQueriesData("*licenses.LicenseEntity", (data) =>
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
