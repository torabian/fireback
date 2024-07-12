import React from 'react';
import {
  View,
  Text,
  TextInput,
  Button,
  TouchableOpacity,
  StyleSheet,
  Alert,
} from 'react-native';

const SignInScreen = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Sign In</Text>
      <Text style={styles.description}>
        Please enter your username and password to continue.
      </Text>

      <TextInput placeholder="Username" style={styles.input} />

      <TextInput placeholder="Password" secureTextEntry style={styles.input} />

      <TouchableOpacity onPress={() => Alert.alert('Forgot Password?')}>
        <Text style={styles.forgotPassword}>Forgot Password?</Text>
      </TouchableOpacity>

      <Button title="Submit" onPress={() => Alert.alert('Submitted!')} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
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

export default SignInScreen;
