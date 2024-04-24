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
export class AppMenuMockProvider {
  @uriMatch("app-menus")
  @method("get")
  async getAppMenus(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("AppMenu", ctx);
  }
  @uriMatch("app-menu/:uniqueId")
  @method("get")
  async getAppMenuByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("AppMenu", ctx);
  }
}