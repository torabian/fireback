import { useT } from "@/modules/fireback/hooks/useT";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { UserInvitationList } from "./UserInvitationList";

export const UserInvitationArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager pageTitle={t.fbMenu.myInvitations}>
        <UserInvitationList />
      </CommonArchiveManager>
    </>
  );
};
