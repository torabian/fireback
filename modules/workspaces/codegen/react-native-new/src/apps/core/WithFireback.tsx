import {useLocale} from '@/modules/fireback/hooks/useLocale';
import {RemoteQueryProvider as FirebackQueryProvider} from '@/modules/fireback/sdk/core/react-tools';
import React from 'react';
import {QueryClient} from 'react-query';

export function WithFireback({
  children,
  queryClient,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
}) {
  const {locale} = useLocale();

  // Temporaty for the demo

  return (
    <FirebackQueryProvider
      token="614d1ca85da6048b708f98e0d1cc22617b15f9bbea12fca0b1fe5088a91d307c"
      preferredAcceptLanguage={locale}
      identifier="fireback"
      queryClient={queryClient}
      remote={'http://localhost:4500/'}>
      {children}
    </FirebackQueryProvider>
  );
}
