import { FormikProps } from "formik";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { WithForm } from "../../components/forms/WithForm";
import { source } from "../../hooks/source";
import { useS } from "../../hooks/useS";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { AuthLoader } from "../auth/AuthLoader";
import { AuthAvailableMethods, AuthMethod } from "./auth.common";
import { strings } from "./strings/translations";
import { usePresenter } from "./Welcome.presenter";

export const WelcomeScreen = () => {
  const {
    onSelect,
    availableOptions,
    totalAvailableMethods,
    isLoadingMethods,
    passportMethodsQuery,
  } = usePresenter();

  // if (passportMethodsQuery.isError || passportMethodsQuery.error) {
  //   return (
  //     <div className="signin-form-container">
  //       <QueryErrorView query={passportMethodsQuery} />
  //     </div>
  //   );
  // }

  // if (totalAvailableMethods === undefined || isLoadingMethods) {
  //   return (
  //     <div className="signin-form-container">
  //       <AuthLoader />
  //     </div>
  //   );
  // }

  // if (totalAvailableMethods === 0) {
  //   return (
  //     <div className="signin-form-container">
  //       <NoMethodAvailable />
  //     </div>
  //   );
  // }

  return (
    <div className="signin-form-container">
      <WithForm
        Form={(props) => (
          <>
            <pre>{JSON.stringify(passportMethodsQuery.data, null, 2)}</pre>
            <Form
              availableOptions={availableOptions}
              onSelect={onSelect}
              {...props}
            />
          </>
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
  const s = useS(strings);

  if (!availableOptions) {
    return null;
  }

  return (
    <>
      <h1>{s.welcomeBack}</h1>
      <p>{s.welcomeBackDescription} </p>
      <div
        role="group"
        aria-label="Login method"
        className="flex gap-2 login-option-buttons"
      >
        {availableOptions.email ? (
          <button id="using-email" onClick={() => onSelect(AuthMethod.Email)}>
            {s.emailMethod}
          </button>
        ) : null}

        {availableOptions.phone ? (
          <button id="using-phone" onClick={() => onSelect(AuthMethod.Phone)}>
            {s.phoneMethod}
          </button>
        ) : null}

        {availableOptions.google ? (
          <button id="using-google" disabled>
            <img className="button-icon" src={source("/common/google.png")} />
            {s.google}
          </button>
        ) : null}
      </div>
    </>
  );
};

const NoMethodAvailable = () => {
  const s = useS(strings);
  return (
    <>
      <h1>{s.noAuthenticationMethod}</h1>
      <p>{s.noAuthenticationMethodDescription}</p>
    </>
  );
};
