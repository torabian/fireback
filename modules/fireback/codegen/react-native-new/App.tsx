import 'react-native-gesture-handler';

import React, {useRef} from 'react';
import {SafeAreaView, StatusBar, useColorScheme} from 'react-native';

import {WithFireback} from '@/apps/core/WithFireback';
import {QueryClient, QueryClientProvider} from 'react-query';

import {FirebackMockServer} from '@/modules/fireback/mock-server';
import {NavigationContainer} from '@react-navigation/native';
import {ApplicationRoutes} from './ApplicationRoutes';
import {BottomSheetProvider} from '@/modules/fireback/hooks/BottomSheetProvider';
import {GestureHandlerRootView} from 'react-native-gesture-handler';

function App(): React.JSX.Element {
  const queryClient = useRef(new QueryClient());
  const mockServer = useRef<any>(FirebackMockServer);

  const isDarkMode = useColorScheme() === 'dark';
  const backgroundStyle = {
    // backgroundColor: 'black',
    // backgroundColor: isDarkMode ? 'Colors.darker : Colors.lighter',
    flex: 1,
  };

  return (
    <NavigationContainer>
      <GestureHandlerRootView style={{flex: 1}}>
        <QueryClientProvider client={queryClient.current}>
          <WithFireback
            mockServer={mockServer}
            queryClient={queryClient.current}>
            <SafeAreaView style={backgroundStyle}>
              <BottomSheetProvider>
                <StatusBar
                  barStyle={isDarkMode ? 'light-content' : 'dark-content'}
                  backgroundColor={backgroundStyle.backgroundColor}
                />
                <ApplicationRoutes />
              </BottomSheetProvider>
            </SafeAreaView>
          </WithFireback>
        </QueryClientProvider>
      </GestureHandlerRootView>
    </NavigationContainer>
  );
}

export default App;
