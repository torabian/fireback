import { GoogleOAuthProvider, useGoogleLogin } from "@react-oauth/google";
import { FormikProps, useFormik } from "formik";
import { useContext } from "react";
import { AuthLoader } from "../../components/auth-loader/AuthLoader";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { source } from "../../hooks/source";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { usePostPassportViaOauth } from "../../sdk/modules/workspaces/usePostPassportViaOauth";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
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
  const form = useFormik({ initialValues: {}, onSubmit: () => {} });

  if (passportMethodsQuery.isError || passportMethodsQuery.error) {
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

  console.log(1, availableOptions);

  return (
    <div className="signin-form-container">
      {availableOptions.googleOAuthClientKey ? (
        <GoogleOAuthProvider clientId={availableOptions.googleOAuthClientKey}>
          <Form
            availableOptions={availableOptions}
            onSelect={onSelect}
            form={form}
          />
        </GoogleOAuthProvider>
      ) : (
        <Form
          availableOptions={availableOptions}
          onSelect={onSelect}
          form={form}
        />
      )}
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
  const { submit, mutation } = usePostPassportViaOauth({});
  const { setSession } = useContext(RemoteQueryContext);
  const { locale } = useLocale();
  const { goBack, state, replace, push } = useRouter();

  const continueWithResult = (token: string) => {
    submit({ service: "google", token })
      .then((res) => {
        setSession(res.data.session);
        if ((window as any).ReactNativeWebView) {
          (window as any).ReactNativeWebView.postMessage(
            JSON.stringify(res.data)
          );
        }
        if (process.env.REACT_APP_DEFAULT_ROUTE) {
          const to = (
            process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
          ).replace("{locale}", locale || "en");
          replace(to, to);
        }
      })
      .catch((err) => {
        alert(err);
      });
  };

  const s = useS(strings);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <h1>{s.welcomeBack}</h1>
      <p>{s.welcomeBackDescription} </p>
      <div
        role="group"
        aria-label="Login method"
        className="flex gap-2 login-option-buttons"
      >
        {availableOptions.email ? (
          <button
            id="using-email"
            type="button"
            onClick={() => onSelect(AuthMethod.Email)}
          >
            {s.emailMethod}
          </button>
        ) : null}

        {availableOptions.phone ? (
          <button
            id="using-phone"
            type="button"
            onClick={() => onSelect(AuthMethod.Phone)}
          >
            {s.phoneMethod}
          </button>
        ) : null}

        {availableOptions.google ? (
          <GoogleLogin continueWithResult={continueWithResult} />
        ) : null}
      </div>
    </form>
  );
};

const GoogleLogin = ({
  continueWithResult,
}: {
  continueWithResult: (value: string) => void;
}) => {
  const s = useS(strings);
  const login = useGoogleLogin({
    onSuccess: (tokenResponse) => {
      continueWithResult(tokenResponse.access_token);
    },
    scope: ["https://www.googleapis.com/auth/userinfo.profile"].join(" "),
  });

  return (
    <>
      <button id="using-google" type="button" onClick={() => login()}>
        <img className="button-icon" src={source("/common/google.png")} />
        {s.google}
      </button>
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
