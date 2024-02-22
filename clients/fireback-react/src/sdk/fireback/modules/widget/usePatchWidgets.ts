// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { WidgetActions } from "./widget-actions";
import * as widget from "./index";
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

export function usePatchWidgets({
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
    ? WidgetActions.fnExec(execFnOverride(options))
    : execFn
    ? WidgetActions.fnExec(execFn(options))
    : WidgetActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchWidgets(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<widget.WidgetEntity>>,
    IResponse<core.BulkRecordRequest<widget.WidgetEntity>>,
    Partial<core.BulkRecordRequest<widget.WidgetEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<core.BulkRecordRequest<widget.WidgetEntity>>,
    response: IResponse<
      core.BulkRecordRequest<core.BulkRecordRequest<widget.WidgetEntity>>
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
    values: Partial<core.BulkRecordRequest<widget.WidgetEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<widget.WidgetEntity>>
    >
  ): Promise<IResponse<core.BulkRecordRequest<widget.WidgetEntity>>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<widget.WidgetEntity>>
        ) {
          queryClient.setQueriesData("*widget.WidgetEntity", (data) =>
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
