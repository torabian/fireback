import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../../../fireback-ui/components/page-title/PageTitle";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useWorkspaceGetActionQuery } from "../../../sdk/abac/WorkspaceGetAction";
import { WorkspaceNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useWorkspaceGetActionQuery({ params: { uniqueId } });
  var d: any | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.name || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(WorkspaceNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: t.wokspaces.name,
              elem: d?.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
