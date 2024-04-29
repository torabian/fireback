import { FormikHelpers } from "formik";
import { useContext } from "react";
import { useMutation } from "react-query";
import { 
  execApiFn,
  IResponse,
  mutationErrorsToFormik,
  IResponseList
} from "../../core/http-tools";
import {
  RemoteQueryContext,
  UseRemoteQuery,
  queryBeforeSend
} from "../../core/react-tools";
import {
    CheckClassicPassportActionReqDto,
    CheckClassicPassportActionResDto,
} from "../workspaces/WorkspacesActionsDto"
export function usePostWorkspacePassportCheck(props?: UseRemoteQuery) {
  let {queryClient, query, execFnOverride} = props || {};
  query = query || {}
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
  const url = "/workspace/passport/check".substr(1);
  let computedUrl = `${url}?${new URLSearchParams(
    queryBeforeSend(query)
  ).toString()}`;
  // Attach the details of the request to the fn
  const fn = (body: any) => rpcFn("POST", computedUrl, body);
  const mutation = useMutation<
    IResponse<CheckClassicPassportActionResDto>,
    IResponse<CheckClassicPassportActionResDto>,
    Partial<CheckClassicPassportActionReqDto>
  >(fn);
  // Only entities are having a store in front-end
  const fnUpdater = (
    data: IResponseList<CheckClassicPassportActionResDto> | undefined,
    item: IResponse<CheckClassicPassportActionResDto>
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }
    // To me it seems this is not a good or any correct strategy to update the store.
    // When we are posting, we want to add it there, that's it. Not updating it.
    // We have patch, but also posting with ID is possible.
    if (data.data && item?.data) {
      data.data.items = [item.data, ...(data?.data?.items || [])];
    }
    return data;
  };
  const submit = (
    values: Partial<CheckClassicPassportActionReqDto>,
    formikProps?: FormikHelpers<Partial<CheckClassicPassportActionResDto>>
  ): Promise<IResponse<CheckClassicPassportActionResDto>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<CheckClassicPassportActionResDto>) {
          queryClient?.setQueryData<IResponseList<CheckClassicPassportActionResDto>>(
            "*workspaces.CheckClassicPassportActionResDto",
            (data) => fnUpdater(data, response) as any
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
