import React from 'react';

import {StyleSheet, Text, View} from 'react-native';

export const WizardHeader = () => {
  return (
    <View style={styles.wrapper}>
      <Text style={styles.title}>Wizard</Text>
      <Text style={styles.description}>
        Please fill out steps and information below.
      </Text>
    </View>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    marginVertical: 50,
    alignContent: 'center',
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    fontSize: 30,
  },
  description: {fontSize: 16, marginHorizontal: 50, textAlign: 'center'},
});
