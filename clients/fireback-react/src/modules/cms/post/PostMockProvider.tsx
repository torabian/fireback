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
export class PostMockProvider {
  @uriMatch("posts")
  @method("get")
  async getPosts(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Post", ctx);
  }
  @uriMatch("post/:uniqueId")
  @method("get")
  async getPostByUniqueId(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Post", ctx);
  }
}
