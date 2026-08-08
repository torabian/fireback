import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useGsmProviderBrowseActionQuery } from "../../../sdk/messaging/GsmProviderBrowseAction";
import { useGsmProviderAwareDeleteAction } from "../../../sdk/messaging/GsmProviderAwareDeleteAction";
import { columns } from "./GsmProviderColumns";
import { GsmProviderDto } from "../../../sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { useS } from "../../../hooks/useS";
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
