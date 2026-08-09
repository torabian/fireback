import { useEffect } from "react";

import { useAuthentication } from "./AuthenticationContext";
import { mapRawSessionToAuthenticationSession } from "./authenticationUtils";

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

    let parsed: unknown;
    try {
      parsed = JSON.parse(raw);
    } catch {
      return;
    }

    const session = mapRawSessionToAuthenticationSession(parsed);
    if (session) setSession(session);
    // Only ever run once, against whatever the URL looked like on first
    // render - setSession is stable for the provider's lifetime, and
    // location.search is read directly rather than via a dependency.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
}
