import { FormikProps } from "formik";
import { useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useGetPassportsAvailableMethods } from "../../sdk/modules/workspaces/useGetPassportsAvailableMethods";
import { usePostWorkspacePassportCheck } from "../../sdk/modules/workspaces/usePostWorkspacePassportCheck";
import { CheckClassicPassportActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { AuthMethod } from "./auth.common";
import { useRecaptcha2 } from "../../hooks/useRecaptcha2";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const usePresenter = ({ method }: { method: AuthMethod }) => {
  const s = useS(strings);
  const { goBack, push, state } = useRouter();
  const { locale } = useLocale();
  const { submit: submitCheck, mutation } = usePostWorkspacePassportCheck();
  const canGoBack = state?.canGoBack === false ? false : true;

  const { query: passportMethodsQuery } = useGetPassportsAvailableMethods({});

  const enabledRecaptcha2 =
    passportMethodsQuery.data?.data?.enabledRecaptcha2 || false;
  const recaptcha2ClientKey =
    passportMethodsQuery.data?.data?.recaptcha2ClientKey || undefined;

  const form = useRef<FormikProps<
    Partial<CheckClassicPassportActionReqDto>
  > | null>();
  const setFormRef = (
    ref: FormikProps<Partial<CheckClassicPassportActionReqDto>>
  ) => {
    form.current = ref;
  };

  let title = s.continueWithEmail;
  let description = s.continueWithEmailDescription;
  if (method === "phone") {
    title = s.continueWithPhone;
    description = s.continueWithPhoneDescription;
  }

  const {
    Component: Recaptcha,
    LegalNotice,
    value,
  } = useRecaptcha2({
    enabled: enabledRecaptcha2,
    sitekey: recaptcha2ClientKey,
  });

  // Update the recaptcha value into the security token.
  useEffect(() => {
    if (!enabledRecaptcha2 || !form.current || !value) {
      return;
    }

    form.current.setFieldValue(
      CheckClassicPassportActionReqDto.Fields.securityToken,
      value
    );
  }, [form.current, value]);

  const submit = (data: Partial<CheckClassicPassportActionReqDto>) => {
    submitCheck(data)
      .then((res) => {
        const { next, flags } = res.data as any;

        // this condition means there is only otp available. So no other chance.
        if (next.includes("otp") && next.length === 1) {
          push(`/${locale}/auth/otp`, undefined, {
            value: data.value,
            type: method,
          });
        } else if (next.includes("signin-with-password")) {
          push(`/${locale}/auth/password`, undefined, {
            value: data.value,
            next,
            flags,
          });
        } else if (next.includes("create-with-password")) {
          push(`/${locale}/auth/complete`, undefined, {
            value: data.value,
            type: method,
            next,
            flags,
          });
        }
      })
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    title,
    mutation,
    canGoBack,
    setFormRef,
    form,
    enabledRecaptcha2,
    recaptcha2ClientKey,
    description,
    Recaptcha,
    LegalNotice,
    s,
    submit,
    goBack,
  };
};
