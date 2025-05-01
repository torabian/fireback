import React from 'react';
import {StyleSheet, View} from 'react-native';
import colors from '~/constants/colors';
import {useKeyboardVisibility} from '~/hooks/useKeyboardVisibility';
import {FormButton} from '../form-button/FormButton';

import {WizardStep} from './WizardHelper';

export const WizardActions = ({
  step,
  steps,
  submitLabel,
  goPrev,
  submit,
  goNext,
  formNeedsValidation,
  canCancelOnFirstStep,
  isSubmitting,
  formikProps,
}: {
  step: number;
  steps: WizardStep[];
  goNext: () => void;
  goPrev: () => void;
  submit: () => void;
  submitLabel?: string;
  canCancelOnFirstStep?: boolean;
  isSubmitting?: boolean;
  formNeedsValidation?: boolean;
  formikProps?: any;
}) => {
  const isLastStep = step === steps.length - 1;
  const {keyboardVisible} = useKeyboardVisibility();

  const checkIfFormIsValid = (): void => {
    if (formNeedsValidation) {
      formikProps.validateForm().then((values: any) => {
        if (Object.keys(values).length === 0) {
          goNext();
        }
      });
    } else {
      goNext();
    }
  };

  return (
    <View style={[styles.wrapper, keyboardVisible && styles.wrapperCompact]}>
      <View style={[styles.action]}>
        <FormButton
          disabled={(step === 0 && !canCancelOnFirstStep) || isSubmitting}
          // buttonStyles={[
          //   styles.button,
          //   keyboardVisible && styles.buttonCompact,
          // ]}
          onPress={goPrev}
          label={step === 0 && canCancelOnFirstStep ? 'Cancel' : 'Previous'}
        />
      </View>
      <View style={[styles.action]}>
        <FormButton
          isSubmitting={isSubmitting}
          // buttonStyles={[
          //   styles.button,
          //   keyboardVisible && styles.buttonCompact,
          // ]}
          onPress={isLastStep ? submit : checkIfFormIsValid}
          label={isLastStep ? submitLabel || 'Complete' : 'Next'}
          // testID={testIds.navigation.wizardButtons.next}
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
    padding: 10,
  },
  wrapperCompact: {
    height: 35,
    paddingVertical: 3,
  },
});
