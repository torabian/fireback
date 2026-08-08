import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { useWorkspaceConfigDistinctGetActionQuery } from "@/modules/fireback/sdk/abac/WorkspaceConfigDistinctGetAction";
import { useWorkspaceConfigDistinctUpdateAction } from "@/modules/fireback/sdk/abac/WorkspaceConfigDistinctUpdateAction";
import { WorkspaceConfigDto } from "@/modules/fireback/sdk/abac/WorkspaceConfigDto";
import { WorkspaceConfigNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { WorkspaceConfigForm } from "./WorkspaceConfigEditForm";
import { strings } from "./strings/translations";

export const WorkspaceConfigEntityManager = ({
  data,
}: DtoEntity<WorkspaceConfigDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceConfigDto>
  >({
    data,
  });
  const getSingleHook = useWorkspaceConfigDistinctGetActionQuery({});
  const patchHook = useWorkspaceConfigDistinctUpdateAction({});

  return (
    <CommonEntityManager
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      disableOnGetFailed
      forceEdit
      onCancel={() => {
        router.goBackOrDefault(
          WorkspaceConfigNavigation.single(undefined, locale),
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceConfigNavigation.single(response.data?.uniqueId, locale)
      }
      customClass="w-100"
      Form={WorkspaceConfigForm}
      onEditTitle={s.workspaceConfigs.editWorkspaceConfig}
      onCreateTitle={s.workspaceConfigs.newWorkspaceConfig}
      data={data}
    />
  );
};
