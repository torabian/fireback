import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useWorkspaceInviteGetActionQuery } from "@/modules/fireback/sdk/abac/WorkspaceInviteGetAction";
import { useWorkspaceInviteCreateAction } from "@/modules/fireback/sdk/abac/WorkspaceInviteCreateAction";
import { useWorkspaceInviteUpdateAction } from "@/modules/fireback/sdk/abac/WorkspaceInviteUpdateAction";
import { WorkspaceInviteDto } from "@/modules/fireback/sdk/abac/WorkspaceInviteDto";
import { WorkspaceInviteForm } from "./WorkspaceInviteForm";

export const WorkspaceInviteEntityManager = ({
  data,
}: DtoEntity<WorkspaceInviteDto>) => {
  const t = useT();
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceInviteDto>
  >({
    data,
  });

  const getSingleHook = useGetWorkspaceInviteByUniqueId({
    query: { uniqueId },
    queryClient,
  });

  const postHook = useWorkspaceInviteCreateAction({});

  const patchHook = useWorkspaceInviteUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(`/${locale}/workspace-invites`);
      }}
      onFinishUriResolver={(_, locale) => `/${locale}/workspace-invites`}
      Form={WorkspaceInviteForm}
      onEditTitle={t.wokspaces.invite.editInvitation}
      onCreateTitle={t.wokspaces.invite.createInvitation}
      data={data}
    />
  );
};
