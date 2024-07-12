import React from 'react';

import {
  View,
  StyleSheet,
  StyleProp,
  ViewStyle,
  TouchableOpacity,
  TouchableHighlightProps,
  Pressable,
} from 'react-native';

export const Card = ({
  children,
  style,
  onPress,
  ...props
}: {
  children: React.ReactNode;
  style?: StyleProp<ViewStyle>;
  onPress?: () => void;
} & TouchableHighlightProps) => {
  const ELEMENT: any = onPress ? Pressable : View;

  return (
    <ELEMENT {...props} onPress={onPress} style={[styles.wrapper, style]}>
      {children}
    </ELEMENT>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    backgroundColor: 'white',
    borderRadius: 5,
    padding: 20,
    shadowColor: '#00000010',
    shadowOffset: {width: 0, height: 1},
    marginVertical: 5,
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
  },
});
