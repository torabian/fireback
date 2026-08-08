import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { UserInvitationList } from "./UserInvitationList";

export const UserInvitationArchiveScreen = () => {
  const s = useS(strings);

  return (
    <>
      <CommonArchiveManager pageTitle={s.menuTitle}>
        <UserInvitationList />
      </CommonArchiveManager>
    </>
  );
};
