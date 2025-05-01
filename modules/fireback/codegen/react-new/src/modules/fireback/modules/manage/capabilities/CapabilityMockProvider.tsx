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
export class CapabilityMockProvider {
  @uriMatch("capabilities")
  @method("get")
  async getCapabilitys(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Capability", ctx);
  }
  @uriMatch("capability/:uniqueId")
  @method("get")
  async getCapabilityByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Capability", ctx);
  }
}
