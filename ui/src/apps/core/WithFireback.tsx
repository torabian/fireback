import { type AppConfig } from "../../modules/fireback-ui/hooks/appConfigTools";

import React, { useContext, useRef } from "react";
import { QueryClient } from "@tanstack/react-query";
import {
  RemoteQueryProvider as FirebackQueryProvider,
  RemoteQueryContext,
} from "../../modules/sdk/core/react-tools";
import { FetchxContext } from "../../modules/sdk/sdk/common/fetchx";
import { FetchxProvider } from "../../modules/sdk/sdk/react/useFetchx";
import { BUILD_VARIABLES } from "../../modules/fireback-ui/hooks/build-variables";

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
    <FirebackQueryProvider
      socket
      preferredAcceptLanguage={locale || config.interfaceLanguage}
      identifier="fireback"
      prefix={prefix}
      queryClient={queryClient}
      remote={BUILD_VARIABLES.REMOTE_SERVICE}
    >
      <WithFetchX children={children} />
    </FirebackQueryProvider>
  );
}

const WithFetchX = ({ children }: { children: any }) => {
  const { options, session } = useContext(RemoteQueryContext);

  const fetchContext = useRef(
    new FetchxContext(BUILD_VARIABLES.REMOTE_SERVICE?.replace(/\/$/, "")),
  );

  fetchContext.current.defaultHeaders = {
    authorization: session?.token,
    "workspace-id": options?.headers["workspace-id"],
  };

  return (
    <FetchxProvider value={fetchContext.current}>{children}</FetchxProvider>
  );
};
