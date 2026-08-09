import { useEffect, useMemo, useState, type ReactNode } from "react";

import {
  AuthenticationContext,
  type AuthenticationContextValue,
  type AuthenticationSession,
  type AuthenticationWorkspace,
} from "./AuthenticationContext";
import {
  buildAuthHeaders,
  buildSelfServiceLoginUrl,
  readStoredValue,
  resolveSelectedWorkspace,
  writeStoredValue,
  SESSION_STORAGE_KEY,
  WORKSPACE_STORAGE_KEY,
} from "./authenticationUtils";

export interface AuthenticationProviderProps<TExtra = unknown> {
  children: ReactNode;
  /**
   * Base URL of the self-service login app (e.g. admin's embedded
   * self-service screens) that `authenticate()` sends the browser to.
   */
  selfServiceUrl: string;
  /** Query param self-service reads to know where to send the user back after signing in. Defaults to "redirect". */
  returnParam?: string;
  /**
   * Configurable "is this session still valid" check - e.g. a POST to a
   * whoami/verify endpoint that relies on a cookie self-service set when it
   * redirected back here. Run once automatically on boot when a persisted
   * session exists, and callable again any time via the `checkValidity`
   * function `useAuthentication()` returns.
   *
   * Should resolve the refreshed session if it's still valid, or null if it
   * isn't (the user is then signed out locally). Left unconfigured, no
   * remote check is ever made and the persisted session is trusted as-is
   * until a request elsewhere fails.
   */
  checkValidity?: (
    session: AuthenticationSession<TExtra>,
  ) => Promise<AuthenticationSession<TExtra> | null>;
}

/**
 * Provides the app's authentication state (session + selected workspace),
 * backed by localStorage so it survives reloads, and exposes the
 * authenticate/signout/checkValidity actions via `useAuthentication()`.
 *
 * This provider does not perform sign-in itself: `authenticate()` redirects
 * to a separate self-service login app, which is expected to authenticate
 * the user and hand control back here (e.g. via a cookie it set, verified by
 * this provider's `checkValidity` on boot, or a page that reads a returned
 * token and calls `setSession(...)` directly).
 */
export function AuthenticationProvider<TExtra = unknown>({
  children,
  selfServiceUrl,
  returnParam = "redirect",
  checkValidity: checkValidityHook,
}: AuthenticationProviderProps<TExtra>) {
  // Lazy initializers read storage synchronously on first render, so there's
  // no "still loading" gap to account for before the persisted session (if
  // any) is available.
  const [session, setSessionState] = useState<
    AuthenticationSession<TExtra> | null
  >(() => readStoredValue<AuthenticationSession<TExtra>>(SESSION_STORAGE_KEY));
  const [explicitWorkspace, setExplicitWorkspace] = useState<
    AuthenticationWorkspace | undefined
  >(
    () =>
      readStoredValue<AuthenticationWorkspace>(WORKSPACE_STORAGE_KEY) ??
      undefined,
  );

  /** Stores a session (or clears it, for null) both in state and in storage. */
  const setSession = (next: AuthenticationSession<TExtra> | null) => {
    writeStoredValue(SESSION_STORAGE_KEY, next);
    setSessionState(next);
  };

  /** Selects (and persists) which workspace/role subsequent requests should be scoped to. */
  const selectWorkspace = (workspace: AuthenticationWorkspace) => {
    writeStoredValue(WORKSPACE_STORAGE_KEY, workspace);
    setExplicitWorkspace(workspace);
  };

  /** Clears session + selected workspace, in both state and storage. */
  const clearSession = () => {
    writeStoredValue(SESSION_STORAGE_KEY, null);
    writeStoredValue(WORKSPACE_STORAGE_KEY, null);
    setSessionState(null);
    setExplicitWorkspace(undefined);
  };

  /** Local-only sign-out: forgets the session without contacting the backend. */
  const signout = () => {
    clearSession();
  };

  /** Redirects to self-service's login screen; see `AuthenticationContextValue.authenticate`. */
  const authenticate = (returnUrl: string = window.location.href) => {
    window.location.href = buildSelfServiceLoginUrl(
      selfServiceUrl,
      returnParam,
      returnUrl,
    );
  };

  /** Runs the configured `checkValidity` prop, if any; see `AuthenticationContextValue.checkValidity`. */
  const checkValidity = async (): Promise<boolean> => {
    if (!session) return false;
    if (!checkValidityHook) return true;
    try {
      const refreshed = await checkValidityHook(session);
      if (refreshed) {
        setSession(refreshed);
        return true;
      }
      clearSession();
      return false;
    } catch {
      // The check request itself failing (network blip, etc.) shouldn't sign
      // the user out - only an explicit "invalid" (null) result should.
      return true;
    }
  };

  useEffect(() => {
    // Boot-time validity check, once, against whatever session was hydrated
    // from storage on first render - intentionally not re-run on every
    // session change (that would immediately re-fire right after
    // authenticate()/setSession() populates a brand new session too).
    //
    // checkValidity only calls setSession/clearSession after its awaited
    // checkValidityHook() call settles, never synchronously - but the
    // set-state-in-effect rule's static analysis can't tell an eventual,
    // async state update apart from an immediate one, so it flags this call
    // regardless. Disabled rather than restructured, since inlining
    // checkValidity's body here would just duplicate the same logic
    // `useAuthentication().checkValidity()` callers rely on elsewhere.
    // eslint-disable-next-line react-hooks/set-state-in-effect
    checkValidity();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const value = useMemo<AuthenticationContextValue<TExtra>>(() => {
    const selectedWorkspace = resolveSelectedWorkspace(
      session,
      explicitWorkspace,
    );
    return {
      session,
      isAuthenticated: session !== null,
      token: session?.token,
      selectedWorkspace,
      headers: buildAuthHeaders(session, selectedWorkspace),
      setSession,
      selectWorkspace,
      authenticate,
      signout,
      checkValidity,
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [session, explicitWorkspace, selfServiceUrl, returnParam, checkValidityHook]);

  return (
    <AuthenticationContext.Provider value={value}>
      {children}
    </AuthenticationContext.Provider>
  );
}
