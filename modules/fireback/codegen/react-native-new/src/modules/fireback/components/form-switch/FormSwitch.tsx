import React, {useEffect, useState} from 'react';
import {Switch} from 'react-native';
import colors from '~/constants/colors';
import {BaseFormElement} from '../base-form-element/BaseFormElement';

export const FormSwitch = (props: {
  value: boolean | null;
  onChange: (value: boolean) => void;
  disabled?: boolean;
  label: string;
}) => {
  const {value, onChange, disabled} = props;
  const [toggleOff, setToggleOff] = useState(false);

  let thumbColor = value ? colors.white : colors.white;
  let trackColor = {
    true: colors.primaryColor,
    false: colors.graySwitchBg,
  };

  useEffect(() => {
    setToggleOff(disabled ?? false);
  }, [disabled]);

  if (disabled) {
    thumbColor = colors.graySwitchBg;
    trackColor = {
      true: colors.grayText,
      false: colors.grayText,
    };
  }

  return (
    <BaseFormElement {...props} labelStyle={{left: 60}} hasAnimation={false}>
      <Switch
        disabled={toggleOff}
        thumbColor={thumbColor}
        trackColor={trackColor}
        value={value || false}
        onValueChange={() => onChange(!value)}
      />
    </BaseFormElement>
  );
};
