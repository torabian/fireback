import React, {useCallback, useEffect, useRef, useState} from 'react';
import {Text} from 'react-native';
import PhoneInput from 'react-native-phone-number-input';
import {FormTextProps} from '../form-text/FormText';

export interface FormPhoneProps extends FormTextProps {}

export const FormPhone = (props: FormPhoneProps) => {
  const {getInputRef, value, onChange, ...restProps} = props;
  const ref = useRef<PhoneInput | null>();

  useEffect(() => {
    if (!value) {
      ref.current?.setState({number: ''});
    }
  }, [value]);
  return (
    // <BaseFormElement focused={focused} onPress={onPress} {...props}>
    <>
      <PhoneInput
        {...restProps}
        ref={el => (ref.current = el)}
        defaultValue={value}
        defaultCode="PL"
        autoFocus
        layout="second"
        value={value !== undefined ? `${value}` : 'x'}
        onChangeFormattedText={onChange}
        withShadow
      />
    </>
    // </BaseFormElement>
  );
};
