import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { PageTagForm } from "./PageTagEditForm";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
import { useGetPageTagByUniqueId } from "src/sdk/fireback/modules/cms/useGetPageTagByUniqueId";
import { usePostPageTag } from "src/sdk/fireback/modules/cms/usePostPageTag";
import { usePatchPageTag } from "src/sdk/fireback/modules/cms/usePatchPageTag";
export const PageTagEntityManager = ({ data }: DtoEntity<PageTagEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<PageTagEntity>
  >({
    data,
  });
  const getSingleHook = useGetPageTagByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPageTag({
    queryClient,
  });
  const patchHook = usePatchPageTag({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PageTagEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PageTagEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PageTagForm}
      onEditTitle={t.pagetags.editpageTag}
      onCreateTitle={t.pagetags.newpageTag}
      data={data}
    />
  );
};
