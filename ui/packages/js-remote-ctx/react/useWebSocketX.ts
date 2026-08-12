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

export function useWebSocketX<TSend = any, TRecv = any, TQuery = any>(
  fn: (
    overrideUrl?: string | undefined,
    qs?: TQuery | undefined
  ) => WebSocketX<TSend, TRecv>
): UseWebSocketResult<TSend, TRecv> {
  const socketRef = useRef<WebSocketX<TSend, TRecv> | null>(null);
  const [state, setState] = useState<UseWebSocketState<TRecv>>({
    messages: [],
    isOpen: false,
    error: undefined,
  });

  const create = useCallback(() => {
    const ws = fn();
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
    const ws = create();
    return () => ws.close();
  }, [create]);

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
