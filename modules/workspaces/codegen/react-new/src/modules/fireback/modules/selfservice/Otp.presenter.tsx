import { FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { usePostWorkspacePassportOtp } from "../../sdk/modules/workspaces/usePostWorkspacePassportOtp";
import {
  ClassicPassportOtpActionReqDto,
  ClassicPassportOtpActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";
import { useCompleteAuth } from "./auth.common";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const s = useS(strings);
  const { submit: singin, mutation } = usePostWorkspacePassportOtp();
  const { onComplete } = useCompleteAuth();
  const { setSession } = useContext(RemoteQueryContext);

  const form = useRef<FormikProps<
    Partial<ClassicPassportOtpActionReqDto>
  > | null>();
  const setFormRef = (
    ref: FormikProps<Partial<ClassicPassportOtpActionReqDto>>
  ) => {
    form.current = ref;
  };

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

  const submit = (values: Partial<ClassicPassportOtpActionReqDto>) => {
    singin({ ...values, value: state.value })
      .then(successful)
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    mutation,
    setFormRef,
    form,
    s,
    submit,
    goBack,
  };
};
