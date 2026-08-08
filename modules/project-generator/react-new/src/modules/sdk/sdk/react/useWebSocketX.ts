// @ts-nocheck
import { useEffect, useRef, useState, useCallback } from "react";
import { WebSocketX } from "../common/WebSocketX";
export type UseWebSocketResult<TSend, TRecv> = {
  send: (msg: TSend) => void;
  close: () => void;
  restart: () => void;
  socket?: WebSocketX<TSend, TRecv>;
} & UseWebSocketState<TRecv>;
interface UseWebSocketState<TRecv> {
  messages: TRecv[];
  isOpen: boolean;
  error?: Event | null;
}
export interface UseWebSocketXOptions {
  // When true, no connection is opened on mount - the first one only opens
  // once you call restart(). Useful for something like a search-as-you-type
  // box that shouldn't hit the server until the user actually types.
  lazy?: boolean;
}
export function useWebSocketX<TSend = any, TRecv = any, TQuery = any>(
  fn: (
    overrideUrl?: string | undefined,
    qs?: TQuery | undefined
  ) => WebSocketX<TSend, TRecv>,
  options?: UseWebSocketXOptions
): UseWebSocketResult<TSend, TRecv> {
  const socketRef = useRef<WebSocketX<TSend, TRecv> | null>(null);
  // fn is re-captured into this ref on every render, so create()/restart()
  // below always reconnect using whatever url/qs the caller most recently
  // passed in. Without this, create (memoized once via useCallback(..., []))
  // would keep closing over the very first fn it ever received, and
  // restart() could never pick up a changed query (e.g. a new search
  // phrase) - it would just reopen the same original connection.
  const fnRef = useRef(fn);
  fnRef.current = fn;
  const [state, setState] = useState<UseWebSocketState<TRecv>>({
    messages: [],
    isOpen: false,
    error: undefined,
  });
  const create = useCallback(() => {
    const ws = fnRef.current();
    socketRef.current = ws;
    setState({
      messages: [],
      error: undefined,
      isOpen: ws.readyState === ws.OPEN,
    });
    ws.addEventListener("message", (ev) => {
      setState((prev) => {
        return {
          ...prev,
          messages: [...prev.messages, ev.data],
        };
      });
    });
    ws.addEventListener("error", (ev) => {
      setState((prev) => {
        return {
          ...prev,
          error: ev,
        };
      });
    });
    ws.addEventListener("open", () => {
      setState((prev) => {
        return {
          ...prev,
          isOpen: true,
        };
      });
    });
    ws.addEventListener("close", () => {
      setState((prev) => {
        return {
          ...prev,
          isOpen: false,
        };
      });
    });
    return ws;
  }, []);
  useEffect(() => {
    if (options?.lazy) {
      return;
    }
    create();
  }, [create]);
  // Always close whatever's in socketRef on unmount, whether it was opened by
  // the eager-mount effect above or later, on demand, via restart() in lazy
  // mode - a single effect with no deps so it registers this exactly once.
  useEffect(() => {
    return () => {
      socketRef.current?.close();
    };
  }, []);
  const send = (msg: TSend) => {
    socketRef.current?.send(msg);
  };
  const close = () => {
    socketRef.current?.close();
  };
  const restart = () => {
    close();
    create();
  };
  return {
    ...state,
    send,
    close,
    restart,
    socket: socketRef.current || undefined,
  };
}
