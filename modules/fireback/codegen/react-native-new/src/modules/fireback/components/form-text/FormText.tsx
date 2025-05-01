import React, {useCallback, useRef, useState} from 'react';
import {
  Platform,
  StyleSheet,
  TextInput,
  TextInputProps,
  ViewStyle,
} from 'react-native';
import {BaseFormElement} from '../base-form-element/BaseFormElement';
import colors from '../../constants/colors';

export type FormTextProps = {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: string) => void;
  secureTextEntry?: boolean;
  Icon?: any;
  errorMessage?: string;
  style?: ViewStyle;
  value?: any | null;
  focused?: boolean;
  getInputRef?: (ref: any) => void;
} & TextInputProps;

export const FormText = (props: FormTextProps) => {
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    style,
    errorMessage,
    onChange,
    value,
    disabled,
    focused: f = false,
    ...restProps
  } = props;

  const [focused, setFocused] = useState(false);
  const ref = useRef<TextInput | null>();
  const onPress = useCallback(() => {
    ref.current?.focus();
  }, [ref.current]);

  return (
    <BaseFormElement focused={focused} onPress={onPress} {...props}>
      <TextInput
        {...restProps}
        ref={el => (ref.current = el)}
        style={styles.input}
        value={value}
        onChangeText={onChange}
        onBlur={() => setFocused(false)}
        onFocus={() => setFocused(true)}
        editable={disabled !== true}
        selectTextOnFocus={disabled !== true}
        placeholderTextColor={colors.placeholderText}
        secureTextEntry={secureTextEntry}
      />
    </BaseFormElement>
  );
};

const left = 10;

const styles = StyleSheet.create({
  input: {
    color: colors.primaryColor,
    fontSize: 14,
    padding: Platform.select({macos: 5, default: 0}),
    position: 'absolute',
    left: Platform.select({android: left - 4, ios: left, macos: 10}),
    top: Platform.select({android: 15, ios: 25, macos: 20}),
    right: 10,
  },
});
