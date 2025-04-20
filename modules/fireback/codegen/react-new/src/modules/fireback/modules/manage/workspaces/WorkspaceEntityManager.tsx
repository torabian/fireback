import { useCommonEntityManager } from "../../../hooks/useCommonEntityManager";

import { usePostWorkspace } from "../../../sdk/modules/abac/usePostWorkspace";

import {
  CommonEntityManager,
  DtoEntity,
} from "../../../components/entity-manager/CommonEntityManager";
import { useT } from "../../../hooks/useT";
import { useGetWorkspaceByUniqueId } from "../../../sdk/modules/abac/useGetWorkspaceByUniqueId";
import { usePatchWorkspace } from "../../../sdk/modules/abac/usePatchWorkspace";
import { WorkspaceEditForm } from "./WorkspaceEditForm";
import { WorkspaceEntity } from "../../../sdk/modules/abac/WorkspaceEntity";

export const WorkspaceEntityManager = ({
  data,
}: DtoEntity<WorkspaceEntity>) => {
  const t = useT();
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceEntity>
  >({
    data,
  });

  const getSingleHook = useGetWorkspaceByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostWorkspace({
    queryClient,
  });

  const patchHook = usePatchWorkspace({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          WorkspaceEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        WorkspaceEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={WorkspaceEditForm}
      onEditTitle={t.wokspaces.editWorkspae}
      onCreateTitle={t.wokspaces.createNewWorkspace}
      data={data}
    />
  );
};
