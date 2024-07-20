import AsyncStorage from '@react-native-async-storage/async-storage';
import {UserSessionDto} from 'fireback-tools/modules/passports';
import {store} from '~/store/Store';
// import decode from 'jwt-decode';

export function resetSession() {
  store.session.next({token: '', user: null});
  return AsyncStorage.setItem('session', '');
}

export function setSession(session: UserSessionDto) {
  store.session.next(session);
  return AsyncStorage.setItem('session', JSON.stringify(session));
}

// export function checkTokenExpiry(token: string) {
//   const now = Math.ceil(Date.now() / 1000);
//   const expiry = decode<{exp: number; [key: string]: unknown}>(token).exp;
//   const timeToExpiry = expiry - now;
//   const tokenExpired = timeToExpiry < 0;

//   return tokenExpired;
// }

export async function getSession() {
  try {
    const content = await AsyncStorage.getItem('session');
    if (!content) {
      return null;
    }

    return JSON.parse(content);
  } catch (err) {
    throw err;
  }
}
