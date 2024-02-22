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

export function usePatchWorkspace({
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

  const fn = (entity: any) => Q().patchWorkspace(entity);

  const mutation = useMutation<
    IResponse<workspaces.WorkspaceEntity>,
    IResponse<workspaces.WorkspaceEntity>,
    Partial<workspaces.WorkspaceEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.WorkspaceEntity> | undefined,
    item: IResponse<workspaces.WorkspaceEntity>
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
    //       WorkspaceActions.isWorkspaceEntityEqual(t, item.data)
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
    values: Partial<workspaces.WorkspaceEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.WorkspaceEntity>>
  ): Promise<IResponse<workspaces.WorkspaceEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceEntity>) {
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
