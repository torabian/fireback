import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { useWorkspaceGetActionQuery } from "@fireback/manage/sdk/abac/WorkspaceGetAction";
import { WorkspaceNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

export const WorkspaceSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useWorkspaceGetActionQuery({ params: { uniqueId } });
  var d: any | undefined = getSingleHook.data?.data?.item;
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
              label: s.name,
              elem: d?.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
