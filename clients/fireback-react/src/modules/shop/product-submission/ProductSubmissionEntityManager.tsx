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
export const ProductSubmissionEntityManager = ({
  data,
}: DtoEntity<ProductSubmissionEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<ProductSubmissionEntity>
  >({
    data,
  });
  const getSingleHook = useGetProductSubmissionByUniqueId({
    query: { uniqueId },
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
      onCancel={() => {
        router.goBackOrDefault(
          ProductSubmissionEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        ProductSubmissionEntity.Navigation.single(
          response.data?.uniqueId,
          locale
        )
      }
      Form={ProductSubmissionForm}
      onEditTitle={t.productsubmissions.editproductSubmission}
      onCreateTitle={t.productsubmissions.newproductSubmission}
      data={data}
    />
  );
};
