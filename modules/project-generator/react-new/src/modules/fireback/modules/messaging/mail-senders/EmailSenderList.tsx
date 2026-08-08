import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useEmailSenderBrowseActionQuery } from "@/modules/fireback/sdk/messaging/EmailSenderBrowseAction";
import { useEmailSenderAwareDeleteAction } from "@/modules/fireback/sdk/messaging/EmailSenderAwareDeleteAction";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailSenderDto } from "@/modules/fireback/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
import { columns } from "./EmailSenderColumns";

export const EmailSenderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useEmailSenderBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailSenderNavigation.single(uniqueId)
        }
        deleteHook={useEmailSenderAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
