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

export function useDeleteWidget({
  execFnOverride,
  queryClient,
}: {
  queryClient: QueryClient;
  execFnOverride?: any;
}) {
  const { options, execFn } = useContext(RemoteQueryContext);
  const fnx = execFnOverride
    ? WidgetActions.fnExec(execFnOverride(options))
    : execFn
    ? WidgetActions.fnExec(execFn(options))
    : WidgetActions.fn(options);
  const Q = () => fnx;

  const fn = (entity) => Q().deleteWidget(entity);

  const mutation = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = () => {};

  const submit = (
    values: core.DeleteRequest,
    formikProps: FormikHelpers<core.DeleteResponse>
  ): Promise<IResponse<core.DeleteResponse>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<core.DeleteResponse>) {
          queryClient.setQueryData<IResponseList<core.DeleteResponse>>(
            "*widget.WidgetEntity",
            (data) => fnUpdater(data, response)
          );

          resolve(response);
        },

        onError(error: any) {
          formikProps.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  return { mutation, submit, fnUpdater };
}
