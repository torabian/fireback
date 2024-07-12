import React from 'react';
import {Image, ScrollView, StyleSheet, Text} from 'react-native';
import Button from '../../components/button/Button';
import {themeLight, useTheme} from '../theme';
import {useNavigation} from '@react-navigation/native';
import ContinueWithEmailScreen from './ContinueWithEmailScreen';

const AuthWelcomeScreen = () => {
  const {theme, setTheme} = useTheme();
  const {navigate} = useNavigation<any>();

  return (
    <ScrollView style={[theme.background, styles.container]}>
      <Image
        style={{
          width: 412 * 0.3,
          height: 489 * 0.3,
          alignSelf: 'center',
          marginBottom: 30,
        }}
        source={require('../../assets/auth_logo.png')}
      />
      <Text style={theme?.h2}>
        Use an account to sync your data across your devices
      </Text>
      <Text style={theme?.p}>
        You can use multiple methods of authenticating and storing your
        information in the cloud
      </Text>

      <Button
        title="Continue with Google"
        icon={require('./assets/google.png')}
        onPress={() => setTheme(themeLight)}
      />
      <Button
        title="Continue with Apple"
        icon={require('./assets/apple.png')}
        onPress={() => setTheme(themeLight)}
      />
      <Button
        title="Continue with Facebook"
        icon={require('./assets/facebook.png')}
        onPress={() => setTheme(themeLight)}
      />
      <Button
        title="Continue with Phone"
        icon={require('./assets/phone.png')}
        onPress={() => setTheme(themeLight)}
      />
      <Button
        title="Continue with Email"
        icon={require('./assets/email.png')}
        onPress={() => navigate(ContinueWithEmailScreen.Name)}
      />
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  h2: {
    fontSize: 22,
    fontWeight: 'bold',
  },
  container: {
    flex: 1,
    padding: 16,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 8,
    textAlign: 'center',
  },
  description: {
    fontSize: 16,
    marginBottom: 24,
    textAlign: 'center',
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    borderRadius: 4,
    marginBottom: 16,
    paddingHorizontal: 8,
  },
  forgotPassword: {
    color: 'blue',
    marginBottom: 24,
    textAlign: 'center',
  },
});

export default AuthWelcomeScreen;
