import React, {useEffect, useRef, useState} from 'react';
import {Animated, Easing, TextProps, TextStyle} from 'react-native';

export function AnimatedText({
  children,
  style,
}: {
  children: string;
} & TextProps) {
  const [current, setCurrent] = useState('');
  const fadeRef = useRef(new Animated.Value(0));

  useEffect(() => {
    Animated.timing(fadeRef.current, {
      toValue: 0,
      duration: 300,
      easing: Easing.linear,
      useNativeDriver: true,
    }).start();

    setTimeout(() => {
      setCurrent(children);
      Animated.timing(fadeRef.current, {
        toValue: 1,
        duration: 300,
        easing: Easing.linear,
        useNativeDriver: true,
      }).start();
    }, 300);
  }, [children]);

  return (
    <Animated.Text style={[style, {opacity: fadeRef.current}]}>
      {current}
    </Animated.Text>
  );
}
