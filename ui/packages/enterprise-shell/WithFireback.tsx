import { type AppConfig } from "@fireback/ui-core/hooks/appConfigTools";

import React from "react";
import { QueryClient } from "@tanstack/react-query";
import { useAuthentication } from "@fireback/auth-client";
import { AuthenticationProvider } from "@fireback/auth-client";
import { FetchxContext } from "@fireback/js-remote-ctx/common/fetchx";
import { FetchxProvider } from "@fireback/js-remote-ctx/react/useFetchx";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
import { useFirebackSocket } from "@fireback/ui-core/hooks/useFirebackSocket";

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

  const fetchContext = React.useRef(
    new FetchxContext(BUILD_VARIABLES.REMOTE_SERVICE?.replace(/\/$/, "")),
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
