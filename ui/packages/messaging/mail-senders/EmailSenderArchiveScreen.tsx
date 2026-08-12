import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { EmailSenderList } from "./EmailSenderList";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { EmailSenderDto } from "@fireback/messaging/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";

export const EmailSenderArchiveScreen = () => {
  const s = useS(strings);

  return (
    <>
      <CommonArchiveManager
        pageTitle={s.menuTitle}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailSenderNavigation.create());
        }}
      >
        <EmailSenderList />
      </CommonArchiveManager>
    </>
  );
};
