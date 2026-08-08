import { useT } from "../../../hooks/useT";
import { EmailSenderList } from "./EmailSenderList";
import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
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
