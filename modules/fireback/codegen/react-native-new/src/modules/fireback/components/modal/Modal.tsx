import React from 'react';
import {
  Dimensions,
  Modal as RNModal,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from 'react-native';
import {KeyboardAwareScrollView} from 'react-native-keyboard-aware-scroll-view';
import {BehaviorSubject} from 'rxjs';
import CrossIcon from '~/assets/icons/cross.svg';
import colors from '~/constants/colors';

export interface ModalDialog {
  title?: string;
  data?: any;
  Component?: () => JSX.Element;
}

export interface ModalDialogInterchange extends ModalDialog {
  visible?: boolean;
}

export function openDialg(dialog: ModalDialog) {
  ModalInterchange.next({
    ...dialog,
    visible: true,
  });
}

export const ModalInterchange = new BehaviorSubject<ModalDialogInterchange>({
  visible: false,
});

export const Modal = ({
  title,
  Body,
  data,
  isVisible,
  onClose,
  testID,
}: {
  title: string;
  data?: any;
  Body?: (props: {data?: any; fnClose: () => void}) => JSX.Element;
  isVisible: boolean;
  onClose(): void;
  testID?: string;
}) => {
  return (
    <RNModal
      visible={isVisible}
      animationType="slide"
      transparent={true}
      testID={testID}>
      <View style={styles.modalBg} />
      <View style={styles.modalBody}>
        <View style={styles.modalHeaderWrapper}>
          <Text style={styles.modalHeader}>{title}</Text>
          <TouchableOpacity
            hitSlop={{bottom: 20, right: 20, top: 20, left: 20}}
            onPress={onClose}>
            <CrossIcon
              color={colors.placeholderText}
              style={styles.modalCrossIcon}
            />
          </TouchableOpacity>
        </View>

        <KeyboardAwareScrollView
          keyboardShouldPersistTaps="always"
          showsVerticalScrollIndicator={false}>
          {Body && (
            <Body
              fnClose={() => {
                setTimeout(() => {
                  onClose();
                }, 100);
              }}
              data={data}
            />
          )}
        </KeyboardAwareScrollView>
      </View>
    </RNModal>
  );
};

const styles = StyleSheet.create({
  modalBg: {
    backgroundColor: 'rgba(55, 63, 76, 0.8)',
    flex: 1,
  },
  modalBody: {
    backgroundColor: colors.white,
    position: 'absolute',
    bottom: 0,
    left: 0,
    maxHeight: Dimensions.get('screen').height - 100,
    right: 0,
    paddingVertical: 5,
    paddingHorizontal: 20,
    borderTopLeftRadius: 5,
    borderTopRightRadius: 5,
  },
  modalHeaderWrapper: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 30,
  },
  modalCrossIcon: {
    alignSelf: 'flex-end',
    width: 50,
    height: 50,
  },
  modalHeader: {
    fontWeight: 'bold',
    fontSize: 18,
    color: colors.white,
  },
});
