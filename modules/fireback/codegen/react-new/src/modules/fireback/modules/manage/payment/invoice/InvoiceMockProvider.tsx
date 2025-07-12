import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse } from "@/modules/fireback/sdk/core/http-tools";
export class InvoiceMockProvider {
  @uriMatch("invoices")
  @method("get")
  async getInvoices(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Invoice", ctx);
  }
  @uriMatch("invoice/:uniqueId")
  @method("get")
  async getInvoiceByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Invoice", ctx);
  }
}
