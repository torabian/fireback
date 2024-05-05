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
export class ShoppingCartMockProvider {
  @uriMatch("shopping-carts")
  @method("get")
  async getShoppingCarts(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("ShoppingCart", ctx);
  }
  @uriMatch("shopping-cart/:uniqueId")
  @method("get")
  async getShoppingCartByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("ShoppingCart", ctx);
  }
}
