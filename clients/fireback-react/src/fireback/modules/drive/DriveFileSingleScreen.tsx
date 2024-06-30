import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/fireback/components/page-title/PageTitle";
import { useRemoteInformation } from "@/fireback/hooks/useEnvironment";
import { useT } from "@/fireback/hooks/useT";
import { useRouter } from "@/Router";
import { FileEntity } from "@/sdk/fireback/modules/workspaces/FileEntity";
import { useGetFileByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetFileByUniqueId";

export const DriveFileSingleScreen = () => {
  const router = useRouter();
  const uniqueId = router.query.uniqueId as string;

  const getSingleHook = useGetFileByUniqueId({ query: { uniqueId } });
  let d: FileEntity | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.name || "");
  const t = useT();
  const { directPath } = useRemoteInformation();

  return (
    <>
      <CommonSingleManager getSingleHook={getSingleHook}>
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: t.drive.name,
              elem: d?.name,
            },
            {
              label: t.drive.size,
              elem: d?.size,
            },
            {
              label: t.drive.type,
              elem: d?.type,
            },
            {
              label: t.drive.virtualPath,
              elem: d?.virtualPath,
            },
            {
              label: t.drive.viewPath,
              elem: <pre>{directPath(d)}</pre>,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
