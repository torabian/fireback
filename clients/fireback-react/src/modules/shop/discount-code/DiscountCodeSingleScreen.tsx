import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import { useT } from "@/fireback/hooks/useT";
import { useGetDiscountCodeByUniqueId } from "src/sdk/fireback/modules/shop/useGetDiscountCodeByUniqueId";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
export const DiscountCodeSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetDiscountCodeByUniqueId({ query: { uniqueId } });
  var d: DiscountCodeEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(DiscountCodeEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.series,
              label: t.discountCodes.series,
            },
            {
              elem: d?.limit,
              label: t.discountCodes.limit,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
