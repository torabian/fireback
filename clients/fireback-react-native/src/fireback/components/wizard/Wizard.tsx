import {FormikProps} from 'formik';
import React, {useEffect, useState} from 'react';
import {ScrollView, StyleSheet, View, ViewStyle} from 'react-native';
import Toast from 'react-native-toast-message';
import colors from '~/constants/colors';
import {useKeyboardVisibility} from '~/hooks/useKeyboardVisibility';
import {WizardActions} from './WizardActions';
import {WizardStep} from './WizardHelper';
import {WizardSteps} from './WizardSteps';

export const Wizard = ({
  steps,
  formik,
  data,
  contentStyle,
  divider,
  Header,
  submitLabel,
  leftOffset,
  requestClose,
  shouldGoToStart,
}: {
  steps: WizardStep[];
  formik: FormikProps<any>;
  data?: any;
  contentStyle?: ViewStyle;
  divider?: false;
  Header?: any;
  submitLabel?: string;
  leftOffset?: number;
  requestClose?: () => void;
  shouldGoToStart?: boolean;
}) => {
  const [step, setStep] = useState(0);
  const [maxStepEver, setMaxStepUserWent] = useState(0);

  const goNext = () => {
    setStep(s => (s < steps.length ? s + 1 : steps.length - 1));
  };

  useEffect(() => {
    setMaxStepUserWent(s => (step > s ? step : s));
  }, [step]);

  const {keyboardVisible} = useKeyboardVisibility();

  const CurrentStepComponent: any = steps[step].component;

  const goToStart = () => {
    setStep(0);
  };
  const currentStepNeedsFormValidation = steps[step]?.formNeedsValidation;

  const goPrev = () => {
    if (step === 0 && requestClose) {
      requestClose();
      return;
    }
    setStep(s => (s > 1 ? s - 1 : 0));
  };

  const onStepPress = (index: number) => {
    setStep(index);
  };

  return (
    <>
      {Header && !keyboardVisible && <Header />}
      {steps.length > 1 && (
        <WizardSteps
          formik={formik}
          leftOffset={leftOffset}
          onStepPress={onStepPress}
          maxStepEver={maxStepEver}
          currentIndex={step}
          steps={steps}
        />
      )}
      <ScrollView style={styles.wrapper}>
        {divider !== false ? <View style={styles.divider} /> : null}

        <View style={[styles.content, contentStyle]}>
          {CurrentStepComponent && (
            <CurrentStepComponent
              data={data}
              form={formik}
              values={formik.values}
            />
          )}
        </View>
      </ScrollView>
      <WizardActions
        submit={() => {
          formik.submitForm().then(() => {
            if (shouldGoToStart) {
              goToStart();
              formik.resetForm({});
            }
            Toast.show({
              type: 'info',
              text1: 'Success',
              text2: 'Architectural request has been saved successfully.',
            });
          });
        }}
        step={step}
        steps={steps}
        canCancelOnFirstStep={!!requestClose}
        isSubmitting={formik.isSubmitting}
        goNext={goNext}
        formNeedsValidation={currentStepNeedsFormValidation}
        submitLabel={submitLabel}
        goPrev={goPrev}
        formikProps={formik}
      />
    </>
  );
};

const styles = StyleSheet.create({
  wrapper: {flex: 1},
  content: {paddingHorizontal: 20},
  divider: {
    marginVertical: 30,
    borderBottomWidth: 2,
    borderBottomColor: colors.gray,
  },
});
