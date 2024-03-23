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
export class DiscountCodeMockProvider {
  @uriMatch("discount-codes")
  @method("get")
  async getDiscountCodes(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("DiscountCode", ctx);
  }
  @uriMatch("discount-code/:uniqueId")
  @method("get")
  async getDiscountCodeByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("DiscountCode", ctx);
  }
}