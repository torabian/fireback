// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: commonprofile
 */

import { FormikHelpers } from "formik";
import React, { useCallback } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { CommonProfileActions } from "./common-profile-actions";
import * as commonprofile from "./index";
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
export function useCommonprofile(
  { options, query, execFn }: { options: RemoteRequestOption; query?: any },
  queryClient: QueryClient,
  execFn?: ExecApi
) {
  const caller = execFn
    ? CommonProfileActions.fnExec(execFn(options))
    : CommonProfileActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const commonProfilesQuery = useQuery(
    ["*[]commonprofile.CommonProfileEntity", options],
    () => Q().getCommonProfiles(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const commonProfilesExportQuery = useQuery(
    ["*[]commonprofile.CommonProfileEntity", options],
    () => Q().getCommonProfilesExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const commonProfileByUniqueIdQuery = useQuery(
    ["*commonprofile.CommonProfileEntity", options],
    (uniqueId: string) => Q().getCommonProfileByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post commonProfile

  const mutationPostCommonProfile = useMutation<
    IResponse<commonprofile.CommonProfileEntity>,
    IResponse<commonprofile.CommonProfileEntity>,
    commonprofile.CommonProfileEntity
  >((entity) => {
    return Q().postCommonProfile(entity);
  });

  // Only entities are having a store in front-end

  const fnPostCommonProfileUpdater = (
    data: IResponseList<commonprofile.CommonProfileEntity> | undefined,
    item: IResponse<commonprofile.CommonProfileEntity>
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
          CommonProfileActions.isCommonProfileEntityEqual(t, item.data)
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

  const submitPostCommonProfile = (
    values: commonprofile.CommonProfileEntity,
    formikProps?: FormikHelpers<commonprofile.CommonProfileEntity>
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostCommonProfile.mutate(values, {
        onSuccess(response: IResponse<commonprofile.CommonProfileEntity>) {
          queryClient.setQueriesData<
            IResponseList<commonprofile.CommonProfileEntity>
          >("*[]commonprofile.CommonProfileEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: commonprofile.CommonProfileEntity) => {
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

  // patch commonProfile

  const mutationPatchCommonProfile = useMutation<
    IResponse<commonprofile.CommonProfileEntity>,
    IResponse<commonprofile.CommonProfileEntity>,
    commonprofile.CommonProfileEntity
  >((entity) => {
    return Q().patchCommonProfile(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchCommonProfileUpdater = (
    data: IResponseList<commonprofile.CommonProfileEntity> | undefined,
    item: IResponse<commonprofile.CommonProfileEntity>
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
          CommonProfileActions.isCommonProfileEntityEqual(t, item.data)
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

  const submitPatchCommonProfile = (
    values: commonprofile.CommonProfileEntity,
    formikProps?: FormikHelpers<commonprofile.CommonProfileEntity>
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchCommonProfile.mutate(values, {
        onSuccess(response: IResponse<commonprofile.CommonProfileEntity>) {
          queryClient.setQueriesData<
            IResponseList<commonprofile.CommonProfileEntity>
          >("*[]commonprofile.CommonProfileEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: commonprofile.CommonProfileEntity) => {
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

  // patch commonProfileDistinct

  const mutationPatchCommonProfileDistinct = useMutation<
    IResponse<commonprofile.CommonProfileEntity>,
    IResponse<commonprofile.CommonProfileEntity>,
    commonprofile.CommonProfileEntity
  >((entity) => {
    return Q().patchCommonProfileDistinct(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchCommonProfileDistinctUpdater = (
    data: IResponseList<commonprofile.CommonProfileEntity> | undefined,
    item: IResponse<commonprofile.CommonProfileEntity>
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
          CommonProfileActions.isCommonProfileEntityEqual(t, item.data)
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

  const submitPatchCommonProfileDistinct = (
    values: commonprofile.CommonProfileEntity,
    formikProps?: FormikHelpers<commonprofile.CommonProfileEntity>
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchCommonProfileDistinct.mutate(values, {
        onSuccess(response: IResponse<commonprofile.CommonProfileEntity>) {
          queryClient.setQueriesData<
            IResponseList<commonprofile.CommonProfileEntity>
          >("*[]commonprofile.CommonProfileEntity", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: commonprofile.CommonProfileEntity) => {
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

  const commonProfileDistinctQuery = useQuery(
    ["*commonprofile.CommonProfileEntity", options],
    () => Q().getCommonProfileDistinct(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // patch commonProfiles

  const mutationPatchCommonProfiles = useMutation<
    IResponse<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>,
    IResponse<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>,
    core.BulkRecordRequest<commonprofile.CommonProfileEntity>
  >((entity) => {
    return Q().patchCommonProfiles(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchCommonProfiles = (
    values: core.BulkRecordRequest<commonprofile.CommonProfileEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<commonprofile.CommonProfileEntity>
    >
  ): Promise<
    IResponse<core.BulkRecordRequest<commonprofile.CommonProfileEntity>>
  > => {
    return new Promise((resolve, reject) => {
      mutationPatchCommonProfiles.mutate(values, {
        onSuccess(
          response: IResponse<
            core.BulkRecordRequest<commonprofile.CommonProfileEntity>
          >
        ) {
          queryClient.setQueriesData<
            IResponseList<
              core.BulkRecordRequest<commonprofile.CommonProfileEntity>
            >
          >(
            "*[]core.BulkRecordRequest[commonprofile.CommonProfileEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (
                  item: core.BulkRecordRequest<commonprofile.CommonProfileEntity>
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
  const mutationDeleteCommonProfile = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deleteCommonProfile();
  });

  const fnDeleteCommonProfileUpdater = (
    data: IResponseList<commonprofile.CommonProfileEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = CommonProfileActions.getCommonProfileEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeleteCommonProfile = (
    values: string[],
    formikProps?: FormikHelpers<commonprofile.CommonProfileEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeleteCommonProfile.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<
            IResponseList<commonprofile.CommonProfileEntity>
          >("*[]commonprofile.CommonProfileEntity", (data) =>
            fnDeleteCommonProfileUpdater(data, values)
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
    commonProfilesQuery,
    commonProfilesExportQuery,
    commonProfileByUniqueIdQuery,
    mutationPostCommonProfile,
    submitPostCommonProfile,
    mutationPatchCommonProfile,
    submitPatchCommonProfile,
    mutationPatchCommonProfileDistinct,
    submitPatchCommonProfileDistinct,
    commonProfileDistinctQuery,
    mutationPatchCommonProfiles,
    submitPatchCommonProfiles,
    mutationDeleteCommonProfile,
    submitDeleteCommonProfile,
  };
}
