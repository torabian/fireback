import { useFormik } from "formik";
import { useEffect, useState } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { IResponse } from "../../sdk/core/http-tools";
import { useGetWorkspacePublicTypes } from "../../sdk/modules/abac/useGetWorkspacePublicTypes";
import { usePostPassportsSignupClassic } from "../../sdk/modules/abac/usePostPassportsSignupClassic";

import { useS } from "../../hooks/useS";
import {
  ClassicSignupActionReqDto,
  ClassicSignupActionResDto,
} from "../../sdk/modules/abac/AbacActionsDto";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const { goBack, state, push } = useRouter();
  const { locale } = useLocale();
  const { onComplete } = useCompleteAuth();
  const { submit: signup, mutation } = usePostPassportsSignupClassic();
  const totpUrl = state?.totpUrl;
  const { items: workspaceTypes, query } = useGetWorkspacePublicTypes({
    unauthorized: true,
  });
  const s = useS(strings);

  // The external service requests specific workspace type.
  const requestedWorkspaceTypeId = sessionStorage.getItem("workspace_type_id");

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
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicSignupActionReqDto>>({
    initialValues: {},
    onSubmit: submit,
  });

  const isLoading = query.isLoading;

  // Previous screen sends the email/phone here
  useEffect(() => {
    form?.setFieldValue(ClassicSignupActionReqDto.Fields.value, state?.value);
  }, [state?.value]);

  // we expect either the account completion is successful this stage
  // only catch is, if the server requires totp (dual factor)
  const successful = (res: IResponse<ClassicSignupActionResDto>) => {
    if (res.data.session) {
      onComplete(res);
    } else if (res.data.continueToTotp) {
      push(`/${locale}/selfservice/totp-setup`, undefined, {
        totpUrl: res.data.totpUrl || totpUrl,
        forcedTotp: res.data.forcedTotp,
        password: form.values.password,
        value: state?.value,
      });
    }
  };

  const [selectedWorkspaceTypeId, setSelectedWorkspaceType] = useState("");
  let workspaceTypeId =
    workspaceTypes.length === 1
      ? workspaceTypes[0].uniqueId
      : selectedWorkspaceTypeId;

  if (requestedWorkspaceTypeId) {
    workspaceTypeId = requestedWorkspaceTypeId;
  }

  return {
    mutation,
    isLoading,
    form,
    setSelectedWorkspaceType,
    totpUrl,
    workspaceTypeId,
    submit,
    goBack,
    s,
    workspaceTypes,
    state,
  };
};
