import { useFormik } from "formik";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { usePostPassportsSigninClassic } from "../../sdk/modules/workspaces/usePostPassportsSigninClassic";
import {
  ClassicSigninActionReqDto,
  ClassicSigninActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { useCompleteAuth } from "./auth.common";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { submit: signin, mutation } = usePostPassportsSigninClassic();
  const { onComplete } = useCompleteAuth();

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const submit = (values: Partial<ClassicSigninActionReqDto>) => {
    signin({ ...values, password, value })
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicSigninActionReqDto>>({
    initialValues: {},
    onSubmit: signin,
  });

  const successful = (res: IResponse<ClassicSigninActionResDto>) => {
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
