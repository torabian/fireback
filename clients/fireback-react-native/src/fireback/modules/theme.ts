import {useState} from 'react';
import {StyleSheet, TextStyle} from 'react-native';

export interface Theme {
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

export const themeLight = StyleSheet.create<Theme>({
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

export const themeDark = StyleSheet.create<Theme>({
  p: {
    marginBottom: 20,
  },
  h2: {
    color: 'black',
    fontSize: 22,
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

export function useTheme() {
  const [theme, setTheme] = useState<Theme>(themeLight);

  return {theme, setTheme};
}
