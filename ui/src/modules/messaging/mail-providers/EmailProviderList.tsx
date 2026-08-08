import { useEmailProviderBrowseActionQuery } from "../../sdk/messaging/EmailProviderBrowseAction";
import { useEmailProviderAwareDeleteAction } from "../../sdk/messaging/EmailProviderAwareDeleteAction";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../fireback-ui/components/strings/translations";
import { columns } from "./EmailProviderColumns";
import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { EmailProviderDto } from "../../sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "../../sdk/navigation/MessagingNavigation";
import { strings } from "./strings/translations";

export const EmailProviderList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={useEmailProviderBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailProviderNavigation.single(uniqueId)
        }
        deleteHook={useEmailProviderAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
