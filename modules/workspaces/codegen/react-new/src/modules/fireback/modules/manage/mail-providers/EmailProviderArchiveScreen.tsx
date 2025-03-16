import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { useT } from "../../../hooks/useT";

import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/workspaces/EmailProviderEntity";
import { EmailProviderList } from "./EmailProviderList";

export const EmailProviderArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.emailProviders}
        newEntityHandler={({ locale, router }) => {
          router.push(EmailProviderEntity.Navigation.create());
        }}
      >
        <EmailProviderList />
      </CommonArchiveManager>
    </>
  );
};
