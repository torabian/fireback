import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGsmProviderBrowseActionQuery } from "@/modules/fireback/sdk/messaging/GsmProviderBrowseAction";
import { useGsmProviderAwareDeleteAction } from "@/modules/fireback/sdk/messaging/GsmProviderAwareDeleteAction";
import { columns } from "./GsmProviderColumns";
import { GsmProviderDto } from "@/modules/fireback/sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "@/modules/fireback/sdk/navigation/MessagingNavigation";
import { useS } from "@/modules/fireback/hooks/useS";
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
