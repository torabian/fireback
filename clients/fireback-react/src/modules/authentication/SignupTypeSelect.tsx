import { useT } from "@/hooks/useT";

import { useContext } from "react";
import { useQueryClient } from "react-query";

import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";
import {
  IResponse,
  UserSessionDto,
  WorkspaceInviteEntity,
} from "src/sdk/fireback";

import Link from "@/components/link/Link";
import { PageSection } from "@/components/page-section/PageSection";
import { AppConfigContext } from "@/hooks/appConfigTools";
import { useGetPublicWorkspaceTypes } from "src/sdk/fireback/modules/workspaces/useGetPublicWorkspaceTypes";

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
        <div className="form-login-ui">
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
                        {item.title}
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
      </PageSection>
    </div>
  );
};
