// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { BackupTableMetaActions } from "./backup-table-meta-actions";
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

export function usePatchBackupTableMetas({
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
    ? BackupTableMetaActions.fnExec(execFnOverride(options))
    : execFn
    ? BackupTableMetaActions.fnExec(execFn(options))
    : BackupTableMetaActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchBackupTableMetas(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>,
    Partial<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<
      core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
    >,
    response: IResponse<
      core.BulkRecordRequest<
        core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
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
    values: Partial<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<workspaces.BackupTableMetaEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
          >
        ) {
          queryClient.setQueriesData(
            "*workspaces.BackupTableMetaEntity",
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
