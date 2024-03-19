import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetTagByUniqueId } from "src/sdk/fireback/modules/shop/useGetTagByUniqueId";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
export const TagSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetTagByUniqueId({ query: { uniqueId } });
  var d: TagEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(TagEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.name,
                label: t.tags.name,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};