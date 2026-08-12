import { ModalContext } from "@fireback/ui-core/components/modal/Modal";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { type FormikProps } from "formik";
import { useContext } from "react";
import { WorkspaceDto } from "@fireback/manage/sdk/abac/WorkspaceDto";
import { EmailProviderEditForm } from "@fireback/messaging/mail-providers/EmailProviderEditForm";

export const MailTemplateForm = ({
  form,
}: {
  form: FormikProps<Partial<WorkspaceDto>>;
}) => {
  const { values, setFieldValue, errors } = form;
  const s = useS(strings);

  return (
    <>
      <PageSection title={s.mailServerConfiguration}>
        <EmailProviderEditForm form={form as any} />
      </PageSection>
      <PageSection title={s.emailSendingConfiguration}>
        <p>{s.emailSendingConfigurationHint}</p>
        <MailTemplateConfiguration />
      </PageSection>
    </>
  );
};

function MailTemplateConfiguration() {
  const useModal = useContext(ModalContext);

  return <div>form here</div>;
}
