import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { WorkspaceTypeEntity } from "src/sdk/fireback";
// import { useGetWorkspaceWorkspaceTypeByUniqueId } from "src/sdk/fireback/modules/passports/useGetWorkspaceWorkspaceTypeByUniqueId";
// import { usePatchWorkspaceWorkspaceType } from "src/sdk/fireback/modules/passports/usePatchWorkspaceWorkspaceType";
// import { usePostWorkspaceWorkspaceType } from "src/sdk/fireback/modules/passports/usePostWorkspaceWorkspaceType";
import { WorkspaceTypeEditForm } from "./WorkspaceTypeEditForm";
import { useGetWorkspaceTypeByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetWorkspaceTypeByUniqueId";
import { usePostWorkspaceType } from "src/sdk/fireback/modules/workspaces/usePostWorkspaceType";
import { usePatchWorkspaceType } from "src/sdk/fireback/modules/workspaces/usePatchWorkspaceType";

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
