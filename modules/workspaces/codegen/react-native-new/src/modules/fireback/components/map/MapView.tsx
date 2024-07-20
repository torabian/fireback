import React, {useEffect, useRef, useState} from 'react';
import {
  Animated,
  Clipboard,
  Easing,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
  ViewStyle,
} from 'react-native';
import MapView, {
  LatLng,
  Marker,
  Polygon,
  PROVIDER_GOOGLE,
  Region,
} from 'react-native-maps';
import Pin from '~/assets/icons/pin.svg';
import colors from '~/constants/colors';

import {GooglePlacesAutocomplete} from 'react-native-google-places-autocomplete';

const GooglePlacesInput = () => {
  return (
    <GooglePlacesAutocomplete
      placeholder="Search"
      onPress={(data, details = null) => {
        // 'details' is provided when fetchDetails = true
      }}
      query={{
        key: 'YOUR API KEY',
        language: 'en',
      }}
    />
  );
};

var mapStyle = [
  {elementType: 'geometry', stylers: [{color: '#f5f5f5'}]},
  {elementType: 'labels.icon', stylers: [{visibility: 'off'}]},
];

export interface FriendlyRegion {
  latitudeDelta?: number;
  longitudeDelta?: number;
  latitude?: number;
  longitude?: number;
}

const DEFAULT_LOCATION = {
  latitude: 37.78825,
  longitude: -122.4324,
};
export function Map({
  onChange,
  value,
  style,
  resetPoint,
  markers,
  mode,
}: {
  value?: LatLng;
  style?: ViewStyle;
  resetPoint?: LatLng;
  onChange?: (data: FriendlyRegion) => void;
  markers?: Array<LatLng>;
  mode?: 'view' | 'edit';
}) {
  const currentRegion = useRef<Region>({
    latitudeDelta: 0.01,
    longitudeDelta: 0.01,
    latitude: value?.latitude || DEFAULT_LOCATION.latitude,
    longitude: value?.longitude || DEFAULT_LOCATION.longitude,
  });
  const fadeRef = useRef(new Animated.Value(0));

  const [points, setPoints] = useState<
    Array<{points: {longitude: string; latitude: string}[]}>
  >([]);

  mode = mode || 'view';

  const onCopy = () => {
    Clipboard.setString(
      `${currentRegion.current?.latitude},${currentRegion.current?.longitude}`,
    );
  };

  const onPanStart = () => {
    if (mode !== 'edit') {
      return;
    }

    Animated.timing(fadeRef.current, {
      toValue: 1,
      duration: 300,
      easing: Easing.ease,
      useNativeDriver: true,
    }).start();
  };

  const onPanFinish = () => {
    if (mode !== 'edit') {
      return;
    }

    Animated.timing(fadeRef.current, {
      toValue: 0,
      duration: 150,
      easing: Easing.linear,
      useNativeDriver: true,
    }).start();
  };

  const onRegionChangeComplete = (region: Region) => {
    currentRegion.current = region;
    onPanFinish();

    if (onChange) {
      onChange(region);
    }
  };

  const onReset = () => {
    if (resetPoint) {
      onRegionChangeComplete({
        ...resetPoint,
        latitudeDelta: currentRegion.current?.latitudeDelta,
        longitudeDelta: currentRegion.current?.longitudeDelta,
      });
    }
  };

  const translateY = fadeRef.current.interpolate({
    inputRange: [0, 1],
    outputRange: [0, -10],
  });

  const opacity = fadeRef.current.interpolate({
    inputRange: [0, 1],
    outputRange: [1, 0.3],
  });

  return (
    <View style={[{height: 360}, style]}>
      <MapView
        // provider={PROVIDER_GOOGLE}
        style={{flex: 1}}
        zoomEnabled
        zoomControlEnabled
        // customMapStyle={mapStyle}
        onTouchStart={onPanStart}
        onRegionChangeComplete={onRegionChangeComplete}
        region={{
          ...(value && value.latitude ? value : DEFAULT_LOCATION),
          latitudeDelta: currentRegion.current?.latitudeDelta,
          longitudeDelta: currentRegion.current?.longitudeDelta,
        }}
        initialRegion={currentRegion.current}>
        {(markers || []).map(marker => (
          <Marker coordinate={marker} />
        ))}

        {points.map(p =>
          p?.points ? <Polygon coordinates={p?.points} /> : null,
        )}
      </MapView>

      <View
        style={{
          position: 'absolute',
          zIndex: 999,
          left: 20,
          right: 20,
          top: 20,
        }}>
        <GooglePlacesInput />
      </View>
      <View style={{position: 'absolute', zIndex: 999, right: 20, bottom: 20}}>
        <TouchableOpacity onPress={onCopy} style={styles.action}>
          <Text style={styles.actionText}>Copy location</Text>
        </TouchableOpacity>
        {resetPoint && (
          <TouchableOpacity
            onPress={onReset}
            style={[styles.action, {marginTop: 5}]}>
            <Text style={styles.actionText}>Reset</Text>
          </TouchableOpacity>
        )}
      </View>

      {mode === 'edit' && (
        <Animated.View
          pointerEvents="none"
          style={[
            {
              left: 0,
              right: 0,
              top: 0,
              bottom: 0,
              alignContent: 'center',
              justifyContent: 'center',
              alignItems: 'center',
              zIndex: 999,
              position: 'absolute',
            },
            {
              transform: [{translateY}],
            },
            {
              opacity,
            },
          ]}>
          <Pin width={40} height={40} />
        </Animated.View>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  action: {
    backgroundColor: '#000000BB',
    paddingVertical: 3,
    paddingHorizontal: 10,
    borderRadius: 10,
  },
  actionText: {
    textAlign: 'center',
    color: colors.white,
  },
});
