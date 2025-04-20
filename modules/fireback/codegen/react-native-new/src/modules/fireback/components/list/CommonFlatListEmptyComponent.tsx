import AnimatedLottieView from 'lottie-react-native';
import React from 'react';
import {StyleSheet, Text, View, ViewStyle} from 'react-native';

import emptyAnimation from '~/assets/animations/empty-list.json';
import errorAnimation from '~/assets/animations/error.json';
import colors from '~/constants/colors';
import t from '~/constants/t';

export const CommonFlatListEmptyComponent = ({
  response,
  style,
}: {
  response?: any;
  style?: ViewStyle;
}) => {
  const hasError = !!response?.message;
  const message = hasError ? response.message : t.flatListCommon.empty;

  return (
    <View style={[styles.emptyCommon, style]}>
      <AnimatedLottieView
        source={hasError ? errorAnimation : emptyAnimation}
        autoPlay={true}
        loop={false}
        style={styles.animation}
      />

      <Text style={styles.header}>{message}</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  animation: {
    width: 160,
    height: 160,
  },

  header: {
    color: colors.grayText,
    fontSize: 16,
    marginTop: 30,
  },
  emptyCommon: {
    flex: 1,
    alignItems: 'center',
    paddingVertical: 50,
  },
});
