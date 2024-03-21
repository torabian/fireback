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
export class PostTagMockProvider {
  @uriMatch("posttags")
  @method("get")
  async getPostTags(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("PostTag", ctx);
  }
  @uriMatch("post-tag/:uniqueId")
  @method("get")
  async getPostTagByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("PostTag", ctx);
  }
}
