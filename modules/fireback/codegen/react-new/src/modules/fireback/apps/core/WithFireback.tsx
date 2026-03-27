import { type AppConfig } from "../../hooks/appConfigTools";

import React, { useRef } from "react";
import { QueryClient } from "react-query";
import { fetchXMock, mockExecFn } from "../../hooks/mock-tools";
import { RemoteQueryProvider as FirebackQueryProvider } from "../../sdk/core/react-tools";
import { FetchxContext } from "../../sdk/sdk/common/fetchx";
import { FetchxProvider } from "../../sdk/sdk/react/useFetchx";

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
    new FetchxContext(process.env.REACT_APP_REMOTE_SERVICE?.replace(/\/$/, "")),
  );

  /// #if process.env.REACT_APP_INACCURATE_MOCK_MODE == "true"
  fetchContext.current.fetchOverrideFn = fetchXMock(mockServer);
  /// #endif

  return (
    <FirebackQueryProvider
      socket
      preferredAcceptLanguage={locale || config.interfaceLanguage}
      identifier="fireback"
      prefix={prefix}
      queryClient={queryClient}
      remote={process.env.REACT_APP_REMOTE_SERVICE}
      /// #if process.env.REACT_APP_INACCURATE_MOCK_MODE == "true"
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
