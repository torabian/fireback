import React from 'react';
import {StyleSheet, Text, View, ViewStyle} from 'react-native';
import colors from '~/constants/colors';
import {ApiError} from '~/interfaces/Error';
import {AnimatedText} from '../animated-text/AnimatedText';

export const ErrorsView = ({
  errors,
  error,
  style,
}: {
  error?: ApiError;
  errors?: any;
  style?: ViewStyle;
}) => {
  if (!error && !errors) {
    return null;
  }

  let errorList: any = {};

  if (error && error.errors) {
    errorList = error.errors;
  } else if (errors) {
    errorList = errors;
  }

  const keys = Object.keys(errorList);

  if (keys?.length === 0 && !(error?.title || error?.message)) {
    return null;
  }

  return (
    <React.Fragment>
      {errors.form && (
        <AnimatedText style={styles.text}>{errors.form}</AnimatedText>
      )}
      {errorList.length && (
        <View style={[styles.errorsViewWrapper, style]}>
          {(error?.title || error?.message) && (
            <Text>{error?.title || error?.message}</Text>
          )}
          {keys.map((key: string) => {
            return (
              <View key={key}>
                <Text>&bull; {(errorList as any)[key]}</Text>
              </View>
            );
          })}
        </View>
      )}
    </React.Fragment>
  );
};

const styles = StyleSheet.create({
  errorsViewWrapper: {
    padding: 10,
    backgroundColor: colors.darkBackground,
  },
  text: {
    color: colors.error,
    textAlign: 'center',
  },
});
