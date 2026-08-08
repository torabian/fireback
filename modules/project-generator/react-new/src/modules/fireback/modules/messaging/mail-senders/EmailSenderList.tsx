import { CommonListManager } from "../../../../fireback-ui/components/entity-manager/CommonListManager";
import { useEmailSenderBrowseActionQuery } from "../../../sdk/messaging/EmailSenderBrowseAction";
import { useEmailSenderAwareDeleteAction } from "../../../sdk/messaging/EmailSenderAwareDeleteAction";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../../../fireback-ui/components/strings/translations";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../../sdk/navigation/MessagingNavigation";
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
