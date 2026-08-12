import { useEmailProviderBrowseActionQuery } from "@fireback/messaging/sdk/messaging/EmailProviderBrowseAction";
import { useEmailProviderAwareDeleteAction } from "@fireback/messaging/sdk/messaging/EmailProviderAwareDeleteAction";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { columns } from "./EmailProviderColumns";
import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { EmailProviderDto } from "@fireback/messaging/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
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
