import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "@/fireback/hooks/mock-tools";
import { IResponse } from "@/sdk/fireback/core/http-tools";
export class OrderMockProvider {
  @uriMatch("orders")
  @method("get")
  async getOrders(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Order", ctx);
  }
  @uriMatch("order/:uniqueId")
  @method("get")
  async getOrderByUniqueId(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Order", ctx);
  }
}
