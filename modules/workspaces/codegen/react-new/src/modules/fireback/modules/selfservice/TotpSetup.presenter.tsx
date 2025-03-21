import { FormikProps } from "formik";
import { useRef } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { usePostPassportTotpConfirm } from "../../sdk/modules/workspaces/usePostPassportTotpConfirm";
import {
  ConfirmClassicPassportTotpActionReqDto,
  ConfirmClassicPassportTotpActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { useCompleteAuth } from "./auth.common";

export const usePresenter = () => {
  const { goBack, state } = useRouter();
  const { submit: confirm, mutation } = usePostPassportTotpConfirm();
  const { onComplete } = useCompleteAuth();

  const form = useRef<FormikProps<
    Partial<ConfirmClassicPassportTotpActionReqDto>
  > | null>();
  const setFormRef = (
    ref: FormikProps<Partial<ConfirmClassicPassportTotpActionReqDto>>
  ) => {
    form.current = ref;
  };

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const successful = (
    res: IResponse<ConfirmClassicPassportTotpActionResDto>
  ) => {
    if (res.data?.session) {
      onComplete(res);
    }
  };

  const submit = (values: Partial<ConfirmClassicPassportTotpActionReqDto>) => {
    confirm({ ...values, password, value })
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
