import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
// import { useGetWorkspaceWorkspaceTypeByUniqueId } from "../../sdk/modules/passports/useGetWorkspaceWorkspaceTypeByUniqueId";
// import { usePatchWorkspaceWorkspaceType } from "../../sdk/modules/passports/usePatchWorkspaceWorkspaceType";
// import { usePostWorkspaceWorkspaceType } from "../../sdk/modules/passports/usePostWorkspaceWorkspaceType";
import { WorkspaceTypeEditForm } from "./WorkspaceTypeEditForm";
import { useGetWorkspaceTypeByUniqueId } from "../../sdk/modules/workspaces/useGetWorkspaceTypeByUniqueId";
import { usePostWorkspaceType } from "../../sdk/modules/workspaces/usePostWorkspaceType";
import { usePatchWorkspaceType } from "../../sdk/modules/workspaces/usePatchWorkspaceType";
import { WorkspaceTypeEntity } from "../../sdk/modules/workspaces/WorkspaceTypeEntity";

export const WorkspaceTypeEntityManager = ({
  data,
}: DtoEntity<WorkspaceTypeEntity>) => {
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<WorkspaceTypeEntity>
  >({
    data,
  });

  const getSingleHook = useGetWorkspaceTypeByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostWorkspaceType({
    queryClient,
  });

  const patchHook = usePatchWorkspaceType({
    queryClient,
  });

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
