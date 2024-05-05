import { useRouter } from "@/Router";
import { useEditAction } from "@/fireback/components/action-menu/ActionMenu";
import { QueryErrorView } from "@/fireback/components/error-view/QueryError";
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

  const { query } = useGetWorkspaceByUniqueId({ query: { uniqueId } });
  // var d: WorkspaceEntity | undefined = query.data?.data;
  var d: any | undefined = query.data?.data;
  usePageTitle(d?.name || "");

  useEditAction(() => {
    router.push(WorkspaceEntity.Navigation.edit(uniqueId, locale));
  });

  return (
    <>
      <QueryErrorView query={query} />
      <GeneralEntityView
        entity={d}
        fields={[
          {
            label: t.wokspaces.invite.name,
            elem: d?.name,
          },
        ]}
      />
    </>
  );
};
