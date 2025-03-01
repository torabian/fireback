import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useGetWorkspaceConfigByUniqueId } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceConfigByUniqueId";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { useGetWorkspaceConfigDistinct } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceConfigDistinct";

export const WorkspaceConfigSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetWorkspaceConfigDistinct({});
  var d: WorkspaceConfigEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(`/${locale}/root/workspace/config/edit`);
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.recaptcha2ServerKey,
              label: t.workspaceConfigs.recaptcha2ServerKey,
            },
            {
              elem: d?.recaptcha2ClientKey,
              label: t.workspaceConfigs.recaptcha2ClientKey,
            },
            {
              elem: d?.enableOtp,
              label: t.workspaceConfigs.enableOtp,
            },
            {
              elem: d?.enableRecaptcha2,
              label: t.workspaceConfigs.enableRecaptcha2,
            },
            {
              elem: d?.enableTotp,
              label: t.workspaceConfigs.enableTotp,
            },
            {
              elem: d?.forcePasswordOnPhone,
              label: t.workspaceConfigs.forcePasswordOnPhone,
            },
            {
              elem: d?.forcePersonNameOnPhone,
              label: t.workspaceConfigs.forcePersonNameOnPhone,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
