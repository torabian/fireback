import { ModalContext } from "../../fireback-ui/components/modal/Modal";
import { PageSection } from "../../fireback-ui/components/page-section/PageSection";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { type FormikProps } from "formik";
import { useContext } from "react";
import { WorkspaceDto } from "../../sdk/abac/WorkspaceDto";
import { EmailProviderEditForm } from "../mail-providers/EmailProviderEditForm";

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
  const { options } = useContext(RemoteQueryContext);
  const useModal = useContext(ModalContext);

  return <div>form here</div>;
}
