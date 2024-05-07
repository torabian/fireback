import { useT } from "@/fireback/hooks/useT";

import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
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
