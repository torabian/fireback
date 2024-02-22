import { ErrorsView } from "@/components/error-view/ErrorView";
import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { useQueryClient } from "react-query";

import Link from "@/components/link/Link";
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

import { FormSelect } from "@/components/forms/form-select/FormSelect";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { getAuthOtpMethods } from "./AuthHooks";
import { AuthLoader } from "./AuthLoader";
import { TimerUntil } from "./TimerUntil";
import { usePostPassportRequestResetMailPassword } from "@/sdk/fireback/modules/workspaces/usePostPassportRequestResetMailPassword";

const initialValues: Partial<OtpAuthenticateDto> = {
  otp: "",
  type: "email",
  value: "",
};

enum TwoFactorState {
  Initial = "initial",
  VerifyCode = "verify",
}

export const OtpPassword = ({
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
    submit(values, formikProps as any).then((response) => {
      if (response.data) {
      }
    });
  };
  let stage: TwoFactorState = TwoFactorState.Initial;

  stage =
    mutation.data || mutation.error
      ? TwoFactorState.VerifyCode
      : TwoFactorState.Initial;

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
        return (
          <form
            className="signup-form"
            onSubmit={(e) => {
              e.preventDefault();

              // formik.submitForm();
            }}
          >
            <div className="signup-wrapper wrapper-center-content">
              <PageSection title="">
                {mutation.isLoading && <AuthLoader />}

                <div className="form-login-ui">
                  <h1 className="signup-title">{t.abac.otpTitle}</h1>
                  <p>{t.abac.otpTitleHint}</p>

                  <ErrorsView errors={formik.errors} />

                  <FormSelect
                    value={values.type}
                    type="verbose"
                    onChange={(value) => setFieldValue("type", value, false)}
                    errorMessage={errors.type}
                    options={getAuthOtpMethods(t)}
                    name="type"
                    label={t.abac.otpResetMethod}
                  />

                  <div className="row">
                    {formik.values.type === "email" ? (
                      <div className="col-12">
                        <FormText
                          disabled={stage === TwoFactorState.VerifyCode}
                          label={t.abac.emailAddress}
                          autoFocus
                          type="email"
                          dir="ltr"
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
                          hint={t.wokspaces.invite.phoneNumberHint}
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

                  {stage !== TwoFactorState.Initial && (
                    <div onClick={resetAll}>
                      <span>{t.abac.otpOrDifferent}</span>
                    </div>
                  )}

                  {showEnterCode ? (
                    <ReactCodeInput
                      className="otp-react-code-input"
                      values={(values.otp || "").split("")}
                      onChange={(value) => setFieldValue("otp", value, false)}
                    />
                  ) : null}

                  {blockedUntil && (
                    <TimerUntil
                      onResend={() => formik.submitForm()}
                      until={blockedUntil}
                    />
                  )}

                  <FormButton
                    disabled={!formik.values.value}
                    isSubmitting={mutation.isLoading}
                    onClick={() => formik.submitForm()}
                    label={t.requestReset}
                  />

                  <div className="auth-form-helper">
                    <span
                      style={{
                        textAlign: "center",
                      }}
                    >
                      {t.alreadyHaveAnAccount}
                    </span>
                    <Link className="btn btn-secondary" href="/signin">
                      {t.signinInstead}
                    </Link>
                  </div>
                </div>
              </PageSection>
            </div>
          </form>
        );
      }}
    </Formik>
  );
};
