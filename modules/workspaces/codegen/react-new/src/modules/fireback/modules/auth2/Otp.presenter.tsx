import { FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { usePostWorkspacePassportOtp } from "../../sdk/modules/workspaces/usePostWorkspacePassportOtp";
import { UserSessionDto } from "../../sdk/modules/workspaces/UserSessionDto";
import {
  ClassicPassportOtpActionReqDto,
  ClassicPassportOtpActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";

export const usePresenter = () => {
  const { goBack, state, replace } = useRouter();
  const { locale } = useLocale();
  const { submit: singin, mutation } = usePostWorkspacePassportOtp();

  const { setSession } = useContext(RemoteQueryContext);

  const form = useRef<FormikProps<
    Partial<ClassicPassportOtpActionReqDto>
  > | null>();
  const setFormRef = (
    ref: FormikProps<Partial<ClassicPassportOtpActionReqDto>>
  ) => {
    form.current = ref;
  };

  // Previous screen sends the email/phone here
  useEffect(() => {
    form.current.setFieldValue(
      ClassicPassportOtpActionReqDto.Fields.value,
      state.value
    );
  }, [state.value, form.current]);

  const successful = (res: IResponse<ClassicPassportOtpActionResDto>) => {
    setSession(res.data?.session);
    if (process.env.REACT_APP_DEFAULT_ROUTE) {
      const to = (
        process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
      ).replace("{locale}", locale || "en");

      replace(to, to);
    }
  };

  const submit = (values: Partial<ClassicPassportOtpActionReqDto>) => {
    singin(values)
      .then(successful)
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    mutation,
    setFormRef,
    form,
    submit,
    goBack,
  };
};
