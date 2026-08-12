import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useEmailSenderBrowseActionQuery } from "@fireback/messaging/sdk/messaging/EmailSenderBrowseAction";
import { useEmailSenderAwareDeleteAction } from "@fireback/messaging/sdk/messaging/EmailSenderAwareDeleteAction";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { EmailSenderDto } from "@fireback/messaging/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { columns } from "./EmailSenderColumns";
import { strings } from "./strings/translations";

export const EmailSenderList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={useEmailSenderBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailSenderNavigation.single(uniqueId)
        }
        deleteHook={useEmailSenderAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
