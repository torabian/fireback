import React, {useState} from 'react';
import {
  ActivityIndicator,
  Image,
  ImageSourcePropType,
  StyleSheet,
  View,
} from 'react-native';

export const Avatar = ({
  imageSource,
}: {
  imageSource: ImageSourcePropType | undefined | string | null;
}) => {
  // I did not find a good UI implementation so far.
  const [loading, setLoading] = useState(false);

  return (
    <View style={styles.container}>
      {loading && (
        <ActivityIndicator
          size="large"
          color="#0000ff"
          style={{width: 5, height: 5, position: 'absolute'}}
        />
      )}
      <Image
        source={{uri: imageSource as string}}
        style={styles.image}
        onPartialLoad={() => setLoading(false)}
        onLoadEnd={() => setLoading(false)}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    position: 'relative',
  },
  image: {
    width: 100,
    height: 100,
    borderRadius: 50,
  },
});
