import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "@/hooks/mock-tools";
import { IResponse } from "@/sdk/fireback/core/http-tools";
export class ProductMockProvider {
  @uriMatch("products")
  @method("get")
  async getProducts(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Product", ctx);
  }
  @uriMatch("product/:uniqueId")
  @method("get")
  async getProductByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Product", ctx);
  }
}
