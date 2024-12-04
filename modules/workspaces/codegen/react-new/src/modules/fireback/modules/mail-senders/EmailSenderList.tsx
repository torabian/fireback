import { useT } from "../../hooks/useT";
import { useDeleteEmailSender } from "../../sdk/modules/workspaces/useDeleteEmailSender";

import { CommonListManager } from "../../components/entity-manager/CommonListManager";
import { useGetEmailSenders } from "../../sdk/modules/workspaces/useGetEmailSenders";
import { columns } from "./EmailSenderColumns";
import { EmailSenderEntity } from "../../sdk/modules/workspaces/EmailSenderEntity";

export const EmailSenderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetEmailSenders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailSenderEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteEmailSender}
      ></CommonListManager>
    </>
  );
};
