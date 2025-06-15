import {Button, Text, View} from 'react-native';

export const ConfirmDrawer = ({
  onConfirm,
  onReject,
  title,
}: {
  onConfirm: () => void;
  onReject: () => void;
  title?: string;
}) => (
  <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
    <Text>{title || 'Confirm this action?'}</Text>
    <Button title="Yes" onPress={onConfirm} />
    <Button title="No" onPress={onReject} />
  </View>
);
