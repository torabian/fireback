import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse } from "@/modules/fireback/sdk/core/http-tools";
export class RegionalContentMockProvider {
  @uriMatch("regional-contents")
  @method("get")
  async getRegionalContents(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("RegionalContent", ctx);
  }
  @uriMatch("regional-content/:uniqueId")
  @method("get")
  async getRegionalContentByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("RegionalContent", ctx);
  }
}
