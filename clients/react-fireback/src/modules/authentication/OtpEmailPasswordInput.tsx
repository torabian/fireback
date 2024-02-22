import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormText } from "@/components/forms/form-text/FormText";
import { useT } from "@/hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { useQueryClient } from "react-query";

import {
  EmailAccountSigninDto,
  IResponse,
  UserSessionDto,
  WorkspaceInviteEntity,
} from "src/sdk/fireback";

import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { usePostPassportSigninEmail } from "src/sdk/fireback/modules/workspaces/usePostPassportSigninEmail";

const initialValues: Partial<EmailAccountSigninDto> = {};

export const OtpEmailPasswordInput = ({
  onSuccess,
  invite,
}: {
  onSuccess?: (d: IResponse<UserSessionDto>) => void;
  invite?: WorkspaceInviteEntity;
}) => {
  const t = useT();
  const { setSession, session, options, isAuthenticated } =
    useContext(RemoteQueryContext);

  const queryClient = useQueryClient();
  const formik = useRef<FormikProps<Partial<EmailAccountSigninDto>> | null>();

  const {
    submit: submitPostPassportSigninEmail,
    mutation: mutationPostPassportSigninEmail,
  } = usePostPassportSigninEmail({ queryClient });

  useEffect(() => {
    formik.current?.setValues({
      ...formik.current.values,
      // firstName: invite?.firstName,
      // lastName: invite?.lastName,
      // email: invite?.email,
    });
  }, []);

  const onSubmit = (
    values: Partial<EmailAccountSigninDto>,
    formikProps: FormikHelpers<Partial<EmailAccountSigninDto>>
  ) => {
    // onRemoteChange && onRemoteChange("remote");
    // setTimeout(() => {
    submitPostPassportSigninEmail(values, formikProps as any).then(
      (response) => {
        if (response.data) {
          // Think about this.
          // if (shouldRemember) {
          //   localStorage.setItem(
          //     "credentials",
          //     JSON.stringify({
          //       email: values.email,
          //       password: values.password,
          //     })
          //   );
          // } else {
          //   formik.current?.setValues({ email: "", password: "" });
          // }
          // setOptions({ headers: { Authorization: response.data.token } });
          setSession(response.data);
          onSuccess && onSuccess(response);
        }
      }
    );
    // }, 300);
  };

  return (
    <Formik
      innerRef={(p) => {
        if (p) formik.current = p;
      }}
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      {(formik: FormikProps<Partial<EmailAccountSigninDto>>) => {
        const { values, setFieldValue, errors } = formik;
        return (
          <form
            onSubmit={(e) => {
              e.preventDefault();
            }}
          >
            <div className="step-header">
              <span>1</span> Password
            </div>
            <FormText
              type="password"
              value={(formik as any).values.password}
              label="Password"
              errorMessage={(formik as any).errors.password}
              onChange={(value) => setFieldValue("password", value, false)}
            />
            <FormButton
              isSubmitting={mutationPostPassportSigninEmail.isLoading}
              onClick={() => formik.submitForm()}
              label={t.loginButton}
            />
          </form>
        );
      }}
    </Formik>
  );
};
