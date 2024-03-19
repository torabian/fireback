import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetProductByUniqueId } from "src/sdk/fireback/modules/shop/useGetProductByUniqueId";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";

export const ProductSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetProductByUniqueId({ query: { uniqueId } });
  var d: ProductEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(ProductEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.name,
              label: t.products.name,
            },
            {
              elem: d?.description,
              label: t.products.description,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
