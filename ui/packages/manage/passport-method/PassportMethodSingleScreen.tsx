import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { usePassportMethodGetActionQuery } from "@fireback/manage/sdk/abac/PassportMethodGetAction";
import { PassportMethodNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
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
          router.push(PassportMethodNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView entity={d} fields={[]} />
      </CommonSingleManager>
    </>
  );
};
