import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useRef, useState } from "react";
import { useQueryClient } from "react-query";

import { QueryErrorView } from "@/components/error-view/QueryError";
import Link from "@/components/link/Link";
import { PageSection } from "@/components/page-section/PageSection";
import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";
import {
  EmailAccountSigninDto,
  IResponse,
  UserSessionDto,
} from "src/sdk/fireback";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { usePostPassportAuthorizeOs } from "src/sdk/fireback/modules/workspaces/usePostPassportAuthorizeOs";
import { usePostPassportSigninEmail } from "src/sdk/fireback/modules/workspaces/usePostPassportSigninEmail";
import { getCachedCredentials, useRememberingLoginForm } from "./AuthHooks";
import { AuthLoader } from "./AuthLoader";
import { UserOsProfileCard, UserProfileCard } from "./UserProfileCard";

export const Signin = ({
  onSuccess,
  onRemoteChange,
}: {
  onSuccess?: (d: IResponse<UserSessionDto>) => void;
  onRemoteChange?: (mode: "ipc" | "remote") => void;
}) => {
  const initialValues: Partial<EmailAccountSigninDto> = getCachedCredentials();

  const t = useT();
  const router = useRouter();
  const [name, setName] = useState("");
  const { locale } = useLocale();
  const queryClient = useQueryClient();
  const formik = useRef<FormikProps<Partial<EmailAccountSigninDto>> | null>();
  const passwordRef = useRef<any | null>();
  // const { setSession, ref, isAuthenticated } = useContext(AuthContext);
  const { RememberSwitch, shouldRemember } = useRememberingLoginForm(formik);
  const { setSession, session, options, isAuthenticated } =
    useContext(RemoteQueryContext);

  const {
    submit: submitPostPassportSigninEmail,
    mutation: mutationPostPassportSigninEmail,
  } = usePostPassportSigninEmail({ queryClient });

  const { submit: osAuthorizeSubmit, mutation: osAuthorizeMutation } =
    usePostPassportAuthorizeOs({ queryClient });

  const onSubmit = (
    values: Partial<EmailAccountSigninDto>,
    formikProps: FormikHelpers<Partial<EmailAccountSigninDto>>
  ) => {
    // onRemoteChange && onRemoteChange("remote");
    // setTimeout(() => {
    submitPostPassportSigninEmail(values, formikProps as any).then(
      (response) => {
        if (response.data) {
          if (shouldRemember) {
            localStorage.setItem(
              "credentials",
              JSON.stringify({
                email: values.email,
                password: values.password,
              })
            );
          } else {
            formik.current?.setValues({ email: "", password: "" });
          }
          // setOptions({ headers: { Authorization: response.data.token } });
          setSession(response.data);
          onSuccess && onSuccess(response);

          if (process.env.REACT_APP_DEFAULT_ROUTE) {
            const to = (
              process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
            ).replace("{locale}", locale || "en");

            router.replace(to, to);
          }
        }
      }
    );
    // }, 300);
  };

  const osSubmit = () => {
    // onRemoteChange && onRemoteChange("ipc");
    // setTimeout(() => {
    osAuthorizeSubmit({}).then((response) => {
      if (response.data) {
        setSession(response.data);
        onSuccess && onSuccess(response);
        if (process.env.REACT_APP_DEFAULT_ROUTE) {
          router.replace(
            process.env.REACT_APP_DEFAULT_ROUTE,
            process.env.REACT_APP_DEFAULT_ROUTE
          );
        }
      }
    });
    // }, 300);
  };

  return (
    <Formik
      innerRef={(p) => {
        if (p) formik.current = p;
      }}
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      {(formik: FormikProps<Partial<EmailAccountSigninDto>>) => {
        return (
          <div className="signup-form">
            <div className="signup-wrapper">
              <PageSection title="">
                {isAuthenticated ? (
                  <UserProfileCard />
                ) : (
                  <div className="form-login-ui">
                    {/* <img className="product-logo" src="/logo.svg" /> */}

                    {process.env.REACT_APP_ALLOW_OS_LOGIN === "true" && (
                      <div style={{ display: "flex", flexDirection: "column" }}>
                        <UserOsProfileCard />
                        <FormButton
                          isSubmitting={osAuthorizeMutation.isLoading}
                          onClick={osSubmit}
                          label={t.loginButtonOs}
                        />{" "}
                        <QueryErrorView query={osAuthorizeMutation} />
                      </div>
                    )}
                    {process.env.REACT_APP_ALLOW_REMOTE_LOGIN !== "false" ? (
                      <form
                        onSubmit={(e) => {
                          e.preventDefault();
                        }}
                        style={{ display: "flex", flexDirection: "column" }}
                      >
                        <QueryErrorView
                          query={mutationPostPassportSigninEmail}
                        />

                        <h1>{t.abac.signin}</h1>

                        {mutationPostPassportSigninEmail.isLoading && (
                          <AuthLoader />
                        )}
                        <FormText
                          label={t.abac.email}
                          autoFocus
                          // pattern="[^ @]*@[^ @]*"
                          type="text"
                          dir="ltr"
                          value={formik.values.email}
                          errorMessage={formik.errors.email}
                          onChange={(value) =>
                            formik.setFieldValue("email", value, false)
                          }
                        />

                        <FormText
                          type="password"
                          value={formik.values.password}
                          dir="ltr"
                          label={t.abac.password}
                          errorMessage={formik.errors.password}
                          onChange={(value) =>
                            formik.setFieldValue("password", value, false)
                          }
                        />

                        <RememberSwitch />

                        <FormButton
                          isSubmitting={
                            mutationPostPassportSigninEmail.isLoading
                          }
                          onClick={() => formik.submitForm()}
                          label={t.loginButton}
                        />

                        <span
                          style={{
                            textAlign: "center",
                          }}
                        >
                          {t.firstTime}
                        </span>
                        <Link className="btn btn-secondary" href="/signup">
                          {t.createAccount}
                        </Link>
                        <Link className="btn btn-secondary mt-3" href="/otp">
                          {t.forgotPassword}
                        </Link>
                      </form>
                    ) : null}
                  </div>
                )}
              </PageSection>
            </div>
          </div>
        );
      }}
    </Formik>
  );
};
