import { useFormik } from "formik";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { IResponse } from "../../sdk/core/http-tools";
import { usePostWorkspacePassportOtp } from "../../sdk/modules/workspaces/usePostWorkspacePassportOtp";
import {
  ClassicPassportOtpActionReqDto,
  ClassicPassportOtpActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
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

  const successful = (res: IResponse<ClassicPassportOtpActionResDto>) => {
    if (res.data?.session) {
      onComplete(res);
    } else if (res.data.continueWithCreation) {
      push(`/${locale}/selfservice/complete`, undefined, {
        value: state.value,
        type: state.type,
        sessionSecret: res.data.sessionSecret,
        totpUrl: res.data.totpUrl,
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
