// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import { useMutation, useQuery, useQueryClient, QueryClient , UseQueryOptions} from "react-query";
import { <%- actionClass %> } from "./<%- sourceAction %>";
import * as <%- data.ModuleName %> from "./index";
import { execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  ExecApi,
  mutationErrorsToFormik,
  IResponseList
} from "../../core/http-tools";
import {RemoteQueryContext } from "../../core/react-tools";

export function FN_NAME({ queryOptions, execFnOverride, query, queryClient, unauthorized }: { query?: any , queryClient?: QueryClient, unauthorized?: boolean, execFnOverride?: any, queryOptions?: UseQueryOptions<any>}) {
  const { options, execFn } = useContext(RemoteQueryContext);
  const fnx = execFnOverride ? <%- actionClass %>.fnExec(execFnOverride(options)) : execFn ? <%- actionClass %>.fnExec(execFn(options)) : <%- actionClass %>.fn(options)

  const Q = () =>
    fnx
      .withPreloads(query?.withPreloads)
      .query(query.query);

  const fn = () => Q().<%-route.ExternFuncName %>(
    <% if (route.Params && route.Params.length) { %>
            
      <% for (const param of route.Params) { %>
          query.<%- param.replace(":", "") %>,
      <% } %>
  <% } %>

  )

  const auth = options?.headers?.authorization
  const hasKey = auth != "undefined" && auth != undefined && auth !=null && auth != "null" && !!auth
  const query$ = useQuery([options, query, "<%-route.ExternFuncName %>"], fn, {
    cacheTime: 1001,
    retry: false,
    keepPreviousData: true,
    enabled: (hasKey || unauthorized ) && !!query?.uniqueId,
    ...(queryOptions || {})
  });

  return { query: query$ };
}
