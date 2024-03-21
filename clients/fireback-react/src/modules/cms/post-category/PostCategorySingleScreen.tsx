import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetPostCategoryByUniqueId } from "src/sdk/fireback/modules/cms/useGetPostCategoryByUniqueId";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
export const PostCategorySingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetPostCategoryByUniqueId({ query: { uniqueId } });
  var d: PostCategoryEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(PostCategoryEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.name,
              label: t.postcategories.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
