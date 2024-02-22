import { useT } from "@/hooks/useT";

import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useRef } from "react";
import { useQueryClient } from "react-query";

import { PageSection } from "@/components/page-section/PageSection";
import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";
import {
  EmailAccountSigninDto,
  IResponse,
  UserSessionDto,
  WorkspaceInviteEntity,
} from "src/sdk/fireback";

import { useGetWorkspaceInviteByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetWorkspaceInviteByUniqueId";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useRememberingLoginForm } from "./AuthHooks";
import { Signup } from "./SignupManager";
import { usePostPassportSignupEmail } from "@/sdk/fireback/modules/workspaces/usePostPassportSignupEmail";

const initialValues: Partial<EmailAccountSigninDto> = {
  email: "",
  password: "",
};

enum InviteState {
  USER_EXISTS_BUT_NOT_LOGGED_IN = "USER_EXISTS_BUT_NOT_LOGGED_IN",
  USER_EXISTS_BUT_LOGGED_IN_DIFFERENTACCOUNT = "USER_EXISTS_BUT_LOGGED_IN_DIFFERENTACCOUNT",
  USER_EXISTS_AND_LOGGED_IN = "USER_EXISTS_AND_LOGGED_IN",
  USER_DOES_NOT_EXISTS = "USER_DOES_NOT_EXISTS",
  USER_DOES_NOT_EXISTS_BUT_LOGGED_IN = "USER_DOES_NOT_EXISTS_BUT_LOGGED_IN",
  UNKNOWN = "UNKNOWN",
}

function determineInviteState(
  invite?: WorkspaceInviteEntity,
  session?: any
): InviteState {
  if (!(invite as any)?.inviteeUserId && !session) {
    return InviteState.USER_DOES_NOT_EXISTS;
  }
  if ((invite as any)?.inviteeUserId && !session) {
    return InviteState.USER_EXISTS_BUT_NOT_LOGGED_IN;
  }

  if ((invite as any)?.inviteeUserId === session?.user?.uniqueId) {
    return InviteState.USER_EXISTS_AND_LOGGED_IN;
  }

  if (
    (invite as any)?.inviteeUserId &&
    (invite as any)?.inviteeUserId !== session?.user?.uniqueId
  ) {
    return InviteState.USER_EXISTS_BUT_LOGGED_IN_DIFFERENTACCOUNT;
  }

  if (!(invite as any)?.inviteeUserId && session) {
    return InviteState.USER_DOES_NOT_EXISTS_BUT_LOGGED_IN;
  }

  return InviteState.UNKNOWN;
}

// queryInvitation.data?.data.inviteeUserId
export const JoinToWorkspace = ({
  onSuccess,
}: {
  onSuccess: (d: IResponse<UserSessionDto>) => void;
}) => {
  const t = useT();
  const router = useRouter();
  const { locale } = useLocale();
  const queryClient = useQueryClient();
  const formik = useRef<FormikProps<Partial<EmailAccountSigninDto>> | null>();
  const passwordRef = useRef<any | null>();
  const { RememberSwitch } = useRememberingLoginForm(formik);
  const { setSession, session, isAuthenticated } =
    useContext(RemoteQueryContext);

  const uniqueId = router.query.uniqueId as string;

  const { query: queryInvitation } = useGetWorkspaceInviteByUniqueId({
    query: { uniqueId },
    queryClient,
  });

  const state = determineInviteState(queryInvitation.data?.data, session);

  const {
    submit: submitPostPassportSignupEmail,
    mutation: mutationPostPassportSignupEmail,
  } = usePostPassportSignupEmail({ queryClient });

  const onSubmit = (
    values: Partial<EmailAccountSigninDto>,
    formikProps: FormikHelpers<Partial<EmailAccountSigninDto>>
  ) => {
    submitPostPassportSignupEmail(values, formikProps as any).then(
      (response) => {
        if (response.data) {
          setSession(response.data);
          onSuccess(response);

          // setSession(response.data);
          // saveCredentials(values);
          // navigation.navigate('app', {screen: Screens.Home});
        }
      }
    );
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
        return (
          // <form
          //   onSubmit={(e) => {
          //     e.preventDefault();

          //     // formik.submitForm();
          //   }}
          // >
          <>
            <div className="join-to-workspace">
              <JoinWorkspaceForm
                invite={queryInvitation.data?.data}
                state={state}
                formik={formik}
              />
              {state === InviteState.USER_DOES_NOT_EXISTS &&
                queryInvitation.data?.data && (
                  <div className="signup-in-workspace">
                    <Signup
                      invite={queryInvitation.data?.data}
                      allowEditEmail={false}
                      onSuccess={() => {}}
                    />
                  </div>
                )}
            </div>
          </>
          // </form>
        );
      }}
    </Formik>
  );
};

export const JoinWorkspaceForm = ({
  formik,
  invite,
  state,
}: {
  formik: FormikProps<any>;
  state: InviteState;
  invite?: WorkspaceInviteEntity;
}) => {
  return (
    <div className="auth-wrapper">
      <PageSection title="">
        {/* {mutationPostPassportSignupEmail.isLoading && <AuthLoader />}
        {isAuthenticated ? (
          <UserProfileCard />
        ) : ( */}
        {/* <strong>{state}</strong> */}
        <div className="form-login-ui">
          <h1>Join workspace</h1>
          <p>Someone has invited you to join their workspace.</p>

          {state === InviteState.USER_DOES_NOT_EXISTS_BUT_LOGGED_IN && (
            <div>
              This invitation is for {invite?.email}, you need to logout first,
              and then create this account.
            </div>
          )}
          {state === InviteState.USER_EXISTS_AND_LOGGED_IN && (
            <>
              <div>
                You can accept the invite and join {invite?.workspace?.name} as{" "}
                {invite?.role?.name}
              </div>
              <div className="mt-4">
                <button className="btn btn-sm btn-primary">Accept</button>{" "}
                <button className="btn btn-sm btn-danger ">Reject</button>
              </div>
            </>
          )}
          {state === InviteState.USER_EXISTS_BUT_LOGGED_IN_DIFFERENTACCOUNT && (
            <div>
              This invitation is for another account. In order to accept it,
              logout first and login again
            </div>
          )}
          {state === InviteState.USER_EXISTS_BUT_NOT_LOGGED_IN && (
            <div>
              You need to login first in order to accept this invitation
            </div>
          )}
        </div>

        {/* )} */}
      </PageSection>
    </div>
  );
};
