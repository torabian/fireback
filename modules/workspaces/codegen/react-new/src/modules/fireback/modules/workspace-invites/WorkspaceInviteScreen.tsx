import { useRouter } from "../../hooks/useRouter";
import { useEditAction } from "../../components/action-menu/ActionMenu";
import { QueryErrorView } from "../../components/error-view/QueryError";
import { GeneralEntityView } from "../../components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../components/page-title/PageTitle";
import { useLocale } from "../../hooks/useLocale";
import { useT } from "../../hooks/useT";
import { useGetWorkspaceByUniqueId } from "../../sdk/modules/workspaces/useGetWorkspaceByUniqueId";
import { WorkspaceEntity } from "../../sdk/modules/workspaces/WorkspaceEntity";

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
