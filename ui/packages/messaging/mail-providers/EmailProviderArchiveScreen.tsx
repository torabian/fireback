import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

import { EmailProviderDto } from "@fireback/messaging/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
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
