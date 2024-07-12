import React from 'react';
import {StyleSheet, View} from 'react-native';
// import SkeletonPlaceholder from 'react-native-skeleton-placeholder';

import colors from '~/constants/colors';
import {Card} from '../card/Card';

export const DefaultCardLoader = ({showAddress}: {showAddress?: boolean}) => {
  return (
    <Card style={styles.container}>
      {/* <SkeletonPlaceholder>
        <SkeletonPlaceholder.Item flexDirection="row" alignItems="center">
          <SkeletonPlaceholder.Item marginLeft={20}>
            <SkeletonPlaceholder.Item
              width={180}
              height={10}
              borderRadius={4}
            />

            {showAddress !== false ? (
              <>
                <SkeletonPlaceholder.Item
                  marginTop={24}
                  width={80}
                  height={8}
                  borderRadius={4}
                />
                <SkeletonPlaceholder.Item
                  marginTop={10}
                  width={100}
                  height={8}
                  borderRadius={4}
                />
              </>
            ) : (
              <></>
            )}
          </SkeletonPlaceholder.Item>
        </SkeletonPlaceholder.Item>
      </SkeletonPlaceholder>
      <View style={styles.dividerStyle} />

      <SkeletonPlaceholder>
        <SkeletonPlaceholder.Item flexDirection="row" alignItems="center">
          <SkeletonPlaceholder.Item marginLeft={20}>
            <SkeletonPlaceholder.Item width={80} height={5} borderRadius={4} />
            <SkeletonPlaceholder.Item
              marginTop={5}
              width={80}
              height={5}
              borderRadius={4}
            />
          </SkeletonPlaceholder.Item>

          <SkeletonPlaceholder.Item
            marginLeft={20}
            alignContent="space-between"
            justifyContent="space-between">
            <SkeletonPlaceholder.Item width={80} height={5} borderRadius={4} />
            <SkeletonPlaceholder.Item
              alignSelf="flex-end"
              marginTop={5}
              width={80}
              height={5}
              borderRadius={4}
            />
          </SkeletonPlaceholder.Item>

          <SkeletonPlaceholder.Item alignSelf="flex-end" marginLeft={20}>
            <SkeletonPlaceholder.Item
              width={25}
              height={25}
              borderRadius={40}
            />
          </SkeletonPlaceholder.Item>
        </SkeletonPlaceholder.Item>
      </SkeletonPlaceholder> */}
    </Card>
  );
};

const styles = StyleSheet.create({
  container: {
    paddingVertical: 20,
    marginVertical: 5,
    marginHorizontal: 5,
  },
  dividerStyle: {
    borderBottomColor: 'silver',
    borderBottomWidth: 2,
    marginTop: 21,
    marginBottom: 29,
    marginHorizontal: 0,
  },
});
