import 'react-native-gesture-handler';

import React, {useRef} from 'react';
import {SafeAreaView, StatusBar, useColorScheme} from 'react-native';

import {Colors} from 'react-native/Libraries/NewAppScreen';
import {CustomerArchiveScreen} from './src/modules/customers/CustomerArchiveScreen';
import {QueryClient, QueryClientProvider} from 'react-query';
import {WithFireback} from '@/apps/core/WithFireback';

import {createDrawerNavigator} from '@react-navigation/drawer';
import {NavigationContainer} from '@react-navigation/native';
import SignInScreen from '@/modules/auth/SigninScreen';

const Drawer = createDrawerNavigator();

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
            <Drawer.Navigator>
              <Drawer.Screen name="Feed" component={CustomerArchiveScreen} />
              <Drawer.Screen name="Sigin" component={SignInScreen} />
            </Drawer.Navigator>
          </SafeAreaView>
        </WithFireback>
      </QueryClientProvider>
    </NavigationContainer>
  );
}

export default App;
