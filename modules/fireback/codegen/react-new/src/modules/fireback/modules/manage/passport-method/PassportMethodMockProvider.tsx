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
export class PassportMethodMockProvider {
  @uriMatch("passport-methods")
  @method("get")
  async getPassportMethods(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("PassportMethod", ctx);
  }
  @uriMatch("passport-method/:uniqueId")
  @method("get")
  async getPassportMethodByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("PassportMethod", ctx);
  }
}
