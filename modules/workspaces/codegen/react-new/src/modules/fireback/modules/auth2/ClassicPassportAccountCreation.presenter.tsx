import { FormikProps } from "formik";
import { useContext, useEffect, useRef, useState } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useGetWorkspacePublicTypes } from "../../sdk/modules/workspaces/useGetWorkspacePublicTypes";
import { usePostPassportsSignupClassic } from "../../sdk/modules/workspaces/usePostPassportsSignupClassic";
import {
  ClassicSignupActionReqDto,
  ClassicSignupActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const { submit: signup, mutation } = usePostPassportsSignupClassic();
  const totpUrl = state?.totpUrl;
  const { items: workspaceTypes, query } = useGetWorkspacePublicTypes({
    unauthorized: true,
  });
  const s = useS(strings);

  const { setSession } = useContext(RemoteQueryContext);

  const form = useRef<FormikProps<Partial<ClassicSignupActionReqDto>> | null>();
  const setFormRef = (ref: FormikProps<Partial<ClassicSignupActionReqDto>>) => {
    form.current = ref;
  };

  const isLoading = query.isLoading;

  // Previous screen sends the email/phone here
  useEffect(() => {
    form.current?.setFieldValue(
      ClassicSignupActionReqDto.Fields.value,
      state?.value
    );
  }, [state?.value, form.current]);

  // we expect either the account completion is successful this stage
  // only catch is, if the server requires totp (dual factor)
  const successful = (res: IResponse<ClassicSignupActionResDto>) => {
    if (res.data.session) {
      setSession(res.data.session);
      if (process.env.REACT_APP_DEFAULT_ROUTE) {
        const to = (
          process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
        ).replace("{locale}", locale || "en");

        replace(to, to);
      }
    } else if (res.data.continueToTotp) {
      push(`/${locale}/auth/totp-setup`, undefined, {
        totpUrl: res.data.totpUrl || totpUrl,
        forcedTotp: res.data.forcedTotp,
        password: form.current.values.password,
        value: state?.value,
      });
    }
  };

  const [selectedWorkspaceId, setSelected] = useState("");
  const workspaceTypeId =
    workspaceTypes.length === 1
      ? workspaceTypes[0].uniqueId
      : selectedWorkspaceId;

  const submit = (values: Partial<ClassicSignupActionReqDto>) => {
    signup({
      ...values,
      value: state?.value,
      workspaceTypeId,
      type: state?.type,
      sessionSecret: state?.sessionSecret,
    })
      .then(successful)
      .catch((error) => {
        form.current?.setErrors(mutationErrorsToFormik(error));
      });
  };

  return {
    mutation,
    setFormRef,
    isLoading,
    form,
    totpUrl,
    submit,
    goBack,
    s,
    workspaceTypes,
    state,
  };
};
