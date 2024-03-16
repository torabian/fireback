import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { CategoryForm } from "./CategoryEditForm";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
import { useGetCategoryByUniqueId } from "src/sdk/fireback/modules/shop/useGetCategoryByUniqueId";
import { usePostCategory } from "src/sdk/fireback/modules/shop/usePostCategory";
import { usePatchCategory } from "src/sdk/fireback/modules/shop/usePatchCategory";
export const CategoryEntityManager = ({ data }: DtoEntity<CategoryEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<CategoryEntity>
  >({
    data,
  });
  const getSingleHook = useGetCategoryByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostCategory({
    queryClient,
  });
  const patchHook = usePatchCategory({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          CategoryEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        CategoryEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={CategoryForm}
      onEditTitle={t.categories.editCategory}
      onCreateTitle={t.categories.newCategory}
      data={data}
    />
  );
};
