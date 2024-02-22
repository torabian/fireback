// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { ActivationKeyActions } from "./activation-key-actions";
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

export function usePatchActivationKeys({
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
    ? ActivationKeyActions.fnExec(execFnOverride(options))
    : execFn
    ? ActivationKeyActions.fnExec(execFn(options))
    : ActivationKeyActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchActivationKeys(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<licenses.ActivationKeyEntity>>,
    IResponse<core.BulkRecordRequest<licenses.ActivationKeyEntity>>,
    Partial<core.BulkRecordRequest<licenses.ActivationKeyEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<licenses.ActivationKeyEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<licenses.ActivationKeyEntity>
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
    values: Partial<core.BulkRecordRequest<licenses.ActivationKeyEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<licenses.ActivationKeyEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<licenses.ActivationKeyEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<licenses.ActivationKeyEntity>
          >
        ) {
          queryClient.setQueriesData("*licenses.ActivationKeyEntity", (data) =>
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
