import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useWorkspaceTypeCreateAction } from "../../../sdk/abac/WorkspaceTypeCreateAction";
import { WorkspaceTypeDto } from "../../../sdk/abac/WorkspaceTypeDto";
import { useWorkspaceTypeGetActionQuery } from "../../../sdk/abac/WorkspaceTypeGetAction";
import { useWorkspaceTypeUpdateAction } from "../../../sdk/abac/WorkspaceTypeUpdateAction";
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
