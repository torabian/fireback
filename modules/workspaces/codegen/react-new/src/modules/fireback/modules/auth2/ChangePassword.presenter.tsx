import { FormikProps } from "formik";
import { useEffect, useRef } from "react";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { usePostPassportChangePassword } from "../../sdk/modules/workspaces/usePostPassportChangePassword";
import {
  ChangePasswordActionReqDto,
  ClassicSigninActionReqDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { strings } from "./strings/translations";

// We extend it because the repeat password is only a front-end validation
export class ChangePasswordDto extends ChangePasswordActionReqDto {
  public password2?: string;
  static Fields = {
    ...ChangePasswordActionReqDto.Fields,
    password2: "password2",
  };
}

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const { submit: changePassword, mutation } = usePostPassportChangePassword();

  const form = useRef<FormikProps<Partial<ChangePasswordDto>> | null>();
  const setFormRef = (ref: FormikProps<Partial<ChangePasswordDto>>) => {
    form.current = ref;
  };

  const submit = () => {
    changePassword(form.current.values).then((res) => {
      alert(JSON.stringify(res));
      // push(`/${locale}/auth/otp`, undefined, {
      //   value: form.current.values.value,
      // });
    });
  };

  // Previous screen sends the email/phone here
  useEffect(() => {
    if (!state?.value) {
      return;
    }

    form.current.setFieldValue(
      ClassicSigninActionReqDto.Fields.value,
      state.value
    );
  }, [state?.value, form.current]);

  return {
    mutation,
    setFormRef,
    form,
    submit,
    goBack,
    s,
  };
};
