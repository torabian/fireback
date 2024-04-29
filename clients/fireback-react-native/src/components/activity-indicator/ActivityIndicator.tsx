import {useEffect, useRef} from 'react';
import {Animated, StyleSheet, View} from 'react-native';
import {usePrevious} from '../../hooks/usePrevious';
import {theme} from '../../themes/theme';

export enum ActivityState {
  Active = 1,
  Inactive = 2,
  Unknow = 3,
}

function colorFromState(state: ActivityState | undefined) {
  return state === ActivityState.Active
    ? theme.activeColor
    : state === ActivityState.Inactive
    ? theme.inactiveColor
    : theme.idleColor;
}

export function ActivityIndicator(props: {state: ActivityState}) {
  const backgroundColor = useRef(new Animated.Value(0)).current;
  const prev = usePrevious<ActivityState>(props.state);

  useEffect(() => {
    backgroundColor.setValue(0);
    Animated.timing(backgroundColor, {
      toValue: 1,
      duration: 1000,
      useNativeDriver: true,
    }).start();

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [props.state]);

  const colorInterpolate = backgroundColor.interpolate({
    inputRange: [0, 1],
    outputRange: [colorFromState(prev), colorFromState(props.state)],
  });

  return (
    <View>
      <Animated.View
        style={[
          styles.wrapper,
          {
            backgroundColor: colorInterpolate,
          },
        ]}></Animated.View>
    </View>
  );
}

const bulletSize = 10;
const styles = StyleSheet.create({
  wrapper: {
    width: bulletSize,
    height: bulletSize,
    borderRadius: bulletSize / 2,
    overflow: 'hidden',
  },
  active: {
    backgroundColor: theme.activeColor,
  },
  inactive: {
    backgroundColor: theme.inactiveColor,
  },
  idle: {
    backgroundColor: theme.idleColor,
  },
});
