import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { EmailSenderList } from "./EmailSenderList";
import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { EmailSenderDto } from "../../../sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "../../../sdk/navigation/MessagingNavigation";

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
