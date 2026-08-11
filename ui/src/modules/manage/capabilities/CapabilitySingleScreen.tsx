import { CommonSingleManager } from "../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { useCapabilityGetActionQuery } from "../../sdk/abac/CapabilityGetAction";
import { CapabilityDto } from "../../sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "../../sdk/navigation/AbacNavigation";
import { strings } from "./strings/translations";
import { usePageTitle } from "@/modules/fireback-ui/components/page-title/PageTitle";

export const CapabilitySingleScreen = () => {
  const { uniqueId } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useCapabilityGetActionQuery({ params: { uniqueId } });
  var d = getSingleHook.data?.data?.item;

  const t = useS(strings);
  usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(CapabilityNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.name,
              label: t.capabilities.name,
            },
            {
              elem: d?.description,
              label: t.capabilities.description,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
