import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useGetWorkspaceInviteByUniqueId } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceInviteByUniqueId";
import { usePatchWorkspaceInvite } from "@/modules/fireback/sdk/modules/workspaces/usePatchWorkspaceInvite";
import { usePostWorkspaceInvite } from "@/modules/fireback/sdk/modules/workspaces/usePostWorkspaceInvite";
import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceInviteEntity";
import { WorkspaceInviteForm } from "./WorkspaceInviteForm";

export const WorkspaceInviteEntityManager = ({
  data,
}: DtoEntity<WorkspaceInviteEntity>) => {
  const t = useT();
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceInviteEntity>
  >({
    data,
  });

  const getSingleHook = useGetWorkspaceInviteByUniqueId({
    query: { uniqueId },
    queryClient,
  });

  const postHook = usePostWorkspaceInvite({
    queryClient,
  });

  const patchHook = usePatchWorkspaceInvite({
    queryClient,
  });

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
