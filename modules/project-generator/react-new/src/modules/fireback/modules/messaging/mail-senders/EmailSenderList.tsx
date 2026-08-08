import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useEmailSenderBrowseActionQuery } from "../../../sdk/messaging/EmailSenderBrowseAction";
import { useEmailSenderAwareDeleteAction } from "../../../sdk/messaging/EmailSenderAwareDeleteAction";
import { useT } from "../../../hooks/useT";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../../sdk/navigation/MessagingNavigation";
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
