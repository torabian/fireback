import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { WithForm } from "../../components/forms/WithForm";
import { useS } from "../../hooks/useS";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { usePresenter } from "./ClassicSigninPassword.presenter";
import { strings } from "./strings/translations";

export const ClassicSigninPassword = ({}: {}) => {
  const {
    goBack,
    submit,
    mutation,
    setFormRef,
    continueWithOtp,
    otpEnabled,
    s,
  } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>{s.enterPassword}</h1>
      <p>{s.enterPasswordDescription}</p>
      <QueryErrorView query={mutation} />
      <WithForm
        setFormRef={setFormRef}
        onSubmit={submit}
        Form={(props) => (
          <Form
            {...props}
            mutation={mutation}
            continueWithOtp={continueWithOtp}
            otpEnabled={otpEnabled}
          />
        )}
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
    <div>
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
        onClick={() => form.submitForm()}
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
    </div>
  );
};
