import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useGetCapabilityByUniqueId } from "@/modules/fireback/sdk/modules/workspaces/useGetCapabilityByUniqueId";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/workspaces/CapabilityEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const CapabilitySingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetCapabilityByUniqueId({ query: { uniqueId } });
  var d: CapabilityEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(CapabilityEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.name,
                label: t.capabilities.name,
              },    
              {
                elem: d?.description,
                label: t.capabilities.description,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};
