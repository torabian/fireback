import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { usePresenter } from "./ClassicSigninPassword.presenter";
import { WithForm } from "../../components/forms/WithForm";
import { QueryErrorView } from "../../components/error-view/QueryError";

export const ClassicSigninPassword = ({}: {}) => {
  const { goBack, submit, mutation, setFormRef, continueWithOtp } =
    usePresenter();

  return (
    <div className="signin-form-container">
      <h1>Enter password</h1>
      <p>Enter your password to continue to the system</p>
      <QueryErrorView query={mutation} />
      <WithForm
        setFormRef={setFormRef}
        onSubmit={submit}
        Form={(props) => (
          <Form
            {...props}
            mutation={mutation}
            continueWithOtp={continueWithOtp}
          />
        )}
      />
      <button
        id="back-to-general-step"
        onClick={goBack}
        className="bg-transparent border-0"
      >
        Change the email address
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
  continueWithOtp,
}: {
  form: FormikProps<Partial<ClassicSigninActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
  continueWithOtp: () => void;
}) => {
  const disabled = !form.values.value || !form.values.password;

  return (
    <div>
      <FormText
        type="password"
        value={form.values.password}
        label="Password"
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
        Continue
      </FormButton>

      <button
        onClick={continueWithOtp}
        className="bg-transparent border-0 mt-3 mb-3"
      >
        Use one time password instead
      </button>
    </div>
  );
};
