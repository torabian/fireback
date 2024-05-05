import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { PageCategoryForm } from "./PageCategoryEditForm";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
import { useGetPageCategoryByUniqueId } from "src/sdk/fireback/modules/cms/useGetPageCategoryByUniqueId";
import { usePostPageCategory } from "src/sdk/fireback/modules/cms/usePostPageCategory";
import { usePatchPageCategory } from "src/sdk/fireback/modules/cms/usePatchPageCategory";
export const PageCategoryEntityManager = ({
  data,
}: DtoEntity<PageCategoryEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<PageCategoryEntity>
  >({
    data,
  });
  const getSingleHook = useGetPageCategoryByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPageCategory({
    queryClient,
  });
  const patchHook = usePatchPageCategory({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PageCategoryEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PageCategoryEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PageCategoryForm}
      onEditTitle={t.pagecategories.editpageCategory}
      onCreateTitle={t.pagecategories.newpageCategory}
      data={data}
    />
  );
};
