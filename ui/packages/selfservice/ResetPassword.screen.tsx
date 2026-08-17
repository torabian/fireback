import { type FormikProps } from "formik";
import { type UseMutationResult } from "@tanstack/react-query";
import { QueryErrorView } from "@fireback/ui-core/components/error-view/QueryError";
import { FormButton } from "@fireback/ui-core/components/forms/form-button/FormButton";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { useS } from "@fireback/ui-core/hooks/useS";
import { CompletePassportPasswordResetActionReq } from "@fireback/selfservice/sdk/abac/CompletePassportPasswordResetAction";
import { usePresenter } from "./ResetPassword.presenter";
import { strings } from "./strings/translations";

export const ResetPasswordScreen = ({}: {}) => {
  const { mutation, form, s } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>{s.resetPassword.title}</h1>
      <p>{s.resetPassword.description}</p>
      <QueryErrorView query={mutation} />
      <Form form={form} mutation={mutation} />
    </div>
  );
};

const Form = ({
  form,
  mutation,
}: {
  form: FormikProps<Partial<CompletePassportPasswordResetActionReq>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
}) => {
  const s = useS(strings);
  const { value, otp, password } = form.values;
  const disabled = !value || !otp || (password?.length || 0) < 6;

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <FormText
        type="text"
        value={value}
        label={s.resetPassword.valueLabel}
        id="value-input"
        errorMessage={form.errors.value}
        onChange={(v) =>
          form.setFieldValue(
            CompletePassportPasswordResetActionReq.Fields.value,
            v,
            false,
          )
        }
      />

      <FormText
        type="text"
        value={otp}
        label={s.resetPassword.otpLabel}
        id="otp-input"
        errorMessage={form.errors.otp}
        onChange={(v) =>
          form.setFieldValue(
            CompletePassportPasswordResetActionReq.Fields.otp,
            v,
            false,
          )
        }
      />

      <FormText
        type="password"
        value={password}
        label={s.resetPassword.newPasswordLabel}
        id="password-input"
        errorMessage={form.errors.password}
        onChange={(v) =>
          form.setFieldValue(
            CompletePassportPasswordResetActionReq.Fields.password,
            v,
            false,
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
