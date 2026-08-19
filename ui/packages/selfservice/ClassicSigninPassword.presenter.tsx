import { useFormik } from "formik";
import { useEffect } from "react";
import { mutationErrorsToFormik } from "@fireback/ui-core/hooks/api";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";

import {
  ClassicSigninActionReq,
  ClassicSigninActionRes,
  useClassicSigninAction,
} from "@fireback/selfservice/sdk/abac/ClassicSigninAction";
import type { GResponse } from "@fireback/js-remote-ctx/envelopes";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";
import {
  ClassicPassportRequestOtpActionReq,
  useClassicPassportRequestOtpAction,
} from "@fireback/selfservice/sdk/abac/ClassicPassportRequestOtpAction";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, state, push } = useRouter();
  const { onComplete } = useCompleteAuth();
  const mutation = useClassicSigninAction();
  const otpEnabled = state?.canContinueOnOtp;
  const requestOtpMutation = useClassicPassportRequestOtpAction();

  const submit = (values: Partial<ClassicSigninActionReq>) => {
    mutation
      .mutateAsync(
        new ClassicSigninActionReq({
          value: values.value,
          password: values.password,
        }),
      )
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicSigninActionReq>>({
    initialValues: {},
    onSubmit: submit,
  });

  const continueWithOtp = () => {
    requestOtpMutation
      .mutateAsync(
        new ClassicPassportRequestOtpActionReq({ value: form.values.value }),
      )
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

    form.setFieldValue(ClassicSigninActionReq.Fields.value, state.value);
  }, [state?.value]);

  const successful = (res: GResponse<ClassicSigninActionRes>) => {
    // here we need to also check if there is another step!!!

    if (res.data.item.session) {
      onComplete(res);
    } else if (res.data.item.next?.includes("enter-totp")) {
      push(`/selfservice/totp-enter`, undefined, {
        value: form.values.value,
        password: form.values.password,
      });
    } else if (res.data.item.next?.includes("setup-totp")) {
      push(`/selfservice/totp-setup`, undefined, {
        totpUrl: res.data.item.totpUrl,

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
