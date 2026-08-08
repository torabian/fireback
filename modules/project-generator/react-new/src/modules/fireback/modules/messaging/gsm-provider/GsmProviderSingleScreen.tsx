import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useGsmProviderGetActionQuery } from "../../../sdk/messaging/GsmProviderGetAction";
import { GsmProviderDto } from "../../../sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
export const GsmProviderSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGsmProviderGetActionQuery({ params: { uniqueId } });
  var d: GsmProviderDto | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(GsmProviderNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.apiKey,
                label: t.gsmProviders.apiKey,
              },    
              {
                elem: d?.mainSenderNumber,
                label: t.gsmProviders.mainSenderNumber,
              },    
              {
                elem: d?.invokeUrl,
                label: t.gsmProviders.invokeUrl,
              },    
              {
                elem: d?.invokeBody,
                label: t.gsmProviders.invokeBody,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};
