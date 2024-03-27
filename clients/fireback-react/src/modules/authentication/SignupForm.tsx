import { ErrorsView } from "@/components/error-view/ErrorView";
import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";

import { FormikProps } from "formik";

import Link from "@/components/link/Link";
import { PageSection } from "@/components/page-section/PageSection";

import { AuthLoader } from "./AuthLoader";
import { UserProfileCard } from "./UserProfileCard";
import { ClassicSignupActionReqDto } from "@/sdk/fireback/modules/workspaces/WorkspacesActionsDto";

export const SignupForm = ({
  allowEditEmail,
  formik,
  loading,
  isAuthenticated,
  RememberSwitch,
  formDescription,
}: {
  loading: boolean;
  isAuthenticated: boolean;
  formik: FormikProps<Partial<ClassicSignupActionReqDto>>;
  allowEditEmail?: boolean;
  RememberSwitch: any;
  formDescription?: string;
}) => {
  const t = useT();

  return (
    <div className="signup-wrapper">
      <PageSection title="">
        {isAuthenticated ? (
          <UserProfileCard />
        ) : (
          <div className="form-login-ui login-form-section">
            {loading && <AuthLoader />}
            {/* <img className="product-logo" src="/logo.svg" /> */}
            <h1 className="signup-title">{t.abac.signup}</h1>
            <p>{formDescription || t.signup.defaultDescription}</p>

            <ErrorsView errors={formik.errors} />
            <div className="row">
              <div className=" col-sm-12">
                <FormText
                  label={t.abac.firstName}
                  autoFocus
                  errorMessage={formik.errors.firstName}
                  value={formik.values.firstName}
                  onChange={(value) =>
                    formik.setFieldValue("firstName", value, false)
                  }
                />
              </div>
              <div className=" col-sm-12">
                <FormText
                  label={t.abac.lastName}
                  errorMessage={formik.errors.lastName}
                  value={formik.values.lastName}
                  onChange={(value) =>
                    formik.setFieldValue("lastName", value, false)
                  }
                />
              </div>
              <div className=" col-sm-12">
                <FormText
                  label={t.abac.email}
                  disabled={allowEditEmail === false}
                  dir="ltr"
                  errorMessage={formik.errors.value}
                  value={formik.values.value}
                  onChange={(value) =>
                    formik.setFieldValue("value", value, false)
                  }
                />
              </div>
              <div className=" col-sm-12">
                <FormText
                  dir="ltr"
                  type="password"
                  value={formik.values.password}
                  label={t.abac.password}
                  errorMessage={formik.errors.password}
                  onChange={(value) =>
                    formik.setFieldValue("password", value, false)
                  }
                />
              </div>
            </div>
            <div className="remember-me">
              <RememberSwitch />
            </div>

            <FormButton
              isSubmitting={loading}
              onClick={() => formik.submitForm()}
              label={t.signupButton}
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
              <Link className="btn btn-secondary mt-3" href="/otp">
                {t.forgotPassword}
              </Link>
            </div>
          </div>
        )}
      </PageSection>
    </div>
  );
};
