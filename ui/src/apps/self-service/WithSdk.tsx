import { type AppConfig } from "@fireback/ui-core/hooks/appConfigTools";

import React from "react";
import { QueryClient } from "@tanstack/react-query";

// Import the Remote Query Provider from your new generated SDK
// import { RemoteQueryProvider as TestQueryProviders } from "@fireback/sdk/projectname/core/react-tools";

export function WithSdk({
  children,
  queryClient,
  config,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
  config: AppConfig;
}) {
  return children;

  // Uncomment the code below, if you want to inject your own project query provider
  // on top of the Fireback. Fireback features have their own provider, you might need this
  // as many as different backends you want to connect
  // return (
  //   <TestQueryProviders
  //     preferredAcceptLanguage={locale || config.interfaceLanguage}
  //     identifier="projectname"
  //     queryClient={queryClient}
  //     remote={BUILD_VARIABLES.REMOTE_SERVICE}
  //     defaultExecFn={() => (options: any) => myCustomExecFn(options)}
  //   >
  //     {children}
  //   </TestQueryProviders>
  // );
}
