import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { WithForm } from "../../components/forms/WithForm";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { AuthMethod } from "./auth.common";
import { usePresenter } from "./ClassicPassport.presenter";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const ClassicPassportScreen = ({ method }: { method: AuthMethod }) => {
  const {
    description,
    title,
    goBack,
    submit,
    mutation,
    setFormRef,
    canGoBack,
    LegalNotice,
    Recaptcha,
    s,
  } = usePresenter({
    method,
  });

  return (
    <div className="signin-form-container">
      <h1>{title}</h1>
      <p>{description}</p>
      <QueryErrorView query={mutation} />
      <WithForm
        setFormRef={setFormRef}
        onSubmit={submit}
        Form={(props) => (
          <Form {...props} method={method} mutation={mutation} />
        )}
      />

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
  form: FormikProps<Partial<ClassicSigninActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
  method: AuthMethod;
  disabled: boolean;
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
    <div>
      <FormText
        autoFocus
        type={inputType}
        id="value-input"
        dir="ltr"
        value={form?.values?.value}
        errorMessage={form?.errors.value}
        onChange={(value) =>
          form.setFieldValue(
            ClassicSigninActionReqDto.Fields.value,
            value,
            false
          )
        }
      />
      <FormButton
        className="btn btn-primary w-100 d-block mb-2"
        onClick={() => form.submitForm()}
        mutation={mutation}
        id="submit-form"
        disabled={disabled}
      >
        {s.continue}
      </FormButton>
    </div>
  );
};
