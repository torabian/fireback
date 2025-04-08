import { useS } from "@/modules/fireback/hooks/useS";
import { useGetUsersInvitations } from "@/modules/fireback/sdk/modules/abac/useGetUsersInvitations";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { strings } from "./strings/translations";
import { userInvitationColumns } from "./UserInvitationColumns";

import { ModalContext } from "@/modules/fireback/components/modal/Modal";
import { useContext } from "react";
import { UserInvitationsQueryColumns } from "@/modules/fireback/sdk/modules/abac/UserInvitationsQueryColumns";
import { usePostUserInvitationAccept } from "@/modules/fireback/sdk/modules/abac/usePostUserInvitationAccept";

export const UserInvitationList = () => {
  const s = useS(strings);

  const useModal = useContext(ModalContext);

  const { submit: acceptInvite } = usePostUserInvitationAccept();

  const onAccept = (dto: UserInvitationsQueryColumns) => {
    useModal.openModal({
      title: s.confirmAcceptTitle,
      confirmButtonLabel: s.acceptBtn,
      component: () => <div>{s.confirmAcceptDescription}</div>,
      onSubmit: async () => {
        return acceptInvite({ invitationUniqueId: dto.uniqueId }).then(
          (res) => {
            console.log("Accept:", res);
          }
        );
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
