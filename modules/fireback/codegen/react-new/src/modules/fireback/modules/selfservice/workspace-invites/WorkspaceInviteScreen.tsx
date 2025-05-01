import { strings } from "./strings/translations";
import { useT } from "@/modules/fireback/hooks/useT";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useS } from "@/modules/fireback/hooks/useS";
import { useGetWorkspaceInviteByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetWorkspaceInviteByUniqueId";
import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceInviteEntity";
import { usePageTitle } from "@/modules/fireback/hooks/authContext";
import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useRouter } from "@/modules/fireback/hooks/useRouter";

export const WorkspaceInviteSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();
  const s = useS(strings);

  const getSingleHook = useGetWorkspaceInviteByUniqueId({
    query: { uniqueId },
  });

  var d: WorkspaceInviteEntity | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.firstName + " " + d?.lastName || "");

  return (
    <>
      <CommonSingleManager
        getSingleHook={getSingleHook}
        editEntityHandler={() =>
          router.push(WorkspaceInviteEntity.Navigation.edit(uniqueId))
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
