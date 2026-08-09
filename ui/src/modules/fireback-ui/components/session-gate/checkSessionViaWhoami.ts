import { WhoamiAction } from "@/modules/sdk/abac/WhoamiAction";
import {
  readStoredValue,
  writeStoredValue,
  SESSION_STORAGE_KEY,
} from "@/modules/fireback-ui/auth/authenticationUtils";
import type { AuthenticationSession } from "@/modules/fireback-ui/auth/AuthenticationContext";

/**
 * Real `checkSession` for SessionGate (see SessionGate.tsx) - calls the
 * whoami endpoint (GET /, see Abac.emi.yml/WhoamiActionImplementation.go) to
 * decide whether the app is safe to render, distinguishing three outcomes:
 *
 *  - No session stored at all (readStoredValue finds nothing under the same
 *    key AuthenticationProvider persists to): there's nothing to verify, so
 *    this resolves immediately and lets the rest of the app render as
 *    normal - its own routes already handle the signed-out state (e.g.
 *    self-service's own welcome/signup screens, manage's "Authentication
 *    Currently Unavailable" welcome page).
 *  - A session IS stored, but whoami rejects it with 401/403: the token has
 *    expired or been revoked. That's an actual "you need to sign in again",
 *    not a "try later" - the stored session is cleared and the browser is
 *    sent to sign in, the same destination ForcedAuthenticated.tsx's manual
 *    "sign in instead" link uses.
 *  - A session is stored, but the request fails for any other reason
 *    (network error, CORS, 5xx, ...): that's the backend being unreachable
 *    or broken, not a verdict on the session - this rejects so SessionGate
 *    keeps retrying, and the stored session is left untouched.
 */
export async function checkSessionViaWhoami(): Promise<void> {
  const { response } = await WhoamiAction.Fetch({});

  if (response.status === 401 || response.status === 403) {
    writeStoredValue(SESSION_STORAGE_KEY, null);
    window.location.href = `/selfservice?redirect=${encodeURIComponent(
      window.location.pathname + window.location.search,
    )}`;
    throw new Error("Session expired - redirecting to sign in.");
  }

  if (!response.ok) {
    // Some other failure (5xx, ...) - same treatment as the request itself
    // throwing (network error, offline, ...): not a verdict on the session,
    // just keep retrying.
    throw new Error(`whoami failed with status ${response.status}`);
  }

  // 2xx - session confirmed valid.
}
