import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { useT } from "../../hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useRef, useState } from "react";
import { useQueryClient } from "react-query";

import { QueryErrorView } from "../../components/error-view/QueryError";
import Link from "../../components/link/Link";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";

import { IResponse } from "../../definitions/JSONStyle";
import { httpErrorHanlder } from "../../hooks/api";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { EmailAccountSigninDto } from "../../sdk/modules/workspaces/EmailAccountSigninDto";
import { usePostPassportAuthorizeOs } from "../../sdk/modules/workspaces/usePostPassportAuthorizeOs";
import { usePostPassportsSigninClassic } from "../../sdk/modules/workspaces/usePostPassportsSigninClassic";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { getCachedCredentials, useRememberingLoginForm } from "./AuthHooks";
import { AuthLoader } from "./AuthLoader";
import { UserOsProfileCard, UserProfileCard } from "./UserProfileCard";

export const Signin = ({
  onSuccess,
  onRemoteChange,
}: {
  onSuccess?: (d: IResponse<any>) => void;
  onRemoteChange?: (mode: "ipc" | "remote") => void;
}) => {
  const initialValues: Partial<EmailAccountSigninDto> = getCachedCredentials();
  console.log(25, initialValues);
  // const initialValues: Partial<EmailAccountSigninDto> = {};

  const t = useT();
  const router = useRouter();
  const [name, setName] = useState("");
  const { locale } = useLocale();
  const queryClient = useQueryClient();
  const formik = useRef<FormikProps<Partial<EmailAccountSigninDto>> | null>();
  const passwordRef = useRef<any | null>();
  // const { setSession, ref, isAuthenticated } = useContext(AuthContext);
  const { RememberSwitch, shouldRemember } = useRememberingLoginForm(formik);
  const { setSession, session, isAuthenticated } =
    useContext(RemoteQueryContext);

  const {
    submit: submitPostPassportSigninEmail,
    mutation: mutationPostPassportSigninEmail,
  } = usePostPassportsSigninClassic({ queryClient });

  const { submit: osAuthorizeSubmit, mutation: osAuthorizeMutation } =
    usePostPassportAuthorizeOs({ queryClient });

  const onSubmit = (
    values: Partial<EmailAccountSigninDto>,
    formikProps: FormikHelpers<Partial<EmailAccountSigninDto>>
  ) => {
    submitPostPassportSigninEmail(values, formikProps as any)
      .then((response) => {
        if (response.data) {
          if (shouldRemember) {
            localStorage.setItem(
              "credentials",
              JSON.stringify({
                value: (values as any).value,
                password: values.password,
              })
            );
          } else {
            formik.current?.setValues({ email: "", password: "" });
          }
          setSession((response as any).data);

          onSuccess && onSuccess(response);

          if (process.env.REACT_APP_DEFAULT_ROUTE) {
            const to = (
              process.env.REACT_APP_DEFAULT_ROUTE || "/{locale}/signin"
            ).replace("{locale}", locale || "en");

            router.replace(to, to);
          }
        }
      })
      .catch((e: any) => httpErrorHanlder(e, t));
  };

  const osSubmit = () => {
    osAuthorizeSubmit({})
      .then((response) => {
        if (response.data) {
          setSession((response as any).data);

          onSuccess && onSuccess(response);
          if (process.env.REACT_APP_DEFAULT_ROUTE) {
            router.replace(
              process.env.REACT_APP_DEFAULT_ROUTE,
              process.env.REACT_APP_DEFAULT_ROUTE
            );
          }
        }
      })
      .catch((e: any) => httpErrorHanlder(e, t));
  };

  return (
    <Formik
      innerRef={(p) => {
        if (p) formik.current = p;
      }}
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      {(formik: FormikProps<Partial<ClassicSigninActionReqDto>>) => {
        if (!formik.values) {
          return null;
        }
        return (
          <div className="signup-form">
            <div className="signup-wrapper">
              {isAuthenticated ? (
                <UserProfileCard />
              ) : (
                <div className="form-login-ui">
                  {/* <img className="product-logo" src="/logo.svg" /> */}

                  {process.env.REACT_APP_ALLOW_OS_LOGIN === "true" && (
                    <div className="login-form-section">
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
                      className="login-form-section"
                    >
                      <QueryErrorView query={mutationPostPassportSigninEmail} />

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
                        value={formik.values?.value}
                        errorMessage={formik.errors.value}
                        onChange={(value) =>
                          formik.setFieldValue(
                            ClassicSigninActionReqDto.Fields.value,
                            value,
                            false
                          )
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
                        isSubmitting={mutationPostPassportSigninEmail.isLoading}
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
                  {/* <AuthDebug /> */}
                </div>
              )}
            </div>
          </div>
        );
      }}
    </Formik>
  );
};
