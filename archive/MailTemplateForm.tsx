import { ModalContext } from "../modules/workspaces/codegen/react-new/src/modules/fireback/components/modal/Modal";
import { PageSection } from "../modules/workspaces/codegen/react-new/src/modules/fireback/components/page-section/PageSection";
import { useT } from "../modules/workspaces/codegen/react-new/src/modules/fireback/hooks/useT";
import { RemoteQueryContext } from "../modules/workspaces/codegen/react-new/src/modules/fireback/sdk/core/react-tools";
import { FormikProps } from "formik";
import { useContext } from "react";
import { WorkspaceEntity } from "../modules/workspaces/codegen/react-new/src/modules/fireback/sdk/modules/abac/WorkspaceEntity";
import { EmailProviderEditForm } from "../modules/workspaces/codegen/react-new/src/modules/fireback/modules/manage/mail-providers/EmailProviderEditForm";

export const MailTemplateForm = ({
  form,
}: {
  form: FormikProps<Partial<WorkspaceEntity>>;
}) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();

  return (
    <>
      <PageSection title={t.wokspaces.mailServerConfiguration}>
        <EmailProviderEditForm form={form as any} />
      </PageSection>
      <PageSection title={t.wokspaces.emailSendingConfiguration}>
        <p>{t.wokspaces.emailSendingConfigurationHint}</p>
        <MailTemplateConfiguration />
      </PageSection>
    </>
  );
};

function MailTemplateConfiguration() {
  const { options } = useContext(RemoteQueryContext);
  const t = useT();
  const useModal = useContext(ModalContext);

  return <div>form here</div>;
}
