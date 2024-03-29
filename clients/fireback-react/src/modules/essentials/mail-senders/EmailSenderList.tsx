import { useT } from "@/hooks/useT";
import { useDeleteEmailSender } from "src/sdk/fireback/modules/workspaces/useDeleteEmailSender";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useGetEmailSenders } from "src/sdk/fireback/modules/workspaces/useGetEmailSenders";
import { columns } from "./EmailSenderColumns";
import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";

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
