import type {
  AuthenticationHeaders,
  AuthenticationSession,
  AuthenticationWorkspace,
} from "./AuthenticationContext";

/**
 * localStorage keys the session/selected workspace are persisted under.
 * Exported (rather than kept private to AuthenticationProvider) so other
 * code that needs to read/clear the same storage without going through the
 * provider - e.g. checkSessionViaWhoami, which runs before any provider is
 * mounted - stays in sync with it instead of duplicating the key string.
 */
export const SESSION_STORAGE_KEY = "nima_ui_session";
export const WORKSPACE_STORAGE_KEY = "nima_ui_selected_workspace";

/**
 * localStorage is only present in a browser. Guarded so this module doesn't
 * blow up if it's ever imported in a non-DOM context (e.g. SSR, tests).
 */
export function hasLocalStorage(): boolean {
  return typeof window !== "undefined" && !!window.localStorage;
}

/**
 * Reads and JSON-parses a value from localStorage.
 * @returns the parsed value, or null when the key is absent or its stored value isn't valid JSON.
 */
export function readStoredValue<T>(key: string): T | null {
  if (!hasLocalStorage()) return null;
  const raw = window.localStorage.getItem(key);
  if (!raw) return null;
  try {
    return JSON.parse(raw) as T;
  } catch {
    // Corrupt/foreign value under our key - treat as absent rather than throw.
    return null;
  }
}

/**
 * JSON-stringifies and writes a value to localStorage under `key`.
 * Passing `null` removes the key instead (there's no such thing as a stored `null`).
 */
export function writeStoredValue<T>(key: string, value: T | null): void {
  if (!hasLocalStorage()) return;
  if (value === null) {
    window.localStorage.removeItem(key);
  } else {
    window.localStorage.setItem(key, JSON.stringify(value));
  }
}

/**
 * Picks which workspace requests should be scoped to: the explicitly
 * selected one if there is one, otherwise the session's only workspace (most
 * users only ever belong to one, and shouldn't have to pick it), otherwise
 * undefined (signed out, or signed in with >1 workspace and none picked yet).
 */
export function resolveSelectedWorkspace<TExtra>(
  session: AuthenticationSession<TExtra> | null,
  explicit: AuthenticationWorkspace | undefined,
): AuthenticationWorkspace | undefined {
  if (explicit) return explicit;
  if (session?.workspaces?.length === 1) return session.workspaces[0];
  return undefined;
}

/**
 * Builds the `{ authorization, workspace-id, role-id }` headers implied by a
 * session/workspace pair - ready to spread into a FetchxContext's
 * `defaultHeaders` so every generated SDK request carries them. Fields with
 * no value (e.g. no workspace selected yet) are simply omitted.
 */
export function buildAuthHeaders<TExtra>(
  session: AuthenticationSession<TExtra> | null,
  workspace: AuthenticationWorkspace | undefined,
): AuthenticationHeaders {
  const headers: AuthenticationHeaders = {};
  if (session?.token) headers.authorization = session.token;
  if (workspace?.workspaceId) headers["workspace-id"] = workspace.workspaceId;
  if (workspace?.roleId) headers["role-id"] = workspace.roleId;
  return headers;
}

/**
 * Builds the URL `authenticate()` sends the browser to: self-service's login
 * screen, with a `returnParam` query param set to `returnUrl` so self-service
 * knows where to send the user back once they're signed in.
 */
export function buildSelfServiceLoginUrl(
  selfServiceUrl: string,
  returnParam: string,
  returnUrl: string,
): string {
  const url = new URL(selfServiceUrl);
  url.searchParams.set(returnParam, returnUrl);
  return url.toString();
}

/**
 * Session payloads arrive in different shapes depending on the code path: an
 * `MOne<UserSessionDto>` wrapper (mutation responses, before `.get()`), a
 * `UserSessionDto` class instance, or a plain object (e.g. already
 * `JSON.parse`d from the self-service redirect's `?session=` query param). A
 * stringify/parse round-trip normalizes all of them to a plain object in one
 * step - the same technique this module's storage helpers rely on, and
 * MCollection/MOne's own `toJSON()` is written to support exactly this.
 */
function normalizeSessionPayload(raw: unknown): Record<string, unknown> | null {
  if (!raw) return null;
  try {
    return JSON.parse(JSON.stringify(raw));
  } catch {
    return null;
  }
}

/**
 * Maps fireback's raw session JSON (a `UserSessionDto`, as handed back by
 * sign-in/sign-up mutations or the self-service `?session=` redirect) onto
 * this app's own, smaller `AuthenticationSession` shape. Returns null when
 * `raw` isn't parseable or doesn't look like a session at all (no token).
 *
 * The full normalized payload is kept under `extra` regardless, so nothing
 * is lost even for fields this mapping doesn't know about.
 */
export function mapRawSessionToAuthenticationSession(
  raw: unknown,
): AuthenticationSession | null {
  const data = normalizeSessionPayload(raw);
  if (!data) return null;

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
    extra: data,
  };
}
