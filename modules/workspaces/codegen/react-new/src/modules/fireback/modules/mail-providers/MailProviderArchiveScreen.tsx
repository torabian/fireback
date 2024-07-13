import { useT } from "@/modules/fireback/hooks/useT";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { EmailProviderList } from "./MailProviderList";
import { EmailProviderEntity } from "../../sdk/modules/workspaces/EmailProviderEntity";

export const EmailProviderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailProviders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailProviderEntity.Navigation.create(locale));
        }}
      >
        <EmailProviderList />
      </CommonArchiveManager>
    </>
  );
};
