import {AuthRouter} from '@/modules/fireback/modules/auth/Router';
import {UserArchiveScreen} from '@/modules/fireback/modules/users/user/UserArchiveScreen';
import {RemoteQueryContext} from '@/modules/fireback/sdk/core/react-tools';
import {createBottomTabNavigator} from '@react-navigation/bottom-tabs';
import {
  DrawerContentScrollView,
  DrawerItem,
  createDrawerNavigator,
} from '@react-navigation/drawer';
import React, {useContext} from 'react';
import {Text, TouchableOpacity, View} from 'react-native';

const Tab = createBottomTabNavigator();

const Drawer = createDrawerNavigator();

function CustomDrawerContent(props: any) {
  const {signout} = useContext(RemoteQueryContext);

  const {navigation} = props;

  return (
    <DrawerContentScrollView {...props}>
      <DrawerItem
        label="Users"
        onPress={() => navigation.navigate(UserArchiveScreen.Name)}
      />

      <DrawerItem label="Logout" onPress={() => signout()} />
    </DrawerContentScrollView>
  );
}

function DrawerNavigator() {
  return (
    <Drawer.Navigator
      drawerContent={props => <CustomDrawerContent {...props} />}
      initialRouteName="HomeTabs">
      <Drawer.Screen
        name={UserArchiveScreen.Name}
        component={UserArchiveScreen}
      />
    </Drawer.Navigator>
  );
}

export function ApplicationRoutes(): React.JSX.Element {
  const {setSession, session, checked, isAuthenticated, signout} =
    useContext(RemoteQueryContext);

  if (!checked) {
    return (
      <View>
        <Text>Loading</Text>
      </View>
    );
  }

  if (!isAuthenticated) {
    return <AuthRouter />;
  }

  return (
    <>
      <DrawerNavigator />
    </>
  );
}
