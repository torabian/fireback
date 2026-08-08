import { CommonSingleManager } from "../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../fireback-ui/components/page-title/PageTitle";
import { useRouter } from "../../fireback-ui/hooks/useRouter";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { useWorkspaceTypeGetActionQuery } from "../../sdk/abac/WorkspaceTypeGetAction";
import { WorkspaceTypeNavigation } from "../../sdk/navigation/AbacNavigation";

export const WorkspaceTypeSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;

  const getSingleHook = useWorkspaceTypeGetActionQuery({
    params: { uniqueId },
  });

  var d = getSingleHook.data?.data?.item;
  usePageTitle(d?.title || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(WorkspaceTypeNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.slug,
              elem: d?.slug,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
