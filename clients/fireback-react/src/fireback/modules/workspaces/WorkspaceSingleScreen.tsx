import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/fireback/components/page-title/PageTitle";
import { useLocale } from "@/fireback/hooks/useLocale";
import { useT } from "@/fireback/hooks/useT";
import { useGetWorkspaceByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetWorkspaceByUniqueId";
import { WorkspaceEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceEntity";

export const WorkspaceSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetWorkspaceByUniqueId({ query: { uniqueId } });
  var d: any | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.name || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(WorkspaceEntity.Navigation.edit(uniqueId, locale));
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
