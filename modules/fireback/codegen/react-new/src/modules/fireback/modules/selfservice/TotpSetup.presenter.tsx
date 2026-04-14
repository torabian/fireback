import { useFormik } from "formik";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useRouter } from "../../hooks/useRouter";
import { ConfirmClassicPassportTotpActionReq, useConfirmClassicPassportTotpAction, type ConfirmClassicPassportTotpActionRes } from "../../sdk/modules/abac/ConfirmClassicPassportTotp";
import type { GResponse } from "../../sdk/sdk/envelopes";
import { useCompleteAuth } from "./auth.common";

export const usePresenter = () => {
  const { goBack, state } = useRouter();
  const mutation = useConfirmClassicPassportTotpAction();
  const { onComplete } = useCompleteAuth();

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const submit = (values: Partial<ConfirmClassicPassportTotpActionReq>) => {
    mutation.mutateAsync(new ConfirmClassicPassportTotpActionReq({ ...values, password, value }))
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ConfirmClassicPassportTotpActionReq>>({
    initialValues: {},
    onSubmit: submit,
  });

  const successful = (
    res: GResponse<ConfirmClassicPassportTotpActionRes>
  ) => {
    if (res.data?.item?.session) {
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
