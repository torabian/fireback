import { WhoamiAction } from "@fireback/ui-core/sdk/abac/WhoamiAction";
import { FetchxContext } from "@fireback/js-remote-ctx/common/fetchx";
import { BUILD_VARIABLES } from "../../hooks/build-variables";
import { SESSION_STORAGE_KEY } from "@fireback/auth-client";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { wasmFetchOverride } from "@fireback/wasm-server/wasmServer";

// The app's real session lives under this key - see
// fireback-ui/auth/AuthenticationProvider.tsx, which persists it there via
// authenticationUtils' writeStoredValue(SESSION_STORAGE_KEY, ...). checkSession
// runs as a plain function (SessionGate can't be a context consumer of
// AuthenticationContext, since it's meant to sit above any provider - see
// SessionGate.tsx), so the real session's storage key is read directly here
// rather than through a hook.
function readRemoteQueryToken(): string | undefined {
  try {
    const raw = localStorage.getItem(SESSION_STORAGE_KEY);
    if (!raw) return undefined;
    const parsed = JSON.parse(raw) as { token?: string };
    return parsed?.token || undefined;
  } catch {
    return undefined;
  }
}

/**
 * Real `checkSession` for SessionGate (see SessionGate.tsx) - calls the
 * whoami endpoint (GET /whoami, see Abac.emi.yml/WhoamiActionImplementation.go)
 * to decide whether the app is safe to render, distinguishing three outcomes:
 *
 *  - No session stored at all: there's nothing to verify, so this resolves
 *    immediately and lets the rest of the app render as normal - its own
 *    routes already handle the signed-out state (e.g. self-service's own
 *    welcome/signup screens, manage's "Authentication Currently Unavailable"
 *    welcome page).
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
  console.log(1);
  const token = readRemoteQueryToken();
  console.log(2, token);

  // if (!token) {
  //   return;
  // }

  const { response } = await WhoamiAction.Fetch(
    { headers: { authorization: token } },
    {
      ctx: new FetchxContext("", null, null, null, wasmFetchOverride()),
    },
  );

  if (response.status === 401 || response.status === 403) {
    localStorage.removeItem(SESSION_STORAGE_KEY);
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
