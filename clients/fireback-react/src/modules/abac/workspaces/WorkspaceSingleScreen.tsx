import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { useGetWorkspaceByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetWorkspaceByUniqueId";
import { WorkspaceNavigationTools } from "src/sdk/fireback/modules/workspaces/workspace-navigation-tools";

export const WorkspaceSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetWorkspaceByUniqueId({ query: { uniqueId } });
  // var d: WorkspaceEntity | undefined = query.data?.data;
  var d: any | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.name || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(WorkspaceNavigationTools.edit(uniqueId, locale));
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
