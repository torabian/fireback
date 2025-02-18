import { FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { usePostPassportsSigninClassic } from "../../sdk/modules/workspaces/usePostPassportsSigninClassic";
import { UserSessionDto } from "../../sdk/modules/workspaces/UserSessionDto";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useLocale } from "../../hooks/useLocale";
import { usePostWorkspacePassportRequestOtp } from "../../sdk/modules/workspaces/usePostWorkspacePassportRequestOtp";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const { submit: singin, mutation } = usePostPassportsSigninClassic();

  const { submit: requestOtp } = usePostWorkspacePassportRequestOtp();

  const { setSession } = useContext(RemoteQueryContext);

  const form = useRef<FormikProps<Partial<ClassicSigninActionReqDto>> | null>();
  const setFormRef = (ref: FormikProps<Partial<ClassicSigninActionReqDto>>) => {
    form.current = ref;
  };

  const continueWithOtp = () => {
    requestOtp({ value: form.current.values.value }).then((res) => {
      console.log(25, res);
      push(`/${locale}/auth/otp`, undefined, {
        value: form.current.values.value,
      });
    });
  };

  // Previous screen sends the email/phone here
  useEffect(() => {
    form.current.setFieldValue(
      ClassicSigninActionReqDto.Fields.value,
      state.value
    );
  }, [state.value, form.current]);

  const successful = (res: IResponse<UserSessionDto>) => {
    setSession(res.data);
    if (process.env.REACT_APP_DEFAULT_ROUTE) {
      const to = (
        process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
      ).replace("{locale}", locale || "en");

      replace(to, to);
    }
  };

  const submit = (values: Partial<ClassicSigninActionReqDto>) => {
    singin({ value: values.value, password: values.password })
      .then(successful)
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    mutation,
    setFormRef,
    continueWithOtp,
    form,
    submit,
    goBack,
  };
};
