import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetProductSubmissionByUniqueId } from "src/sdk/fireback/modules/shop/useGetProductSubmissionByUniqueId";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
export const ProductSubmissionSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetProductSubmissionByUniqueId({ query: { uniqueId } });
  var d: ProductSubmissionEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(ProductSubmissionEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.description,
                label: t.productsubmissions.description,
              },    
              {
                elem: d?.sku,
                label: t.productsubmissions.sku,
              },    
              {
                elem: d?.brand,
                label: t.productsubmissions.brand,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};