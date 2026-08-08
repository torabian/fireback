import { CommonSingleManager } from "../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { usePassportMethodGetActionQuery } from "../../sdk/abac/PassportMethodGetAction";
import { PassportMethodNavigation } from "../../sdk/navigation/AbacNavigation";
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
