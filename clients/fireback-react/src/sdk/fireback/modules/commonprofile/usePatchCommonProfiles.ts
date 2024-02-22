// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { CommonProfileActions } from "./common-profile-actions";
import * as commonprofile from "./index";
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

export function usePatchCommonProfiles({
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
    ? CommonProfileActions.fnExec(execFnOverride(options))
    : execFn
    ? CommonProfileActions.fnExec(execFn(options))
    : CommonProfileActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchCommonProfiles(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>,
    IResponse<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>,
    Partial<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<commonprofile.CommonProfileEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<commonprofile.CommonProfileEntity>
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
    values: Partial<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<commonprofile.CommonProfileEntity>
          >
        ) {
          queryClient.setQueriesData(
            "*commonprofile.CommonProfileEntity",
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
