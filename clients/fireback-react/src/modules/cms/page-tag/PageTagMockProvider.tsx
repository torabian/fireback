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
export class PageTagMockProvider {
  @uriMatch("pagetags")
  @method("get")
  async getPageTags(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("PageTag", ctx);
  }
  @uriMatch("page-tags/:uniqueId")
  @method("get")
  async getPageTagByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("PageTag", ctx);
  }
}
