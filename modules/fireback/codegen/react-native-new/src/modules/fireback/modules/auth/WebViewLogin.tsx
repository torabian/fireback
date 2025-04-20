import React, {useContext, useState} from 'react';
import {Modal, StyleSheet, View, Button as ButtonNative} from 'react-native';
import {WebView, WebViewMessageEvent} from 'react-native-webview';
import {RemoteQueryContext} from '../../sdk/core/react-tools';
import Button from '../../components/button/Button';

export const WebViewLogin = () => {
  const [modalVisible, setModalVisible] = useState(false);
  const {setSession} = useContext(RemoteQueryContext);

  const onMessage = (event: WebViewMessageEvent) => {
    try {
      const data = JSON.parse(event.nativeEvent.data);
      if (data.session) {
        setSession(data.session);
        setModalVisible(false);
      }
    } catch (e) {
      console.error('Error parsing WebView message', e);
    }
  };

  return (
    <View style={styles.container}>
      <Button
        title="Continue with Fireback"
        onPress={() => setModalVisible(true)}
      />
      <Modal visible={modalVisible} animationType="slide">
        <View style={styles.modalContainer}>
          <ButtonNative title="Close" onPress={() => setModalVisible(false)} />
          <WebView
            source={{
              uri: 'http://192.168.1.10:4508/selfservice/#/en/passports',
            }}
            onMessage={onMessage}
            javaScriptEnabled={true}
            style={{flex: 1}}
          />
        </View>
      </Modal>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {flex: 1, justifyContent: 'center', alignItems: 'center'},
  modalContainer: {flex: 1, marginTop: 50},
});
