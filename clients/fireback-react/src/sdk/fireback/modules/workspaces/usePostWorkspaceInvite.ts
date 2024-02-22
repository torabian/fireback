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

export function usePostWorkspaceInvite({
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

  const fn = (entity: any) => Q().postWorkspaceInvite(entity);

  const mutation = useMutation<
    IResponse<workspaces.WorkspaceInviteEntity>,
    IResponse<workspaces.WorkspaceInviteEntity>,
    Partial<workspaces.WorkspaceInviteEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.WorkspaceInviteEntity> | undefined,
    item: IResponse<workspaces.WorkspaceInviteEntity>
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
    //       WorkspaceInviteActions.isWorkspaceInviteEntityEqual(t, item.data)
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
    values: Partial<workspaces.WorkspaceInviteEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.WorkspaceInviteEntity>>
  ): Promise<IResponse<workspaces.WorkspaceInviteEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.WorkspaceInviteEntity>) {
          queryClient.setQueryData<
            IResponseList<workspaces.WorkspaceInviteEntity>
          >("*workspaces.WorkspaceInviteEntity", (data) =>
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
