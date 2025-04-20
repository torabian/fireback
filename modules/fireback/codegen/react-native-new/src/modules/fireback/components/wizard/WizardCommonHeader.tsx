import React from 'react';

import {StyleSheet, Text, View} from 'react-native';
import colors from '~/constants/colors';
import {useKeyboardVisibility} from '~/hooks/useKeyboardVisibility';

export const WizardCommonHeader = ({
  title,
  description,
  Icon,
}: {
  title: string;
  description?: string;
  Icon?: any;
}) => {
  const {keyboardVisible} = useKeyboardVisibility();

  return (
    <View>
      <View style={styles.wrapper}>
        {Icon && <Icon width={60} height={60} />}
        {!keyboardVisible && <Text style={styles.title}>{title}</Text>}
        {description && <Text style={styles.description}>{description}</Text>}
      </View>
      <View
        style={[styles.divider, keyboardVisible && styles.dividerCompact]}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  dividerCompact: {
    marginVertical: 5,
  },
  divider: {
    borderBottomColor: colors.gray,
    borderBottomWidth: 1,
    marginVertical: 20,
  },
  wrapper: {
    marginVertical: 30,
    alignContent: 'center',
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    marginTop: 10,
    fontSize: 24,
    marginHorizontal: 30,
    textAlign: 'center',
  },
  description: {
    fontSize: 15,
    color: colors.grayText,
    marginTop: 20,
    marginHorizontal: 50,
    textAlign: 'center',
  },
});
