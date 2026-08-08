import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

import { EmailProviderDto } from "../../../sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "../../../sdk/navigation/MessagingNavigation";
import { EmailProviderList } from "./EmailProviderList";

export const EmailProviderArchiveScreen = () => {
  const s = useS(strings);

  return (
    <>
      <CommonArchiveManager
        pageTitle={s.menuTitle}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailProviderNavigation.create());
        }}
      >
        <EmailProviderList />
      </CommonArchiveManager>
    </>
  );
};
