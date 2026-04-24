import { type AppConfig } from "../../hooks/appConfigTools";

import React, { useContext, useRef } from "react";
import { QueryClient } from "react-query";
import { fetchXMock, mockExecFn } from "../../hooks/mock-tools";
import {
  RemoteQueryProvider as FirebackQueryProvider,
  RemoteQueryContext,
} from "../../sdk/core/react-tools";
import { FetchxContext } from "../../sdk/sdk/common/fetchx";
import { FetchxProvider } from "../../sdk/sdk/react/useFetchx";
import { BUILD_VARIABLES } from "../../hooks/build-variables";

export function WithFireback({
  children,
  queryClient,
  prefix,
  mockServer,
  config,
  locale,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
  config: AppConfig;
  mockServer: any;
  prefix?: string;
  locale?: string;
}) {
  return (
    <FirebackQueryProvider
      socket
      preferredAcceptLanguage={locale || config.interfaceLanguage}
      identifier="fireback"
      prefix={prefix}
      queryClient={queryClient}
      remote={BUILD_VARIABLES.REMOTE_SERVICE}
      defaultExecFn={
        BUILD_VARIABLES.INACCURATE_MOCK_MODE === "true"
          ? () => {
              console.log("Mock mode enabled on legacy code gen");
              return (options: any) => mockExecFn(options, mockServer.current);
            }
          : undefined
      }
    >
      <WithFetchX children={children} mockServer={mockServer} />
    </FirebackQueryProvider>
  );
}

const WithFetchX = ({
  children,
  mockServer,
}: {
  children: any;
  mockServer: any;
}) => {
  const { options, session } = useContext(RemoteQueryContext);

  const fetchContext = useRef(
    new FetchxContext(BUILD_VARIABLES.REMOTE_SERVICE?.replace(/\/$/, "")),
  );

  fetchContext.current.defaultHeaders = {
    authorization: session?.token,
    "workspace-id": options?.headers["workspace-id"],
  };

  if (BUILD_VARIABLES.INACCURATE_MOCK_MODE === "true") {
    console.log(
      "Inaccurate mock mode is enabled. All requests are being routed out.",
    );
    fetchContext.current.fetchOverrideFn = fetchXMock(mockServer);
  }

  return (
    <FetchxProvider value={fetchContext.current}>{children}</FetchxProvider>
  );
};
