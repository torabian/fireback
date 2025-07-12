import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: "uniqueId",
    width: 200,
  },
  {
    name: InvoiceEntity.Fields.title,
    title: t.invoices.title,
    width: 100,
  },
  {
    name: InvoiceEntity.Fields.amount,
    title: t.invoices.amount,
    width: 100,
    getCellValue: (entity: InvoiceEntity) => entity.amount?.formatted,
  },
  {
    name: InvoiceEntity.Fields.finalStatus,
    title: t.invoices.finalStatus,
    width: 100,
  },
];
