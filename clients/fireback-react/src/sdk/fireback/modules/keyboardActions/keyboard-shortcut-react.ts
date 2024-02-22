// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: keyboardActions
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
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
/**
 * Gives you formik forms, all mutations, submit actions, and error handling,
 * and provides internal store for all changes happens through this
 * for modules
 */
export function useKeyboardActions(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? KeyboardShortcutActions.fnExec(execFn(options))
    : KeyboardShortcutActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const keyboardShortcutsQuery = useQuery(
    ["*[]keyboardActions.KeyboardShortcutEntity", options],
    () => Q().getKeyboardShortcuts(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const keyboardShortcutsExportQuery = useQuery(
    ["*[]keyboardActions.KeyboardShortcutEntity", options],
    () => Q().getKeyboardShortcutsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const keyboardShortcutByUniqueIdQuery = useQuery(
    ["*keyboardActions.KeyboardShortcutEntity", options],
    (uniqueId: string) => Q().getKeyboardShortcutByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post keyboardShortcut

  const mutationPostKeyboardShortcut = useMutation<
    IResponse<keyboardActions.KeyboardShortcutEntity>,
    IResponse<keyboardActions.KeyboardShortcutEntity>,
    keyboardActions.KeyboardShortcutEntity
  >((entity) => {
    return Q().postKeyboardShortcut(entity);
  });

  // Only entities are having a store in front-end

  const fnPostKeyboardShortcutUpdater = (
    data: IResponseList<keyboardActions.KeyboardShortcutEntity> | undefined,
    item: IResponse<keyboardActions.KeyboardShortcutEntity>
  ) => {
    return [];

    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items && item.data) {
      data.data.items = data.data.items.map((t) => {
        if (
          item.data !== undefined &&
          KeyboardShortcutActions.isKeyboardShortcutEntityEqual(t, item.data)
        ) {
          return item.data;
        }

        return t;
      });
    } else if (data?.data && item.data) {
      data.data.items = [item.data, ...(data?.data?.items || [])];
    }

    return data;
  };

  const submitPostKeyboardShortcut = (
    values: keyboardActions.KeyboardShortcutEntity,
    formikProps?: FormikHelpers<keyboardActions.KeyboardShortcutEntity>
  ): Promise<IResponse<keyboardActions.KeyboardShortcutEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostKeyboardShortcut.mutate(values, {
        onSuccess(response: IResponse<keyboardActions.KeyboardShortcutEntity>) {
          queryClient.setQueriesData<
            IResponseList<keyboardActions.KeyboardShortcutEntity>
          >("*[]keyboardActions.KeyboardShortcutEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: keyboardActions.KeyboardShortcutEntity) => {
                if (item.uniqueId === response.data?.uniqueId) {
                  return response.data;
                }

                return item;
              }
            );

            return data;
          });

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  // patch keyboardShortcut

  const mutationPatchKeyboardShortcut = useMutation<
    IResponse<keyboardActions.KeyboardShortcutEntity>,
    IResponse<keyboardActions.KeyboardShortcutEntity>,
    keyboardActions.KeyboardShortcutEntity
  >((entity) => {
    return Q().patchKeyboardShortcut(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchKeyboardShortcutUpdater = (
    data: IResponseList<keyboardActions.KeyboardShortcutEntity> | undefined,
    item: IResponse<keyboardActions.KeyboardShortcutEntity>
  ) => {
    return [];

    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items && item.data) {
      data.data.items = data.data.items.map((t) => {
        if (
          item.data !== undefined &&
          KeyboardShortcutActions.isKeyboardShortcutEntityEqual(t, item.data)
        ) {
          return item.data;
        }

        return t;
      });
    } else if (data?.data && item.data) {
      data.data.items = [item.data, ...(data?.data?.items || [])];
    }

    return data;
  };

  const submitPatchKeyboardShortcut = (
    values: keyboardActions.KeyboardShortcutEntity,
    formikProps?: FormikHelpers<keyboardActions.KeyboardShortcutEntity>
  ): Promise<IResponse<keyboardActions.KeyboardShortcutEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchKeyboardShortcut.mutate(values, {
        onSuccess(response: IResponse<keyboardActions.KeyboardShortcutEntity>) {
          queryClient.setQueriesData<
            IResponseList<keyboardActions.KeyboardShortcutEntity>
          >("*[]keyboardActions.KeyboardShortcutEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: keyboardActions.KeyboardShortcutEntity) => {
                if (item.uniqueId === response.data?.uniqueId) {
                  return response.data;
                }

                return item;
              }
            );

            return data;
          });

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  // patch keyboardShortcuts

  const mutationPatchKeyboardShortcuts = useMutation<
    IResponse<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>,
    IResponse<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>,
    core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
  >((entity) => {
    return Q().patchKeyboardShortcuts(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchKeyboardShortcuts = (
    values: core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchKeyboardShortcuts.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
            >
          >(
            "*[]core.BulkRecordRequest[keyboardActions.KeyboardShortcutEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
                ) => {
                  if (item.uniqueId === response.data?.uniqueId) {
                    return response.data;
                  }

                  return item;
                }
              );

              return data;
            }
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

  // Deleting an entity
  const mutationDeleteKeyboardShortcut = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteKeyboardShortcut();
  });

  const fnDeleteKeyboardShortcutUpdater = (
    data: IResponseList<keyboardActions.KeyboardShortcutEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key =
          KeyboardShortcutActions.getKeyboardShortcutEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteKeyboardShortcut = (
    values: string[],
    formikProps?: FormikHelpers<keyboardActions.KeyboardShortcutEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteKeyboardShortcut.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<keyboardActions.KeyboardShortcutEntity>
          >("*[]keyboardActions.KeyboardShortcutEntity", (data) =>
            fnDeleteKeyboardShortcutUpdater(data, values)
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

  return {
    queryClient,
    keyboardShortcutsQuery,
    keyboardShortcutsExportQuery,
    keyboardShortcutByUniqueIdQuery,
    mutationPostKeyboardShortcut,
    submitPostKeyboardShortcut,
    mutationPatchKeyboardShortcut,
    submitPatchKeyboardShortcut,
    mutationPatchKeyboardShortcuts,
    submitPatchKeyboardShortcuts,
    mutationDeleteKeyboardShortcut,
    submitDeleteKeyboardShortcut,
  };
}
