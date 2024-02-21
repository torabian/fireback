import { FormikHelpers } from "formik";
import { useContext } from "react";
import { useMutation, QueryClient } from "react-query";
import {
  execApiFn,
  IResponse,
  mutationErrorsToFormik,
  IResponseList
} from "../../core/http-tools";

import { RemoteQueryContext, queryBeforeSend, PatchProps } from "../../core/react-tools";


{{ range $key, $value := .imports }}
{{ if $value.Items}}
import {
  {{ range $value.Items }}
    {{ .}},
  {{ end }}

} from "{{ $key}}"
{{ end }}
{{ end }}


export function use{{ .r.GetFuncNameUpper}}(props?: PatchProps) {
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
  const url = "{{ .r.Url}}".substr(1);

  let computedUrl = `${url}?${new URLSearchParams(
    queryBeforeSend(query)
  ).toString()}`;

  {{ template "routeUrl" .r }}

  // Attach the details of the request to the fn
  const fn = (body: any) => rpcFn("{{ .r.Method }}", computedUrl, body);

  const mutation = useMutation<
    IResponse<{{ .r.ResponseEntityComputed}}>,
    IResponse<{{ .r.ResponseEntityComputed}}>,
    Partial<{{ .r.RequestEntityComputed}}>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<{{ .r.ResponseEntityComputed}}> | undefined,
    item: IResponse<{{ .r.ResponseEntityComputed}}>
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
    values: Partial<{{ .r.RequestEntityComputed}}>,
    formikProps?: FormikHelpers<Partial<{{ .r.ResponseEntityComputed}}>>
  ): Promise<IResponse<{{ .r.ResponseEntityComputed}}>> => {
    return new Promise((resolve, reject) => {
      
      mutation.mutate(values, {
        onSuccess(response: IResponse<{{ .r.ResponseEntityComputed}}>) {
          queryClient?.setQueriesData("{{ .r.EntityKey }}", (data: any) =>
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
