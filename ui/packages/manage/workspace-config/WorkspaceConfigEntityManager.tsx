import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useWorkspaceConfigDistinctGetActionQuery } from "@fireback/manage/sdk/abac/WorkspaceConfigDistinctGetAction";
import { useWorkspaceConfigDistinctUpdateAction } from "@fireback/manage/sdk/abac/WorkspaceConfigDistinctUpdateAction";
import { WorkspaceConfigDto } from "@fireback/manage/sdk/abac/WorkspaceConfigDto";
import { WorkspaceConfigNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
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
