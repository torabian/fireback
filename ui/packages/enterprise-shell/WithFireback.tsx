import { type AppConfig } from "@fireback/ui-core/hooks/appConfigTools";

import React from "react";
import { QueryClient } from "@tanstack/react-query";
import { useAuthentication } from "@fireback/auth-client";
import { AuthenticationProvider } from "@fireback/auth-client";
import { FetchxContext } from "@fireback/js-remote-ctx/common/fetchx";
import { FetchxProvider } from "@fireback/js-remote-ctx/react/useFetchx";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
import { useFirebackSocket } from "@fireback/ui-core/hooks/useFirebackSocket";
import { useWasmFetchOverride } from "@fireback/wasm-server/WasmFetchContext";

export function WithFireback({
  children,
  queryClient,
  prefix,
  config,
  locale,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
  config: AppConfig;
  prefix?: string;
  locale?: string;
}) {
  return (
    <AuthenticationProvider selfServiceUrl="/selfservice">
      <WithFetchX queryClient={queryClient} children={children} />
    </AuthenticationProvider>
  );
}

const WithFetchX = ({
  children,
  queryClient,
}: {
  children: any;
  queryClient: QueryClient;
}) => {
  const { token, selectedWorkspace } = useAuthentication();

  // When an app mounts WithWasmServer (@fireback/wasm-server) around its
  // router with VITE_USE_WASM_SERVER on, it has already downloaded and
  // booted a fireback server compiled to wasm by the time this mounts (it
  // gates rendering of its children around that) — so the base URL becomes
  // irrelevant and every request instead goes through wasmFetchOverride to
  // the in-tab server, picked up here via WasmFetchOverrideContext (see that
  // file's own doc comment for why it's a context and not a direct import of
  // @fireback/wasm-server: this file is shared by every app, and a static
  // import of the real wasm/pglite code here would ship it in all of them).
  // src/apps/wasm-demo/App.tsx is the one app that does this today (see its
  // own doc comment) - for projectname/self-service there's no Provider
  // mounted, so this stays undefined and behaves exactly as before.
  const wasmFetchOverride = useWasmFetchOverride();

  const fetchContext = React.useRef(
    new FetchxContext(
      BUILD_VARIABLES.REMOTE_SERVICE?.replace(/\/$/, ""),
      {},
      undefined,
      undefined,
      wasmFetchOverride,
    ),
  );

  fetchContext.current.defaultHeaders = {
    authorization: token,
    "workspace-id": selectedWorkspace?.workspaceId,
  };

  useFirebackSocket(
    BUILD_VARIABLES.REMOTE_SERVICE,
    token,
    selectedWorkspace?.workspaceId,
    queryClient,
  );

  return (
    <FetchxProvider value={fetchContext.current}>{children}</FetchxProvider>
  );
};
