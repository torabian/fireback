import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { WithForm } from "../../components/forms/WithForm";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import ReactCodeInput from "../../thirdparty/react-verification-code-input";
import { usePresenter } from "./Otp.presenter";

export const OtpScreen = ({}: {}) => {
  const { goBack, submit, mutation, setFormRef } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>Enter OTP</h1>
      <p>We have sent you an one time password, please enter to continue.</p>
      <QueryErrorView query={mutation} />
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
        Try another account
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
    <>
      <ReactCodeInput
        values={form.values.password?.split("")}
        onChange={(value) =>
          form.setFieldValue(
            ClassicSigninActionReqDto.Fields.password,
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
        Continue
      </FormButton>
    </>
  );
};
