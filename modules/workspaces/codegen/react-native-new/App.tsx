import 'react-native-gesture-handler';

import React, {useRef} from 'react';
import {SafeAreaView, StatusBar, useColorScheme} from 'react-native';

import {WithFireback} from '@/apps/core/WithFireback';
import {Colors} from 'react-native/Libraries/NewAppScreen';
import {QueryClient, QueryClientProvider} from 'react-query';

import {NavigationContainer} from '@react-navigation/native';
import {AuthRouter} from '@/modules/fireback/modules/auth/Router';

function App(): React.JSX.Element {
  const queryClient = useRef(new QueryClient());

  const isDarkMode = useColorScheme() === 'dark';

  const backgroundStyle = {
    backgroundColor: isDarkMode ? Colors.darker : Colors.lighter,
    flex: 1,
  };

  return (
    <NavigationContainer>
      <QueryClientProvider client={queryClient.current}>
        <WithFireback queryClient={queryClient.current}>
          <SafeAreaView style={backgroundStyle}>
            <StatusBar
              barStyle={isDarkMode ? 'light-content' : 'dark-content'}
              backgroundColor={backgroundStyle.backgroundColor}
            />
            {/* Put component, router, stack, etc here, or make it conditinail based on 
            using authentication situation */}
            <AuthRouter />
          </SafeAreaView>
        </WithFireback>
      </QueryClientProvider>
    </NavigationContainer>
  );
}

export default App;
