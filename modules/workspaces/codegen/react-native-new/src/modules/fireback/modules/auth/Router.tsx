import 'react-native-gesture-handler';

import React from 'react';

import AuthWelcomeScreen from '@/modules/fireback/modules/auth/AuthWelcomeScreen';
import ContinueWithEmailScreen from '@/modules/fireback/modules/auth/ContinueWithEmailScreen';

import EnterPasswordScreen from '@/modules/fireback/modules/auth/EnterPasswordScreen';
import FinishSignup from '@/modules/fireback/modules/auth/FinishSignup';
import {TransitionPresets, createStackNavigator} from '@react-navigation/stack';
const Stack = createStackNavigator();

export const AuthRouter = () => {
  return (
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
  );
};
