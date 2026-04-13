import { useFormik } from "formik";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useRouter } from "../../hooks/useRouter";

import { ClassicSigninActionReq, ClassicSigninActionRes, useClassicSigninAction } from "../../sdk/modules/abac/ClassicSignin";
import type { GResponse } from "../../sdk/sdk/envelopes";
import { useCompleteAuth } from "./auth.common";


export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const mutation = useClassicSigninAction();
  const { onComplete } = useCompleteAuth();

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const submit = (values: Partial<ClassicSigninActionReq>) => {
    mutation.mutateAsync(new ClassicSigninActionReq({ ...values, password, value }))
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicSigninActionReq>>({
    initialValues: {},
    onSubmit: (values, helpers) => {
      mutation.mutateAsync(new ClassicSigninActionReq(values))
    },
  });

  const successful = (res: GResponse<ClassicSigninActionRes>) => {
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
