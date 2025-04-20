import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "{{ .FirebackUiDir }}/hooks/mock-tools";
import { IResponse } from "{{ .SdkDir }}/core/http-tools";

export class {{ .Template }}MockProvider {
  @uriMatch("{{ .templatesDashed }}")
  @method("get")
  async get{{ .Template }}s(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("{{ .Template }}", ctx);
  }
  @uriMatch("{{ .templateDashed }}/:uniqueId")
  @method("get")
  async get{{ .Template }}ByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("{{ .Template }}", ctx);
  }
}
