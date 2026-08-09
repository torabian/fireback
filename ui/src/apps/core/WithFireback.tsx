import { type AppConfig } from "../../modules/fireback-ui/hooks/appConfigTools";

import React from "react";
import { QueryClient } from "@tanstack/react-query";
import { useAuthentication } from "../../modules/fireback-ui/auth/AuthenticationContext";
import { AuthenticationProvider } from "../../modules/fireback-ui/auth/AuthenticationProvider";
import { FetchxContext } from "../../modules/sdk/sdk/common/fetchx";
import { FetchxProvider } from "../../modules/sdk/sdk/react/useFetchx";
import { BUILD_VARIABLES } from "../../modules/fireback-ui/hooks/build-variables";
import { useFirebackSocket } from "../../modules/fireback-ui/hooks/useFirebackSocket";

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
