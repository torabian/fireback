import { strings } from "./strings/translations";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { useWorkspaceInviteGetActionQuery } from "../../../sdk/abac/WorkspaceInviteGetAction";
import { WorkspaceInviteDto } from "../../../sdk/abac/WorkspaceInviteDto";
import { WorkspaceInviteNavigation } from "../../../sdk/navigation/AbacNavigation";
import { usePageTitle } from "../../../../fireback-ui/hooks/authContext";
import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useRouter } from "../../../../fireback-ui/hooks/useRouter";

export const WorkspaceInviteSingleScreen = () => {
  const router = useRouter();
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
