import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import { useT } from "@/fireback/hooks/useT";
import { useGetOrderByUniqueId } from "src/sdk/fireback/modules/shop/useGetOrderByUniqueId";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
export const OrderSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetOrderByUniqueId({ query: { uniqueId } });
  var d: OrderEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(OrderEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.shippingAddress,
              label: t.orders.shippingAddress,
            },
            {
              elem: d?.invoiceNumber,
              label: t.orders.invoiceNumber,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
