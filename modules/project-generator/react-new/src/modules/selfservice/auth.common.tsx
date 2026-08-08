import { useContext, useEffect } from "react";
import { useLocale } from "../fireback-ui/hooks/useLocale";
import { useRouter } from "../fireback-ui/hooks/useRouter";
import { useS } from "../fireback-ui/hooks/useS";
import { RemoteQueryContext } from "../sdk/core/react-tools";
import { strings } from "./strings/translations";
import type { ClassicPassportOtpActionRes } from "../sdk/abac/ClassicPassportOtpAction";
import type { ClassicSigninActionRes } from "../sdk/abac/ClassicSigninAction";
import type { ClassicSignupActionRes } from "../sdk/abac/ClassicSignupAction";
import type { GResponse } from "../sdk/sdk/envelopes";
import type { MOne } from "../sdk/sdk/common/operators";
import type { UserSessionDto } from "../sdk/abac/UserSessionDto";

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
  const { setSession, selectUrw, selectedUrw } = useContext(RemoteQueryContext);
  const { locale } = useLocale();
  const { replace } = useRouter();
  const s = useS(strings);

  const onComplete = (
    res: GResponse<
      | ClassicSigninActionRes
      | ClassicSignupActionRes
      | ClassicPassportOtpActionRes
    >,
  ) => {
    // Handle React Native WebView
    if ((window as any).ReactNativeWebView) {
      (window as any).ReactNativeWebView.postMessage(JSON.stringify(res.data));
    }

    // Get the "redirect" query param
    const urlParams = new URLSearchParams(window.location.search);
    const redirectUrl = urlParams.get("redirect");

    // check also, if there is localstorage to redirect regardless
    const redirect2 = sessionStorage.getItem("redirect_temporary");

    const session = (res?.data?.item?.session as MOne<UserSessionDto>)?.get();

    // Get the token from session response
    const token = session.token;

    if (!token) {
      alert(s.authenticationFailed);
      return;
    }

    setSession(session);

    // Clean up url options which are set earlier.
    sessionStorage.removeItem("redirect_temporary");
    sessionStorage.removeItem("workspace_type_id");

    if (redirect2) {
      window.location.href = redirect2;
    } else if (redirectUrl) {
      // Append the token to the redirect URL
      const finalUrl = new URL(redirectUrl);
      finalUrl.searchParams.set(
        "session",
        JSON.stringify(res.data.item.session),
      );

      // Redirect to the final URL
      window.location.href = finalUrl.toString();
    } else {
      // Fallback to the default route
      const to = "/{locale}/dashboard".replace("{locale}", locale || "en");

      replace(to, to);
    }
  };

  return { onComplete };
};

/**
 * @description A lot of times, we need to set an option for authentication,
 * from external source such as query params, such as redirect url, selected workspace.
 * This function would read that from query params or hash, and store it in session storage.
 * @param key
 */
export const useTemporaryParamOptions = (keys: string[]) => {
  useEffect(() => {
    const searchParams = new URLSearchParams(window.location.search);
    const hashIndex = window.location.hash.indexOf("?");
    const hashParams =
      hashIndex !== -1
        ? new URLSearchParams(window.location.hash.slice(hashIndex))
        : new URLSearchParams();

    keys.forEach((key) => {
      const value = searchParams.get(key) || hashParams.get(key);
      if (value) {
        sessionStorage.setItem(key, value);
      }
    });
  }, [keys.join(",")]); // join to avoid useEffect missing changes
};
