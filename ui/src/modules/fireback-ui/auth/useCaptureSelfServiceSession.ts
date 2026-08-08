import { useEffect } from "react";

import { useAuthentication, type AuthenticationSession } from "./AuthenticationContext";

const SESSION_QUERY_PARAM = "session";

/**
 * Captures the `?session=<json>` query param self-service appends when it
 * redirects back after a successful login (see
 * `AuthenticationContextValue.authenticate` / `AuthenticationProvider`'s doc
 * comments for the full flow): parses it, stores it via `setSession`, and
 * strips the param from the URL so it doesn't linger in the address bar or
 * get re-processed on a later reload or if the link is shared/bookmarked.
 *
 * Mount this once near the app root, inside an `AuthenticationProvider` -
 * self-service's `redirect` param can point back at any page (not just
 * "/"), so this needs to run regardless of which route is currently active.
 */
export function useCaptureSelfServiceSession(): void {
  const { setSession } = useAuthentication();

  useEffect(() => {
    const url = new URL(window.location.href);
    const raw = url.searchParams.get(SESSION_QUERY_PARAM);
    if (!raw) return;

    // Strip the param unconditionally, even if parsing fails below - a
    // session that can't be understood shouldn't get re-attempted on every
    // reload, and it shouldn't stay sitting in the address bar either way.
    url.searchParams.delete(SESSION_QUERY_PARAM);
    window.history.replaceState(null, "", url.toString());

    const session = parseSelfServiceSession(raw);
    if (session) setSession(session);
    // Only ever run once, against whatever the URL looked like on first
    // render - setSession is stable for the provider's lifetime, and
    // location.search is read directly rather than via a dependency.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
}

/**
 * Maps fireback's raw session JSON (a `UserSessionDto`, as handed back via
 * the `?session=` redirect) onto this app's own, smaller
 * `AuthenticationSession` shape. Returns null when `raw` isn't parseable JSON
 * or doesn't look like a session at all (no token).
 *
 * The full raw payload is kept under `extra` regardless, so nothing is lost
 * even for fields this mapping doesn't know about.
 */
function parseSelfServiceSession(raw: string): AuthenticationSession | null {
  let parsed: unknown;
  try {
    parsed = JSON.parse(raw);
  } catch {
    return null;
  }
  if (!parsed || typeof parsed !== "object") return null;

  const data = parsed as Record<string, unknown>;
  const token = data.token;
  if (typeof token !== "string" || !token) return null;

  const rawUser = data.user as Record<string, unknown> | null | undefined;
  const user =
    rawUser && typeof rawUser.uniqueId === "string"
      ? {
          uniqueId: rawUser.uniqueId,
          name: typeof rawUser.name === "string" ? rawUser.name : undefined,
          email:
            typeof rawUser.email === "string" ? rawUser.email : undefined,
        }
      : undefined;

  const rawWorkspaces = Array.isArray(data.userWorkspaces)
    ? (data.userWorkspaces as Record<string, unknown>[])
    : [];
  const workspaces = rawWorkspaces
    .filter(
      (workspace) =>
        typeof workspace.workspaceId === "string" && workspace.workspaceId,
    )
    .map((workspace) => ({
      workspaceId: workspace.workspaceId as string,
      roleId:
        typeof workspace.roleId === "string" ? workspace.roleId : undefined,
      name: typeof workspace.name === "string" ? workspace.name : undefined,
    }));

  return {
    token,
    user,
    workspaces: workspaces.length > 0 ? workspaces : undefined,
    extra: parsed,
  };
}
