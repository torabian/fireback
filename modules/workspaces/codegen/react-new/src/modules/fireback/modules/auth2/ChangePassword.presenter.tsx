import { FormikProps } from "formik";
import { useEffect, useRef } from "react";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { usePostPassportChangePassword } from "../../sdk/modules/workspaces/usePostPassportChangePassword";
import { ChangePasswordActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
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
  const { goBack, state, replace, push, query } = useRouter();
  const { submit: changePassword, mutation } = usePostPassportChangePassword();
  const uniqueId = query?.uniqueId as string;

  const form = useRef<FormikProps<
    Partial<ChangePasswordActionReqDto>
  > | null>();
  const setFormRef = (
    ref: FormikProps<Partial<ChangePasswordActionReqDto>>
  ) => {
    form.current = ref;
  };

  const submit = () => {
    changePassword(form.current.values).then((res) => {
      goBack();
    });
  };

  useEffect(() => {
    if (!uniqueId || !form.current) {
      return;
    }

    form.current.setFieldValue(
      ChangePasswordActionReqDto.Fields.uniqueId,
      uniqueId
    );
  }, [uniqueId, form.current]);

  return {
    mutation,
    setFormRef,
    form,
    submit,
    goBack,
    s,
  };
};
