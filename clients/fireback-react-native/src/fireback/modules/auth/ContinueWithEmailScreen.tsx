import {FormText} from '@/fireback/components/form-text/FormText';
import {usePostWorkspacePassportCheck} from '@/sdk/fireback/modules/workspaces/usePostWorkspacePassportCheck';
import React from 'react';
import {Alert, ScrollView, StyleSheet, Text, View} from 'react-native';
import Button from '../../components/button/Button';
import {useTheme} from '../theme';
import {FormManager} from './FormManager';
import {useNavigation} from '@react-navigation/native';
import EnterPasswordScreen from './EnterPasswordScreen';
import {CheckClassicPassportActionReqDto} from '@/sdk/fireback/modules/workspaces/WorkspacesActionsDto';
import FinishSignup from './FinishSignup';

const ContinueWithEmailScreen = () => {
  const {theme} = useTheme();
  const {submit, mutation} = usePostWorkspacePassportCheck({});
  const {navigate} = useNavigation<any>();

  const onSubmit = (values: CheckClassicPassportActionReqDto) => {
    submit(values)
      .then(res => {
        if (res.data?.exists) {
          navigate(EnterPasswordScreen.Name);
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
