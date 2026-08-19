import { type FormikProps } from "formik";
import { type UseMutationResult } from "@tanstack/react-query";
import { QueryErrorView } from "@fireback/ui-core/components/error-view/QueryError";
import { FormButton } from "@fireback/ui-core/components/forms/form-button/FormButton";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { useS } from "@fireback/ui-core/hooks/useS";
import { AuthMethod } from "./auth.common";
import { usePresenter } from "./ClassicPassport.presenter";
import { LanguageSwitcher } from "./LanguageSwitcher";
import { strings } from "./strings/translations";
import { ClassicSigninActionReq } from "@fireback/selfservice/sdk/abac/ClassicSigninAction";

export const ClassicPassportScreen = ({ method }: { method: AuthMethod }) => {
  const {
    description,
    title,
    goBack,
    submit,
    mutation,
    form,
    canGoBack,
    LegalNotice,
    Recaptcha,
    s,
  } = usePresenter({
    method,
  });

  return (
    <div className="signin-form-container">
      <div className="card auth-card">
        <div className="card-body">
          {/* Welcome.screen.tsx normally has the first (and only) crack at this -
              but Welcome.presenter.tsx auto-skips straight here, with no user
              interaction, whenever exactly one auth method is configured (see its
              own comment) - which makes this screen the real first landing page
              for that (common) case, and Welcome's own LanguageSwitcher
              unreachable. Shown here too so switching language never depends on
              how many auth methods happen to be enabled. */}
          <LanguageSwitcher />
          <h1>{title}</h1>
          <p>{description}</p>
          <QueryErrorView query={mutation} />

          <Form form={form} method={method} mutation={mutation} />

          <Recaptcha />

          {canGoBack ? (
            <button
              id="go-back-button"
              className="btn bg-transparent w-100 mt-4"
              onClick={goBack}
            >
              {s.chooseAnotherMethod}
            </button>
          ) : null}

          <LegalNotice />
        </div>
      </div>
    </div>
  );
};

const isValidEmail = (email) => {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
};

const Form = ({
  form,
  mutation,
  method,
}: {
  form: FormikProps<Partial<ClassicSigninActionReq>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
  method: AuthMethod;
  disabled?: boolean;
}) => {
  let inputType: "phonenumber" | "email" = "email";
  if (method === AuthMethod.Phone) {
    inputType = "phonenumber";
  }

  let disabled = !form?.values?.value;
  if (AuthMethod.Email === method) {
    disabled = !isValidEmail(form?.values?.value);
  }

  const s = useS(strings);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <FormText
        autoFocus
        type={inputType}
        id="value-input"
        dir="ltr"
        value={form?.values?.value}
        errorMessage={form?.errors.value}
        onChange={(value) =>
          form.setFieldValue(ClassicSigninActionReq.Fields.value, value, false)
        }
      />
      <FormButton
        className="btn btn-primary w-100 d-block mb-2"
        mutation={mutation}
        id="submit-form"
        disabled={disabled}
      >
        {s.continue}
      </FormButton>
    </form>
  );
};
