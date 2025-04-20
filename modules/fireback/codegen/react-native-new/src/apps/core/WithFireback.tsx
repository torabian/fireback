import {mockExecFn} from '@/modules/fireback/hooks/mock-tools';
import {useLocale} from '@/modules/fireback/hooks/useLocale';
import {
  CredentialStorage,
  RemoteQueryProvider as FirebackQueryProvider,
} from '@/modules/fireback/sdk/core/react-tools';
import AsyncStorage from '@react-native-async-storage/async-storage';
import React from 'react';
import {Alert} from 'react-native';
import {QueryClient} from 'react-query';

export function WithFireback({
  children,
  mockServer,
  queryClient,
}: {
  children: React.ReactNode;
  queryClient: QueryClient;
  mockServer: any;
}) {
  const {locale} = useLocale();

  // Temporaty for the demo

  return (
    <FirebackQueryProvider
      preferredAcceptLanguage={locale}
      identifier="fireback"
      queryClient={queryClient}
      credentialStorage={new ReactNativeStorage()}
      /// #if process.env.REACT_APP_INACCURATE_MOCK_MODE == "true"
      defaultExecFn={() => {
        return (options: any) => mockExecFn(options, mockServer.current);
      }}
      /// #endif
      remote={'http://localhost:4500/'}>
      {children}
    </FirebackQueryProvider>
  );
}

class ReactNativeStorage implements CredentialStorage {
  async setItem(key: string, value: string) {
    return await AsyncStorage.setItem(key, value);
  }
  async getItem(key: string) {
    return await AsyncStorage.getItem(key);
  }
  async removeItem(key: string) {
    return await AsyncStorage.removeItem(key);
  }
}
