import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";

import { useWorkspaceCreateAction } from "../../../sdk/abac/WorkspaceCreateAction";

import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useWorkspaceGetActionQuery } from "../../../sdk/abac/WorkspaceGetAction";
import { useWorkspaceUpdateAction } from "../../../sdk/abac/WorkspaceUpdateAction";
import { WorkspaceEditForm } from "./WorkspaceEditForm";
import { WorkspaceDto } from "../../../sdk/abac/WorkspaceDto";
import { WorkspaceNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceEntityManager = ({
  data,
}: DtoEntity<WorkspaceDto>) => {
  const t = useT();
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
        router.goBackOrDefault(
          WorkspaceNavigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceNavigation.single(response.data?.uniqueId, locale)
      }
      Form={WorkspaceEditForm}
      onEditTitle={t.wokspaces.editWorkspae}
      onCreateTitle={t.wokspaces.createNewWorkspace}
      data={data}
    />
  );
};
