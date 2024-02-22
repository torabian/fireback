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
import { WidgetAreaActions } from "./widget-area-actions";
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
    ? WidgetAreaActions.fnExec(execFn(options))
    : WidgetAreaActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const widgetAreasQuery = useQuery(
    ["*[]widget.WidgetAreaEntity", options],
    () => Q().getWidgetAreas(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const widgetAreasExportQuery = useQuery(
    ["*[]widget.WidgetAreaEntity", options],
    () => Q().getWidgetAreasExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const widgetAreaByUniqueIdQuery = useQuery(
    ["*widget.WidgetAreaEntity", options],
    (uniqueId: string) => Q().getWidgetAreaByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post widgetArea

  const mutationPostWidgetArea = useMutation<
    IResponse<widget.WidgetAreaEntity>,
    IResponse<widget.WidgetAreaEntity>,
    widget.WidgetAreaEntity
  >((entity) => {
    return Q().postWidgetArea(entity);
  });

  // Only entities are having a store in front-end

  const fnPostWidgetAreaUpdater = (
    data: IResponseList<widget.WidgetAreaEntity> | undefined,
    item: IResponse<widget.WidgetAreaEntity>
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
          WidgetAreaActions.isWidgetAreaEntityEqual(t, item.data)
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

  const submitPostWidgetArea = (
    values: widget.WidgetAreaEntity,
    formikProps?: FormikHelpers<widget.WidgetAreaEntity>
  ): Promise<IResponse<widget.WidgetAreaEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostWidgetArea.mutate(values, {
        onSuccess(response: IResponse<widget.WidgetAreaEntity>) {
          queryClient.setQueriesData<IResponseList<widget.WidgetAreaEntity>>(
            "*[]widget.WidgetAreaEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: widget.WidgetAreaEntity) => {
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

  // patch widgetArea

  const mutationPatchWidgetArea = useMutation<
    IResponse<widget.WidgetAreaEntity>,
    IResponse<widget.WidgetAreaEntity>,
    widget.WidgetAreaEntity
  >((entity) => {
    return Q().patchWidgetArea(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchWidgetAreaUpdater = (
    data: IResponseList<widget.WidgetAreaEntity> | undefined,
    item: IResponse<widget.WidgetAreaEntity>
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
          WidgetAreaActions.isWidgetAreaEntityEqual(t, item.data)
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

  const submitPatchWidgetArea = (
    values: widget.WidgetAreaEntity,
    formikProps?: FormikHelpers<widget.WidgetAreaEntity>
  ): Promise<IResponse<widget.WidgetAreaEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWidgetArea.mutate(values, {
        onSuccess(response: IResponse<widget.WidgetAreaEntity>) {
          queryClient.setQueriesData<IResponseList<widget.WidgetAreaEntity>>(
            "*[]widget.WidgetAreaEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: widget.WidgetAreaEntity) => {
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

  // patch widgetAreas

  const mutationPatchWidgetAreas = useMutation<
    IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>,
    IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>,
    core.BulkRecordRequest<widget.WidgetAreaEntity>
  >((entity) => {
    return Q().patchWidgetAreas(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchWidgetAreas = (
    values: core.BulkRecordRequest<widget.WidgetAreaEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<widget.WidgetAreaEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchWidgetAreas.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<widget.WidgetAreaEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<widget.WidgetAreaEntity>>
          >(
            "*[]core.BulkRecordRequest[widget.WidgetAreaEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<widget.WidgetAreaEntity>) => {
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
  const mutationDeleteWidgetArea = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteWidgetArea();
  });

  const fnDeleteWidgetAreaUpdater = (
    data: IResponseList<widget.WidgetAreaEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = WidgetAreaActions.getWidgetAreaEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteWidgetArea = (
    values: string[],
    formikProps?: FormikHelpers<widget.WidgetAreaEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteWidgetArea.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<widget.WidgetAreaEntity>>(
            "*[]widget.WidgetAreaEntity",
            (data) => fnDeleteWidgetAreaUpdater(data, values)
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
    widgetAreasQuery,
    widgetAreasExportQuery,
    widgetAreaByUniqueIdQuery,
    mutationPostWidgetArea,
    submitPostWidgetArea,
    mutationPatchWidgetArea,
    submitPatchWidgetArea,
    mutationPatchWidgetAreas,
    submitPatchWidgetAreas,
    mutationDeleteWidgetArea,
    submitDeleteWidgetArea,
  };
}
