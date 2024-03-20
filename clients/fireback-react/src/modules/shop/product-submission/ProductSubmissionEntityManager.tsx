import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { ProductSubmissionForm } from "./ProductSubmissionEditForm";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
import { useGetProductSubmissionByUniqueId } from "src/sdk/fireback/modules/shop/useGetProductSubmissionByUniqueId";
import { usePostProductSubmission } from "src/sdk/fireback/modules/shop/usePostProductSubmission";
import { usePatchProductSubmission } from "src/sdk/fireback/modules/shop/usePatchProductSubmission";
import { uuidv4 } from "@/helpers/api";
export const ProductSubmissionEntityManager = ({
  data,
}: DtoEntity<ProductSubmissionEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<ProductSubmissionEntity>
  >({
    data,
  });
  const getSingleHook = useGetProductSubmissionByUniqueId({
    query: {
      uniqueId,
      deep: true,
      withPreloads: "Price.Variations,Price.Variations.Currency",
    },
  });
  const postHook = usePostProductSubmission({
    queryClient,
  });
  const patchHook = usePatchProductSubmission({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      customClass=""
      onCancel={() => {
        router.goBackOrDefault(
          ProductSubmissionEntity.Navigation.query(undefined, locale)
        );
      }}
      // beforeSubmit={(data) => {
      //   data.price.uniqueId = uuidv4();
      //   return data;
      // }}
      onFinishUriResolver={(response, locale) => {
        return ProductSubmissionEntity.Navigation.single(
          response.data?.uniqueId,
          locale
        );
      }}
      Form={ProductSubmissionForm}
      onEditTitle={t.productsubmissions.editproductSubmission}
      onCreateTitle={t.productsubmissions.newproductSubmission}
      data={data}
    />
  );
};
