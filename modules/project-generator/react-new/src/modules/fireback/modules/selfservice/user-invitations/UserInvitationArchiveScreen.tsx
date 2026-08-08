import { useT } from "../../../../fireback-ui/hooks/useT";

import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
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
