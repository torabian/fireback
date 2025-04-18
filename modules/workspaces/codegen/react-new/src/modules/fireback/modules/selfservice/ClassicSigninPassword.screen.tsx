import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { useS } from "../../hooks/useS";
import { usePresenter } from "./ClassicSigninPassword.presenter";
import { strings } from "./strings/translations";
import { ClassicSigninActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";

export const ClassicSigninPassword = ({}: {}) => {
  const { goBack, mutation, form, continueWithOtp, otpEnabled, s } =
    usePresenter();

  return (
    <div className="signin-form-container">
      <h1>{s.enterPassword}</h1>
      <p>{s.enterPasswordDescription}</p>
      <QueryErrorView query={mutation} />

      <Form
        form={form}
        mutation={mutation}
        continueWithOtp={continueWithOtp}
        otpEnabled={otpEnabled}
      />

      <button
        id="go-back-button"
        onClick={goBack}
        className="btn bg-transparent w-100 mt-4"
      >
        {s.anotherAccount}
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
  otpEnabled,
  continueWithOtp,
}: {
  form: FormikProps<Partial<ClassicSigninActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
  continueWithOtp: () => void;
  otpEnabled: boolean;
}) => {
  const s = useS(strings);
  const disabled = !form.values.value || !form.values.password;

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <FormText
        type="password"
        value={form.values.password}
        label={s.password}
        id="password-input"
        autoFocus
        errorMessage={form.errors.password}
        onChange={(value) =>
          form.setFieldValue(
            ClassicSigninActionReqDto.Fields.password,
            value,
            false
          )
        }
      />

      <FormButton
        className="btn btn-primary w-100 d-block mb-2"
        mutation={mutation}
        id="submit-form"
        disabled={disabled}
      >
        {s.continue}
      </FormButton>

      {otpEnabled && (
        <button
          onClick={continueWithOtp}
          className="bg-transparent border-0 mt-3 mb-3"
        >
          {s.useOneTimePassword}
        </button>
      )}
    </form>
  );
};
