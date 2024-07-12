import {useTheme} from '../../modules/theme';
import React from 'react';
import {TouchableOpacity, Text, Image, StyleSheet} from 'react-native';

interface ButtonProps {
  title: string;
  icon?: any;
  onPress?: () => void;
}
const Button = ({title, icon, onPress}: ButtonProps) => {
  const {theme} = useTheme();
  return (
    <TouchableOpacity style={styles.button} onPress={onPress}>
      <Image source={icon} style={styles.icon} />
      <Text style={styles.title}>{title}</Text>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  button: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 10,
    paddingTop: 13,
    paddingBottom: 13,
    borderColor: 'black',
    borderWidth: 1,
    borderRadius: 25,
    marginTop: 4,
    marginBottom: 4,
  },
  icon: {
    position: 'absolute',
    width: 24,
    height: 24,
    marginRight: 10,
    left: 10,
  },
  title: {
    fontSize: 14,
    textAlign: 'center',
    alignSelf: 'center',
    flex: 1,
    fontWeight: 'bold',
  },
});

export default Button;
