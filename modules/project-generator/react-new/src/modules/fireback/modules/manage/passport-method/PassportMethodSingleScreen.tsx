import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { usePassportMethodGetActionQuery } from "@/modules/fireback/sdk/abac/PassportMethodGetAction";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/abac/PassportMethodEntity";
import { strings } from "./strings/translations";

export const PassportMethodSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = usePassportMethodGetActionQuery({
    params: { uniqueId },
  });
  var d = getSingleHook.data?.data?.item;
  const t = useS(strings);

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
