import { FormikHelpers } from "formik";
import React, { useContext } from "react";
import { useMutation } from "react-query";
import {
  execApiFn,
  IDeleteResponse,
  mutationErrorsToFormik,
  DeleteRequest
} from "../../core/http-tools";
import { DeleteProps, RemoteQueryContext, queryBeforeSend } from "../../core/react-tools";
export function useDeleteTableViewSizing(props?: DeleteProps) {
  const {execFnOverride, queryClient, query} = (props || {})
  const { options, execFn } = useContext(RemoteQueryContext);
  // Calculare the function which will do the remote calls.
  // We consider to use global override, this specific override, or default which
  // comes with the sdk.
  const rpcFn = execFnOverride
    ? execFnOverride(options)
    : execFn
    ? execFn(options)
    : execApiFn(options);
  // Url of the remote affix.
  const url = "/table-view-sizing".substr(1);
  let computedUrl = `${url}?${new URLSearchParams(
    queryBeforeSend(query)
  ).toString()}`;
  // Attach the details of the request to the fn
  const fn = (body: any) => rpcFn("DELETE", computedUrl, body);
  const mutation = useMutation<
    IDeleteResponse,
    IDeleteResponse,
    DeleteRequest
  >(fn);
  const fnUpdater = (
    data: IDeleteResponse | undefined,
    item: IDeleteResponse
  ) => {
    return data;
  };
  const submit = (
    values: DeleteRequest,
    formikProps?: FormikHelpers<any>
  ): Promise<IDeleteResponse> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IDeleteResponse) {
          /*
          * Here we need an improvement here. We are cleanning the data, but
          * We may not have to actually
          */
          queryClient?.setQueryData<IDeleteResponse>(
            "*workspaces.TableViewSizingEntity",
            (data) => fnUpdater(data, response) as any
          );
          queryClient?.invalidateQueries("*workspaces.TableViewSizingEntity");
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
