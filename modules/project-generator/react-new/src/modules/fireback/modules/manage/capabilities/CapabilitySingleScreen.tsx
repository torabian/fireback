import { CommonSingleManager } from "../../../components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "../../../hooks/useCommonEntityManager";
import { useS } from "../../../hooks/useS";
import { useCapabilityGetActionQuery } from "../../../sdk/abac/CapabilityGetAction";
import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "../../../sdk/navigation/AbacNavigation";
import { strings } from "./strings/translations";
import { usePageTitle } from "../../../hooks/authContext";

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
