import React from 'react';
import {StyleSheet, Text, View} from 'react-native';
import {ListItem} from '~/interfaces/UI';

export const RowValue = (props: ListItem) => {
  return (
    <View style={styles.row}>
      <Text style={styles.label}>{props.label}</Text>
      <Text>{props.value}</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  row: {flexDirection: 'row', justifyContent: 'space-between'},
  label: {fontWeight: 'bold'},
});
