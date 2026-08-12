import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useWorkspaceTypeCreateAction } from "@fireback/manage/sdk/abac/WorkspaceTypeCreateAction";
import { WorkspaceTypeDto } from "@fireback/manage/sdk/abac/WorkspaceTypeDto";
import { useWorkspaceTypeGetActionQuery } from "@fireback/manage/sdk/abac/WorkspaceTypeGetAction";
import { useWorkspaceTypeUpdateAction } from "@fireback/manage/sdk/abac/WorkspaceTypeUpdateAction";
import { WorkspaceTypeEditForm } from "./WorkspaceTypeEditForm";

export const WorkspaceTypeEntityManager = ({
  data,
}: DtoEntity<WorkspaceTypeDto>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceTypeDto>
  >({
    data,
  });
  const s = useS(strings);

  const getSingleHook = useWorkspaceTypeGetActionQuery({
    params: { uniqueId },
  });

  const postHook = useWorkspaceTypeCreateAction({});

  const patchHook = useWorkspaceTypeUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(`/${locale}/workspace-types`);
      }}
      onFinishUriResolver={(response, locale) =>
        `/${locale}/workspace-type/${response.data?.uniqueId}`
      }
      Form={WorkspaceTypeEditForm}
      onEditTitle={s.editWorkspaceType}
      onCreateTitle={s.newWorkspaceType}
      data={data}
    />
  );
};
