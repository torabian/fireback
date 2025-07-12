import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { columns } from "./InvoiceColumns";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
import { useGetInvoices } from "@/modules/fireback/sdk/modules/payment/useGetInvoices";
import { useDeleteInvoice } from "@/modules/fireback/sdk/modules/payment/useDeleteInvoice";
import { useS } from "@/modules/fireback/hooks/useS";
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
        deleteHook={useDeleteInvoice}
      ></CommonListManager>
    </>
  );
};
