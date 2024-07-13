import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useGetWorkspaceInviteByUniqueId } from "../../sdk/modules/workspaces/useGetWorkspaceInviteByUniqueId";
import { usePatchWorkspaceInvite } from "../../sdk/modules/workspaces/usePatchWorkspaceInvite";
import { usePostWorkspaceInvite } from "../../sdk/modules/workspaces/usePostWorkspaceInvite";
import { WorkspaceInviteForm } from "./WorkspaceInviteForm";
import { WorkspaceInviteEntity } from "../../sdk/modules/workspaces/WorkspaceInviteEntity";

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
