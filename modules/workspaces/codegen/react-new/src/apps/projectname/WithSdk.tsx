import { AppConfig } from "@/modules/fireback/hooks/appConfigTools";
import { useT } from "@/modules/fireback/hooks/useT";

import { mockExecFn } from "@/modules/fireback/hooks/mock-tools";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import React from "react";
import { QueryClient } from "react-query";

// Import the Remote Query Provider from your new generated SDK
// import { RemoteQueryProvider as TestQueryProviders } from "../../modules/sdk/projectname/core/react-tools";

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

  return children;

  // Uncomment the code below, if you want to inject your own project query provider
  // on top of the Fireback. Fireback features have their own provider, you might need this
  // as many as different backends you want to connect
  // return (
  //   <TestQueryProviders
  //     preferredAcceptLanguage={locale || config.interfaceLanguage}
  //     identifier="projectname"
  //     queryClient={queryClient}
  //     remote={process.env.REACT_APP_REMOTE_SERVICE}
  //     /// #if process.env.REACT_APP_INACCURATE_MOCK_MODE == "true"
  //     defaultExecFn={() => {
  //       return (options: any) => mockExecFn(options, mockServer.current);
  //     }}
  //     /// #endif
  //     // defaultExecFn={() => (options: any) =>
  //     //   mockExecFn(options, mockServer.current, t)}
  //   >
  //     {children}
  //   </TestQueryProviders>
  // );
}
