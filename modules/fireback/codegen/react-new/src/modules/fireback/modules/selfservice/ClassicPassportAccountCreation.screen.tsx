import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { useS } from "../../hooks/useS";
import { usePresenter } from "./ClassicPassportAccountCreation.presenter";
import { strings } from "./strings/translations";
import { AuthLoader } from "../../components/auth-loader/AuthLoader";
import { ClassicSignupActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";

export const ClassicPassportAccountCreation = ({}: {}) => {
  const {
    goBack,
    submit,
    mutation,
    form,
    state,
    workspaceTypes,
    workspaceTypeId,
    totpUrl,
    isLoading,
    setSelectedWorkspaceType,
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

  // If there are more than single workspace, we need to ask the user to choose one.
  if (workspaceTypes.length >= 2 && !workspaceTypeId) {
    return (
      <div
        className="signin-form-container fadein"
        style={{ animation: "fadein 1s" }}
      >
        <h1>{s.completeYourAccount}</h1>
        <p>{s.completeYourAccountDescription}</p>
        <div className=" ">
          {workspaceTypes.map((workspaceType) => (
            <div className="mt-3">
              <h2>{workspaceType.title}</h2>
              <p>{workspaceType.description}</p>
              <button
                key={workspaceType.uniqueId}
                className="btn btn-outline-primary w-100"
                onClick={() => {
                  setSelectedWorkspaceType(workspaceType.uniqueId);
                }}
              >
                Select
              </button>
            </div>
          ))}
        </div>
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
      <Form form={form} mutation={mutation} />
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
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
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
        mutation={mutation}
        id="submit-form"
        disabled={disabled}
      >
        {s.continue}
      </FormButton>
    </form>
  );
};
