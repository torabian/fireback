/**
 * Tools for authentication based on fireback ABAC plugin
 */

import { UserSessionDto } from "src/sdk/fireback";
import React, { useContext, useEffect, useState } from "react";

export interface IAuthContext {
  token: string | null;
}

export type SetAuthStateFn = (token: string) => void;

export interface IAuthContextProvider {
  ref: IAuthContext;
  setToken: SetAuthStateFn;
  setSession: (session: UserSessionDto) => void;
  signout: () => void;
  isAuthenticated: boolean;
}

export const AuthContext = React.createContext<IAuthContextProvider>({
  setToken() {},
  setSession() {},
  signout() {},
  ref: {
    token: "",
  },
  isAuthenticated: false,
});

export interface PageTitleOptions {
  onTrigger: (actionKey: string) => void;
}

function getAuthState() {
  const userConfig = localStorage.getItem("app_auth_state");
  if (!userConfig) {
    return;
  }

  try {
    const cnf = JSON.parse(userConfig);
    if (cnf) {
      return { ...cnf };
    }
    return {};
  } catch (error) {}
  return {};
}

const appAuthInitialState = getAuthState();
export function usePageTitle(title?: string) {
  const t = useContext(AuthContext);

  useEffect(() => {
    t.setToken(title || "");
  }, [title]);
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [authData, setAuthData] = useState<IAuthContext>(appAuthInitialState);

  const signout = () => {
    setAuthData({ token: "" });
    localStorage.removeItem("app_auth_state");
  };

  const setSession = (session: UserSessionDto) => {
    const newConf = { ...authData, ...session };
    setAuthData(newConf);
    localStorage.setItem("app_auth_state", JSON.stringify(newConf));
  };

  const setToken: SetAuthStateFn = (token) => {
    const newConf = { ...authData, token };
    setAuthData(newConf);
    localStorage.setItem("app_auth_state", JSON.stringify(newConf));
  };

  const isAuthenticated = !!authData?.token;

  return (
    <AuthContext.Provider
      value={{
        signout,
        setSession,
        isAuthenticated,
        ref: authData,
        setToken,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
