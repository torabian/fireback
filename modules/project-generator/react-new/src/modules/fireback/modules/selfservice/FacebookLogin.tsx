import { useEffect } from "react";
import { source } from "../../hooks/source";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const FacebookLogin = ({
  continueWithResult,
  facebookAppId,
}: {
  continueWithResult: (value: string) => void;
  facebookAppId: string;
}) => {
  const s = useS(strings);

  // Load Facebook SDK
  useEffect(() => {
    if ((window as any).FB) return; // Already loaded

    const script = document.createElement("script");
    script.src = "https://connect.facebook.net/en_US/sdk.js";
    script.async = true;
    script.onload = () => {
      (window as any).FB.init({
        appId: facebookAppId,
        cookie: true,
        xfbml: false,
        version: "v19.0",
      });
    };
    document.body.appendChild(script);
  }, []);

  const loginWithFacebook = () => {
    const FB = (window as any).FB;
    if (!FB) {
      alert("Facebook SDK not loaded");
      return;
    }

    FB.login(
      (response: any) => {
        console.log("Facebook:", response);
        if (response.authResponse?.accessToken) {
          continueWithResult(response.authResponse.accessToken);
        } else {
          alert("Facebook login failed");
        }
      },
      { scope: "email,public_profile" }
    );
  };

  return (
    <button id="using-facebook" type="button" onClick={loginWithFacebook}>
      <img className="button-icon" src={source("/common/facebook.png")} />
      Facebook
    </button>
  );
};
