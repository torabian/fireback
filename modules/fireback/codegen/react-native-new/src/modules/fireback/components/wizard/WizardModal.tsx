import React from 'react';

import {Modal, StyleSheet, TouchableOpacity, View} from 'react-native';
import Toast from 'react-native-toast-message';
import CloseSVG from '~/assets/icons/close-modal-icon.svg';

export const WizardModal = ({
  showModal,
  setModalVisiblity,
  children,
}: {
  showModal: boolean;
  setModalVisiblity: (v?: any) => void;
  children: any;
}) => {
  return (
    <Modal visible={showModal}>
      <View style={styles.wrapper}>
        <TouchableOpacity
          onPress={() => {
            setModalVisiblity(undefined);
          }}
          hitSlop={{bottom: 10, top: 10, right: 10, left: 10}}
          style={{position: 'absolute', top: 30, right: 20, zIndex: 999}}>
          <CloseSVG />
        </TouchableOpacity>

        {children}
      </View>
      <Toast />
    </Modal>
  );
};

const styles = StyleSheet.create({
  wrapper: {flex: 1},
});
