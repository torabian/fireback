import { useT } from "@/modules/fireback/hooks/useT";

import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { userInvitationColumns } from "./UserInvitationColumns";
import { useGetWorkspaceInvites } from "../../sdk/modules/workspaces/useGetWorkspaceInvites";

export const UserInvitationList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={userInvitationColumns(t)}
        queryHook={useGetWorkspaceInvites}
      ></CommonListManager>
      {/* <SmartHead title={t.course.title} />
      {items.length === 0 && <span>{t.noPendingInvite}</span>}
      {items.map((item: any) => (
        <UserInvitationItem key={(item as any).uniqueId} invite={item} />
      ))} */}
    </>
  );
};
