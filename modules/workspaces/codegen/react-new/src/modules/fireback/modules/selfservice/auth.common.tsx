import { useContext } from "react";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useRouter } from "../../hooks/useRouter";
import { useLocale } from "../../hooks/useLocale";

export enum AuthMethod {
  Email = "email",
  Phone = "phone",
  Google = "google",
}

export interface AuthAvailableMethods {
  email: boolean;
  phone: boolean;
  google: boolean;
}

export const useCompleteAuth = () => {
  const { setSession } = useContext(RemoteQueryContext);
  const { locale } = useLocale();
  const { goBack, state, replace, push } = useRouter();
  const onComplete = (res) => {
    setSession(res.data.session);

    // Handle React Native WebView
    if ((window as any).ReactNativeWebView) {
      (window as any).ReactNativeWebView.postMessage(JSON.stringify(res.data));
    }

    // Get the "redirect" query param
    const urlParams = new URLSearchParams(window.location.search);
    const redirectUrl = urlParams.get("redirect");

    // Get the token from session response
    const token = res.data?.session?.token; // Adjust based on your API response

    if (redirectUrl && token) {
      // Append the token to the redirect URL
      const finalUrl = new URL(redirectUrl);
      finalUrl.searchParams.set("session", JSON.stringify(res.data.session));

      // Redirect to the final URL
      window.location.href = finalUrl.toString();
    } else {
      // Fallback to the default route
      const to = (
        process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
      ).replace("{locale}", locale || "en");

      replace(to, to);
    }
  };

  return { onComplete };
};
