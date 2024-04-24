import {StyleSheet, Text, View} from 'react-native';
import {
  ActivityIndicator,
  ActivityState,
} from '../../components/activity-indicator/ActivityIndicator';
import {Avatar} from '../../components/avatar/Avatar';

import {
  CustomerAddress,
  CustomerEntity,
} from '@/sdk/fireback/modules/demo/CustomerEntity';
import {useActivityState} from '@/sdk/fireback/modules/demo/useGetCustomers';
import {themeEl} from '../../themes/theme';
import {useEffect} from 'react';

export const CustomerCard = (props: {entity: CustomerEntity}) => {
  const {query} = useActivityState();
  const thisOne = (query.data?.data?.items || []).find(
    f => f.uniqueId === props.entity.uniqueId,
  );

  // I live this for testing
  useEffect(() => {
    console.log('Created: ', props.entity.firstName, props.entity.lastName);
    return () =>
      console.log('Removed: ', props.entity.firstName, props.entity.lastName);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <View style={styles.wrapper}>
      <View style={styles.avatarSection}>
        <Avatar imageSource={props.entity.avatar} />

        <View style={styles.statusContainer}>
          <ActivityIndicator state={thisOne?.status || ActivityState.Unknow} />
        </View>
      </View>
      <View style={styles.textualSection}>
        <Text>
          {props.entity.firstName} {props.entity.lastName}
        </Text>
        <AddressView entity={props.entity.address} />
        <View style={themeEl.keyPairRow}>
          <Text style={themeEl.textLabel}>Id:</Text>
          <Text>{props.entity.uniqueId || ''}</Text>
        </View>
      </View>
    </View>
  );
};

const AddressView = (props: {entity: CustomerAddress | null | undefined}) => {
  const address = props.entity || {};

  return (
    <View>
      <View style={themeEl.keyPairRow}>
        <Text style={themeEl.textLabel}>Country:</Text>
        <Text style={{width: 120}} numberOfLines={1} ellipsizeMode="tail">
          {address?.country || ''}
        </Text>
      </View>
      <View style={themeEl.keyPairRow}>
        <Text style={themeEl.textLabel}>City:</Text>
        <Text numberOfLines={1} ellipsizeMode="tail">
          {address?.city || ''}
        </Text>
      </View>
      <View style={themeEl.keyPairRow}>
        <Text style={themeEl.textLabel}>Street:</Text>
        <Text numberOfLines={1} ellipsizeMode="tail">
          {address?.street || ''}
        </Text>
      </View>
      <View style={themeEl.keyPairRow}>
        <Text style={themeEl.textLabel}>Zip Code:</Text>
        <Text numberOfLines={1} ellipsizeMode="tail">
          {address?.zipCode || ''}
        </Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    borderWidth: 1,
    padding: 20,
    flexDirection: 'row',
    margin: 10,
    borderRadius: 5,
    borderColor: 'silver',
  },
  avatarSection: {width: 120},
  textualSection: {flex: 1},
  statusContainer: {
    // To me the initial position felt good actually.
    // position: "absolute",
    // right: 0,
    // top: 0,
  },
});
