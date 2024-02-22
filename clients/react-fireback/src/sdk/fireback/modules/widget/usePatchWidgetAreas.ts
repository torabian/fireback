// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { WidgetAreaActions } from "./widget-area-actions";
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

export function usePatchWidgetAreas({
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
    ? WidgetAreaActions.fnExec(execFnOverride(options))
    : execFn
    ? WidgetAreaActions.fnExec(execFn(options))
    : WidgetAreaActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchWidgetAreas(entity);

  const mutation = useMutation<
    IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>,
    IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>,
    Partial<core.BulkRecordRequest<widget.WidgetAreaEntity>>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = (
    data: PossibleStoreData<core.BulkRecordRequest<widget.WidgetAreaEntity>>,
    response: IResponse<
      core.BulkRecordRequest<core.BulkRecordRequest<widget.WidgetAreaEntity>>
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
    values: Partial<core.BulkRecordRequest<widget.WidgetAreaEntity>>,
    formikProps?: FormikHelpers<
      Partial<core.BulkRecordRequest<widget.WidgetAreaEntity>>
    >
  ): Promise<IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>
        ) {
          queryClient.setQueriesData("*widget.WidgetAreaEntity", (data) =>
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
