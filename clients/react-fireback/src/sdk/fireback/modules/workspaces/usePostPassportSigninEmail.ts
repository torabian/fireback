// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { PassportActions } from "./passport-actions";
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

export function usePostPassportSigninEmail({
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
    ? PassportActions.fnExec(execFnOverride(options))
    : execFn
    ? PassportActions.fnExec(execFn(options))
    : PassportActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postPassportSigninEmail(entity);

  const mutation = useMutation<
    IResponse<workspaces.UserSessionDto>,
    IResponse<workspaces.UserSessionDto>,
    Partial<workspaces.EmailAccountSigninDto>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater: any = () => {};

  const submit = (
    values: Partial<workspaces.EmailAccountSigninDto>,
    formikProps?: FormikHelpers<Partial<workspaces.UserSessionDto>>
  ): Promise<IResponse<workspaces.UserSessionDto>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.UserSessionDto>) {
          queryClient.setQueryData<IResponseList<workspaces.UserSessionDto>>(
            "*workspaces.UserSessionDto",
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
