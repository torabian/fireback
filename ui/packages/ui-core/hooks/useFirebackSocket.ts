import { useEffect, useState } from "react";
import type { QueryClient } from "@tanstack/react-query";

export interface SocketNotification<T = any> {
  cacheKey?: string;
  name: string;
  payload: T;
}

/**
 * Opens the app's websocket connection (query-invalidation notifications from
 * the backend) for as long as a token/workspace pair is available. Mounted
 * once, near the app root (see `WithFireback.tsx`) - moved here verbatim from
 * the old `sdk/core/react-tools.tsx`'s `useSocket`, just re-sourced from
 * `useAuthentication()` instead of `RemoteQueryContext`.
 */
export function useFirebackSocket(
  remote: string | undefined,
  token: string | undefined,
  workspaceId: string | undefined,
  queryClient: QueryClient | undefined,
  enabled = true,
) {
  const [socketState, setSocketState] = useState({ state: "unknown" });

  useEffect(() => {
    if (
      !remote ||
      !token ||
      token === "undefined" ||
      // No workspace resolved yet (e.g. a multi-workspace account sitting on
      // the workspace picker, before selectWorkspace() has run) - wait for
      // one rather than opening a socket scoped to the literal string
      // "undefined" (template-literal-stringified from an actual `undefined`
      // workspaceId), which the backend now accepts as a value rather than
      // rejecting outright, so this would otherwise silently subscribe to a
      // bogus workspace instead of failing loudly.
      !workspaceId ||
      workspaceId === "undefined" ||
      enabled === false
    ) {
      return;
    }
    const wsRemote = remote.replace("https", "wss").replace("http", "ws");
    let conn: WebSocket;
    try {
      conn = new WebSocket(
        `${wsRemote}ws?token=${token}&workspaceId=${workspaceId}`,
      );
      conn.onerror = function (evt) {
        setSocketState({ state: "error" });
      };
      conn.onclose = function (evt) {
        setSocketState({ state: "closed" });
      };
      conn.onmessage = function (evt: any) {
        try {
          const msg: SocketNotification = JSON.parse(evt.data);

          // Old style event messages
          if (msg?.cacheKey) {
            // react-query's types want a filters object here, but this
            // string-key call style is the same one used elsewhere in the
            // app (e.g. WorkspaceList.tsx) and matches the old
            // react-tools.tsx's behavior (that file just had `@ts-nocheck`
            // suppressing the mismatch) - preserved as-is rather than
            // changed to `{ queryKey: [...] }`, which would be a real
            // behavior change outside this migration's scope.
            queryClient?.invalidateQueries(msg?.cacheKey as any);
          }
        } catch (e: any) {
          // The socket isn't guaranteed to send JSON - the server also
          // pushes plain-text status messages, e.g. "Authorization general
          // failed" when the token/workspace it was opened with gets
          // rejected. That's not a parsing bug, so it shouldn't be logged
          // (or read) as one - surface it as socket state instead, so
          // callers can actually react to it (re-authenticate, show a
          // banner, etc.) instead of it only ever showing up as a console
          // error.
          if (typeof evt?.data === "string") {
            console.warn("Socket sent a non-JSON message:", evt.data);
            setSocketState({ state: "error", message: evt.data } as any);
          } else {
            console.error("Socket message parsing error", evt);
          }
        }
      };
      conn.onopen = function (evt) {
        setSocketState({ state: "connected" });
      };
    } catch (err) {}

    return () => {
      if (conn?.readyState === 1) {
        conn.close();
      }
    };
  }, [token, workspaceId]);

  return { socketState };
}
