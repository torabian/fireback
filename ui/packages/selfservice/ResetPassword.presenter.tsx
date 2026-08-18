import { useFormik } from "formik";
import { useEffect } from "react";
import { useSearchParams } from "react-router-dom";
import type { GResponse } from "@fireback/js-remote-ctx/envelopes";
import { mutationErrorsToFormik } from "@fireback/ui-core/hooks/api";
import { useS } from "@fireback/ui-core/hooks/useS";
import {
  CompletePassportPasswordResetActionReq,
  type CompletePassportPasswordResetActionRes,
  useCompletePassportPasswordResetAction,
} from "@fireback/selfservice/sdk/abac/CompletePassportPasswordResetAction";
import { useCompleteAuth } from "./auth.common";
import { strings } from "./strings/translations";

// Lands here from the email/SMS link ClassicPassportRequestOtpAction/
// SendPassportResetEmailAction send (modules/abac/ClassicPassportRequestOtpActionImplementation.go
// builds it from config.SelfServiceBaseUrl) - a real page load, not in-app
// navigation, so the passport value comes from the URL's own query string
// (?value=...) rather than router state the way most other selfservice screens
// read it (see e.g. Otp.presenter.tsx's state.value).
export const usePresenter = () => {
  const s = useS(strings);
  const [searchParams] = useSearchParams();
  const mutation = useCompletePassportPasswordResetAction();
  const { onComplete } = useCompleteAuth();

  const submit = (values: Partial<CompletePassportPasswordResetActionReq>) => {
    mutation
      .mutateAsync(new CompletePassportPasswordResetActionReq(values))
      .then(successful)
      .catch((error) => {
        form?.setErrors(mutationErrorsToFormik(error));
      });
  };

  const successful = (
    res: GResponse<CompletePassportPasswordResetActionRes>,
  ) => {
    if (res?.data?.item?.reset) {
      onComplete(res);
    }
  };

  const form = useFormik<Partial<CompletePassportPasswordResetActionReq>>({
    initialValues: {},
    onSubmit: submit,
  });

  useEffect(() => {
    const value = searchParams.get("value");
    if (value) {
      form.setFieldValue(
        CompletePassportPasswordResetActionReq.Fields.value,
        value,
        false,
      );
    }
    // Only meant to prefill once, on load - not every time form identity changes.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams]);

  return {
    mutation,
    form,
    submit,
    s,
  };
};
