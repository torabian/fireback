import {useState} from 'react';
import {ITheme} from '../styling/theming';
import {themeLight} from '@/themes/theme';

export function useTheme() {
  const [theme, setTheme] = useState<ITheme>(themeLight);

  return {theme, setTheme};
}
