import {Route} from 'react-router-dom';
import {UserArchiveScreen} from './UserArchiveScreen';
import {UserEntityManager} from './UserEntityManager';
import {UserSingleScreen} from './UserSingleScreen';
import {UserEntity} from 'src/sdk/fireback/modules/fireback/UserEntity';
export function useUserRoutes() {
  return (
    <>
      <Route
        element={<UserEntityManager />}
        path={UserEntity.Navigation.Rcreate}
      />
      <Route
        element={<UserSingleScreen />}
        path={UserEntity.Navigation.Rsingle}></Route>
      <Route
        element={<UserEntityManager />}
        path={UserEntity.Navigation.Redit}></Route>
      <Route
        element={<UserArchiveScreen />}
        path={UserEntity.Navigation.Rquery}></Route>
    </>
  );
}
