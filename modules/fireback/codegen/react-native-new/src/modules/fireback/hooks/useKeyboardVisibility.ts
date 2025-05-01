import {useCallback, useEffect, useState} from 'react';
import {Dimensions, Keyboard, KeyboardEvent} from 'react-native';

export const useKeyboardVisibility = () => {
  const [isKeyboadVisible, setIsKeyboadVisible] = useState({
    keyboardVisible: false,
    height: 0,
    viewportHeight: Dimensions.get('screen').height,
  });

  const _keyboardDidShow = useCallback((e: KeyboardEvent) => {
    setIsKeyboadVisible({
      keyboardVisible: true,
      height: e.endCoordinates.height,
      viewportHeight: Dimensions.get('window').height - e.endCoordinates.height,
    });
  }, []);

  const _keyboardDidHide = useCallback(() => {
    setIsKeyboadVisible({
      keyboardVisible: false,
      height: 0,
      viewportHeight: Dimensions.get('screen').height,
    });
  }, []);

  useEffect(() => {
    Keyboard.addListener('keyboardDidShow', _keyboardDidShow);
    Keyboard.addListener('keyboardDidHide', _keyboardDidHide);

    return () => {
      Keyboard.addListener('keyboardDidShow', _keyboardDidShow);
      Keyboard.addListener('keyboardDidHide', _keyboardDidHide);
    };
  }, [_keyboardDidHide, _keyboardDidShow]);

  return isKeyboadVisible;
};
