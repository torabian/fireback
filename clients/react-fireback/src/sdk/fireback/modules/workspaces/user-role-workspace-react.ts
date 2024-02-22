// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: workspaces
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { UserRoleWorkspaceActions } from "./user-role-workspace-actions";
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
/**
 * Gives you formik forms, all mutations, submit actions, and error handling,
 * and provides internal store for all changes happens through this
 * for modules
 */
export function useWorkspaces(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? UserRoleWorkspaceActions.fnExec(execFn(options))
    : UserRoleWorkspaceActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const userRoleWorkspacesQuery = useQuery(
    ["*[]workspaces.UserRoleWorkspaceEntity", options],
    () => Q().getUserRoleWorkspaces(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  return { queryClient, userRoleWorkspacesQuery };
}
