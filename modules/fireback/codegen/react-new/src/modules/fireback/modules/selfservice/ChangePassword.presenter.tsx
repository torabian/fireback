import { useFormik } from "formik";
import { useEffect } from "react";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { ChangePasswordActionReq, useChangePasswordAction } from "../../sdk/modules/abac/ChangePassword";
import { strings } from "./strings/translations";


export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, state, replace, push, query } = useRouter();
  const mutation = useChangePasswordAction();
  const uniqueId = query?.uniqueId as string;

  const submit = () => {
    mutation.mutateAsync(new ChangePasswordActionReq(form.values)).then((res) => {
      goBack();
    });
  };

  const form = useFormik<ChangePasswordActionReq>({
    initialValues: {},
    onSubmit: submit,
  });

  useEffect(() => {
    if (!uniqueId || !form) {
      return;
    }

    form.setFieldValue(ChangePasswordActionReq.Fields.uniqueId, uniqueId);
  }, [uniqueId]);

  return {
    mutation,
    form,
    submit,
    goBack,
    s,
  };
};
