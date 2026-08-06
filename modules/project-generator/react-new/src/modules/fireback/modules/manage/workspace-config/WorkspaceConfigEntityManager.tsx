import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { useWorkspaceConfigDistinctGetActionQuery } from "@/modules/fireback/sdk/abac/WorkspaceConfigDistinctGetAction";
import { useWorkspaceConfigDistinctUpdateAction } from "@/modules/fireback/sdk/abac/WorkspaceConfigDistinctUpdateAction";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceConfigEntity";
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
          WorkspaceConfigEntity.Navigation.single(undefined, locale),
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
