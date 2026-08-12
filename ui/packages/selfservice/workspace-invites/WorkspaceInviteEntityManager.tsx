import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { useWorkspaceInviteGetActionQuery } from "@fireback/selfservice/sdk/abac/WorkspaceInviteGetAction";
import { useInviteToWorkspaceAction } from "@fireback/selfservice/sdk/abac/InviteToWorkspaceAction";
import { useWorkspaceInviteUpdateAction } from "@fireback/selfservice/sdk/abac/WorkspaceInviteUpdateAction";
import { WorkspaceInviteDto } from "@fireback/selfservice/sdk/abac/WorkspaceInviteDto";
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

  const getSingleHook = useWorkspaceInviteGetActionQuery({
    params: { uniqueId },
  });

  // Creating an invite has to go through InviteToWorkspaceAction (not the plain CRUD
  // WorkspaceInviteCreateAction) - it's the action that actually sends the invite
  // email/SMS, see InviteToWorkspaceActionImplementation.go. WorkspaceInvitationDto
  // (its body type) carries every field this form sets
  // (firstName/lastName/email/phonenumber/roleId/coverLetter/targetUserLocale/
  // forceEmailAddress/forcePhoneNumber), so no form changes are needed.
  //
  // creatorFn is overridden to the identity function as a workaround: unlike an
  // entity's generated CreateAction (whose Fetch always wraps the response in
  // GResponse.inject(), see e.g. GsmProviderCreateAction.ts), a hand-declared
  // "dto in/out" action like this one's generated Fetch constructs its creatorFn's
  // result (by default `new WorkspaceInvitationDto(item)`) directly from the raw
  // response body - even though the backend actually sends the same
  // {"data":{"item":...}}/{"error":...} envelope either way (see
  // InviteToWorkspaceActionImplementation.go's fireback.GResponseSingleItem). Since
  // none of WorkspaceInvitationDto's own fields are named "data"/"error", that default
  // construction silently discards the whole payload. CommonEntityManager.tsx's
  // response.data?.item/response.error handling needs the raw envelope shape, so this
  // passes it through unchanged instead.
  const postHook = useInviteToWorkspaceAction({ creatorFn: (item: unknown) => item as any });

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
