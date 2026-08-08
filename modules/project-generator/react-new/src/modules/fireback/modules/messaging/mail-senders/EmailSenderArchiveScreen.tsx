import { useT } from "../../../../fireback-ui/hooks/useT";
import { EmailSenderList } from "./EmailSenderList";
import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../../sdk/navigation/MessagingNavigation";

export const EmailSenderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailSenders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailSenderNavigation.create());
        }}
      >
        <EmailSenderList />
      </CommonArchiveManager>
    </>
  );
};
