import { type AppConfig } from "../../hooks/appConfigTools";

import React, { useRef } from "react";
import { QueryClient } from "react-query";
import { fetchXMock, mockExecFn } from "../../hooks/mock-tools";
import { RemoteQueryProvider as FirebackQueryProvider } from "../../sdk/core/react-tools";
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
  const fetchContext = useRef(
    new FetchxContext(BUILD_VARIABLES.REMOTE_SERVICE?.replace(/\/$/, "")),
  );

  if (BUILD_VARIABLES.INACCURATE_MOCK_MODE === "true") {
    fetchContext.current.fetchOverrideFn = fetchXMock(mockServer);
  }

  return (
    <FirebackQueryProvider
      socket
      preferredAcceptLanguage={locale || config.interfaceLanguage}
      identifier="fireback"
      prefix={prefix}
      queryClient={queryClient}
      remote={BUILD_VARIABLES.REMOTE_SERVICE}
      /// #if BUILD_VARIABLES.INACCURATE_MOCK_MODE == "true"
      defaultExecFn={() => {
        return (options: any) => mockExecFn(options, mockServer.current);
      }}
      /// #endif
      // defaultExecFn={() => (options: any) =>
      //   mockExecFn(options, mockServer.current, t)}
    >
      <FetchxProvider value={fetchContext.current}>{children}</FetchxProvider>
    </FirebackQueryProvider>
  );
}
