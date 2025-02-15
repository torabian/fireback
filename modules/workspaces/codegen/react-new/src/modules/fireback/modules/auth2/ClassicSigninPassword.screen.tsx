import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { usePresenter } from "./ClassicSigninPassword.presenter";
import { WithForm } from "../../components/forms/WithForm";

export const ClassicSigninPassword = ({}: {}) => {
  const { goBack, submit, mutation, setFormRef } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>Enter password</h1>
      <p>Enter your password to continue to the system</p>
      <WithForm
        setFormRef={setFormRef}
        onSubmit={submit}
        Form={(props) => <Form {...props} mutation={mutation} />}
      />
      <button
        id="back-to-general-step"
        className="btn btn-secondary w-100 d-block"
        onClick={goBack}
      >
        Change the email address
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
}: {
  form: FormikProps<Partial<ClassicSigninActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
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
    </div>
  );
};
