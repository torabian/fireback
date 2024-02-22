// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { KeyboardShortcutActions } from "./keyboard-shortcut-actions";
import * as keyboardActions from "./index";
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

export function usePatchKeyboardShortcuts({
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
    ? KeyboardShortcutActions.fnExec(execFnOverride(options))
    : execFn
    ? KeyboardShortcutActions.fnExec(execFn(options))
    : KeyboardShortcutActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchKeyboardShortcuts(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>,
    IResponse<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>,
    Partial<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
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
    values: Partial<
      core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
    >,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
          >
        ) {
          queryClient.setQueriesData(
            "*keyboardActions.KeyboardShortcutEntity",
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
