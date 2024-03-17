import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetCategoryByUniqueId } from "src/sdk/fireback/modules/shop/useGetCategoryByUniqueId";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
export const CategorySingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetCategoryByUniqueId({ query: { uniqueId } });
  var d: CategoryEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(CategoryEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.name,
                label: t.categories.name,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};