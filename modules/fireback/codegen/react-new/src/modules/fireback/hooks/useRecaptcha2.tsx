import { createRef, useEffect, useRef, useState } from "react";
const ReCAPTCHA2 = require("react-google-recaptcha").default;

export const useRecaptcha2 = ({
  sitekey,
  enabled,
  invisible,
}: {
  sitekey: string;
  enabled: boolean;
  invisible?: boolean;
}) => {
  // By default let's use invisible.
  invisible = invisible === undefined ? true : invisible;

  const [value, setValue] = useState();
  const [forceVisible, setForceVisible] = useState(false);
  const recaptcha2Ref = createRef<typeof ReCAPTCHA2>();
  const refValue = useRef("");

  useEffect(() => {
    if (enabled && recaptcha2Ref.current) {
      recaptcha2Ref.current?.execute();
      recaptcha2Ref.current?.reset();
    }
  }, [enabled, recaptcha2Ref.current]);

  useEffect(() => {
    setTimeout(() => {
      if (!refValue.current) {
        setForceVisible(true);
      }
    }, 2000);
  }, []);

  // Place it in the form near the Submit button
  const Component = () => {
    if (!enabled || !sitekey) {
      return null;
    }
    return (
      <>
        <ReCAPTCHA2
          sitekey={sitekey}
          size={invisible && !forceVisible ? "invisible" : undefined}
          ref={recaptcha2Ref}
          onChange={(value) => {
            setValue(value);
            refValue.current = value;
          }}
        />
      </>
    );
  };

  // This is needed on the footer to make it compliant with google license if using invisible
  const LegalNotice = () => {
    if (!invisible || !enabled) {
      return null;
    }

    return (
      <div className="mt-5 recaptcha-closure">
        This site is protected by reCAPTCHA and the Google
        <a target="_blank" href="https://policies.google.com/privacy">
          {" "}
          Privacy Policy{" "}
        </a>{" "}
        and
        <a target="_blank" href="https://policies.google.com/terms">
          {" "}
          Terms of Service{" "}
        </a>{" "}
        apply.
      </div>
    );
  };

  return { value, Component, LegalNotice };
};
