import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetPostTagByUniqueId } from "src/sdk/fireback/modules/cms/useGetPostTagByUniqueId";
import { PostTagEntity } from "src/sdk/fireback/modules/cms/PostTagEntity";
export const PostTagSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetPostTagByUniqueId({ query: { uniqueId } });
  var d: PostTagEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(PostTagEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.name,
              label: t.posttags.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
