import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { useS } from "../../hooks/useS";
import ReactCodeInput from "../../thirdparty/react-verification-code-input";
import { strings } from "./strings/translations";
import { usePresenter } from "./TotpEnter.presenter";
import { ConfirmClassicPassportTotpActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";

export const TotpEnter = ({}: {}) => {
  const { goBack, mutation, form } = usePresenter();
  const s = useS(strings);

  return (
    <div className="signin-form-container">
      <h1>{s.enterTotp}</h1>
      <p>{s.enterTotpDescription}</p>
      <QueryErrorView query={mutation} />

      <Form form={form} mutation={mutation} />

      <button
        id="go-back-button"
        className="btn  w-100 d-block"
        onClick={goBack}
      >
        {s.anotherAccount}
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
}: {
  form: FormikProps<Partial<ConfirmClassicPassportTotpActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
}) => {
  const disabled = !form.values.totpCode || form.values.totpCode.length != 6;
  const s = useS(strings);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
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
        className="btn btn-success w-100 d-block mb-2"
        mutation={mutation}
        id="submit-form"
        disabled={disabled}
      >
        {s.continue}
      </FormButton>
    </form>
  );
};
