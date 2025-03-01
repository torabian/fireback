import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { WorkspaceConfigForm } from "./WorkspaceConfigEditForm";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
import { useGetWorkspaceConfigByUniqueId } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceConfigByUniqueId";
import { usePostWorkspaceConfig } from "@/modules/fireback/sdk/modules/workspaces/usePostWorkspaceConfig";
import { usePatchWorkspaceConfig } from "@/modules/fireback/sdk/modules/workspaces/usePatchWorkspaceConfig";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { useGetWorkspaceConfigDistinct } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceConfigDistinct";
import { usePatchWorkspaceConfigDistinct } from "@/modules/fireback/sdk/modules/workspaces/usePatchWorkspaceConfigDistinct";
export const WorkspaceConfigEntityManager = ({
  data,
}: DtoEntity<WorkspaceConfigEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceConfigEntity>
  >({
    data,
  });
  const getSingleHook = useGetWorkspaceConfigDistinct({
    query: { uniqueId },
  });

  const patchHook = usePatchWorkspaceConfigDistinct({
    queryClient,
  });
  return (
    <CommonEntityManager
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      forceEdit
      onCancel={() => {
        router.goBackOrDefault(
          WorkspaceConfigEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceConfigEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={WorkspaceConfigForm}
      onEditTitle={s.workspaceConfigs.editWorkspaceConfig}
      onCreateTitle={s.workspaceConfigs.newWorkspaceConfig}
      data={data}
    />
  );
};
