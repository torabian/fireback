import { CommonSingleManager } from "../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import { useRegionalContentGetActionQuery } from "../../sdk/abac/RegionalContentGetAction";
import { RegionalContentDto } from "../../sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "../../sdk/navigation/AbacNavigation";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
export const RegionalContentSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useRegionalContentGetActionQuery({
    params: { uniqueId },
  });
  var d: RegionalContentDto | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(RegionalContentNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.region,
              label: t.regionalContents.region,
            },
            {
              elem: d?.title,
              label: t.regionalContents.title,
            },
            {
              elem: d?.languageId,
              label: t.regionalContents.languageId,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
