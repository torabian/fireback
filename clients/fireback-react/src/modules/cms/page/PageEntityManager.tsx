import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { PageForm } from "./PageEditForm";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
import { useGetPageByUniqueId } from "src/sdk/fireback/modules/cms/useGetPageByUniqueId";
import { usePostPage } from "src/sdk/fireback/modules/cms/usePostPage";
import { usePatchPage } from "src/sdk/fireback/modules/cms/usePatchPage";
export const PageEntityManager = ({ data }: DtoEntity<PageEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<PageEntity>
  >({
    data,
  });
  const getSingleHook = useGetPageByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPage({
    queryClient,
  });
  const patchHook = usePatchPage({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(PageEntity.Navigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        PageEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PageForm}
      onEditTitle={t.pages.editpage}
      onCreateTitle={t.pages.newpage}
      data={data}
    />
  );
};
