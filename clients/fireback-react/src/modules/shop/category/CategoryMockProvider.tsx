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
export class CategoryMockProvider {
  @uriMatch("categories")
  @method("get")
  async getCategorys(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Category", ctx);
  }
  @uriMatch("<no value>/:uniqueId")
  @method("get")
  async getCategoryByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Category", ctx);
  }
}