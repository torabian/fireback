import React from 'react';
import {Platform, ScrollView, StyleSheet, ViewStyle} from 'react-native';
import {ListItem} from '~/interfaces/UI';
import {FormButton} from '../form-button/FormButton';

export const TagsList = ({
  value,
  multiple,
  items,
  setValue,
  forceSelected,
  style,
  itemStyle,
}: {
  value: string[] | string;
  items: ListItem[];
  multiple?: boolean;
  setValue: (stringArr: string[]) => void;
  forceSelected?: boolean;
  itemStyle?: ViewStyle;
  style?: ViewStyle;
}) => {
  return (
    <ScrollView
      showsHorizontalScrollIndicator={false}
      horizontal
      style={[styles.buttons, style]}>
      {items.map(item => {
        const isActive = value.includes(item.value);
        return (
          <FormButton
            size="small"
            key={item.value}
            style={styles.button}
            label={item.label}
            disabled={item.disabled}
            onPress={() => {
              if (multiple && isActive && Array.isArray(value)) {
                setValue(value.filter(val => val !== item.value));
              } else if (multiple && !isActive) {
                setValue([...value, item.value]);
              } else if (isActive && !forceSelected) {
                setValue([]);
              } else {
                setValue([item.value]);
              }
            }}
          />
        );
      })}
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  button: {
    marginRight: 5,
  },
  stateItem: {
    width: 'auto',
  },
  containerStyle: {flex: 1},
});
