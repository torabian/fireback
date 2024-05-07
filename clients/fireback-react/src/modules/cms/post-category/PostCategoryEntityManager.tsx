import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { PostCategoryForm } from "./PostCategoryEditForm";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
import { useGetPostCategoryByUniqueId } from "src/sdk/fireback/modules/cms/useGetPostCategoryByUniqueId";
import { usePostPostCategory } from "src/sdk/fireback/modules/cms/usePostPostCategory";
import { usePatchPostCategory } from "src/sdk/fireback/modules/cms/usePatchPostCategory";
export const PostCategoryEntityManager = ({
  data,
}: DtoEntity<PostCategoryEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<PostCategoryEntity>
  >({
    data,
  });
  const getSingleHook = useGetPostCategoryByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPostCategory({
    queryClient,
  });
  const patchHook = usePatchPostCategory({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PostCategoryEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PostCategoryEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PostCategoryForm}
      onEditTitle={t.postcategories.editpostCategory}
      onCreateTitle={t.postcategories.newpostCategory}
      data={data}
    />
  );
};
