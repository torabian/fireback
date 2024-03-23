import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { OrderForm } from "./OrderEditForm";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
import { useGetOrderByUniqueId } from "src/sdk/fireback/modules/shop/useGetOrderByUniqueId";
import { usePostOrder } from "src/sdk/fireback/modules/shop/usePostOrder";
import { usePatchOrder } from "src/sdk/fireback/modules/shop/usePatchOrder";
export const OrderEntityManager = ({ data }: DtoEntity<OrderEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<OrderEntity>
  >({
    data,
  });
  const getSingleHook = useGetOrderByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostOrder({
    queryClient,
  });
  const patchHook = usePatchOrder({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          OrderEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        OrderEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ OrderForm }
      onEditTitle={t.orders.editOrder }
      onCreateTitle={t.orders.newOrder }
      data={data}
    />
  );
};