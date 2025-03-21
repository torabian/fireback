import { FormikProps } from "formik";
import { useRef } from "react";
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

  const form = useRef<FormikProps<Partial<ClassicSigninActionReqDto>> | null>();
  const setFormRef = (ref: FormikProps<Partial<ClassicSigninActionReqDto>>) => {
    form.current = ref;
  };

  const totpUrl = state?.totpUrl;
  const forcedTotp = state?.forcedTotp;
  const password = state?.password;
  const value = state?.value;

  const successful = (res: IResponse<ClassicSigninActionResDto>) => {
    if (res.data?.session) {
      onComplete(res);
    }
  };

  const submit = (values: Partial<ClassicSigninActionReqDto>) => {
    signin({ ...values, password, value })
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
