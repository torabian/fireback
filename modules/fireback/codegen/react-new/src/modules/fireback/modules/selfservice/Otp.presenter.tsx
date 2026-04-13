import { useFormik } from "formik";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { usePostWorkspacePassportOtp } from "../../sdk/modules/abac/usePostWorkspacePassportOtp";

import {
  ClassicPassportOtpActionReqDto,
  ClassicPassportOtpActionResDto,
} from "../../sdk/modules/abac/AbacActionsDto";
import type { GResponse } from "../../sdk/sdk/envelopes";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const s = useS(strings);
  const { submit: signin, mutation } = usePostWorkspacePassportOtp();
  const { onComplete } = useCompleteAuth();

  const submit = (values: Partial<ClassicPassportOtpActionReqDto>) => {
    signin({ ...values, value: state.value })
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicPassportOtpActionReqDto>>({
    initialValues: {},
    onSubmit: submit,
  });

  const successful = (res: GResponse<ClassicPassportOtpActionResDto>) => {
    if (res.data?.item.session) {
      onComplete(res);
    } else if (res.data?.item?.continueWithCreation) {
      push(`/${locale}/selfservice/complete`, undefined, {
        value: state.value,
        type: state.type,
        sessionSecret: res.data.item?.sessionSecret,
        totpUrl: res.data.item?.totpUrl,
      });
    }
  };

  return {
    mutation,
    form,
    s,
    submit,
    goBack,
  };
};
