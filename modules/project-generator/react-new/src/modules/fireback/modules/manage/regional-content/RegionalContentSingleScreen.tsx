import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useGetRegionalContentByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetRegionalContentByUniqueId";
import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const RegionalContentSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetRegionalContentByUniqueId({ query: { uniqueId } });
  var d: RegionalContentEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(RegionalContentEntity.Navigation.edit(uniqueId));
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
