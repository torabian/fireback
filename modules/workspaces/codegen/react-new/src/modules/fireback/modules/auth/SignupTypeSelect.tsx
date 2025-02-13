import { useT } from "../../hooks/useT";

import { useQueryClient } from "react-query";

import Link from "../../components/link/Link";
import { useGetPublicWorkspaceTypes } from "../../sdk/modules/workspaces/useGetPublicWorkspaceTypes";
import { IResponse } from "../../sdk/core/http-tools";
import { UserSessionDto } from "../../sdk/modules/workspaces/UserSessionDto";
import { WorkspaceInviteEntity } from "../../sdk/modules/workspaces/WorkspaceInviteEntity";

export const SignupTypeSelect = ({
  onSuccess,
  allowEditEmail,
  invite,
}: {
  onSuccess?: (d: IResponse<UserSessionDto>) => void;
  invite?: WorkspaceInviteEntity;
  allowEditEmail?: boolean;
}) => {
  const t = useT();

  const queryClient = useQueryClient();
  const { query } = useGetPublicWorkspaceTypes({
    queryClient,
    unauthorized: true,
  });

  const items = query.data?.data?.items || [];

  return (
    <div className="signup-wrapper">
      <div className="form-login-ui ">
        <div className="login-form-section">
          <h1>{t.abac.signupType}</h1>
          {items.length === 0 && <p>{t.noSignupType}</p>}

          {items.length > 0 && (
            <div>
              <p>{t.abac.signupTypeHint}</p>
              {items.map((item) => {
                return (
                  <div className="signup-workspace-type" key={item.uniqueId}>
                    <h2>
                      <Link href={`/signup/${item.uniqueId}`}>
                        {item.title || item.uniqueId}
                      </Link>
                    </h2>
                    <p>{(item as any).description}</p>
                  </div>
                );
              })}
            </div>
          )}
          <Link href={`/signin`}>{t.signinInstead}</Link>
        </div>
      </div>
    </div>
  );
};
