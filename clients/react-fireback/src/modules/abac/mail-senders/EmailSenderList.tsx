import { useT } from "@/hooks/useT";
import { EmailSenderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-sender-navigation-tools";
import { useDeleteEmailSender } from "src/sdk/fireback/modules/workspaces/useDeleteEmailSender";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useGetEmailSenders } from "src/sdk/fireback/modules/workspaces/useGetEmailSenders";
import { columns } from "./EmailSenderColumns";

export const EmailSenderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetEmailSenders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailSenderNavigationTools.single(uniqueId)
        }
        deleteHook={useDeleteEmailSender}
      ></CommonListManager>
    </>
  );
};
