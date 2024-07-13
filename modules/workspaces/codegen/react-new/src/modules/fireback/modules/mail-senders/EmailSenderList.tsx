import { useT } from "@/modules/fireback/hooks/useT";
import { useDeleteEmailSender } from "../../sdk/modules/workspaces/useDeleteEmailSender";

import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
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
