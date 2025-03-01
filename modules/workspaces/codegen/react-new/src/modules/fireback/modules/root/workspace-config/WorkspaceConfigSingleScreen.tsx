import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useS } from "@/modules/fireback/hooks/useS";
import { useGetWorkspaceConfigDistinct } from "@/modules/fireback/sdk/modules/workspaces/useGetWorkspaceConfigDistinct";
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
import { strings } from "./strings/translations";

export const WorkspaceConfigSingleScreen = () => {
  const getSingleHook = useGetWorkspaceConfigDistinct({});
  var d: WorkspaceConfigEntity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(`/${locale}/root/workspace/config/edit`);
        }}
        noBack
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          title={t.workspaceConfigs.title}
          description={t.workspaceConfigs.description}
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
              elem: d?.requireOtpOnSignin,
              label: t.workspaceConfigs.requireOtpOnSignin,
            },
            {
              elem: d?.requireOtpOnSignup,
              label: t.workspaceConfigs.requireOtpOnSignup,
            },
            {
              elem: d?.enableTotp,
              label: t.workspaceConfigs.enableTotp,
            },
            {
              elem: d?.forceTotp,
              label: t.workspaceConfigs.forceTotp,
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
