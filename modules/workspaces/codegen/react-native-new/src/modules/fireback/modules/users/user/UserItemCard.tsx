import {StyleSheet, Text, View} from 'react-native';
import {
  ActivityIndicator,
  ActivityState,
} from '../../../components/activity-indicator/ActivityIndicator';

import {themeEl} from '../../../../../themes/theme';
import {Avatar} from '../../../components/avatar/Avatar';
import {UserEntity} from '../../../sdk/core/react-tools';

export const UserItemCard = (props: {entity: UserEntity}) => {
  return (
    <View style={styles.wrapper}>
      <View style={styles.avatarSection}>
        <Avatar imageSource={props.entity.photo} />

        <View style={styles.statusContainer}>
          <ActivityIndicator state={ActivityState.Unknow} />
        </View>
      </View>
      <View style={styles.textualSection}>
        <Text>
          {props.entity.firstName} {props.entity.lastName}
        </Text>
        <View style={themeEl.keyPairRow}>
          <Text style={themeEl.textLabel}>Id:</Text>
          <Text>{props.entity.uniqueId || ''}</Text>
        </View>
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
  statusContainer: {},
});
