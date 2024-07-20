import React from 'react';

import {
  StyleSheet,
  TouchableOpacity,
  Text,
  ViewStyle,
  TextStyle,
  ActivityIndicator,
} from 'react-native';
import colors from '~/constants/colors';

export const FormButton = ({
  style,
  textStyle,
  isSubmitting,
  label,
  type,
  size,
  onPress,
  disabled,
  Icon,
}: {
  onPress?: () => void;
  label: string;
  isSubmitting?: boolean;
  disabled?: boolean;
  style?: ViewStyle;
  textStyle?: TextStyle;
  Icon?: React.FC;
  size?: 'default' | 'small';
  type?: 'primary' | 'secondary';
}) => {
  return (
    <TouchableOpacity
      onPress={onPress}
      disabled={disabled}
      style={[
        styles.wrapper,
        type === 'secondary' && styles.secondaryWrapper,
        size === 'small' && styles.wrapperSmall,
        disabled && styles.wrapperDisabled,
        style,
      ]}>
      {isSubmitting && (
        <ActivityIndicator
          color={type === 'secondary' ? colors.primaryColor : colors.white}
        />
      )}
      {Icon ? <Icon /> : null}
      <Text
        style={[
          styles.label,
          type === 'secondary' && styles.secondaryLabel,
          textStyle,
        ]}>
        {label}
      </Text>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    backgroundColor: colors.primaryColor,
    borderRadius: 15,
    height: 50,
    marginVertical: 10,
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
  },
  wrapperSmall: {height: 20, borderRadius: 3},
  secondaryWrapper: {
    backgroundColor: colors.white,
    borderColor: colors.primaryColor,
    borderWidth: 1,
  },
  label: {
    textAlign: 'center',
    color: colors.white,
    fontWeight: 'bold',
    marginHorizontal: 10,
  },
  secondaryLabel: {
    color: colors.primaryColor,
  },
  wrapperDisabled: {
    backgroundColor: colors.graySwitchBg,
  },
});
