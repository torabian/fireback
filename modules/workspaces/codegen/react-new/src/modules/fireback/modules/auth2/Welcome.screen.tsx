import { FormikProps } from "formik";
import { WithForm } from "../../components/forms/WithForm";
import { source } from "../../hooks/source";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { AuthLoader } from "../auth/AuthLoader";
import { AuthAvailableMethods, AuthMethod } from "./auth.common";
import { usePresenter } from "./Welcome.presenter";
import { QueryErrorView } from "../../components/error-view/QueryError";

export const WelcomeScreen = () => {
  const {
    onSelect,
    availableOptions,
    totalAvailableMethods,
    isLoadingMethods,
    passportMethodsQuery,
  } = usePresenter();

  if (passportMethodsQuery.isError) {
    return (
      <div className="signin-form-container">
        <QueryErrorView query={passportMethodsQuery} />
      </div>
    );
  }

  if (totalAvailableMethods === undefined || isLoadingMethods) {
    return (
      <div className="signin-form-container">
        <AuthLoader />
      </div>
    );
  }

  if (totalAvailableMethods === 0) {
    return (
      <div className="signin-form-container">
        <NoMethodAvailable />
      </div>
    );
  }

  return (
    <div className="signin-form-container">
      <WithForm
        Form={(props) => (
          <Form
            availableOptions={availableOptions}
            onSelect={onSelect}
            {...props}
          />
        )}
      />
    </div>
  );
};

const Form = ({
  form,
  onSelect,
  availableOptions,
}: {
  form: FormikProps<Partial<ClassicSigninActionReqDto>>;
  onSelect: (method: AuthMethod) => void;
  availableOptions: AuthAvailableMethods;
}) => {
  return (
    <>
      <h1>Welcome back</h1>
      <p>Select any option to continue. </p>
      <div
        role="group"
        aria-label="Login method"
        className="flex gap-2 login-option-buttons"
      >
        {availableOptions.email ? (
          <button id="using-email" onClick={() => onSelect(AuthMethod.Email)}>
            Email
          </button>
        ) : null}

        {availableOptions.phone ? (
          <button id="using-phone" onClick={() => onSelect(AuthMethod.Phone)}>
            Phone number
          </button>
        ) : null}

        {availableOptions.google ? (
          <button id="using-google" disabled>
            <img className="button-icon" src={source("/common/google.png")} />
            Google
          </button>
        ) : null}
      </div>
    </>
  );
};

const NoMethodAvailable = () => {
  return (
    <>
      <h1>Authentication Currently Unavailable</h1>
      <p>
        Sign-in and registration are not available in your region at this time.
        If you believe this is an error or need access, please contact the
        administrator.
      </p>
    </>
  );
};
