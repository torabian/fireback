import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { InvoiceList } from "./InvoiceList";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
export const InvoiceArchiveScreen = () => {
  const s = useS(strings);
  return (
    <CommonArchiveManager
      pageTitle={s.invoices.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(InvoiceEntity.Navigation.create());
      }}
    >
      <InvoiceList />
    </CommonArchiveManager>
  );
};
