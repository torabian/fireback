import { AppConfig } from "@/modules/fireback/hooks/appConfigTools";
import { useT } from "@/modules/fireback/hooks/useT";

import { mockExecFn } from "@/modules/fireback/hooks/mock-tools";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import React from "react";
import { QueryClient } from "react-query";
import { RemoteQueryProvider as TestQueryProviders } from "../../modules/sdk/projectname/core/react-tools";

export function WithSdk({
  children,
  queryClient,
  mockServer,
  config,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
  config: AppConfig;
  mockServer: any;
}) {
  const { locale } = useLocale();
  const t = useT();

  return (
    <TestQueryProviders
      preferredAcceptLanguage={locale || config.interfaceLanguage}
      identifier="projectname"
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
    </TestQueryProviders>
  );
}
