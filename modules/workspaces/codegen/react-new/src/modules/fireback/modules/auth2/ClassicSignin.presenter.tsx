import { FormikProps } from "formik";
import { useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { usePostWorkspacePassportCheck } from "../../sdk/modules/workspaces/usePostWorkspacePassportCheck";
import { CheckClassicPassportActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { AuthMethod } from "./auth.common";

export const usePresenter = ({ method }: { method: AuthMethod }) => {
  const { goBack, push } = useRouter();
  const { locale } = useLocale();
  const { submit: submitCheck, mutation } = usePostWorkspacePassportCheck();

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

  const submit = (data: Partial<CheckClassicPassportActionReqDto>) => {
    submitCheck({ value: data.value })
      .then((res) => {
        // Here we need bunch of logic actually.
        if (res.data.exists) {
          push(`/${locale}/auth/${method}/password`, undefined, {
            value: data.value,
          });
        } else {
          alert("Not exists!");
        }
      })
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    title,
    mutation,
    setFormRef,
    form,
    description,
    submit,
    goBack,
  };
};
