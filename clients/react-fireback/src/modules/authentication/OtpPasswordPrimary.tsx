import { ErrorsView } from "@/components/error-view/ErrorView";
import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect, useRef, useState } from "react";
import { useQueryClient } from "react-query";

import { PageSection } from "@/components/page-section/PageSection";
import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";
import ReactCodeInput from "react-verification-code-input";

import {
  IResponse,
  OtpAuthenticateDto,
  UserSessionDto,
  WorkspaceInviteEntity,
} from "src/sdk/fireback";

import { source } from "@/helpers/source";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { usePostPassportRequestResetMailPassword } from "src/sdk/fireback/modules/workspaces/usePostPassportRequestResetMailPassword";
import { getAuthOtpMethods } from "./AuthHooks";
import { AuthLoader } from "./AuthLoader";
import { TimerUntil } from "./TimerUntil";
import { OtpEmailPasswordInput } from "./OtpEmailPasswordInput";

const initialValues: Partial<OtpAuthenticateDto> = {
  otp: "",
  type: "",
  value: "",
};

enum TwoFactorState {
  Initial = "initial",
  VerifyCode = "verify",
  PasswordOrOtp = "passwordOrOtp",
}

export const OtpPasswordPrimary = ({
  onSuccess,
  invite,
}: {
  onSuccess?: (d: IResponse<UserSessionDto>) => void;
  invite?: WorkspaceInviteEntity;
}) => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();
  const queryClient = useQueryClient();
  const formik = useRef<FormikProps<Partial<OtpAuthenticateDto>> | null>();
  const passwordRef = useRef<any | null>();
  const { setSession, session, isAuthenticated } =
    useContext(RemoteQueryContext);

  const [state, setState] = useState<"VALUE_ENTERED" | null>(null);

  const { submit, mutation } = usePostPassportRequestResetMailPassword({
    queryClient,
  });

  useEffect(() => {
    formik.current?.setValues({
      ...formik.current.values,
      // firstName: invite?.firstName,
      // lastName: invite?.lastName,
      // email: invite?.email,
    });
  }, []);

  const resetAll = () => {
    mutation.reset();
    formik.current?.resetForm();
  };

  const onSubmit = (
    values: Partial<OtpAuthenticateDto>,
    formikProps: FormikHelpers<Partial<OtpAuthenticateDto>>
  ) => {
    submit(values, formikProps as any)
      .then((response) => {
        if (response.data) {
        }
      })
      .catch((res) => {
        if (res.error.httpCode === 404) {
          router.push(
            `/${locale}/signup/email?email=${formik.current?.values.value}`
          );
        }
      });
  };
  let stage: TwoFactorState = TwoFactorState.Initial;

  stage =
    mutation.data || mutation.error
      ? TwoFactorState.VerifyCode
      : TwoFactorState.Initial;

  if (state === "VALUE_ENTERED") {
    stage = TwoFactorState.PasswordOrOtp;
  }
  const blockedUntil =
    mutation.error?.data?.request?.blockedUntil ||
    mutation.data?.data?.request?.blockedUntil;

  const showEnterCode =
    mutation.error?.data ||
    mutation.data?.data ||
    stage === TwoFactorState.VerifyCode;

  return (
    <Formik
      innerRef={(p) => {
        if (p) formik.current = p;
      }}
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      {(formik: FormikProps<Partial<OtpAuthenticateDto>>) => {
        const { values, setFieldValue, errors } = formik;
        const hasActiveOtp = formik.values.type;
        return (
          <div className="signup-form">
            <div className="signup-wrapper wrapper-center-content">
              <PageSection title="">
                {mutation.isLoading && <AuthLoader />}

                <div className="form-login-ui">
                  {stage === TwoFactorState.Initial && (
                    <>
                      <h1 className="signup-title">
                        Get access to your account
                      </h1>
                      <p>
                        Select a method of login below, and then you would be
                        able to access your account. If you do not have account
                        yet, here will automatically build one for you.
                      </p>
                    </>
                  )}

                  <ErrorsView errors={formik.errors} />

                  {!hasActiveOtp && stage === TwoFactorState.Initial && (
                    <ul className="otp-methods">
                      {getAuthOtpMethods(t).map((method) => {
                        return (
                          <li
                            key={method.value}
                            className="otp-method"
                            onClick={() => setFieldValue("type", method.value)}
                          >
                            <img
                              src={source("common/" + method.value + ".svg")}
                              alt=""
                            />
                            <span>{method.label}</span>
                          </li>
                        );
                      })}
                    </ul>
                  )}

                  <div className="row">
                    {formik.values.type === "email" &&
                    stage !== TwoFactorState.VerifyCode &&
                    stage !== TwoFactorState.PasswordOrOtp ? (
                      <div className="col-12">
                        <FormText
                          label="Email address"
                          autoFocus
                          errorMessage={formik.errors.value}
                          value={formik.values.value}
                          onChange={(value) =>
                            formik.setFieldValue("value", value, false)
                          }
                        />
                      </div>
                    ) : null}
                    {formik.values.type === "sms" ? (
                      <div className="col-12">
                        <FormText
                          value={formik.values.value}
                          disabled={stage === TwoFactorState.VerifyCode}
                          onChange={(value) =>
                            formik.setFieldValue("value", value, false)
                          }
                          errorMessage={formik.errors.value}
                          type="phonenumber"
                          label={t.wokspaces.invite.phoneNumber}
                        />
                        {/* <Form
                          label="Phone number"
                          autoFocus
                          errorMessage={formik.errors.value}
                          value={formik.values.value}
                          onChange={(value) =>
                            formik.setFieldValue("value", value, false)
                          }
                        /> */}
                      </div>
                    ) : null}
                  </div>

                  {/* {stage !== TwoFactorState.Initial && (
                    <div onClick={resetAll}>
                      <span>Or try a different account instead</span>
                    </div>
                  )} */}

                  {stage === TwoFactorState.PasswordOrOtp ? (
                    <>
                      <OtpEmailPasswordInput />
                      <br />
                      <br />
                      <br />
                      <div className="step-header">
                        <span>2</span> One time code
                      </div>
                      <span>
                        If you do not remember the password, you can click to
                        send an otp code, and type it in to continue:
                      </span>

                      <FormButton
                        disabled={!formik.values.value || !!blockedUntil}
                        isSubmitting={mutation.isLoading}
                        onClick={() => formik.submitForm()}
                        label={t.requestReset}
                        className="mt-3"
                      />
                      <ReactCodeInput
                        className="otp-react-code-input"
                        values={(values.otp || "").split("")}
                        onChange={(value) => setFieldValue("otp", value, false)}
                      />
                    </>
                  ) : null}

                  {blockedUntil && (
                    <TimerUntil
                      onResend={() => formik.submitForm()}
                      until={blockedUntil}
                    />
                  )}

                  {stage === TwoFactorState.Initial && formik.values.type && (
                    <FormButton
                      disabled={!formik.values.value}
                      isSubmitting={mutation.isLoading}
                      onClick={() => {
                        setState("VALUE_ENTERED");
                        formik.submitForm();
                      }}
                      label={t.continue}
                    />
                  )}

                  {hasActiveOtp && (
                    <FormButton
                      onClick={() => resetAll()}
                      label={"Other method"}
                      className="mt-5"
                      type="secondary"
                    />
                  )}
                </div>
              </PageSection>
            </div>
          </div>
        );
      }}
    </Formik>
  );
};
