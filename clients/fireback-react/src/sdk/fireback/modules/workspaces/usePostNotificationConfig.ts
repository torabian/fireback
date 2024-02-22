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

export function usePostNotificationConfig({
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

  const fn = (entity: any) => Q().postNotificationConfig(entity);

  const mutation = useMutation<
    IResponse<workspaces.NotificationConfigEntity>,
    IResponse<workspaces.NotificationConfigEntity>,
    Partial<workspaces.NotificationConfigEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.NotificationConfigEntity> | undefined,
    item: IResponse<workspaces.NotificationConfigEntity>
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
    //       NotificationConfigActions.isNotificationConfigEntityEqual(t, item.data)
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
    values: Partial<workspaces.NotificationConfigEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.NotificationConfigEntity>>
  ): Promise<IResponse<workspaces.NotificationConfigEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.NotificationConfigEntity>) {
          queryClient.setQueryData<
            IResponseList<workspaces.NotificationConfigEntity>
          >("*workspaces.NotificationConfigEntity", (data) =>
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
