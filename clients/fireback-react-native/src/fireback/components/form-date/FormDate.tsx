import React, {useEffect, useState} from 'react';
import DatePickerModal from 'react-native-date-picker';
import moment from 'moment';
import {
  BaseFormElement,
  CommonFormElementProps,
} from '../base-form-element/BaseFormElement';

export const FormDate = (props: CommonFormElementProps<string | null>) => {
  const {value, onChange, disabled} = props;
  const [toggleOff, setToggleOff] = useState(false);
  const [modalOpen, setOpenModal] = useState(false);

  const onConfirm = (newDate: Date) => {
    setOpenModal(false);
    props?.onChange(moment(newDate).toISOString());
  };
  const onCancel = () => {
    setOpenModal(false);
  };
  const onOpen = () => {
    setOpenModal(true);
  };

  return (
    <BaseFormElement {...props} onPress={onOpen} displayValue={value}>
      <DatePickerModal
        modal
        open={modalOpen}
        date={moment(props.value).toDate()}
        mode="date"
        onConfirm={onConfirm}
        onCancel={onCancel}
      />
    </BaseFormElement>
  );
};
