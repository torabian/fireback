import { useS } from "@/modules/fireback/hooks/useS";
import { useGetUsersInvitations } from "@/modules/fireback/sdk/modules/workspaces/useGetUsersInvitations";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { strings } from "./strings/translations";
import { userInvitationColumns } from "./UserInvitationColumns";
import { UserInvitationsActionResDto } from "@/modules/fireback/sdk/modules/workspaces/WorkspacesActionsDto";
import { ModalContext } from "@/modules/fireback/components/modal/Modal";
import { useContext } from "react";

export const UserInvitationList = () => {
  const s = useS(strings);

  const useModal = useContext(ModalContext);

  const onAccept = (dto: UserInvitationsActionResDto) => {
    useModal.openModal({
      title: s.confirmAcceptTitle,
      confirmButtonLabel: s.acceptBtn,
      component: () => <div>{s.confirmAcceptDescription}</div>,
      onSubmit: async () => {
        return true;
      },
    });
  };

  const onReject = (dto: UserInvitationsActionResDto) => {
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
