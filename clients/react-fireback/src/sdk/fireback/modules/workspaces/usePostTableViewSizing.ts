// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { TableViewSizingActions } from "./table-view-sizing-actions";
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

export function usePostTableViewSizing({
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
    ? TableViewSizingActions.fnExec(execFnOverride(options))
    : execFn
    ? TableViewSizingActions.fnExec(execFn(options))
    : TableViewSizingActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postTableViewSizing(entity);

  const mutation = useMutation<
    IResponse<workspaces.TableViewSizingEntity>,
    IResponse<workspaces.TableViewSizingEntity>,
    Partial<workspaces.TableViewSizingEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.TableViewSizingEntity> | undefined,
    item: IResponse<workspaces.TableViewSizingEntity>
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
    //       TableViewSizingActions.isTableViewSizingEntityEqual(t, item.data)
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
    values: Partial<workspaces.TableViewSizingEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.TableViewSizingEntity>>
  ): Promise<IResponse<workspaces.TableViewSizingEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.TableViewSizingEntity>) {
          queryClient.setQueryData<
            IResponseList<workspaces.TableViewSizingEntity>
          >("*workspaces.TableViewSizingEntity", (data) =>
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
