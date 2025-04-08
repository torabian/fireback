import { FormikProps, useFormik } from "formik";
import { useEffect, useRef } from "react";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { usePostPassportChangePassword } from "../../sdk/modules/abac/usePostPassportChangePassword";
import { strings } from "./strings/translations";
import { ChangePasswordActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";

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
  const { goBack, state, replace, push, query } = useRouter();
  const { submit: changePassword, mutation } = usePostPassportChangePassword();
  const uniqueId = query?.uniqueId as string;

  const submit = () => {
    changePassword(form.values).then((res) => {
      goBack();
    });
  };

  const form = useFormik<ChangePasswordActionReqDto>({
    initialValues: {},
    onSubmit: submit,
  });

  useEffect(() => {
    if (!uniqueId || !form) {
      return;
    }

    form.setFieldValue(ChangePasswordActionReqDto.Fields.uniqueId, uniqueId);
  }, [uniqueId]);

  return {
    mutation,
    form,
    submit,
    goBack,
    s,
  };
};
