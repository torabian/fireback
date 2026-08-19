import { useAuthentication } from "@fireback/auth-client";
import { mapRawSessionToAuthenticationSession } from "@fireback/auth-client";
import type { MOne } from "@fireback/js-remote-ctx/common/operators";
import { MArray, MCollection } from "@fireback/js-remote-ctx/common/operators";
import { GResponse } from "@fireback/js-remote-ctx/envelopes";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import type { ClassicPassportOtpActionRes } from "@fireback/selfservice/sdk/abac/ClassicPassportOtpAction";
import type { ClassicSigninActionRes } from "@fireback/selfservice/sdk/abac/ClassicSigninAction";
import type { ClassicSignupActionRes } from "@fireback/selfservice/sdk/abac/ClassicSignupAction";
import type { CompletePassportPasswordResetActionRes } from "@fireback/selfservice/sdk/abac/CompletePassportPasswordResetAction";
import type { UserSessionDto } from "@fireback/selfservice/sdk/abac/UserSessionDto";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useQueryUserRoleWorkspacesActionQuery } from "@fireback/ui-core/sdk/abac/QueryUserRoleWorkspacesAction";
import { WhoamiAction } from "@fireback/ui-core/sdk/abac/WhoamiAction";
import { useEffect } from "react";
import { strings } from "./strings/translations";

export enum AuthMethod {
  Email = "email",
  Phone = "phone",
  Google = "google",
  Facebook = "facebook",
}

export interface AuthAvailableMethods {
  email: boolean;
  phone: boolean;
  google: boolean;
  googleOAuthClientKey?: string;
  facebookAppId?: string;
  facebook: boolean;
}

export const useCompleteAuth = () => {
  const { setSession } = useAuthentication();
  const { replace } = useRouter();
  const s = useS(strings);
  const { selectWorkspace } = useAuthentication();
  const ctx = useFetchxContext();
  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    enabled: false,
  });

  const onComplete = (
    res: GResponse<
      | ClassicSigninActionRes
      | ClassicSignupActionRes
      | ClassicPassportOtpActionRes
      | CompletePassportPasswordResetActionRes
    >,
  ) => {
    // Handle React Native WebView
    if ((window as any).ReactNativeWebView) {
      (window as any).ReactNativeWebView.postMessage(JSON.stringify(res.data));
    }

    // Get the "redirect" query param
    const urlParams = new URLSearchParams(window.location.search);
    const redirectUrl = urlParams.get("redirect");

    // check also, if there is localstorage to redirect regardless - captured on
    // Welcome mount by useTemporaryParamOptions(["redirect", ...]) (see
    // Welcome.presenter.tsx), which also catches a ?redirect= placed *after* the
    // hash-router's own "#" (e.g. .../selfservice/#/selfservice/welcome?redirect=...),
    // unlike the plain window.location.search read below, which only ever sees a
    // redirect param placed before the "#".
    const redirect2 = sessionStorage.getItem("redirect");

    const session = (res?.data?.item?.session as MOne<UserSessionDto>)?.get();

    // Get the token from session response
    const token = session.token;

    if (!token) {
      alert(s.authenticationFailed);
      return;
    }

    if (session.userWorkspaces instanceof MCollection) {
      if (session.userWorkspaces.len() === 1) {
        const workspace = session.userWorkspaces.get()[0];
        const headers = {
          authorization: token,
          ["workspace-id"]: workspace.uniqueId,
        };

        WhoamiAction.Fetch({ headers }, { ctx })
          .then((res) => {
            return res.response.result;
          })
          .then((all) => {
            if (all instanceof GResponse) {
              if (all.data.item.workspaces instanceof MArray) {
                const workspaces = all.data.item.workspaces.get();
                if (
                  workspaces.length === 1 &&
                  workspaces[0].roles instanceof MArray
                ) {
                  if (workspaces[0].roles.len() === 1) {
                    selectWorkspace({
                      workspaceId: workspaces[0].uniqueId,
                      name: workspaces[0].name,
                      roleId: workspaces[0].roles.get()[0].uniqueId,
                    });
                  }
                }
              }
            }
          });
      }
    }

    setSession(mapRawSessionToAuthenticationSession(session));

    // Clean up url options which are set earlier.
    sessionStorage.removeItem("redirect");
    sessionStorage.removeItem("workspace_type_id");

    if (redirect2) {
      window.location.href = redirect2;
    } else if (redirectUrl) {
      // Append the token to the redirect URL
      const finalUrl = new URL(redirectUrl);
      finalUrl.searchParams.set(
        "session",
        JSON.stringify(mapRawSessionToAuthenticationSession(session)),
      );

      // Redirect to the final URL
      window.location.href = finalUrl.toString();
    } else {
      // Fallback to the default route
      const to = "/dashboard";

      replace(to, to);
    }
  };

  return { onComplete };
};

/**
 * Reads `key` from wherever a query param can actually show up in this
 * (hash-routed) app: before the "#" (plain window.location.search - the shape
 * ForcedAuthenticated.tsx's own /selfservice?redirect=... link uses) or after
 * it, embedded in the hash route's own query (.../#/selfservice/welcome?
 * redirect=...). Shared by useTemporaryParamOptions below and by App.tsx's own
 * already-authenticated redirect check, so both agree on where a param can live.
 */
export function readUrlParam(key: string): string | null {
  const searchParams = new URLSearchParams(window.location.search);
  const hashIndex = window.location.hash.indexOf("?");
  const hashParams =
    hashIndex !== -1
      ? new URLSearchParams(window.location.hash.slice(hashIndex))
      : new URLSearchParams();
  return searchParams.get(key) || hashParams.get(key);
}

/**
 * @description A lot of times, we need to set an option for authentication,
 * from external source such as query params, such as redirect url, selected workspace.
 * This function would read that from query params or hash, and store it in session storage.
 * @param key
 */
export const useTemporaryParamOptions = (keys: string[]) => {
  useEffect(() => {
    keys.forEach((key) => {
      const value = readUrlParam(key);
      if (value) {
        sessionStorage.setItem(key, value);
      }
    });
  }, [keys.join(",")]); // join to avoid useEffect missing changes
};
