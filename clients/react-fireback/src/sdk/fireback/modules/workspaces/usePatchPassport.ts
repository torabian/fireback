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

export function usePatchPassport({
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

  const fn = (entity: any) => Q().patchPassport(entity);

  const mutation = useMutation<
    IResponse<workspaces.PassportEntity>,
    IResponse<workspaces.PassportEntity>,
    Partial<workspaces.PassportEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.PassportEntity> | undefined,
    item: IResponse<workspaces.PassportEntity>
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    // To me it seems this is not a good or any correct strategy to update the store.
    // When we are posting, we want to add it there, that's it. Not updating it.
    // We have patch, but also posting with ID is possible.

    // if (data?.data?.items && item.data) {
    //   data.data.items = data.data.items.map((t) => {
    //     if (
    //       item.data !== undefined &&
    //       PassportActions.isPassportEntityEqual(t, item.data)
    //     ) {
    //       return item.data;
    //     }

    //     return t;
    //   });
    // } else if (data?.data && item.data) {
    //   data.data.items = [item.data, ...(data?.data?.items || [])];
    // }

    data.data.items = [item.data, ...(data?.data?.items || [])];

    return data;
  };

  const submit = (
    values: Partial<workspaces.PassportEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.PassportEntity>>
  ): Promise<IResponse<workspaces.PassportEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.PassportEntity>) {
          queryClient.setQueriesData("*workspaces.PassportEntity", (data) =>
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
