import { useS } from "@/modules/fireback/hooks/useS";
import { useGetUsersInvitations } from "@/modules/fireback/sdk/modules/workspaces/useGetUsersInvitations";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { strings } from "./strings/translations";
import { userInvitationColumns } from "./UserInvitationColumns";
// Fireback doesn't export queries data types into javascript unfortunately.
// Add this, and later change it
// import { UserInvitationsQueryColumns } from "@/modules/fireback/sdk/modules/workspaces/WorkspacesActionsDto";
import { ModalContext } from "@/modules/fireback/components/modal/Modal";
import { useContext } from "react";

type UserInvitationsQueryColumns = any;
export const UserInvitationList = () => {
  const s = useS(strings);

  const useModal = useContext(ModalContext);

  const onAccept = (dto: UserInvitationsQueryColumns) => {
    useModal.openModal({
      title: s.confirmAcceptTitle,
      confirmButtonLabel: s.acceptBtn,
      component: () => <div>{s.confirmAcceptDescription}</div>,
      onSubmit: async () => {
        return true;
      },
    });
  };

  const onReject = (dto: UserInvitationsQueryColumns) => {
    useModal.openModal({
      title: s.confirmRejectTitle,
      confirmButtonLabel: s.acceptBtn,
      component: () => <div>{s.confirmRejectDescription}</div>,
      onSubmit: async () => {
        return true;
      },
    });
  };

  return (
    <>
      <CommonListManager
        selectable={false}
        columns={userInvitationColumns(s, onAccept, onReject)}
        queryHook={useGetUsersInvitations}
      ></CommonListManager>
    </>
  );
};
