import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useRegionalContentGetActionQuery } from "@/modules/fireback/sdk/abac/RegionalContentGetAction";
import { RegionalContentDto } from "@/modules/fireback/sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const RegionalContentSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useRegionalContentGetActionQuery({ params: { uniqueId } });
  var d: RegionalContentDto | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(RegionalContentNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.region,
                label: t.regionalContents.region,
              },    
              {
                elem: d?.title,
                label: t.regionalContents.title,
              },    
              {
                elem: d?.languageId,
                label: t.regionalContents.languageId,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};
