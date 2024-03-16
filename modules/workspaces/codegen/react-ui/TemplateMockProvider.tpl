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

export class {{ .Template }}MockProvider {
  @uriMatch("{{ .templates }}")
  @method("get")
  async get{{ .Template }}s(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("{{ .Template }}", ctx);
  }
  @uriMatch("{{ .template }}/:uniqueId")
  @method("get")
  async get{{ .Template }}ByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("{{ .Template }}", ctx);
  }
}
