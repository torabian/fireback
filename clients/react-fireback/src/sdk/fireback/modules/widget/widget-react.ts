// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: widget
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { WidgetActions } from "./widget-actions";
import * as widget from "./index";
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
export function useWidget(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? WidgetActions.fnExec(execFn(options))
    : WidgetActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const widgetsQuery = useQuery(
    ["*[]widget.WidgetEntity", options],
    () => Q().getWidgets(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const widgetsExportQuery = useQuery(
    ["*[]widget.WidgetEntity", options],
    () => Q().getWidgetsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const widgetByUniqueIdQuery = useQuery(
    ["*widget.WidgetEntity", options],
    (uniqueId: string) => Q().getWidgetByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post widget

  const mutationPostWidget = useMutation<
    IResponse<widget.WidgetEntity>,
    IResponse<widget.WidgetEntity>,
    widget.WidgetEntity
  >((entity) => {
    return Q().postWidget(entity);
  });

  // Only entities are having a store in front-end

  const fnPostWidgetUpdater = (
    data: IResponseList<widget.WidgetEntity> | undefined,
    item: IResponse<widget.WidgetEntity>
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
          WidgetActions.isWidgetEntityEqual(t, item.data)
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

  const submitPostWidget = (
    values: widget.WidgetEntity,
    formikProps?: FormikHelpers<widget.WidgetEntity>
  ): Promise<IResponse<widget.WidgetEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostWidget.mutate(values, {
        onSuccess(response: IResponse<widget.WidgetEntity>) {
          queryClient.setQueriesData<IResponseList<widget.WidgetEntity>>(
            "*[]widget.WidgetEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: widget.WidgetEntity) => {
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

  // patch widget

  const mutationPatchWidget = useMutation<
    IResponse<widget.WidgetEntity>,
    IResponse<widget.WidgetEntity>,
    widget.WidgetEntity
  >((entity) => {
    return Q().patchWidget(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchWidgetUpdater = (
    data: IResponseList<widget.WidgetEntity> | undefined,
    item: IResponse<widget.WidgetEntity>
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
          WidgetActions.isWidgetEntityEqual(t, item.data)
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

  const submitPatchWidget = (
    values: widget.WidgetEntity,
    formikProps?: FormikHelpers<widget.WidgetEntity>
  ): Promise<IResponse<widget.WidgetEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWidget.mutate(values, {
        onSuccess(response: IResponse<widget.WidgetEntity>) {
          queryClient.setQueriesData<IResponseList<widget.WidgetEntity>>(
            "*[]widget.WidgetEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: widget.WidgetEntity) => {
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

  // patch widgets

  const mutationPatchWidgets = useMutation<
    IResponse<core.BulkRecordRequest<widget.WidgetEntity>>,
    IResponse<core.BulkRecordRequest<widget.WidgetEntity>>,
    core.BulkRecordRequest<widget.WidgetEntity>
  >((entity) => {
    return Q().patchWidgets(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchWidgets = (
    values: core.BulkRecordRequest<widget.WidgetEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<widget.WidgetEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<widget.WidgetEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWidgets.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<widget.WidgetEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<widget.WidgetEntity>>
          >("*[]core.BulkRecordRequest[widget.WidgetEntity]", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: core.BulkRecordRequest<widget.WidgetEntity>) => {
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

  // Deleting an entity
  const mutationDeleteWidget = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteWidget();
  });

  const fnDeleteWidgetUpdater = (
    data: IResponseList<widget.WidgetEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = WidgetActions.getWidgetEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteWidget = (
    values: string[],
    formikProps?: FormikHelpers<widget.WidgetEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteWidget.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<widget.WidgetEntity>>(
            "*[]widget.WidgetEntity",
            (data) => fnDeleteWidgetUpdater(data, values)
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
    widgetsQuery,
    widgetsExportQuery,
    widgetByUniqueIdQuery,
    mutationPostWidget,
    submitPostWidget,
    mutationPatchWidget,
    submitPatchWidget,
    mutationPatchWidgets,
    submitPatchWidgets,
    mutationDeleteWidget,
    submitDeleteWidget,
  };
}
