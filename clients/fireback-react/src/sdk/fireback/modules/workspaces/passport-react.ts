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
import { PassportActions } from "./passport-actions";
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
    ? PassportActions.fnExec(execFn(options))
    : PassportActions.fn(options);
  const Q = () =>
    caller
      .startIndex(query?.startIndex)
      .deep(query?.deep)
      .itemsPerPage(query?.itemsPerPage)
      .query(query?.query);

  const passportsQuery = useQuery(
    ["*[]workspaces.PassportEntity", options],
    () => Q().getPassports(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const passportsExportQuery = useQuery(
    ["*[]workspaces.PassportEntity", options],
    () => Q().getPassportsExport(),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  const passportByUniqueIdQuery = useQuery(
    ["*workspaces.PassportEntity", options],
    (uniqueId: string) => Q().getPassportByUniqueId(uniqueId),
    {
      cacheTime: 1000,
      enabled: !!options?.headers?.authorization,
    }
  );

  // post passport

  const mutationPostPassport = useMutation<
    IResponse<workspaces.PassportEntity>,
    IResponse<workspaces.PassportEntity>,
    workspaces.PassportEntity
  >((entity) => {
    return Q().postPassport(entity);
  });

  // Only entities are having a store in front-end

  const fnPostPassportUpdater = (
    data: IResponseList<workspaces.PassportEntity> | undefined,
    item: IResponse<workspaces.PassportEntity>
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
          PassportActions.isPassportEntityEqual(t, item.data)
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

  const submitPostPassport = (
    values: workspaces.PassportEntity,
    formikProps?: FormikHelpers<workspaces.PassportEntity>
  ): Promise<IResponse<workspaces.PassportEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPostPassport.mutate(values, {
        onSuccess(response: IResponse<workspaces.PassportEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.PassportEntity>>(
            "*[]workspaces.PassportEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.PassportEntity) => {
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

  // patch passport

  const mutationPatchPassport = useMutation<
    IResponse<workspaces.PassportEntity>,
    IResponse<workspaces.PassportEntity>,
    workspaces.PassportEntity
  >((entity) => {
    return Q().patchPassport(entity);
  });

  // Only entities are having a store in front-end

  const fnPatchPassportUpdater = (
    data: IResponseList<workspaces.PassportEntity> | undefined,
    item: IResponse<workspaces.PassportEntity>
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
          PassportActions.isPassportEntityEqual(t, item.data)
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

  const submitPatchPassport = (
    values: workspaces.PassportEntity,
    formikProps?: FormikHelpers<workspaces.PassportEntity>
  ): Promise<IResponse<workspaces.PassportEntity>> => {
    return new Promise((resolve, reject) => {
      mutationPatchPassport.mutate(values, {
        onSuccess(response: IResponse<workspaces.PassportEntity>) {
          queryClient.setQueriesData<IResponseList<workspaces.PassportEntity>>(
            "*[]workspaces.PassportEntity",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: workspaces.PassportEntity) => {
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

  // patch passports

  const mutationPatchPassports = useMutation<
    IResponse<core.BulkRecordRequest<workspaces.PassportEntity>>,
    IResponse<core.BulkRecordRequest<workspaces.PassportEntity>>,
    core.BulkRecordRequest<workspaces.PassportEntity>
  >((entity) => {
    return Q().patchPassports(entity);
  });

  // Only entities are having a store in front-end

  const submitPatchPassports = (
    values: core.BulkRecordRequest<workspaces.PassportEntity>,
    formikProps?: FormikHelpers<
      core.BulkRecordRequest<workspaces.PassportEntity>
    >
  ): Promise<IResponse<core.BulkRecordRequest<workspaces.PassportEntity>>> => {
    return new Promise((resolve, reject) => {
      mutationPatchPassports.mutate(values, {
        onSuccess(
          response: IResponse<core.BulkRecordRequest<workspaces.PassportEntity>>
        ) {
          queryClient.setQueriesData<
            IResponseList<core.BulkRecordRequest<workspaces.PassportEntity>>
          >(
            "*[]core.BulkRecordRequest[workspaces.PassportEntity]",
            (data: any) => {
              if (!(data?.data?.items?.length > 0)) {
                return data;
              }

              data.data.items = data.data.items.map(
                (item: core.BulkRecordRequest<workspaces.PassportEntity>) => {
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
  const mutationDeletePassport = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    core.DeleteRequest
  >(() => {
    return Q().deletePassport();
  });

  const fnDeletePassportUpdater = (
    data: IResponseList<workspaces.PassportEntity> | undefined,
    deleteItemsList: string[]
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    if (data?.data?.items) {
      data.data.items = data.data.items.filter((t) => {
        const key = PassportActions.getPassportEntityPrimaryKey(t);

        if (!deleteItemsList.includes(key)) {
          return true;
        }

        return false;
      });
    }

    return data;
  };

  const submitDeletePassport = (
    values: string[],
    formikProps?: FormikHelpers<workspaces.PassportEntity>
  ) => {
    return new Promise((resolve, reject) => {
      mutationDeletePassport.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          queryClient.setQueryData<IResponseList<workspaces.PassportEntity>>(
            "*[]workspaces.PassportEntity",
            (data) => fnDeletePassportUpdater(data, values)
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

  // post passport/signup/email

  const mutationPostPassportSignupEmail = useMutation<
    IResponse<workspaces.UserSessionDto>,
    IResponse<workspaces.UserSessionDto>,
    workspaces.EmailAccountSignupDto
  >((entity) => {
    return Q().postPassportSignupEmail(entity);
  });

  // Only entities are having a store in front-end

  const submitPostPassportSignupEmail = (
    values: workspaces.EmailAccountSignupDto,
    formikProps?: FormikHelpers<workspaces.EmailAccountSignupDto>
  ): Promise<IResponse<workspaces.UserSessionDto>> => {
    return new Promise((resolve, reject) => {
      mutationPostPassportSignupEmail.mutate(values, {
        onSuccess(response: IResponse<workspaces.UserSessionDto>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailAccountSignupDto>
          >("*[]workspaces.EmailAccountSignupDto", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailAccountSignupDto) => {
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

  // post passport/signin/email

  const mutationPostPassportSigninEmail = useMutation<
    IResponse<workspaces.UserSessionDto>,
    IResponse<workspaces.UserSessionDto>,
    workspaces.EmailAccountSigninDto
  >((entity) => {
    return Q().postPassportSigninEmail(entity);
  });

  // Only entities are having a store in front-end

  const submitPostPassportSigninEmail = (
    values: workspaces.EmailAccountSigninDto,
    formikProps?: FormikHelpers<workspaces.EmailAccountSigninDto>
  ): Promise<IResponse<workspaces.UserSessionDto>> => {
    return new Promise((resolve, reject) => {
      mutationPostPassportSigninEmail.mutate(values, {
        onSuccess(response: IResponse<workspaces.UserSessionDto>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailAccountSigninDto>
          >("*[]workspaces.EmailAccountSigninDto", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailAccountSigninDto) => {
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

  // post passport/authorizeOs

  const mutationPostPassportAuthorizeOs = useMutation<
    IResponse<workspaces.UserSessionDto>,
    IResponse<workspaces.UserSessionDto>,
    workspaces.EmailAccountSigninDto
  >((entity) => {
    return Q().postPassportAuthorizeOs(entity);
  });

  // Only entities are having a store in front-end

  const submitPostPassportAuthorizeOs = (
    values: workspaces.EmailAccountSigninDto,
    formikProps?: FormikHelpers<workspaces.EmailAccountSigninDto>
  ): Promise<IResponse<workspaces.UserSessionDto>> => {
    return new Promise((resolve, reject) => {
      mutationPostPassportAuthorizeOs.mutate(values, {
        onSuccess(response: IResponse<workspaces.UserSessionDto>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.EmailAccountSigninDto>
          >("*[]workspaces.EmailAccountSigninDto", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.EmailAccountSigninDto) => {
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

  // post passport/request-reset-mail-password

  const mutationPostPassportRequestResetMailPassword = useMutation<
    IResponse<workspaces.EmailOtpResponse>,
    IResponse<workspaces.EmailOtpResponse>,
    workspaces.OtpAuthenticateDto
  >((entity) => {
    return Q().postPassportRequestResetMailPassword(entity);
  });

  // Only entities are having a store in front-end

  const submitPostPassportRequestResetMailPassword = (
    values: workspaces.OtpAuthenticateDto,
    formikProps?: FormikHelpers<workspaces.OtpAuthenticateDto>
  ): Promise<IResponse<workspaces.EmailOtpResponse>> => {
    return new Promise((resolve, reject) => {
      mutationPostPassportRequestResetMailPassword.mutate(values, {
        onSuccess(response: IResponse<workspaces.EmailOtpResponse>) {
          queryClient.setQueriesData<
            IResponseList<workspaces.OtpAuthenticateDto>
          >("*[]workspaces.OtpAuthenticateDto", (data: any) => {
            if (!(data?.data?.items?.length > 0)) {
              return data;
            }

            data.data.items = data.data.items.map(
              (item: workspaces.OtpAuthenticateDto) => {
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
    passportsQuery,
    passportsExportQuery,
    passportByUniqueIdQuery,
    mutationPostPassport,
    submitPostPassport,
    mutationPatchPassport,
    submitPatchPassport,
    mutationPatchPassports,
    submitPatchPassports,
    mutationDeletePassport,
    submitDeletePassport,
    mutationPostPassportSignupEmail,
    submitPostPassportSignupEmail,
    mutationPostPassportSigninEmail,
    submitPostPassportSigninEmail,
    mutationPostPassportAuthorizeOs,
    submitPostPassportAuthorizeOs,
    mutationPostPassportRequestResetMailPassword,
    submitPostPassportRequestResetMailPassword,
  };
}
