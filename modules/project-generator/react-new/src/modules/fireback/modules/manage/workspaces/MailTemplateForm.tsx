import { ModalContext } from "../../../components/modal/Modal";
import { PageSection } from "../../../components/page-section/PageSection";
import { useT } from "../../../hooks/useT";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { type FormikProps } from "formik";
import { useContext } from "react";
import { WorkspaceDto } from "../../../sdk/abac/WorkspaceDto";
import { EmailProviderEditForm } from "../mail-providers/EmailProviderEditForm";

export const MailTemplateForm = ({
  form,
}: {
  form: FormikProps<Partial<WorkspaceDto>>;
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
