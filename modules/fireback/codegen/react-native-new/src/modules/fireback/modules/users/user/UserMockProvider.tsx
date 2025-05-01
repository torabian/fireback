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
export class UserMockProvider {
  @uriMatch("users")
  @method("get")
  async getUsers(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("User", ctx);
  }
  @uriMatch("user/:uniqueId")
  @method("get")
  async getUserByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("User", ctx);
  }
}