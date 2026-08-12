import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useGsmProviderBrowseActionQuery } from "@fireback/messaging/sdk/messaging/GsmProviderBrowseAction";
import { useGsmProviderAwareDeleteAction } from "@fireback/messaging/sdk/messaging/GsmProviderAwareDeleteAction";
import { columns } from "./GsmProviderColumns";
import { GsmProviderDto } from "@fireback/messaging/sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
export const GsmProviderList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGsmProviderBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          GsmProviderNavigation.single(uniqueId)
        }
        deleteHook={useGsmProviderAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
