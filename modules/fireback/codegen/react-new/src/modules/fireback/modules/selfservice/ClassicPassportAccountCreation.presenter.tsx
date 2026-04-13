import { useFormik } from "formik";
import { useEffect, useState } from "react";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";

import { useS } from "../../hooks/useS";

import { ClassicSignupActionReq, ClassicSignupActionRes, useClassicSignupAction } from "../../sdk/modules/abac/ClassicSignup";
import { QueryWorkspaceTypesPubliclyActionRes, useQueryWorkspaceTypesPubliclyActionQuery } from "../../sdk/modules/abac/QueryWorkspaceTypesPublicly";
import type { GResponse } from "../../sdk/sdk/envelopes";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const { goBack, state, push } = useRouter();
  const { locale } = useLocale();
  const { onComplete } = useCompleteAuth();
  const mutation = useClassicSignupAction();
  const totpUrl = state?.totpUrl;
  // const { items: workspaceTypes, query } = useGetWorkspacePublicTypes({
  //   unauthorized: true,
  // });

  const { data, isLoading } = useQueryWorkspaceTypesPubliclyActionQuery({});
  const workspaceTypes: QueryWorkspaceTypesPubliclyActionRes[] = data?.data?.items || []

  const s = useS(strings);

  // The external service requests specific workspace type.
  const requestedWorkspaceTypeId = sessionStorage.getItem("workspace_type_id");

  const submit = (values: Partial<ClassicSignupActionReq>) => {
    mutation.mutateAsync(new ClassicSignupActionReq({
      ...values,
      value: state?.value,
      workspaceTypeId,
      type: state?.type,
      sessionSecret: state?.sessionSecret,
    }))
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicSignupActionReq>>({
    initialValues: {},
    onSubmit: submit,
  });


  // Previous screen sends the email/phone here
  useEffect(() => {
    form?.setFieldValue(ClassicSignupActionReq.Fields.value, state?.value);
  }, [state?.value]);

  // we expect either the account completion is successful this stage
  // only catch is, if the server requires totp (dual factor)
  const successful = (res: GResponse<ClassicSignupActionRes>) => {
    if (res.data.item.session) {
      onComplete(res);
    } else if (res.data.item.continueToTotp) {
      push(`/${locale}/selfservice/totp-setup`, undefined, {
        totpUrl: res.data.item.totpUrl || totpUrl,
        forcedTotp: res.data.item.forcedTotp,
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
