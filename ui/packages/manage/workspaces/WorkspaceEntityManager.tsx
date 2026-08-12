import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";

import { useWorkspaceCreateAction } from "@fireback/manage/sdk/abac/WorkspaceCreateAction";

import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { useWorkspaceGetActionQuery } from "@fireback/manage/sdk/abac/WorkspaceGetAction";
import { useWorkspaceUpdateAction } from "@fireback/manage/sdk/abac/WorkspaceUpdateAction";
import { WorkspaceEditForm } from "./WorkspaceEditForm";
import { WorkspaceDto } from "@fireback/manage/sdk/abac/WorkspaceDto";
import { WorkspaceNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

export const WorkspaceEntityManager = ({ data }: DtoEntity<WorkspaceDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceDto>
  >({
    data,
  });

  const getSingleHook = useWorkspaceGetActionQuery({
    params: { uniqueId },
  });

  const postHook = useWorkspaceCreateAction({});

  const patchHook = useWorkspaceUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(WorkspaceNavigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceNavigation.single(response.data?.uniqueId, locale)
      }
      Form={WorkspaceEditForm}
      onEditTitle={s.editWorkspae}
      onCreateTitle={s.createNewWorkspace}
      data={data}
    />
  );
};
