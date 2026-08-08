import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { useWorkspaceInviteGetActionQuery } from "../../../sdk/abac/WorkspaceInviteGetAction";
import { useWorkspaceInviteCreateAction } from "../../../sdk/abac/WorkspaceInviteCreateAction";
import { useWorkspaceInviteUpdateAction } from "../../../sdk/abac/WorkspaceInviteUpdateAction";
import { WorkspaceInviteDto } from "../../../sdk/abac/WorkspaceInviteDto";
import { WorkspaceInviteForm } from "./WorkspaceInviteForm";

export const WorkspaceInviteEntityManager = ({
  data,
}: DtoEntity<WorkspaceInviteDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<WorkspaceInviteDto>
  >({
    data,
  });

  const getSingleHook = useGetWorkspaceInviteByUniqueId({
    query: { uniqueId },
    queryClient,
  });

  const postHook = useWorkspaceInviteCreateAction({});

  const patchHook = useWorkspaceInviteUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(`/${locale}/workspace-invites`);
      }}
      onFinishUriResolver={(_, locale) => `/${locale}/workspace-invites`}
      Form={WorkspaceInviteForm}
      onEditTitle={s.editInvitation}
      onCreateTitle={s.createInvitation}
      data={data}
    />
  );
};
