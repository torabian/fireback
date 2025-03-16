import { FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { usePostPassportsSigninClassic } from "../../sdk/modules/workspaces/usePostPassportsSigninClassic";
import { usePostWorkspacePassportRequestOtp } from "../../sdk/modules/workspaces/usePostWorkspacePassportRequestOtp";
import {
  ClassicSigninActionReqDto,
  ClassicSigninActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const { submit: singin, mutation } = usePostPassportsSigninClassic();
  const otpEnabled = state?.canContinueOnOtp;
  const { submit: requestOtp } = usePostWorkspacePassportRequestOtp();

  const { setSession } = useContext(RemoteQueryContext);

  const form = useRef<FormikProps<Partial<ClassicSigninActionReqDto>> | null>();
  const setFormRef = (ref: FormikProps<Partial<ClassicSigninActionReqDto>>) => {
    form.current = ref;
  };

  const continueWithOtp = () => {
    requestOtp({ value: form.current.values.value }).then((res) => {
      push(`/${locale}/auth/otp`, undefined, {
        value: form.current.values.value,
      });
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

  const successful = (res: IResponse<ClassicSigninActionResDto>) => {
    // here we need to also check if there is another step!!!

    if (res.data.session) {
      setSession(res.data.session);
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
    } else if (res.data.next?.includes("enter-totp")) {
      push(`/${locale}/selfservice/totp-enter`, undefined, {
        value: form.current.values.value,
        password: form.current.values.password,
      });
    } else if (res.data.next?.includes("setup-totp")) {
      push(`/${locale}/selfservice/totp-setup`, undefined, {
        totpUrl: res.data.totpUrl,

        // since we do not allow user to join, it means it's forced :)
        forcedTotp: true,
        password: form.current.values.password,
        value: state?.value,
      });
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
    otpEnabled,
    continueWithOtp,
    form,
    submit,
    goBack,
    s,
  };
};
