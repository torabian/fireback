import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { WorkspaceInviteDto } from "@fireback/selfservice/sdk/abac/WorkspaceInviteDto";
import { useWorkspaceInviteGetActionQuery } from "@fireback/selfservice/sdk/abac/WorkspaceInviteGetAction";
import { WorkspaceInviteNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { strings } from "./strings/translations";

export const WorkspaceInviteSingleScreen = () => {
  const router = useRouter();
  const uniqueId = router.query.uniqueId as string;
  const s = useS(strings);

  const getSingleHook = useWorkspaceInviteGetActionQuery({
    params: { uniqueId },
  });

  var d: WorkspaceInviteDto | undefined = getSingleHook.data?.data?.item;
  usePageTitle(d?.firstName + " " + d?.lastName || "");

  return (
    <>
      <CommonSingleManager
        getSingleHook={getSingleHook}
        editEntityHandler={() =>
          router.push(WorkspaceInviteNavigation.edit(uniqueId))
        }
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.firstName,
              elem: d?.firstName,
            },
            {
              label: s.lastName,
              elem: d?.lastName,
            },
            {
              label: s.email,
              elem: d?.email,
            },
            {
              label: s.phoneNumber,
              elem: d?.phonenumber,
            },
            {
              label: s.forcedEmailAddress,
              elem: d?.forceEmailAddress,
            },
            {
              label: s.forcedPhone,
              elem: d?.forcePhoneNumber,
            },
            {
              label: s.targetLocale,
              elem: d?.targetUserLocale,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
