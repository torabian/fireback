import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { WorkspaceTypeEntity } from "src/sdk/fireback";
import { useGetWorkspaceTypeByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetWorkspaceTypeByUniqueId";

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
          router.push(`/${locale}/workspace-type/edit/${uniqueId}`);
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
            {
              label: t.role.name,
              elem: d?.role?.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
