import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
import { useGetInvoices } from "@/modules/fireback/sdk/modules/payment/useGetInvoices";
import { usePostInvoiceRemove } from "@/modules/fireback/sdk/modules/payment/usePostInvoiceRemove";
import { columns } from "./InvoiceColumns";
import { strings } from "./strings/translations";
export const InvoiceList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGetInvoices}
        uniqueIdHrefHandler={(uniqueId: string) =>
          InvoiceEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostInvoiceRemove}
      ></CommonListManager>
    </>
  );
};
