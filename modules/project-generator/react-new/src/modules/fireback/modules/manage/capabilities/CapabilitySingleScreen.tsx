import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { useCapabilityGetActionQuery } from "@/modules/fireback/sdk/abac/CapabilityGetAction";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/fireback/CapabilityEntity";
import { strings } from "./strings/translations";
import { usePageTitle } from "@/modules/fireback/hooks/authContext";

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
          router.push(CapabilityEntity.Navigation.edit(uniqueId));
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
