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
export class TagMockProvider {
  @uriMatch("tags")
  @method("get")
  async getTags(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Tag", ctx);
  }
  @uriMatch("tag/:uniqueId")
  @method("get")
  async getTagByUniqueId(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Tag", ctx);
  }
}
