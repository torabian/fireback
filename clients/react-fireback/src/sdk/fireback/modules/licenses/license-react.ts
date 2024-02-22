// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: licenses
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { LicenseActions } from "./license-actions";
import * as licenses from "./index";
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
export function useLicenses(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? LicenseActions.fnExec(execFn(options))
    : LicenseActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const licensesQuery = useQuery(
    ["*[]licenses.LicenseEntity", options],
    () => Q().getLicenses(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const licensesExportQuery = useQuery(
    ["*[]licenses.LicenseEntity", options],
    () => Q().getLicensesExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const licenseByUniqueIdQuery = useQuery(
    ["*licenses.LicenseEntity", options],
    (uniqueId: string) => Q().getLicenseByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post license

  const mutationPostLicense = useMutation<
    IResponse<licenses.LicenseEntity>,
    IResponse<licenses.LicenseEntity>,
    licenses.LicenseEntity
  >((entity) => {
    return Q().postLicense(entity);
  });

  // Only entities are having a store in front-end

  const fnPostLicenseUpdater = (
    data: IResponseList<licenses.LicenseEntity> | undefined,
    item: IResponse<licenses.LicenseEntity>
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
          LicenseActions.isLicenseEntityEqual(t, item.data)
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

  const submitPostLicense = (
    values: licenses.LicenseEntity,
    formikProps?: FormikHelpers<licenses.LicenseEntity>
  ): Promise<IResponse<licenses.LicenseEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostLicense.mutate(values, {
        onSuccess(response: IResponse<licenses.LicenseEntity>) {
          queryClient.setQueriesData<IResponseList<licenses.LicenseEntity>>(
            "*[]licenses.LicenseEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: licenses.LicenseEntity) => {
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

  // patch license

  const mutationPatchLicense = useMutation<
    IResponse<licenses.LicenseEntity>,
    IResponse<licenses.LicenseEntity>,
    licenses.LicenseEntity
  >((entity) => {
    return Q().patchLicense(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchLicenseUpdater = (
    data: IResponseList<licenses.LicenseEntity> | undefined,
    item: IResponse<licenses.LicenseEntity>
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
          LicenseActions.isLicenseEntityEqual(t, item.data)
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

  const submitPatchLicense = (
    values: licenses.LicenseEntity,
    formikProps?: FormikHelpers<licenses.LicenseEntity>
  ): Promise<IResponse<licenses.LicenseEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchLicense.mutate(values, {
        onSuccess(response: IResponse<licenses.LicenseEntity>) {
          queryClient.setQueriesData<IResponseList<licenses.LicenseEntity>>(
            "*[]licenses.LicenseEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: licenses.LicenseEntity) => {
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

  // patch licenses

  const mutationPatchLicenses = useMutation<
    IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>,
    IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>,
    core.BulkRecordRequest<licenses.LicenseEntity>
  >((entity) => {
    return Q().patchLicenses(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchLicenses = (
    values: core.BulkRecordRequest<licenses.LicenseEntity>,
    formikProps?: FormikHelpers<core.BulkRecordRequest<licenses.LicenseEntity>>
  ): Promise<IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchLicenses.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<licenses.LicenseEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<licenses.LicenseEntity>>
          >(
            "*[]core.BulkRecordRequest[licenses.LicenseEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<licenses.LicenseEntity>) => {
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
  const mutationDeleteLicense = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteLicense();
  });

  const fnDeleteLicenseUpdater = (
    data: IResponseList<licenses.LicenseEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = LicenseActions.getLicenseEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteLicense = (
    values: string[],
    formikProps?: FormikHelpers<licenses.LicenseEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteLicense.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<licenses.LicenseEntity>>(
            "*[]licenses.LicenseEntity",
            (data) => fnDeleteLicenseUpdater(data, values)
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

  // post license/from-plan/:uniqueId

  const mutationPostLicenseFromPlanByUniqueId = useMutation<
    IResponse<licenses.LicenseEntity>,
    IResponse<licenses.LicenseEntity>,
    licenses.LicenseFromPlanIdDto
  >(
    (
      uniqueId: string,

      entity
    ) => {
      return Q().postLicenseFromPlanByUniqueId(
        uniqueId,

        entity
      );
    }
  );

  // Only entities are having a store in front-end

  const submitPostLicenseFromPlanByUniqueId = (
    values: licenses.LicenseFromPlanIdDto,
    formikProps?: FormikHelpers<licenses.LicenseFromPlanIdDto>
  ): Promise<IResponse<licenses.LicenseEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostLicenseFromPlanByUniqueId.mutate(values, {
        onSuccess(response: IResponse<licenses.LicenseEntity>) {
          queryClient.setQueriesData<
            IResponseList<licenses.LicenseFromPlanIdDto>
          >("*[]licenses.LicenseFromPlanIdDto", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: licenses.LicenseFromPlanIdDto) => {
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

  return {
    queryClient,
    licensesQuery,
    licensesExportQuery,
    licenseByUniqueIdQuery,
    mutationPostLicense,
    submitPostLicense,
    mutationPatchLicense,
    submitPatchLicense,
    mutationPatchLicenses,
    submitPatchLicenses,
    mutationDeleteLicense,
    submitDeleteLicense,
    mutationPostLicenseFromPlanByUniqueId,
    submitPostLicenseFromPlanByUniqueId,
  };
}
