import { useCommonEntityManager } from "../../../hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../components/entity-manager/CommonEntityManager";
// import { useGetWorkspaceWorkspaceTypeByUniqueId } from "../../sdk/modules/passports/useGetWorkspaceWorkspaceTypeByUniqueId";
// import { usePatchWorkspaceWorkspaceType } from "../../sdk/modules/passports/usePatchWorkspaceWorkspaceType";
// import { usePostWorkspaceWorkspaceType } from "../../sdk/modules/passports/usePostWorkspaceWorkspaceType";
import { WorkspaceTypeEditForm } from "./WorkspaceTypeEditForm";
import { useWorkspaceTypeGetActionQuery } from "../../../sdk/abac/WorkspaceTypeGetAction";
import { useWorkspaceTypeCreateAction } from "../../../sdk/abac/WorkspaceTypeCreateAction";
import { useWorkspaceTypeUpdateAction } from "../../../sdk/abac/WorkspaceTypeUpdateAction";
import { WorkspaceTypeDto } from "../../../sdk/abac/WorkspaceTypeDto";

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
