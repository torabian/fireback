import { FormikProps, useFormik } from "formik";
import { useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useGetPassportsAvailableMethods } from "../../sdk/modules/abac/useGetPassportsAvailableMethods";
import { usePostWorkspacePassportCheck } from "../../sdk/modules/abac/usePostWorkspacePassportCheck";
import { AuthMethod } from "./auth.common";
import { useRecaptcha2 } from "../../hooks/useRecaptcha2";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";
import { CheckClassicPassportActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";

export const usePresenter = ({ method }: { method: AuthMethod }) => {
  const s = useS(strings);
  const { goBack, push, state } = useRouter();
  const { locale } = useLocale();
  const { submit: submitCheck, mutation } = usePostWorkspacePassportCheck();
  const canGoBack = state?.canGoBack === false ? false : true;

  const { query: passportMethodsQuery } = useGetPassportsAvailableMethods({
    unauthorized: true,
  });

  const enabledRecaptcha2 =
    passportMethodsQuery.data?.data?.enabledRecaptcha2 || false;
  const recaptcha2ClientKey =
    passportMethodsQuery.data?.data?.recaptcha2ClientKey || undefined;

  const submit = (data: Partial<CheckClassicPassportActionReqDto>) => {
    submitCheck(data)
      .then((res) => {
        const { next, flags } = res.data as any;

        // this condition means there is only otp available. So no other chance.
        if (next.includes("otp") && next.length === 1) {
          push(`/${locale}/selfservice/otp`, undefined, {
            value: data.value,
            type: method,
          });
        } else if (next.includes("signin-with-password")) {
          push(`/${locale}/selfservice/password`, undefined, {
            value: data.value,
            next,
            canContinueOnOtp: next?.includes("otp"),
            flags,
          });
        } else if (next.includes("create-with-password")) {
          push(`/${locale}/selfservice/complete`, undefined, {
            value: data.value,
            type: method,
            next,
            flags,
          });
        }
      })
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<CheckClassicPassportActionReqDto>>({
    initialValues: {},
    onSubmit: submit,
  });

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
    if (!enabledRecaptcha2 || !value) {
      return;
    }

    form.setFieldValue(
      CheckClassicPassportActionReqDto.Fields.securityToken,
      value
    );
  }, [value]);

  return {
    title,
    mutation,
    canGoBack,
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
