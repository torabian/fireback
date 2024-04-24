import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetAppMenuByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetAppMenuByUniqueId";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
export const AppMenuSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetAppMenuByUniqueId({ query: { uniqueId } });
  var d: AppMenuEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(AppMenuEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.href,
                label: t.appMenus.href,
              },    
              {
                elem: d?.icon,
                label: t.appMenus.icon,
              },    
              {
                elem: d?.label,
                label: t.appMenus.label,
              },    
              {
                elem: d?.activeMatcher,
                label: t.appMenus.activeMatcher,
              },    
              {
                elem: d?.applyType,
                label: t.appMenus.applyType,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};