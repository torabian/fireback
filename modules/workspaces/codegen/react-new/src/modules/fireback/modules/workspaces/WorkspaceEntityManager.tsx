import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";

import { usePostWorkspace } from "../../sdk/modules/workspaces/usePostWorkspace";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useGetWorkspaceByUniqueId } from "../../sdk/modules/workspaces/useGetWorkspaceByUniqueId";
import { usePatchWorkspace } from "../../sdk/modules/workspaces/usePatchWorkspace";
import { WorkspaceEditForm } from "./WorkspaceEditForm";
import { WorkspaceEntity } from "../../sdk/modules/workspaces/WorkspaceEntity";

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
          WorkspaceEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={WorkspaceEditForm}
      onEditTitle={t.wokspaces.editWorkspae}
      onCreateTitle={t.wokspaces.createNewWorkspace}
      data={data}
    />
  );
};
