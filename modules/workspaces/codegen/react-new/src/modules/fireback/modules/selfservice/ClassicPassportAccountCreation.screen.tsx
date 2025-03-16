import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { WithForm } from "../../components/forms/WithForm";
import { useS } from "../../hooks/useS";
import { ClassicSignupActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { usePresenter } from "./ClassicPassportAccountCreation.presenter";
import { strings } from "./strings/translations";
import { AuthLoader } from "../../components/auth-loader/AuthLoader";

export const ClassicPassportAccountCreation = ({}: {}) => {
  const {
    goBack,
    submit,
    mutation,
    setFormRef,
    state,
    workspaceTypes,
    totpUrl,
    isLoading,
    s,
  } = usePresenter();

  if (isLoading) {
    return (
      <div className="signin-form-container">
        <AuthLoader />
      </div>
    );
  }

  if (workspaceTypes.length === 0) {
    return (
      <div className="signin-form-container">
        <h1>{s.registerationNotPossible}</h1>
        <p>{s.registerationNotPossibleLine1}</p>
        <p>{s.registerationNotPossibleLine2}</p>
      </div>
    );
  }

  return (
    <div
      className="signin-form-container fadein"
      style={{ animation: "fadein 1s" }}
    >
      <h1>{s.completeYourAccount}</h1>
      <p>{s.completeYourAccountDescription}</p>

      <QueryErrorView query={mutation} />
      <WithForm
        setFormRef={setFormRef}
        onSubmit={submit}
        Form={(props) => (
          <Form {...props} mutation={mutation} totpUrl={totpUrl} />
        )}
      />

      <button
        id="go-step-back"
        onClick={goBack}
        className="bg-transparent border-0"
      >
        {s.cancelStep}
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
}: {
  form: FormikProps<Partial<ClassicSignupActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
}) => {
  const s = useS(strings);
  const disabled =
    !form.values.firstName ||
    !form.values.lastName ||
    !form.values.password ||
    form.values.password.length < 6;

  return (
    <div>
      <FormText
        value={form.values.firstName}
        label={s.firstName}
        id="first-name-input"
        autoFocus
        errorMessage={form.errors.firstName}
        onChange={(value) =>
          form.setFieldValue(
            ClassicSignupActionReqDto.Fields.firstName,
            value,
            false
          )
        }
      />
      <FormText
        value={form.values.lastName}
        label={s.lastName}
        id="last-name-input"
        errorMessage={form.errors.lastName}
        onChange={(value) =>
          form.setFieldValue(
            ClassicSignupActionReqDto.Fields.lastName,
            value,
            false
          )
        }
      />

      <FormText
        type="password"
        value={form.values.password}
        label={s.password}
        id="password-input"
        errorMessage={form.errors.password}
        onChange={(value) =>
          form.setFieldValue(
            ClassicSignupActionReqDto.Fields.password,
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
