import { useT } from "@/modules/fireback/hooks/useT";
import { useDeleteEmailProvider } from "../../sdk/modules/workspaces/useDeleteEmailProvider";

import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGetEmailProviders } from "../../sdk/modules/workspaces/useGetEmailProviders";
import { columns } from "./MailProviderColumns";
import { EmailProviderEntity } from "../../sdk/modules/workspaces/EmailProviderEntity";

export const EmailProviderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetEmailProviders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailProviderEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteEmailProvider}
      ></CommonListManager>
    </>
  );
};
