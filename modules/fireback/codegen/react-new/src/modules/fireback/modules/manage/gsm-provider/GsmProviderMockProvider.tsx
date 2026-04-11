import {
  type Context,
  type DeepPartial,
  getItemUid,
  getJson,
  method,
  uriMatch
} from "@/modules/fireback/hooks/mock-tools";
import { type IResponse } from "@/modules/fireback/sdk/core/http-tools";
export class GsmProviderMockProvider {
  @uriMatch("gsm-providers")
  @method("get")
  async getGsmProviders(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("GsmProvider", ctx);
  }
  @uriMatch("gsm-provider/:uniqueId")
  @method("get")
  async getGsmProviderByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("GsmProvider", ctx);
  }
}
