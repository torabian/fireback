import { FormikProps } from "formik";
import { UseMutationResult } from "react-query";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { FormButton } from "../../components/forms/form-button/FormButton";

import { useS } from "../../hooks/useS";
import { ClassicPassportOtpActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import ReactCodeInput from "../../thirdparty/react-verification-code-input";
import { usePresenter } from "./Otp.presenter";
import { strings } from "./strings/translations";

export const OtpScreen = ({}: {}) => {
  const { goBack, submit, mutation, form, s } = usePresenter();

  return (
    <div className="signin-form-container">
      <h1>{s.enterOtp}</h1>
      <p>{s.enterOtpDescription}</p>
      <QueryErrorView query={mutation} />
      <Form form={form} mutation={mutation} />

      <button
        id="go-back-button"
        className="btn bg-transparent w-100 mt-4"
        onClick={goBack}
      >
        {s.anotherAccount}
      </button>
    </div>
  );
};

const Form = ({
  form,
  mutation,
}: {
  form: FormikProps<Partial<ClassicPassportOtpActionReqDto>>;
  mutation: UseMutationResult<any, any, Partial<any>, any>;
}) => {
  const disabled = !form.values.otp;
  const s = useS(strings);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        form.submitForm();
      }}
    >
      <ReactCodeInput
        values={form.values.otp?.split("")}
        onChange={(value) =>
          form.setFieldValue(
            ClassicPassportOtpActionReqDto.Fields.otp,
            value,
            false
          )
        }
        className="otp-react-code-input"
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
