import 'react-native-gesture-handler';

import React, { useRef } from 'react';
import { SafeAreaView, StatusBar, useColorScheme } from 'react-native';

import { WithFireback } from '@/apps/core/WithFireback';
import { Colors } from 'react-native/Libraries/NewAppScreen';
import { QueryClient, QueryClientProvider } from 'react-query';

import AuthWelcomeScreen from '@/fireback/modules/auth/AuthWelcomeScreen';
import ContinueWithEmailScreen from '@/fireback/modules/auth/ContinueWithEmailScreen';


import EnterPasswordScreen from '@/fireback/modules/auth/EnterPasswordScreen';
import FinishSignup from '@/fireback/modules/auth/FinishSignup';
import { createStackNavigator, TransitionPresets } from '@react-navigation/stack';

 const Stack = createStackNavigator();

import { NavigationContainer } from '@react-navigation/native';


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

            <Stack.Navigator
              screenOptions={{
                ...TransitionPresets.SlideFromRightIOS, // or another preset like FadeFromBottomAndroid
              }}>
              <Stack.Screen
                name={'AuthWelcomeScreen'}
                options={{headerShown: false}}
                component={AuthWelcomeScreen}
              />
              <Stack.Screen
                options={{headerTitle: 'Login or signup'}}
                name={ContinueWithEmailScreen.Name}
                component={ContinueWithEmailScreen}
              />
              <Stack.Screen
                options={{headerTitle: 'Login'}}
                name={EnterPasswordScreen.Name}
                component={EnterPasswordScreen}
              />
              <Stack.Screen
                options={{headerTitle: 'Finish signing up'}}
                name={FinishSignup.Name}
                component={FinishSignup}
              />
            </Stack.Navigator>

            {/* <Drawer.Navigator
              screenOptions={{headerTitle: '', headerShown: false}}>
              <Drawer.Screen name="Sigin" component={AuthWelcomeScreen} />
              <Drawer.Screen
                name={ContinueWithEmailScreen.Name}
                component={ContinueWithEmailScreen}
              />
            </Drawer.Navigator> */}
          </SafeAreaView>
        </WithFireback>
      </QueryClientProvider>
    </NavigationContainer>
  );
}

export default App;
