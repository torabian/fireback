import { useFormik } from "formik";
import { mutationErrorsToFormik } from "../../hooks/api";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";



import type { GResponse } from "../../sdk/sdk/envelopes";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";
import { ClassicPassportOtpActionReq, useClassicPassportOtpAction, type ClassicPassportOtpActionRes } from "../../sdk/modules/abac/ClassicPassportOtp";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const s = useS(strings);
  const mutation = useClassicPassportOtpAction({});
  const { onComplete } = useCompleteAuth();

  const submit = (values: Partial<ClassicPassportOtpActionReq>) => {
    mutation.mutateAsync(new ClassicPassportOtpActionReq({ ...values, value: state.value }))
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const form = useFormik<Partial<ClassicPassportOtpActionReq>>({
    initialValues: {},
    onSubmit: submit,
  });

  const successful = (res: GResponse<ClassicPassportOtpActionRes>) => {
    if (res.data?.item.session) {
      onComplete(res);
    } else if (res.data?.item?.continueWithCreation) {
      push(`/${locale}/selfservice/complete`, undefined, {
        value: state.value,
        type: state.type,
        sessionSecret: res.data.item?.sessionSecret,
        totpUrl: res.data.item?.totpUrl,
      });
    }
  };

  return {
    mutation,
    form,
    s,
    submit,
    goBack,
  };
};
