import { FormCheckbox } from "@/fireback/components/forms/form-switch/FormSwitch";
import { PageSection } from "@/fireback/components/page-section/PageSection";
import { useT } from "@/fireback/hooks/useT";

import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";

import { FormikProps } from "formik";
import { useContext } from "react";
import { MailTemplateConfiguration } from "./MailTemplateConfiguration";
import { NotificationConfigEntity } from "@/sdk/fireback/modules/workspaces/NotificationConfigEntity";

export const WorkspaceNotificationForm = ({
  form,
}: {
  form: FormikProps<Partial<NotificationConfigEntity>>;
}) => {
  const { values, setFieldValue, setValues, errors } = form;
  const t = useT();
  const { options } = useContext(RemoteQueryContext);

  return (
    <>
      <PageSection title={t.wokspaces.mailServerConfiguration}>
        <FormCheckbox
          onChange={(value: boolean) =>
            setFieldValue(
              NotificationConfigEntity.Fields.cascadeToSubWorkspaces,
              value
            )
          }
          value={values.cascadeToSubWorkspaces}
          label={t.wokspaces.cascadeNotificationConfig}
          hint={t.wokspaces.cascadeNotificationConfigHint}
        />
        <FormCheckbox
          onChange={(value: boolean) =>
            setFieldValue(
              NotificationConfigEntity.Fields.cascadeToSubWorkspaces,
              value
            )
          }
          value={values.cascadeToSubWorkspaces}
          label={t.wokspaces.forceSubWorkspaceUseConfig}
        />
        {/* <FormEntitySelect2
          label={t.wokspaces.generalMailProvider}
          fnLoadOptions={async (keyword) => {
            return (
              (
                await EmailProviderActions.fn(options)
                  .query(`name %"${keyword}"%`)
                  .getEmailProviders()
              ).data?.items || []
            );
          }}
          value={form.values.generalEmailProvider}
          onChange={(entity: EmailProviderEntity) => {
            form.setValues({
              ...values,
              generalEmailProvider: entity,
              generalEmailProviderId: entity.uniqueId,
            });
          }}
          labelFn={(t: EmailProviderEntity) =>
            [
              t.type,
              ,
              " - ",
              t?.uniqueId,
              t.apiKey?.substring(0, 10),
              "…",
            ].join(" ")
          }
        /> */}
      </PageSection>

      <PageSection title={t.wokspaces.publicSignup}>
        <p>{t.wokspaces.publicSignupHint}</p>
        <FormCheckbox
          onChange={(value: boolean) =>
            setFieldValue(
              NotificationConfigEntity.Fields.cascadeToSubWorkspaces,
              value
            )
          }
          value={values.cascadeToSubWorkspaces}
          label={t.wokspaces.disablePublicSignup}
          hint={t.wokspaces.disablePublicSignupHint}
        />
      </PageSection>
      <PageSection title={t.wokspaces.emailSendingConfig}>
        <p>{t.wokspaces.emailSendingConfigHint}</p>

        <FormCheckbox
          onChange={(value: boolean) =>
            setFieldValue(
              NotificationConfigEntity.Fields.cascadeToSubWorkspaces,
              value
            )
          }
          value={values.cascadeToSubWorkspaces}
          label={t.wokspaces.forceEmailConfigToSubWorkspaces}
          hint={t.wokspaces.forceEmailConfigToSubWorkspacesHint}
        />
        <MailTemplateConfiguration form={form} />
      </PageSection>
    </>
  );
};
