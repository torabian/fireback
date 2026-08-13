import { mutationErrorsToFormik } from "@fireback/ui-core/hooks/api";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useFormik } from "formik";

import type { GResponse } from "@fireback/js-remote-ctx/envelopes";
import {
  ClassicPassportOtpActionReq,
  useClassicPassportOtpAction,
  type ClassicPassportOtpActionRes,
} from "@fireback/selfservice/sdk/abac/ClassicPassportOtpAction";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";

import type { MOne } from "@fireback/js-remote-ctx/common/operators";
import type { UserSessionDto } from "./sdk/abac/UserSessionDto";

export const usePresenter = () => {
  const { goBack, state, replace, push } = useRouter();
  const { locale } = useLocale();
  const s = useS(strings);
  const mutation = useClassicPassportOtpAction({});
  const { onComplete } = useCompleteAuth();

  const submit = (values: Partial<ClassicPassportOtpActionReq>) => {
    mutation
      .mutateAsync(
        new ClassicPassportOtpActionReq({ ...values, value: state.value }),
      )
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
    const v = (res.data?.item.session as MOne<UserSessionDto>)?.get().token;
    if (v) {
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
