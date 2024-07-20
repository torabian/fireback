import React from 'react';
import {StyleSheet, Text, TextStyle, View, ViewStyle} from 'react-native';

import colors from '~/constants/colors';

export function PageTitle({
  title,
  children,
  style,
  textStyle,
  descriptionStyle,
  description,
  SideAction,
}: {
  style?: ViewStyle;
  title?: string;
  children?: any;
  description?: string;
  textStyle?: TextStyle;
  descriptionStyle?: TextStyle;
  SideAction?: any;
}) {
  return (
    <View style={[styles.container, style]}>
      <View style={[styles.withSideAction]}>
        <View>
          {title && <Text style={[styles.text, textStyle]}>{title}</Text>}
          {description && (
            <Text style={[styles.description, descriptionStyle]}>
              {description}
            </Text>
          )}
        </View>
        {SideAction && <SideAction />}
      </View>
      <View>{children}</View>
    </View>
  );
}

const styles = StyleSheet.create({
  withSideAction: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: 15,
  },
  container: {
    padding: 10,
    backgroundColor: 'white',
    borderBottomWidth: 1,
    borderColor: colors.white,
  },
  description: {
    fontSize: 14,
    marginTop: 5,
    marginHorizontal: 10,
  },
  text: {
    fontSize: 24,
    marginHorizontal: 10,
    marginTop: 10,
    color: colors.black,
    fontWeight: 'bold',
  },
});
