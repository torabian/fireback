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
export class PageCategoryMockProvider {
  @uriMatch("pagecategories")
  @method("get")
  async getPageCategorys(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("PageCategory", ctx);
  }
  @uriMatch("<no value>/:uniqueId")
  @method("get")
  async getPageCategoryByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("PageCategory", ctx);
  }
}