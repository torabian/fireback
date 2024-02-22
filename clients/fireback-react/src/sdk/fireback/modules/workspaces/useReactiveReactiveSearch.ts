// @ts-nocheck

import { FormikHelpers } from "formik";
import React, {
  useCallback,
  useContext,
  useState,
  useRef,
  useEffect,
} from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
  UseQueryOptions,
} from "react-query";
import { WorkspaceActions } from "./workspace-actions";
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
import { RemoteQueryContext } from "../../core/react-tools";

interface QueryData {
  phrase: string;
}

export function useReactiveReactiveSearch({
  queryOptions,
  execFnOverride,
  query,
  queryClient,
  unauthorized,
}: {
  query?: any;
  queryClient?: QueryClient;
  unauthorized?: boolean;
  execFnOverride?: any;
  queryOptions?: UseQueryOptions<any>;
}) {
  /*
  {
  "Method": "REACTIVE",
  "Url": "reactiveSearch",
  "ExternFuncName": "reactiveReactiveSearch",
  "RequestEntity": "",
  "TargetEntity": "",
  "ResponseEntity": "*core.ReactiveSearchResultDto",
  "Action": "",
  "Params": []
}

  */
  const { options } = useContext(RemoteQueryContext);
  const remote = options.prefix;
  const token = options.headers?.authorization;
  const workspaceId = (options.headers as any)["workspace-id"];
  const conneciton = useRef<WebSocket>();
  const [result, setResult] = useState([]);
  const appendResult = (result: IReactiveSearchResult) => {
    setResult((v) => [...v, result]);
  };

  const operate = (value: QueryData) => {
    setResult([]);

    const wsRemote = remote.replace("https", "wss").replace("http", "ws");

    let conn = new WebSocket(
      `${wsRemote}reactiveSearch?acceptLanguage=${
        options.headers["accept-language"]
      }&token=${token}&workspaceId=${workspaceId}&${new URLSearchParams(value)}`
    );

    conneciton.current = conn;

    conn.onmessage = function (evt: any) {
      try {
        const msg = JSON.parse(evt.data);
        if (msg) {
          appendResult(msg);
        }
      } catch (e: any) {
        // Intenrionnaly left blank
      }
    };
  };

  useEffect(() => {
    return () => {
      if (conneciton.current?.readyState === 1) {
        conneciton.current?.close();
      }
    };
  }, []);

  return { operate, data: result };
}
