import { ErrorsView } from "../../components/error-view/ErrorView";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { useT } from "../../hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { useQueryClient } from "react-query";

import Link from "../../components/link/Link";
import { PageSection } from "../../components/page-section/PageSection";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import ReactCodeInput from "../../thirdparty/react-verification-code-input";

import { IResponse } from "../../sdk/core/http-tools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { OtpAuthenticateDto } from "../../sdk/modules/workspaces/OtpAuthenticateDto";
import { usePostPassportRequestResetMailPassword } from "../../sdk/modules/workspaces/usePostPassportRequestResetMailPassword";
import { UserSessionDto } from "../../sdk/modules/workspaces/UserSessionDto";
import { WorkspaceInviteEntity } from "../../sdk/modules/workspaces/WorkspaceInviteEntity";
import { getAuthOtpMethods } from "./AuthHooks";
import { AuthLoader } from "./AuthLoader";
import { TimerUntil } from "./TimerUntil";
import { FormSelect } from "../../components/forms/form-select/FormSelect";
import { createQuerySource } from "../../hooks/useAsQuery";

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

  const methods = getAuthOtpMethods(t);
  const otpOptionsSource = createQuerySource(methods);

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
      {(form: FormikProps<Partial<OtpAuthenticateDto>>) => {
        const { values, setFieldValue, errors } = form;
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

                <div className="form-login-ui login-form-section">
                  <h1 className="signup-title">{t.abac.otpTitle}</h1>
                  <p>{t.abac.otpTitleHint}</p>

                  <ErrorsView errors={form.errors} />

                  <FormSelect
                    value={methods.find((item) => values.type === item.value)}
                    type="verbose"
                    onChange={(value) => setFieldValue("type", value, false)}
                    errorMessage={errors.type}
                    querySource={otpOptionsSource}
                    formEffect={{
                      field: "type",
                      form,
                      beforeSet(item) {
                        return item.value;
                      },
                    }}
                    name="type"
                    label={t.abac.otpResetMethod}
                  />

                  <div className="row">
                    {form.values.type === "email" ? (
                      <div className="col-12">
                        <FormText
                          disabled={stage === TwoFactorState.VerifyCode}
                          label={t.abac.emailAddress}
                          autoFocus
                          type="email"
                          dir="ltr"
                          errorMessage={form.errors.value}
                          value={form.values.value}
                          onChange={(value) =>
                            form.setFieldValue("value", value, false)
                          }
                        />
                      </div>
                    ) : null}
                    {form.values.type === "sms" ? (
                      <div className="col-12">
                        <FormText
                          value={form.values.value}
                          disabled={stage === TwoFactorState.VerifyCode}
                          onChange={(value) =>
                            form.setFieldValue("value", value, false)
                          }
                          errorMessage={form.errors.value}
                          type="phonenumber"
                          label={t.wokspaces.invite.phoneNumber}
                          hint={t.wokspaces.invite.phoneNumberHint}
                        />
                        {/* <Form
                          label="Phone number"
                          autoFocus
                          errorMessage={formik.errors.value}
                          value={form.values.value}
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
                      onResend={() => form.submitForm()}
                      until={blockedUntil}
                    />
                  )}

                  <FormButton
                    disabled={!form.values.value}
                    isSubmitting={mutation.isLoading}
                    onClick={() => form.submitForm()}
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
