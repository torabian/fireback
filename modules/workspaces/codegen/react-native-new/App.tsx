import 'react-native-gesture-handler';

import React, {useRef} from 'react';
import {SafeAreaView, StatusBar, useColorScheme} from 'react-native';

import {WithFireback} from '@/apps/core/WithFireback';
import {Colors} from 'react-native/Libraries/NewAppScreen';
import {QueryClient, QueryClientProvider} from 'react-query';

import {NavigationContainer} from '@react-navigation/native';
import {AuthRouter} from '@/modules/fireback/modules/auth/Router';
import {ApplicationRoutes} from './ApplicationRoutes';
import {FirebackMockServer} from '@/modules/fireback/mock-server';

function App(): React.JSX.Element {
  const queryClient = useRef(new QueryClient());
  const mockServer = useRef<any>(FirebackMockServer);

  const isDarkMode = useColorScheme() === 'dark';
  const backgroundStyle = {
    backgroundColor: isDarkMode ? Colors.darker : Colors.lighter,
    flex: 1,
  };

  return (
    <NavigationContainer>
      <QueryClientProvider client={queryClient.current}>
        <WithFireback mockServer={mockServer} queryClient={queryClient.current}>
          <SafeAreaView style={backgroundStyle}>
            <StatusBar
              barStyle={isDarkMode ? 'light-content' : 'dark-content'}
              backgroundColor={backgroundStyle.backgroundColor}
            />
            <ApplicationRoutes />
          </SafeAreaView>
        </WithFireback>
      </QueryClientProvider>
    </NavigationContainer>
  );
}

export default App;
