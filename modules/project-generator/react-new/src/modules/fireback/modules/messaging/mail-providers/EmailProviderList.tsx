import { useT } from "../../../hooks/useT";
import { useEmailProviderBrowseActionQuery } from "../../../sdk/messaging/EmailProviderBrowseAction";
import { useEmailProviderAwareDeleteAction } from "../../../sdk/messaging/EmailProviderAwareDeleteAction";
import { useS } from "../../../hooks/useS";
import { columns } from "./EmailProviderColumns";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { EmailProviderDto } from "../../../sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { strings } from "./strings/translations";

export const EmailProviderList = () => {
  const t = useT();
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(t, s)}
        queryHook={useEmailProviderBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailProviderNavigation.single(uniqueId)
        }
        deleteHook={useEmailProviderAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
