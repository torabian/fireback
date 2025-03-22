import { useFormik } from "formik";
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

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const submit = (values: Partial<ConfirmClassicPassportTotpActionReqDto>) => {
    confirm({ ...values, password, value })
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ConfirmClassicPassportTotpActionReqDto>>({
    initialValues: {},
    onSubmit: submit,
  });

  const successful = (
    res: IResponse<ConfirmClassicPassportTotpActionResDto>
  ) => {
    if (res.data?.session) {
      onComplete(res);
    }
  };

  return {
    mutation,
    totpUrl,
    forcedTotp,
    form,
    submit,
    goBack,
  };
};
