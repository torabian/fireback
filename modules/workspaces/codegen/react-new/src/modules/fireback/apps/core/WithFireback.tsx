import { AppConfig } from "../../hooks/appConfigTools";
import { useT } from "../../hooks/useT";

import { mockExecFn } from "../../hooks/mock-tools";
import { useLocale } from "../../hooks/useLocale";
import React from "react";
import { QueryClient } from "react-query";
import { RemoteQueryProvider as FirebackQueryProvider } from "../../sdk/core/react-tools";

export function WithFireback({
  children,
  queryClient,
  prefix,
  mockServer,
  config,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
  config: AppConfig;
  mockServer: any;
  prefix?: string;
}) {
  const { locale } = useLocale();

  const t = useT();
  return (
    <FirebackQueryProvider
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
      {children}
    </FirebackQueryProvider>
  );
}
