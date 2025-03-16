import { useRouter } from "../../../hooks/useRouter";
import { CommonSingleManager } from "../../../components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../../components/page-title/PageTitle";
import { useLocale } from "../../../hooks/useLocale";
import { useT } from "../../../hooks/useT";
import { useGetWorkspaceTypeByUniqueId } from "../../../sdk/modules/workspaces/useGetWorkspaceTypeByUniqueId";
import { WorkspaceTypeEntity } from "../../../sdk/modules/workspaces/WorkspaceTypeEntity";

export const WorkspaceTypeSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetWorkspaceTypeByUniqueId({
    query: { uniqueId },
  });

  var d: WorkspaceTypeEntity | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.title || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(WorkspaceTypeEntity.Navigation.edit(uniqueId));
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
