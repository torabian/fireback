// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { KeyboardShortcutActions } from "./keyboard-shortcut-actions";
import * as keyboardActions from "./index";
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

export function usePostKeyboardShortcut({
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
    ? KeyboardShortcutActions.fnExec(execFnOverride(options))
    : execFn
    ? KeyboardShortcutActions.fnExec(execFn(options))
    : KeyboardShortcutActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postKeyboardShortcut(entity);

  const mutation = useMutation<
    IResponse<keyboardActions.KeyboardShortcutEntity>,
    IResponse<keyboardActions.KeyboardShortcutEntity>,
    Partial<keyboardActions.KeyboardShortcutEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<keyboardActions.KeyboardShortcutEntity> | undefined,
    item: IResponse<keyboardActions.KeyboardShortcutEntity>
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
    //       KeyboardShortcutActions.isKeyboardShortcutEntityEqual(t, item.data)
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
    values: Partial<keyboardActions.KeyboardShortcutEntity>,
    formikProps?: FormikHelpers<Partial<keyboardActions.KeyboardShortcutEntity>>
  ): Promise<IResponse<keyboardActions.KeyboardShortcutEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<keyboardActions.KeyboardShortcutEntity>) {
          queryClient.setQueryData<
            IResponseList<keyboardActions.KeyboardShortcutEntity>
          >("*keyboardActions.KeyboardShortcutEntity", (data) =>
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
