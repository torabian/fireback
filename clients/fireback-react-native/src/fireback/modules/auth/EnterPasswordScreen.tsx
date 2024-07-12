import {FormText} from '@/fireback/components/form-text/FormText';
import {usePostWorkspacePassportCheck} from '@/sdk/fireback/modules/workspaces/usePostWorkspacePassportCheck';
import React from 'react';
import {Alert, ScrollView, StyleSheet, Text, View} from 'react-native';
import Button from '../../components/button/Button';
import {useTheme} from '../theme';
import {FormManager} from './FormManager';

const EnterPasswordScreen = () => {
  const {theme} = useTheme();
  const {submit, mutation} = usePostWorkspacePassportCheck({});

  const onSubmit = () => {
    submit({value: 'adasd'})
      .then(res => {
        res.data?.exists;
        Alert.alert('Res' + JSON.stringify(res));
      })
      .catch(err => {
        Alert.alert('Error: ' + err);
      });
  };

  return (
    <ScrollView style={[theme.background, styles.container]}>
      <Text style={theme?.h2}>Welcome back!</Text>
      <Text style={theme?.p}>Enter your password to login.</Text>

      <FormManager
        Form={({form, isEditing}: any) => {
          const {values, setValues, setFieldValue, errors} = form;

          return (
            <View>
              <Text>{JSON.stringify(form.values)}</Text>
              <FormText
                secureTextEntry
                value={values.title}
                onChange={value => setFieldValue('title', value, false)}
              />
            </View>
          );
        }}></FormManager>

      <Button title="Continue" onPress={onSubmit} />
    </ScrollView>
  );
};

EnterPasswordScreen.Name = 'EnterPasswordScreen';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
  },
});

export default EnterPasswordScreen;
