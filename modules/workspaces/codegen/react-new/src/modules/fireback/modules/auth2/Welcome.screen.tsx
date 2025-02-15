import { FormikProps } from "formik";
import { source } from "../../hooks/source";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { usePresenter } from "./Welcome.presenter";
import { WithForm } from "../../components/forms/WithForm";
import { AuthMethod } from "./auth.common";

export const WelcomeScreen = () => {
  const { onSelect } = usePresenter();

  return <WithForm Form={(props) => <Form onSelect={onSelect} {...props} />} />;
};

const Form = ({
  form,
  onSelect,
}: {
  form: FormikProps<Partial<ClassicSigninActionReqDto>>;
  onSelect: (method: AuthMethod) => void;
}) => {
  return (
    <div className="signin-form-container">
      <h1>Welcome back</h1>
      <p>Select any option to continue. </p>
      <div
        role="group"
        aria-label="Login method"
        className="flex gap-2 login-option-buttons"
      >
        <button id="using-email" onClick={() => onSelect(AuthMethod.Email)}>
          Email
        </button>
        <button id="using-phone" onClick={() => onSelect(AuthMethod.Phone)}>
          Phone number
        </button>
        <button id="using-google" disabled>
          <img className="button-icon" src={source("/common/google.png")} />
          Google
        </button>
      </div>
    </div>
  );
};
