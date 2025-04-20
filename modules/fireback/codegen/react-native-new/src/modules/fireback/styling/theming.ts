/*
 * Here in fireback we define a theme interface, which you can implement and fireback
 * and components from it will follow. Feel free to edit this on your useTheme version,
 * or extend it. If you are not using fireback module folder, you can completely ignore this.
 **/
import {StyleSheet, TextStyle} from 'react-native';

export interface ITheme {
  h2: TextStyle;
  p: TextStyle;
  background: {
    backgroundColor: string;
  };
  border: {
    borderRadius: number;
    color: string;
    borderWidth: number;
  };
}
