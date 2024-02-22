import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";

import { WorkspaceEntity } from "src/sdk/fireback";
import { usePostWorkspace } from "src/sdk/fireback/modules/workspaces/usePostWorkspace";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetWorkspaceByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetWorkspaceByUniqueId";
import { usePatchWorkspace } from "src/sdk/fireback/modules/workspaces/usePatchWorkspace";
import { WorkspaceNavigationTools } from "src/sdk/fireback/modules/workspaces/workspace-navigation-tools";
import { WorkspaceEditForm } from "./WorkspaceEditForm";

export const WorkspaceEntityManager = ({
  data,
}: DtoEntity<WorkspaceEntity>) => {
  const t = useT();
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceEntity>
  >({
    data,
  });

  const getSingleHook = useGetWorkspaceByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostWorkspace({
    queryClient,
  });

  const patchHook = usePatchWorkspace({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          WorkspaceNavigationTools.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceNavigationTools.single(response.data?.uniqueId, locale)
      }
      Form={WorkspaceEditForm}
      onEditTitle={t.wokspaces.editWorkspae}
      onCreateTitle={t.wokspaces.createNewWorkspace}
      data={data}
    />
  );
};
