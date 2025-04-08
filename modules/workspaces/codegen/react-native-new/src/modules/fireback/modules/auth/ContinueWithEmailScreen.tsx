import {FormText} from '@/modules/fireback/components/form-text/FormText';
import {usePostWorkspacePassportCheck} from '@/modules/fireback/sdk/modules/abac/usePostWorkspacePassportCheck';
import {CheckClassicPassportActionReqDto} from '@/modules/fireback/sdk/modules/abac/WorkspacesActionsDto';
import {useNavigation} from '@react-navigation/native';
import React from 'react';
import {Alert, ScrollView, StyleSheet, Text, View} from 'react-native';
import {useTheme} from '../theme';
import EnterPasswordScreen from './EnterPasswordScreen';
import FinishSignup from './FinishSignup';
import {FormManager} from './FormManager';

const ContinueWithEmailScreen = () => {
  const {theme} = useTheme();
  const {submit, mutation} = usePostWorkspacePassportCheck({});
  const {navigate} = useNavigation<any>();

  const onSubmit = (values: CheckClassicPassportActionReqDto) => {
    submit(values)
      .then(res => {
        if (res.data?.exists) {
          navigate(EnterPasswordScreen.Name, {value: values.value});
        } else {
          navigate(FinishSignup.Name);
        }
      })
      .catch(err => {
        Alert.alert('Error: ' + err);
      });
  };

  return (
    <ScrollView style={[theme.background, styles.container]}>
      <Text style={theme?.h2}>Enter an email to get started</Text>
      <Text style={theme?.p}>
        We'll check if you have an account and provide next steps.
      </Text>
      <FormManager
        onSubmit={onSubmit}
        Form={({form, isEditing}: any) => {
          const {values, setValues, setFieldValue, errors} = form;

          return (
            <View>
              <FormText
                value={values.value}
                onChange={value =>
                  setFieldValue(
                    CheckClassicPassportActionReqDto.Fields.value,
                    value,
                    false,
                  )
                }
              />
            </View>
          );
        }}></FormManager>
    </ScrollView>
  );
};

ContinueWithEmailScreen.Name = 'ContinueWithEmailScreen';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
  },
});

export default ContinueWithEmailScreen;
