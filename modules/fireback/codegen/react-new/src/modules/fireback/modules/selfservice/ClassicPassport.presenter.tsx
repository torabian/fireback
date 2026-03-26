import { useFormik } from "formik";
import { useEffect } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRecaptcha2 } from "../../hooks/useRecaptcha2";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { CheckClassicPassportActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";
import {
  CheckPassportMethodsActionRes,
  useCheckPassportMethodsActionQuery,
} from "../../sdk/modules/abac/CheckPassportMethods";
import { usePostWorkspacePassportCheck } from "../../sdk/modules/abac/usePostWorkspacePassportCheck";
import { GResponse } from "../../sdk/sdk/envelopes";
import { AuthMethod } from "./auth.common";
import { strings } from "./strings/translations";

export const usePresenter = ({ method }: { method: AuthMethod }) => {
  const s = useS(strings);
  const { goBack, push, state } = useRouter();
  const { locale } = useLocale();
  const { submit: submitCheck, mutation } = usePostWorkspacePassportCheck();
  const canGoBack = state?.canGoBack === false ? false : true;

  let enabledRecaptcha2 = false;
  let recaptcha2ClientKey = "";

  const { data } = useCheckPassportMethodsActionQuery({});

  if (
    data instanceof GResponse &&
    data.data.item instanceof CheckPassportMethodsActionRes
  ) {
    enabledRecaptcha2 = data?.data?.item.enabledRecaptcha2;
    recaptcha2ClientKey = data?.data?.item?.recaptcha2ClientKey;
  } else {
    // There isn't an error checking or validation mechanism to tell user
    // that it has been failed.
  }

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
      value,
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
