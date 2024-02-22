// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { WorkspaceActions } from "./workspace-actions";
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

export function usePatchWorkspaces({
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
    ? WorkspaceActions.fnExec(execFnOverride(options))
    : execFn
    ? WorkspaceActions.fnExec(execFn(options))
    : WorkspaceActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchWorkspaces(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceEntity>>,
    Partial<core.BulkRecordRequest<workspaces.WorkspaceEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<core.BulkRecordRequest<workspaces.WorkspaceEntity>>,
    response: IResponse<
      core.BulkRecordRequest<core.BulkRecordRequest<workspaces.WorkspaceEntity>>
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
    values: Partial<core.BulkRecordRequest<workspaces.WorkspaceEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<workspaces.WorkspaceEntity>>
    >
  ): Promise<IResponse<core.BulkRecordRequest<workspaces.WorkspaceEntity>>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.WorkspaceEntity>
          >
        ) {
          queryClient.setQueriesData("*workspaces.WorkspaceEntity", (data) =>
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
