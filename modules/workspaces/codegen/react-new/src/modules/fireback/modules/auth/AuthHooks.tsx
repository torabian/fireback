import { FormCheckbox } from "@/modules/fireback/components/forms/form-switch/FormSwitch";
import { KeyValue } from "@/modules/fireback/definitions/definitions";
import { useT } from "@/modules/fireback/hooks/useT";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { enTranslations } from "../../translations/en";
import { FormikProps } from "formik";
import React, { useEffect, useState } from "react";

export function getCachedCredentials(): any {
  const state = localStorage.getItem("remember_credentials");
  if (state !== "true") {
    return {};
  }

  const cred = localStorage.getItem("credentials");

  try {
    if (!cred) {
      return {};
    }

    const d = JSON.parse(cred);

    if (d && d.value && d.password) {
      return d as any;
    }
  } catch (error) {
    // Intentially left blank. No need to handle this type of error
    return {};
  }
}
export function useRememberingLoginForm(
  formik: React.MutableRefObject<
    FormikProps<Partial<ClassicSigninActionReqDto>> | null | undefined
  >
) {
  const t = useT();
  const [remember, setRememberState] = useState(false);

  const rememberCredentials = () => {
    setRememberState((r) => !r);
    Promise.resolve(
      localStorage.setItem("remember_credentials", `${!remember}`)
    );

    if (remember) {
      localStorage.removeItem("credentials");
    }
  };

  const bootScreen = async () => {
    const state = await localStorage.getItem("remember_credentials");
    if (state !== "true") {
      return;
    }

    setRememberState(true);

    const cred = await localStorage.getItem("credentials");

    try {
      if (!cred) {
        return;
      }

      const d = JSON.parse(cred);

      if (d && d.value && d.password) {
        formik.current?.setValues({ value: d.value, password: d.password });
      }
    } catch (error) {
      // Intentially left blank. No need to handle this type of error
    }
  };

  useEffect(() => {
    bootScreen();
  }, []);

  const RememberSwitch = () => (
    <div style={{ flexDirection: "row", justifyContent: "center" }}>
      <FormCheckbox
        label={t.abac.remember}
        value={remember || false}
        onChange={rememberCredentials}
      />
    </div>
  );

  return { RememberSwitch, shouldRemember: remember };
}

export const getAuthOtpMethods = (t: typeof enTranslations): KeyValue[] => {
  return [
    {
      label: t.abac.viaEmail,
      value: "email",
    },
    {
      label: t.abac.viaSms,
      value: "sms",
    },
  ];
};
