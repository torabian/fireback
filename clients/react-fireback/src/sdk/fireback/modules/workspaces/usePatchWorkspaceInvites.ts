// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { WorkspaceInviteActions } from "./workspace-invite-actions";
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

export function usePatchWorkspaceInvites({
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
    ? WorkspaceInviteActions.fnExec(execFnOverride(options))
    : execFn
    ? WorkspaceInviteActions.fnExec(execFn(options))
    : WorkspaceInviteActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchWorkspaceInvites(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>,
    Partial<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
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
    values: Partial<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.WorkspaceInviteEntity>
          >
        ) {
          queryClient.setQueriesData(
            "*workspaces.WorkspaceInviteEntity",
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
