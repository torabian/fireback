import { strings } from "./strings/translations";
import { useT } from "../../../hooks/useT";
import { useLocale } from "../../../hooks/useLocale";
import { useS } from "../../../hooks/useS";
import { useWorkspaceInviteGetActionQuery } from "../../../sdk/abac/WorkspaceInviteGetAction";
import { WorkspaceInviteDto } from "../../../sdk/abac/WorkspaceInviteDto";
import { WorkspaceInviteNavigation } from "../../../sdk/navigation/AbacNavigation";
import { usePageTitle } from "../../../hooks/authContext";
import { CommonSingleManager } from "../../../components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../components/general-entity-view/GeneralEntityView";
import { useRouter } from "../../../hooks/useRouter";

export const WorkspaceInviteSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();
  const s = useS(strings);

  const getSingleHook = useGetWorkspaceInviteByUniqueId({
    query: { uniqueId },
  });

  var d: WorkspaceInviteDto | undefined = getSingleHook.query.data?.data;
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
              label: t.wokspaces.invite.firstName,
              elem: d?.firstName,
            },
            {
              label: t.wokspaces.invite.lastName,
              elem: d?.lastName,
            },
            {
              label: t.wokspaces.invite.email,
              elem: d?.email,
            },
            {
              label: t.wokspaces.invite.phoneNumber,
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
