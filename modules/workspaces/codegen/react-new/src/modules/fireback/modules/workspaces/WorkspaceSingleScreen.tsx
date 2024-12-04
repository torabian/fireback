import { useRouter } from "../../hooks/useRouter";
import { CommonSingleManager } from "../../components/entity-manager/CommonSingleManager";
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
