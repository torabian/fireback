import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailSenderEntity } from "@/modules/fireback/sdk/modules/abac/EmailSenderEntity";
import { useGetEmailSenders } from "@/modules/fireback/sdk/modules/abac/useGetEmailSenders";
import { usePostEmailSenderRemove } from "@/modules/fireback/sdk/modules/abac/usePostEmailSenderRemove";
import { columns } from "./EmailSenderColumns";

export const EmailSenderList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetEmailSenders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          EmailSenderEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostEmailSenderRemove}
      ></CommonListManager>
    </>
  );
};
