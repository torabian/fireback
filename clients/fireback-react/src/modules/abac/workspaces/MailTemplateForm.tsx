import { ModalContext } from "@/components/modal/Modal";
import { PageSection } from "@/components/page-section/PageSection";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { FormikProps } from "formik";
import { useContext } from "react";
import { EmailProviderEditForm } from "../mail-providers/MailProviderEditForm";
import { WorkspaceEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceEntity";

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
