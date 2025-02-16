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

export const usePresenter = ({ method }: { method: AuthMethod }) => {
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

  let title = "Continue with Email";
  let description = "Enter your email address to continue.";
  if (method === "phone") {
    title = "Continue with phone";
    description = "Enter your phone number to continue";
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
        // Here we need bunch of logic actually.
        if (res.data.continueWithPassword === true) {
          push(`/${locale}/auth/${method}/password`, undefined, {
            value: data.value,
          });
        } else if (res.data.continueWithPassword === false) {
          push(`/${locale}/auth/otp`, undefined, {
            value: data.value,
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
    submit,
    goBack,
  };
};
