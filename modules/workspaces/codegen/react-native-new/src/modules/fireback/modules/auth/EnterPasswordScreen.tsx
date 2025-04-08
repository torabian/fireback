import { FormText } from '@/modules/fireback/components/form-text/FormText';
import { useNavigation } from '@react-navigation/native';
import { useContext } from 'react';
import { Alert, ScrollView, StyleSheet, Text, View } from 'react-native';
import { RemoteQueryContext } from '../../sdk/core/react-tools';
import { usePostPassportsSigninClassic } from '../../sdk/modules/abac/usePostPassportsSigninClassic';
import { useTheme } from '../theme';
import { UserArchiveScreen } from '../users/user/UserArchiveScreen';
import { FormManager } from './FormManager';

interface EnterPasswordScreenProps {
  route: {
    params: {
      value: string;
    };
  };
}

const EnterPasswordScreen = ({route}: EnterPasswordScreenProps) => {
  const {theme} = useTheme();
  const {submit, mutation} = usePostPassportsSigninClassic({});
  const {setSession, session, isAuthenticated} = useContext(RemoteQueryContext);
  const {navigate} = useNavigation<any>();

  const onSubmit = (data: any) => {
    submit({...data, value: route.params.value})
      .then(res => {
        setSession(res.data);
        navigate(UserArchiveScreen.Name);
      })
      .catch(err => {
        Alert.alert('Error: ' + JSON.stringify(err));
      });
  };

  return (
    <ScrollView style={[theme.background, styles.container]}>
      <Text style={theme?.h2}>Welcome back!</Text>
      <Text style={theme?.p}>Enter your password to login.</Text>

      <FormManager
        onSubmit={onSubmit}
        Form={({form, isEditing}: any) => {
          const {values, setValues, setFieldValue, errors} = form;

          return (
            <View>
              <Text>{JSON.stringify(form.values)}</Text>
              <FormText
                secureTextEntry
                value={values.password}
                onChange={value => setFieldValue('password', value, false)}
              />
            </View>
          );
        }}></FormManager>
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
