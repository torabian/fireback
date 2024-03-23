import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { DiscountCodeForm } from "./DiscountCodeEditForm";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
import { useGetDiscountCodeByUniqueId } from "src/sdk/fireback/modules/shop/useGetDiscountCodeByUniqueId";
import { usePostDiscountCode } from "src/sdk/fireback/modules/shop/usePostDiscountCode";
import { usePatchDiscountCode } from "src/sdk/fireback/modules/shop/usePatchDiscountCode";
export const DiscountCodeEntityManager = ({ data }: DtoEntity<DiscountCodeEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<DiscountCodeEntity>
  >({
    data,
  });
  const getSingleHook = useGetDiscountCodeByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostDiscountCode({
    queryClient,
  });
  const patchHook = usePatchDiscountCode({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          DiscountCodeEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        DiscountCodeEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ DiscountCodeForm }
      onEditTitle={t.discountCodes.editDiscountCode }
      onCreateTitle={t.discountCodes.newDiscountCode }
      data={data}
    />
  );
};