import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useGetPassportMethodByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetPassportMethodByUniqueId";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/abac/PassportMethodEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const PassportMethodSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetPassportMethodByUniqueId({ query: { uniqueId } });
  var d: PassportMethodEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(PassportMethodEntity.Navigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView entity={d} fields={[]} />
      </CommonSingleManager>
    </>
  );
};
