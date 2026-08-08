import {
  CommonEntityManager,
  type DtoEntity,
} from "../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { useWorkspaceConfigDistinctGetActionQuery } from "../../sdk/abac/WorkspaceConfigDistinctGetAction";
import { useWorkspaceConfigDistinctUpdateAction } from "../../sdk/abac/WorkspaceConfigDistinctUpdateAction";
import { WorkspaceConfigDto } from "../../sdk/abac/WorkspaceConfigDto";
import { WorkspaceConfigNavigation } from "../../sdk/navigation/AbacNavigation";
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
