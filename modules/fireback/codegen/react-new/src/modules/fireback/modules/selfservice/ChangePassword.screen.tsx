import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";
import { FormText } from "../../components/forms/form-text/FormText";
import { useS } from "../../hooks/useS";
import { ChangePasswordDto, usePresenter } from "./ChangePassword.presenter";
import { strings } from "./strings/translations";

export const ChangePasswordScreen = ({}: {}) => {
  const { mutation, form, s } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>{s.changePassword.title}</h1>
      <p>{s.changePassword.description}</p>
      <QueryErrorView query={mutation} />
      <Form form={form} mutation={mutation} />
    </div>
  );
};

const Form = ({
  form,
  mutation,
}: {
  form: FormikProps<Partial<ChangePasswordDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
}) => {
  const s = useS(strings);
  const { password2, password } = form.values;
  const disabled = password !== password2 || (password?.length || 0) < 6;

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <FormText
        type="password"
        value={form.values.password}
        label={s.changePassword.pass1Label}
        id="password-input"
        errorMessage={form.errors.password}
        onChange={(value) =>
          form.setFieldValue(ChangePasswordDto.Fields.password, value, false)
        }
      />

      <FormText
        type="password"
        value={form.values.password2}
        label={s.changePassword.pass2Label}
        id="password-input-2"
        errorMessage={form.errors.password}
        onChange={(value) =>
          form.setFieldValue(ChangePasswordDto.Fields.password2, value, false)
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
