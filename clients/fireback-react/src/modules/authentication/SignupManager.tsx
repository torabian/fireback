import { useT } from "@/hooks/useT";
import { usePostPassportSignupEmail } from "src/sdk/fireback/modules/workspaces/usePostPassportSignupEmail";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { useQueryClient } from "react-query";

import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";
import {
  EmailAccountSignupDto,
  IResponse,
  UserSessionDto,
  WorkspaceInviteEntity,
} from "src/sdk/fireback";

import { AppConfigContext } from "@/hooks/appConfigTools";
import { useGetWorkspaceInviteByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetWorkspaceInviteByUniqueId";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useGetPublicWorkspaceTypes } from "src/sdk/fireback/modules/workspaces/useGetPublicWorkspaceTypes";
import { useRememberingLoginForm } from "./AuthHooks";
import { SignupForm } from "./SignupForm";

const initialValues: Partial<EmailAccountSignupDto> = {
  email: "",
  password: "",
};

export const Signup = ({
  onSuccess,
  allowEditEmail,
  invite,
}: {
  onSuccess?: (d: IResponse<UserSessionDto>) => void;
  invite?: WorkspaceInviteEntity;
  allowEditEmail?: boolean;
}) => {
  const t = useT();
  const router = useRouter();
  const { config } = useContext(AppConfigContext);

  const { locale } = useLocale();
  const queryClient = useQueryClient();
  const { query } = useGetPublicWorkspaceTypes({
    queryClient,
    query: {},
  });

  const formik = useRef<FormikProps<Partial<EmailAccountSignupDto>> | null>();
  const passwordRef = useRef<any | null>();
  const { RememberSwitch } = useRememberingLoginForm(formik);
  const { setSession, session, isAuthenticated } =
    useContext(RemoteQueryContext);

  const {
    submit: submitPostPassportSignupEmail,
    mutation: mutationPostPassportSignupEmail,
  } = usePostPassportSignupEmail({ queryClient });

  const { query: queryJoinKey } = useGetWorkspaceInviteByUniqueId({
    query: {
      uniqueId: router.query.joinKey,
    },
  });

  useEffect(() => {
    formik.current?.setValues({
      ...formik.current.values,
      firstName: invite?.firstName,
      lastName: invite?.lastName,
      email: invite?.email,
    });
  }, []);

  const onSubmit = (
    values: Partial<EmailAccountSignupDto>,
    formikProps: FormikHelpers<Partial<EmailAccountSignupDto>>
  ) => {
    submitPostPassportSignupEmail(
      {
        ...values,
        inviteId: invite?.uniqueId,
        publicJoinKeyId: router.query.joinKey,
        workspaceTypeId: router.query.workspaceTypeId,
      } as any,
      formikProps as any
    ).then((response) => {
      if (response.data) {
        setSession(response.data);
        onSuccess && onSuccess(response);
      }
    });
  };

  // Since sometimes we signup directly to other workspaces,
  // We might want to show a custom title message, logo, etc
  let formDescription = t.signup.defaultDescription;
  if (queryJoinKey.data?.data?.uniqueId) {
    const d = queryJoinKey.data?.data;
    formDescription = t.signup.signupToWorkspace
      .replace("{workspaceName}", d.workspace?.name || "")
      .replace("{roleName}", d.role?.name || "");
  }

  const joinKeyUnAvailable = queryJoinKey.data?.data?.uniqueId;

  return (
    <Formik
      innerRef={(p) => {
        if (p) formik.current = p;
      }}
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      {(formik: FormikProps<Partial<EmailAccountSignupDto>>) => {
        return (
          <form
            className="signup-form"
            onSubmit={(e) => {
              e.preventDefault();

              // formik.submitForm();
            }}
          >
            {/* <QueryErrorView query={mutationPostPassportSignupEmail} /> */}
            <SignupForm
              RememberSwitch={RememberSwitch}
              formik={formik}
              isAuthenticated={isAuthenticated}
              formDescription={formDescription}
              loading={mutationPostPassportSignupEmail.isLoading}
              allowEditEmail={allowEditEmail}
            />
          </form>
        );
      }}
    </Formik>
  );
};
