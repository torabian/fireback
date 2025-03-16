import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
import { useGetWorkspaceConfigDistinct } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceConfigDistinct";
import { usePatchWorkspaceConfigDistinct } from "@/modules/fireback/sdk/modules/workspaces/usePatchWorkspaceConfigDistinct";
import { WorkspaceConfigForm } from "./WorkspaceConfigEditForm";
import { strings } from "./strings/translations";
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
      disableOnGetFailed
      forceEdit
      onCancel={() => {
        router.goBackOrDefault(
          WorkspaceConfigEntity.Navigation.single(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceConfigEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      customClass="w-100"
      Form={WorkspaceConfigForm}
      onEditTitle={s.workspaceConfigs.editWorkspaceConfig}
      onCreateTitle={s.workspaceConfigs.newWorkspaceConfig}
      data={data}
    />
  );
};
