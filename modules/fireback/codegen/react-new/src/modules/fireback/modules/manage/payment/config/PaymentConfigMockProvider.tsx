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
export class PaymentConfigMockProvider {
  @uriMatch("payment-configs")
  @method("get")
  async getPaymentConfigs(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("PaymentConfig", ctx);
  }
  @uriMatch("payment-config/:uniqueId")
  @method("get")
  async getPaymentConfigByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("PaymentConfig", ctx);
  }
}
