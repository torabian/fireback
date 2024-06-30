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
export class BrandMockProvider {
  @uriMatch("brands")
  @method("get")
  async getBrands(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Brand", ctx);
  }
  @uriMatch("<no value>/:uniqueId")
  @method("get")
  async getBrandByUniqueId(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Brand", ctx);
  }
}
