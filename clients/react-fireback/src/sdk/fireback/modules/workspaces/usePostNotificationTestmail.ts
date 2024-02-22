// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { NotificationConfigActions } from "./notification-config-actions";
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

export function usePostNotificationTestmail({
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
    ? NotificationConfigActions.fnExec(execFnOverride(options))
    : execFn
    ? NotificationConfigActions.fnExec(execFn(options))
    : NotificationConfigActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postNotificationTestmail(entity);

  const mutation = useMutation<
    IResponse<core.OkayResponse>,
    IResponse<core.OkayResponse>,
    Partial<workspaces.TestMailDto>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = () => {};

  const submit = (
    values: Partial<workspaces.TestMailDto>,
    formikProps?: FormikHelpers<Partial<core.OkayResponse>>
  ): Promise<IResponse<core.OkayResponse>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<core.OkayResponse>) {
          queryClient.setQueryData<IResponseList<core.OkayResponse>>(
            "*core.OkayResponse",
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
