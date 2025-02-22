import { FormikProps } from "formik";
import { QRCodeSVG } from "qrcode.react";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { WithForm } from "../../components/forms/WithForm";
import { useS } from "../../hooks/useS";
import { ConfirmClassicPassportTotpActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import ReactCodeInput from "../../thirdparty/react-verification-code-input";
import { strings } from "./strings/translations";
import { usePresenter } from "./TotpSetup.presenter";

export const TotpSetup = ({}: {}) => {
  const { goBack, submit, mutation, setFormRef, totpUrl, forcedTotp } =
    usePresenter();

  const s = useS(strings);

  return (
    <div className="signin-form-container">
      <h1>{s.setupTotp}</h1>
      <p>{s.setupTotpDescription}</p>
      <QueryErrorView query={mutation} />
      <WithForm
        setFormRef={setFormRef}
        onSubmit={submit}
        Form={(props) => (
          <Form
            {...props}
            totpUrl={totpUrl}
            mutation={mutation}
            forcedTotp={forcedTotp}
          />
        )}
      />

      <button
        id="go-back-button"
        className="btn  w-100 d-block"
        onClick={goBack}
      >
        Try another account
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
  form: FormikProps<Partial<ConfirmClassicPassportTotpActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
  totpUrl: string;
  forcedTotp: boolean;
}) => {
  const s = useS(strings);
  const disabled = !form.values.totpCode || form.values.totpCode.length != 6;

  return (
    <>
      <center>
        <QRCodeSVG value={totpUrl} width={200} height={200} />
      </center>

      <ReactCodeInput
        values={form.values.totpCode?.split("")}
        onChange={(value) =>
          form.setFieldValue(
            ConfirmClassicPassportTotpActionReqDto.Fields.totpCode,
            value,
            false
          )
        }
        className="otp-react-code-input"
      />

      <FormButton
        className="btn btn-primary w-100 d-block mb-2"
        onClick={() => form.submitForm()}
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
    </>
  );
};
