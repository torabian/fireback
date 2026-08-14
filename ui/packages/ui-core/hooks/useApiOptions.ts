import { useAuthentication } from "@fireback/auth-client";
import { BUILD_VARIABLES } from "./build-variables";
import { useLocale } from "./useLocale";

export interface ApiOptionsHeaders {
  authorization?: string;
  "workspace-id"?: string;
  "role-id"?: string;
  "accept-language"?: string;
}

export interface ApiOptions {
  headers: ApiOptionsHeaders;
  prefix: string;
}

/**
 * `{ headers, prefix }` for code that talks to the backend directly (raw
 * `fetch`/`XMLHttpRequest`) instead of through the generated SDK's `fetchx`
 * plumbing - e.g. file export/download links, the reactive search websocket,
 * `QueryError`'s error-message formatting. Replaces the `options` field the
 * old `RemoteQueryContext` used to expose for the same purpose.
 *
 * The generated SDK itself doesn't need this - its requests get their
 * headers from the `FetchxContext` `WithFireback.tsx` wires up straight from
 * `useAuthentication()`.
 */
export function useApiOptions(): ApiOptions {
  const { headers } = useAuthentication();
  const { locale } = useLocale();

  return {
    headers: {
      ...headers,
      "accept-language": locale,
    },
    prefix: BUILD_VARIABLES.REMOTE_SERVICE || "",
  };
}
