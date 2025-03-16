import { FormikProps } from "formik";
import { useContext, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { usePostPassportsSigninClassic } from "../../sdk/modules/workspaces/usePostPassportsSigninClassic";
import {
  ClassicSigninActionReqDto,
  ClassicSigninActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const { submit: signin, mutation } = usePostPassportsSigninClassic();

  const { setSession } = useContext(RemoteQueryContext);

  const form = useRef<FormikProps<Partial<ClassicSigninActionReqDto>> | null>();
  const setFormRef = (ref: FormikProps<Partial<ClassicSigninActionReqDto>>) => {
    form.current = ref;
  };

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const successful = (res: IResponse<ClassicSigninActionResDto>) => {
    if (res.data?.session) {
      setSession(res.data?.session);
      if ((window as any).ReactNativeWebView) {
        (window as any).ReactNativeWebView.postMessage(
          JSON.stringify(res.data)
        );
      }

      if (process.env.REACT_APP_DEFAULT_ROUTE) {
        const to = (
          process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
        ).replace("{locale}", locale || "en");

        replace(to, to);
      }
    }
  };

  const submit = (values: Partial<ClassicSigninActionReqDto>) => {
    signin({ ...values, password, value })
      .then(successful)
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    mutation,
    totpUrl,
    forcedTotp,
    setFormRef,
    form,
    submit,
    goBack,
  };
};
