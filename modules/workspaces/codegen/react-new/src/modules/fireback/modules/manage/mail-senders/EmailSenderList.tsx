import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { columns } from "./EmailSenderColumns";
import { useGetEmailSenders } from "@/modules/fireback/sdk/modules/workspaces/useGetEmailSenders";
import { EmailSenderEntity } from "@/modules/fireback/sdk/modules/workspaces/EmailSenderEntity";
import { useDeleteEmailSender } from "@/modules/fireback/sdk/modules/workspaces/useDeleteEmailSender";

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
