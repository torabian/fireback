import { type FormikProps } from "formik";
import { QRCodeSVG } from "qrcode.react";
import { type UseMutationResult } from "@tanstack/react-query";
import { QueryErrorView } from "../fireback-ui/components/error-view/QueryError";
import { FormButton } from "../fireback-ui/components/forms/form-button/FormButton";
import { useS } from "../fireback-ui/hooks/useS";
import ReactCodeInput from "./components/react-verification-code-input";
import { strings } from "./strings/translations";
import { usePresenter } from "./TotpSetup.presenter";
import { ConfirmClassicPassportTotpActionReq } from "../sdk/abac/ConfirmClassicPassportTotpAction";

export const TotpSetup = ({}: {}) => {
  const { goBack, submit, mutation, form, totpUrl, forcedTotp } =
    usePresenter();

  const s = useS(strings);

  return (
    <div className="signin-form-container">
      <h1>{s.setupTotp}</h1>
      <p>{s.setupTotpDescription}</p>
      <QueryErrorView query={mutation} />

      <Form
        form={form}
        totpUrl={totpUrl}
        mutation={mutation}
        forcedTotp={forcedTotp}
      />

      <button
        id="go-back-button"
        className="btn  w-100 d-block"
        onClick={goBack}
      >
        {s.tryAnotherAccount}
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
  forcedTotp,
  totpUrl,
}: {
  form: FormikProps<Partial<ConfirmClassicPassportTotpActionReq>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
  totpUrl: string;
  forcedTotp: boolean;
}) => {
  const s = useS(strings);
  const disabled = !form.values.totpCode || form.values.totpCode.length != 6;

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <center>
        <QRCodeSVG value={totpUrl} width={200} height={200} />
      </center>

      <ReactCodeInput
        values={form.values.totpCode?.split("")}
        onChange={(value) =>
          form.setFieldValue(
            ConfirmClassicPassportTotpActionReq.Fields.totpCode,
            value,
            false,
          )
        }
        className="otp-react-code-input"
      />

      <FormButton
        className="btn btn-primary w-100 d-block mb-2"
        mutation={mutation}
        id="submit-form"
        disabled={disabled}
      >
        {s.continue}
      </FormButton>

      {forcedTotp !== true && (
        <>
          <p className="mt-4">{s.skipTotpInfo}</p>
          <button className="btn btn-warning w-100 d-block mb-2">
            {s.skipTotpButton}
          </button>
        </>
      )}
    </form>
  );
};
