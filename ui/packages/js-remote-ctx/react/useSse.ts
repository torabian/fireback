import { useEffect, useRef, useState } from "react";
import type { TypedRequestInit } from "../common/fetchx";

export interface UseSSEResult<T> {
  messages: Array<T>;
  error?: Error | null;
  cancel: () => void;
  restart: () => void;
}

type SseFetchFn<Q, H, R> = (
  onMessage?: (ev: MessageEvent) => void,
  qs?: Q,
  init?: TypedRequestInit<R, H>,
  overrideUrl?: string
) => Promise<any>;

export function useSse<T = string, Q = any, H = any, R = any>(
  fetchFn: SseFetchFn<Q, H, R>,
  props?: {
    qs?: Q;
    init?: TypedRequestInit<R, H>;
    overrideUrl?: string;
  }
) {
  const acRef = useRef<AbortController | null>(null);

  const [state, setState] = useState<Partial<UseSSEResult<T>>>({
    messages: [],
  });

  const create = () => {
    const ac = new AbortController();
    acRef.current = ac;

    setState({ messages: [], error: null }); // reset on new stream

    fetchFn(
      (ev) => {
        setState((value) => {
          const next = ev.data as T;
          if ((value.messages || []).includes(next)) return value;
          return { ...value, messages: [...(value.messages || []), next] };
        });
      },
      props?.qs,
      { ...(props?.init || {}), signal: ac.signal },
      props?.overrideUrl
    ).catch((err) => setState((v) => ({ ...v, error: err as Error })));

    return () => acRef.current?.abort();
  };

  useEffect(() => {
    return create();
  }, []);

  const cancel = () => {
    acRef.current?.abort();
  };

  const restart = () => {
    cancel();
    create();
  };

  return { ...state, cancel, restart, messages: state.messages as T[] };
}
