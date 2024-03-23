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
export class PostCategoryMockProvider {
  @uriMatch("postcategories")
  @method("get")
  async getPostCategorys(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("PostCategory", ctx);
  }
  @uriMatch("post-categories/:uniqueId")
  @method("get")
  async getPostCategoryByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("PostCategory", ctx);
  }
}
