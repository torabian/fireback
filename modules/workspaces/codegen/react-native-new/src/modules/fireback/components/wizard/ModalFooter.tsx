import {FormikErrors} from 'formik';
import React, {useEffect, useState} from 'react';
import {StyleSheet, View} from 'react-native';
import colors from '~/constants/colors';
import {useKeyboardVisibility} from '~/hooks/useKeyboardVisibility';
import {FormButton} from '../form-button/FormButton';

export const ModalFooter = ({
  submitLabel,
  onCancel,
  isSubmitting,
  onSubmit,
  isValid,
  errors,
  cancelLabel,
}: {
  submitLabel?: string;
  isValid?: boolean;
  cancelLabel?: string;
  errors?: FormikErrors<any>;
  onCancel?: () => void;
  onSubmit?: () => void;
  isSubmitting?: boolean;
}) => {
  const normalLabel = submitLabel || 'Confirm';
  const [label, setLabel] = useState('');

  const {keyboardVisible} = useKeyboardVisibility();
  const errorsLength = errors ? Object.keys(errors).length : 0;

  const indicateError = () => {
    if (errorsLength === 0) {
      setLabel(normalLabel);
      return;
    }

    setLabel(`${errorsLength} errors`);
    setTimeout(() => {
      setLabel(normalLabel);
    }, 3000);
  };

  useEffect(() => {
    indicateError();
  }, [errors]);

  return (
    <View style={[styles.wrapper, keyboardVisible && styles.wrapperCompact]}>
      <View style={[styles.action]}>
        <FormButton
          // buttonStyles={[
          //   styles.button,
          //   keyboardVisible && styles.buttonCompact,
          // ]}
          onPress={onCancel}
          label={cancelLabel || 'Cancel'}
        />
      </View>
      <View style={[styles.action]}>
        <FormButton
          // isActive={isValid !== false && errorsLength === 0}
          disabled={isValid === false || errorsLength > 0}
          // loading={isSubmitting}
          // buttonStyles={[
          //   styles.button,
          //   keyboardVisible && styles.buttonCompact,
          // ]}
          onPress={onSubmit}
          label={label || normalLabel}
        />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  action: {flex: 1},
  button: {height: 40, borderRadius: 20},
  buttonCompact: {height: 25},
  wrapper: {
    height: 60,
    flexDirection: 'row',
    backgroundColor: colors.gray,
    padding: 10,
  },
  wrapperCompact: {
    height: 35,
    paddingVertical: 3,
  },
});
