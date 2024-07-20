import { FormikHelpers } from "formik";
import React, { 
  useCallback,
  useContext,
  useState,
  useRef,
  useEffect
} from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
  UseQueryOptions
} from "react-query";
import { RemoteQueryContext } from "../../core/react-tools";
interface ReactiveQueryProps {
  query?: any ,
  queryClient?: QueryClient,
  unauthorized?: boolean,
  execFnOverride?: any,
  queryOptions?: UseQueryOptions<any>,
  onMessage?: (msg: string) => void;
  presistResult?: boolean;
}
export function useReactiveReactiveSearch({ 
  queryOptions,
  execFnOverride,
  query,
  queryClient,
  unauthorized,
  onMessage,
  presistResult,
 }: ReactiveQueryProps) {
  const { options } = useContext(RemoteQueryContext);
  const remote = options.prefix;
  const token = options.headers?.authorization;
  const workspaceId = (options.headers as any)["workspace-id"];
  const connection = useRef<WebSocket>();
  const [result, setResult] = useState<any[]>([]);
  const appendResult = (result: any) => {
    setResult((v) => [...v, result]);
  };
  const [connected, setConnected] = useState(false);
  const close = () => {
    if (connection.current?.readyState === 1) {
      connection.current?.close();
    }
    setConnected(false);
  };
  const write = (data: string | ArrayBufferLike | Blob | ArrayBufferView) => {
    connection.current?.send(data);
  };
  /*
  * Creates the connection and tries to establish the connection
  */
  const operate = (value: any) => {
    if (connection.current?.readyState === 1) {
      connection.current?.close();
    }
    setResult([]);
    const wsRemote = remote?.replace("https", "wss").replace("http", "ws");
    const remoteUrl = `/reactiveSearch`.substr(1)
    let url = `${wsRemote}${remoteUrl}?acceptLanguage=${
      (options as any).headers["accept-language"]
    }&token=${token}&workspaceId=${workspaceId}&${new URLSearchParams(
      value
    )}&${new URLSearchParams(query || {})}`;
    url = url.replace(":uniqueId", query?.uniqueId);
    let conn = new WebSocket(url);
    connection.current = conn;
    conn.onopen = function () {
      setConnected(true);
    };
    conn.onmessage = function (evt: any) {
      try {
        const msg = JSON.parse(evt.data);
        if (msg) {
          onMessage && onMessage(msg);
          if (presistResult !== false) {
            appendResult(msg);
          }
        }
      } catch (e: any) {
        // Intenrionnaly left blank
      }
    };
  };
  useEffect(() => {
    return () => {
      close();
    };
  }, []);
  return { operate, data: result,  close, connected, write };
}
