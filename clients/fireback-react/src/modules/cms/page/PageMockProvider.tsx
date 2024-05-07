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
export class PageMockProvider {
  @uriMatch("pages")
  @method("get")
  async getPages(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Page", ctx);
  }
  @uriMatch("page/:uniqueId")
  @method("get")
  async getPageByUniqueId(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Page", ctx);
  }
}
