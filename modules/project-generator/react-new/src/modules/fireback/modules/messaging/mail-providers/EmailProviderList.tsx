import { useT } from "@/modules/fireback/hooks/useT";
import { useEmailProviderBrowseActionQuery } from "@/modules/fireback/sdk/messaging/EmailProviderBrowseAction";
import { useEmailProviderAwareDeleteAction } from "@/modules/fireback/sdk/messaging/EmailProviderAwareDeleteAction";
import { useS } from "@/modules/fireback/hooks/useS";
import { columns } from "./EmailProviderColumns";
import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { EmailProviderDto } from "@/modules/fireback/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
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
