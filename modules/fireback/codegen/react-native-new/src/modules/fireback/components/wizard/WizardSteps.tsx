import {FormikProps} from 'formik';
import React, {useRef} from 'react';

import {StyleSheet, Text, TouchableOpacity, View} from 'react-native';
import colors from '~/constants/colors';
import {useKeyboardVisibility} from '~/hooks/useKeyboardVisibility';
import {WizardStep} from './WizardHelper';

export const WizardSteps = ({
  steps,
  currentIndex,
  onStepPress,
  leftOffset,
  maxStepEver,
  formik,
}: {
  steps: WizardStep[];
  currentIndex: number;
  maxStepEver: number;
  leftOffset?: number;
  onStepPress: (index: number) => void;
  formik: FormikProps<any>;
}) => {
  const refs = useRef<View[]>([]);
  const {keyboardVisible} = useKeyboardVisibility();

  return (
    <View style={[styles.wrapper, keyboardVisible && styles.wrapperCompact]}>
      {(steps || []).map((step, index) => {
        const isInvalid = step.isValid ? !step.isValid(formik) : false;
        const isDisabled = index > maxStepEver;

        return (
          <TouchableOpacity
            style={[styles.step]}
            key={step.label}
            onPress={() => onStepPress(index)}
            disabled={isDisabled}>
            <View
              ref={el => (refs.current[index] = el)}
              style={[
                styles.circle,
                index <= currentIndex && styles.completedStep,
                isInvalid && styles.stepInvalid,
                keyboardVisible && styles.circleCompact,
              ]}
            />
            <Text
              style={[styles.label, keyboardVisible && styles.labelCompact]}>
              {step.label}
            </Text>
          </TouchableOpacity>
        );
      })}
    </View>
  );
};

const circleSize = 25;

const styles = StyleSheet.create({
  completedStep: {
    backgroundColor: colors.primaryColor,
    borderColor: colors.primaryColor,
  },
  step: {
    flexDirection: 'column',
    alignItems: 'center',
  },
  stepInvalid: {
    backgroundColor: colors.error,
    borderColor: colors.error,
  },
  wrapper: {
    paddingHorizontal: 20,
    justifyContent: 'space-between',
    flexDirection: 'row',
  },
  wrapperCompact: {
    marginVertical: 5,
    marginRight: 40,
    top: 10,
    marginBottom: 20,
  },
  labelCompact: {
    fontSize: 10,
  },
  label: {
    color: colors.primaryColor,
  },
  line: {
    width: 20,
    borderBottomColor: colors.gray,
    borderBottomWidth: 2,
    marginTop: 15,
    marginHorizontal: 3,
  },
  circle: {
    width: circleSize,
    height: circleSize,
    borderRadius: circleSize / 2,
    borderWidth: 1,
    borderColor: colors.gray,
  },
  circleCompact: {
    height: 3,
    borderRadius: 0,
  },
});
