import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../../../fireback-ui/components/page-title/PageTitle";
import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useWorkspaceTypeGetActionQuery } from "../../../sdk/abac/WorkspaceTypeGetAction";
import { WorkspaceTypeNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceTypeSingleScreen = () => {
  const router = useRouter();
  const t = useT();
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
              label: t.wokspaces.slug,
              elem: d?.slug,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
