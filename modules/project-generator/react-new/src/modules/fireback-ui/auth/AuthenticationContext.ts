import { createContext, useContext } from "react";

/**
 * A workspace the signed-in user belongs to, together with their role in it.
 * Mirrors the `workspaceId`/`roleId` pair fireback's self-service flow returns
 * per entry of a session's workspace list.
 */
export interface AuthenticationWorkspace {
  workspaceId: string;
  roleId?: string;
  name?: string;
}

/** Minimal identity info about the signed-in user, as needed for display purposes. */
export interface AuthenticationUser {
  uniqueId: string;
  name?: string;
  email?: string;
}

/**
 * The authenticated session as handed back by fireback's self-service login
 * flow: a bearer token, who the user is, and which workspaces they can act
 * as. Generic over `TExtra` so an app can attach whatever additional fields
 * its own backend returns (plan, tenant flags, etc.) without forking this
 * type - stash them under `extra`.
 */
export interface AuthenticationSession<TExtra = unknown> {
  token: string;
  user?: AuthenticationUser;
  workspaces?: AuthenticationWorkspace[];
  /** Capability strings granted to the user, if the backend included them on login. */
  capabilities?: string[];
  /** App-specific session data beyond the common shape above. */
  extra?: TExtra;
}

/** Auth headers derived from the current session/selected workspace, ready to merge into a FetchxContext's defaultHeaders. */
export interface AuthenticationHeaders {
  authorization?: string;
  "workspace-id"?: string;
  "role-id"?: string;
}

export interface AuthenticationContextValue<TExtra = unknown> {
  /** The current session, or null when signed out. */
  session: AuthenticationSession<TExtra> | null;
  /** Convenience flag, equivalent to `session !== null`. */
  isAuthenticated: boolean;
  /** The bearer token to send with API requests, or undefined when signed out. */
  token: string | undefined;
  /**
   * Which of the session's workspaces requests are currently scoped to.
   * Falls back to the session's only workspace when it has exactly one and
   * none has been explicitly selected yet - most users only ever belong to
   * one workspace, and shouldn't have to pick it.
   */
  selectedWorkspace: AuthenticationWorkspace | undefined;
  /** `{ authorization, workspace-id, role-id }`, ready to spread into a FetchxContext's defaultHeaders. */
  headers: AuthenticationHeaders;
  /**
   * Stores a freshly-obtained session (e.g. after self-service redirects back
   * with a token, or after a successful `checkValidity()`) and persists it.
   * Pass null to clear it without also clearing the selected workspace (use
   * `signout` for a full reset).
   */
  setSession: (session: AuthenticationSession<TExtra> | null) => void;
  /** Selects (and persists) which workspace/role subsequent requests should be scoped to. */
  selectWorkspace: (workspace: AuthenticationWorkspace) => void;
  /**
   * Sends the browser to the self-service login app, remembering `returnUrl`
   * (defaults to the current page) so self-service knows where to redirect
   * the user back to once they're signed in.
   */
  authenticate: (returnUrl?: string) => void;
  /** Clears the session and selected workspace, both in memory and in storage. Local only - does not call the backend. */
  signout: () => void;
  /**
   * Re-checks whether the current session is still valid, via whichever
   * function the `AuthenticationProvider` was configured with (its
   * `checkValidity` prop - e.g. a POST to a "whoami"/verify endpoint that
   * relies on a cookie self-service set). Also run automatically once on
   * boot when a persisted session exists.
   *
   * Resolves `true` if the session is (still) considered valid, `false` if
   * it was found invalid and cleared as a result. If the provider wasn't
   * configured with a `checkValidity` function, nothing can be verified
   * remotely, so this just resolves the current `isAuthenticated` value
   * without making a request.
   */
  checkValidity: () => Promise<boolean>;
}

// The Context object itself can't be generic (React.createContext needs one
// fixed type), so it's declared with `any` here and re-specialized to the
// caller's TExtra via a cast in useAuthentication below - the one place that
// escape hatch is allowed in this file.
// eslint-disable-next-line @typescript-eslint/no-explicit-any
type AnyAuthenticationContextValue = AuthenticationContextValue<any>;

export const AuthenticationContext = createContext<
  AnyAuthenticationContextValue | undefined
>(undefined);

/**
 * Reads the current authentication state. Must be called under an
 * `AuthenticationProvider`.
 *
 * `TExtra` should match whatever the app's `AuthenticationProvider` was set
 * up with - it's a type-level convenience only (there's exactly one provider
 * per app, so exactly one true `TExtra`), and isn't and can't be verified at
 * runtime.
 */
export function useAuthentication<
  TExtra = unknown,
>(): AuthenticationContextValue<TExtra> {
  const ctx = useContext(AuthenticationContext);
  if (!ctx) {
    throw new Error(
      "useAuthentication must be used within an AuthenticationProvider",
    );
  }
  return ctx as AuthenticationContextValue<TExtra>;
}
