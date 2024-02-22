// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { EmailSenderActions } from "./email-sender-actions";
import * as workspaces from "./index";
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

export function usePatchEmailSenders({
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
    ? EmailSenderActions.fnExec(execFnOverride(options))
    : execFn
    ? EmailSenderActions.fnExec(execFn(options))
    : EmailSenderActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchEmailSenders(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.EmailSenderEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.EmailSenderEntity>>,
    Partial<core.BulkRecordRequest<workspaces.EmailSenderEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<workspaces.EmailSenderEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<workspaces.EmailSenderEntity>
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
    values: Partial<core.BulkRecordRequest<workspaces.EmailSenderEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<workspaces.EmailSenderEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.EmailSenderEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.EmailSenderEntity>
          >
        ) {
          queryClient.setQueriesData("*workspaces.EmailSenderEntity", (data) =>
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
