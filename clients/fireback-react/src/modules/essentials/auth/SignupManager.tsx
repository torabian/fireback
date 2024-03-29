import { useT } from "@/hooks/useT";
import { usePostPassportSignupEmail } from "src/sdk/fireback/modules/workspaces/usePostPassportSignupEmail";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect, useRef } from "react";
import { useQueryClient } from "react-query";

import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";

import { AppConfigContext } from "@/hooks/appConfigTools";
import { useGetWorkspaceInviteByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetWorkspaceInviteByUniqueId";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useGetPublicWorkspaceTypes } from "src/sdk/fireback/modules/workspaces/useGetPublicWorkspaceTypes";
import { useRememberingLoginForm } from "./AuthHooks";
import { SignupForm } from "./SignupForm";
import { usePostPassportsSignupClassic } from "@/sdk/fireback/modules/workspaces/usePostPassportsSignupClassic";
import { ClassicSignupActionReqDto } from "@/sdk/fireback/modules/workspaces/WorkspacesActionsDto";
import { IResponse } from "@/definitions/JSONStyle";
import { UserSessionDto } from "@/sdk/fireback/modules/workspaces/UserSessionDto";
import { WorkspaceInviteEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceInviteEntity";

const initialValues: Partial<ClassicSignupActionReqDto> = {
  value: "",
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

  const formik = useRef<FormikProps<
    Partial<ClassicSignupActionReqDto>
  > | null>();
  const passwordRef = useRef<any | null>();
  const { RememberSwitch } = useRememberingLoginForm(formik);
  const { setSession, session, isAuthenticated } =
    useContext(RemoteQueryContext);

  const {
    submit: submitPostPassportSignupEmail,
    mutation: mutationPostPassportSignupEmail,
  } = usePostPassportsSignupClassic({ queryClient });

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
      value: invite?.value,
    });
  }, []);

  const onSubmit = (
    values: Partial<ClassicSignupActionReqDto>,
    formikProps: FormikHelpers<Partial<ClassicSignupActionReqDto>>
  ) => {
    submitPostPassportSignupEmail(
      {
        ...values,
        inviteId: invite?.uniqueId,
        publicJoinKeyId: router.query.joinKey,
        workspaceTypeId: router.query.workspaceTypeId,
        type: "email",
      } as any,
      formikProps as any
    ).then((response) => {
      if (response.data) {
        setSession(response.data);
        onSuccess && onSuccess(response as any);
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
      {(formik: FormikProps<Partial<ClassicSignupActionReqDto>>) => {
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
