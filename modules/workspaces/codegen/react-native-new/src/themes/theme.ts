import {ITheme} from '@/modules/fireback/styling/theming';
import {StyleSheet} from 'react-native';

export const theme = {
  activeColor: '#00ff00',
  inactiveColor: '#ff0000',
  idleColor: '#ffff00',
};

export const themeEl = StyleSheet.create({
  textLabel: {
    fontWeight: 'bold',
    marginRight: 5,
  },
  keyPairRow: {
    flexDirection: 'row',
  },
  h1: {
    fontSize: 26,
    fontWeight: 'bold',
    marginLeft: 10,
    marginRight: 10,
    marginTop: 10,
    marginBottom: 16,
  },
});

/// New theming system

export const themeLight = StyleSheet.create<ITheme>({
  h2: {
    color: 'black',
    fontSize: 19,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  p: {
    marginBottom: 20,
  },
  background: {
    backgroundColor: 'white',
  },
  border: {
    borderRadius: 25,
    color: 'black',
    borderWidth: 1,
  },
});

export const themeDark = StyleSheet.create<ITheme>({
  p: {
    marginBottom: 20,
    color: 'white',
  },
  h2: {
    color: 'white',
    fontSize: 19,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  background: {
    backgroundColor: 'black',
  },
  border: {
    borderRadius: 25,
    color: 'white',
    borderWidth: 1,
  },
});
