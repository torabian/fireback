import React, {useEffect, useRef, useState} from 'react';
import {
  Alert,
  Animated,
  Easing,
  Platform,
  StyleSheet,
  Text,
  TextInput,
  TextStyle,
  TouchableOpacity,
  TouchableWithoutFeedback,
  View,
  ViewStyle,
} from 'react-native';
import colors from '../../constants/colors';
import t from '../../constants/t';
import {AnimatedText} from '../animated-text/AnimatedText';
// import CrossIcon from '~/assets/icons/cross.svg';

export interface CommonFormElementProps<T> {
  value?: T;
  onChange?: (value: T) => void;
  disabled?: boolean;
  label?: string;
}

export interface BaseFormElementProps extends CommonFormElementProps<any> {
  Icon?: any;
  errorMessage?: string;
  style?: ViewStyle;
  value?: any | null;
  labelStyle?: TextStyle;
  hasAnimation?: boolean;
  focused?: boolean;
  onPress?: () => void;
  getInputRef?: (ref: any) => void;
  children?: React.ReactNode;
  displayValue?: string | null;
}

function errorMessageAsString(possibleErrorMessage: string | Object) {
  if (typeof possibleErrorMessage === 'string') {
    return possibleErrorMessage;
  }

  if (Array.isArray(possibleErrorMessage)) {
    return possibleErrorMessage.join(', ');
  }

  return JSON.stringify(possibleErrorMessage);
}

export const BaseFormElement = ({
  label,
  getInputRef,
  displayValue,
  Icon,
  style,
  children,
  errorMessage,
  value,
  onPress,
  onChange,
  focused = false,
  labelStyle,
  disabled,
  hasAnimation,
}: BaseFormElementProps) => {
  const ref = useRef<TextInput | null>();
  const fadeRef = useRef(new Animated.Value(0));

  useEffect(() => {
    if (getInputRef && ref.current) {
      getInputRef(ref.current);
    }
  }, [getInputRef, ref.current]);

  useEffect(() => {
    if (hasAnimation === false) {
      return;
    }

    Animated.timing(fadeRef.current, {
      toValue: focused || value || value === 0 ? 1 : 0,
      duration: 150,
      easing: Easing.linear,
      useNativeDriver: true,
    }).start();
  }, [focused, value, hasAnimation]);

  const position = fadeRef.current.interpolate({
    inputRange: [0, 1],
    outputRange: [0, -13],
  });

  const errorDisplayValue = errorMessageAsString(
    (errorMessage && (t as any).api[errorMessage]) || errorMessage || '',
  );

  const onDiscard = () => {
    if (onChange) onChange(null);
  };

  const liftLabel =
    (focused || value || value === 0) && hasAnimation !== false ? true : false;

  return (
    <TouchableWithoutFeedback onPress={onPress}>
      <View>
        <View style={[baseFormElementStyle.wrapper, style]}>
          {Icon ? (
            <View style={{width: 30}}>
              <Icon style={baseFormElementStyle.icon} />
            </View>
          ) : null}
          <View
            style={{
              flex: 1,
              height: 60,
            }}>
            {Platform.OS === 'macos' ? (
              <Text
                style={[
                  baseFormElementStyle.label,
                  labelStyle,
                  {top: liftLabel ? 0 : 20},
                ]}>
                {label}
              </Text>
            ) : (
              <Animated.Text
                style={[
                  baseFormElementStyle.label,
                  labelStyle,
                  {transform: [{translateY: position}]},
                ]}>
                {label}
              </Animated.Text>
            )}

            {children}
            {displayValue ? (
              <Text style={baseFormElementStyle.input}>{displayValue}</Text>
            ) : null}
            {value && disabled !== true ? (
              <TouchableOpacity
                style={baseFormElementStyle.modalCrossIconWrapper}
                hitSlop={{bottom: 20, right: 20, top: 20, left: 20}}
                onPress={onDiscard}>
                {/* <CrossIcon
                  color={colors.placeholderText}
                  style={baseFormElementStyle.modalCrossIcon}
                /> */}
              </TouchableOpacity>
            ) : null}
          </View>
        </View>

        <AnimatedText style={baseFormElementStyle.errorMessage}>
          {errorDisplayValue}
        </AnimatedText>
      </View>
    </TouchableWithoutFeedback>
  );
};

const left = 10;

export const baseFormElementStyle = StyleSheet.create({
  errorMessage: {
    color: colors.error,
    fontSize: 11,
    minHeight: 15,
  },
  label: {
    left,
    top: 23,
    fontSize: 13,
    color: colors.placeholder,
  },
  icon: {
    width: 20,
    height: 20,
  },
  wrapper: {
    borderWidth: 1,
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: colors.inputBackground,
    padding: 10,
    paddingBottom: 0,
    height: 60,
    borderRadius: 10,
    marginVertical: 10,
  },
  input: {
    color: colors.primaryColor,
    fontSize: 14,
    position: 'absolute',
    left: Platform.select({android: left - 4, ios: left}),
    top: Platform.select({android: 15, ios: 25}),
    right: 10,
  },
  modalCrossIconWrapper: {
    position: 'absolute',
    right: 0,
    top: 20,
  },
  modalCrossIcon: {
    width: 50,
    height: 50,
  },
});
