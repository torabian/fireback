import { useT } from "@/fireback/hooks/useT";

import { useQueryClient } from "react-query";

import Link from "@/fireback/components/link/Link";
import { PageSection } from "@/fireback/components/page-section/PageSection";
import { useGetPublicWorkspaceTypes } from "src/sdk/fireback/modules/workspaces/useGetPublicWorkspaceTypes";
import { IResponse } from "@/sdk/fireback/core/http-tools";
import { UserSessionDto } from "@/sdk/fireback/modules/workspaces/UserSessionDto";
import { WorkspaceInviteEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceInviteEntity";

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
      <PageSection title="">
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
      </PageSection>
    </div>
  );
};
