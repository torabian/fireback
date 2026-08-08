import { useCommonEntityManager } from "../../../hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../components/entity-manager/CommonEntityManager";
import { useWorkspaceTypeCreateAction } from "../../../sdk/abac/WorkspaceTypeCreateAction";
import { WorkspaceTypeDto } from "../../../sdk/abac/WorkspaceTypeDto";
import { useWorkspaceTypeGetActionQuery } from "../../../sdk/abac/WorkspaceTypeGetAction";
import { useWorkspaceTypeUpdateAction } from "../../../sdk/abac/WorkspaceTypeUpdateAction";
import { WorkspaceTypeEditForm } from "./WorkspaceTypeEditForm";

export const WorkspaceTypeEntityManager = ({
  data,
}: DtoEntity<WorkspaceTypeDto>) => {
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<WorkspaceTypeDto>
  >({
    data,
  });

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
      onEditTitle={t.fb.editWorkspaceType}
      onCreateTitle={t.fb.newWorkspaceType}
      data={data}
    />
  );
};
