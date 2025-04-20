import { FormikProps, useFormik } from "formik";
import { useContext, useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { usePostPassportsSigninClassic } from "../../sdk/modules/abac/usePostPassportsSigninClassic";
import { usePostWorkspacePassportRequestOtp } from "../../sdk/modules/abac/usePostWorkspacePassportRequestOtp";

import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";
import { useCompleteAuth } from "./auth.common";
import {
  ClassicSigninActionReqDto,
  ClassicSigninActionResDto,
} from "../../sdk/modules/abac/AbacActionsDto";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const { onComplete } = useCompleteAuth();
  const { submit: singin, mutation } = usePostPassportsSigninClassic();
  const otpEnabled = state?.canContinueOnOtp;
  const { submit: requestOtp } = usePostWorkspacePassportRequestOtp();

  const { setSession } = useContext(RemoteQueryContext);

  const submit = (values: Partial<ClassicSigninActionReqDto>) => {
    singin({ value: values.value, password: values.password })
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicSigninActionReqDto>>({
    initialValues: {},
    onSubmit: submit,
  });

  const continueWithOtp = () => {
    requestOtp({ value: form.values.value })
      .then((res) => {
        push(`../otp`, undefined, {
          value: form.values.value,
        });
      })
      .catch((res) => {
        // @todo code gen can send us the entire messages as well
        // so in front-end also we have all of the possible error messages
        if (res.error.message === "OtaRequestBlockedUntil") {
          // Fireback might request otp already if sees the next option might be
          // the otp only.
          push(`../otp`, undefined, {
            value: form.values.value,
          });
        }
      });
  };

  // Previous screen sends the email/phone here
  useEffect(() => {
    if (!state?.value) {
      return;
    }

    form.setFieldValue(ClassicSigninActionReqDto.Fields.value, state.value);
  }, [state?.value]);

  const successful = (res: IResponse<ClassicSigninActionResDto>) => {
    // here we need to also check if there is another step!!!

    if (res.data.session) {
      onComplete(res);
    } else if (res.data.next?.includes("enter-totp")) {
      push(`/${locale}/selfservice/totp-enter`, undefined, {
        value: form.values.value,
        password: form.values.password,
      });
    } else if (res.data.next?.includes("setup-totp")) {
      push(`/${locale}/selfservice/totp-setup`, undefined, {
        totpUrl: res.data.totpUrl,

        // since we do not allow user to join, it means it's forced :)
        forcedTotp: true,
        password: form.values.password,
        value: state?.value,
      });
    }
  };

  return {
    mutation,
    otpEnabled,
    continueWithOtp,
    form,
    submit,
    goBack,
    s,
  };
};
