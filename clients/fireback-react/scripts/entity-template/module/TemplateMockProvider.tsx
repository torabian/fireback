import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "@/hooks/mock-tools";
import { IResponse } from "src/sdk/fireback";

export class TemplateMockProvider {
  @uriMatch("templates")
  @method("get")
  async getTemplates(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("Template", ctx);
  }
  @uriMatch("template/:uniqueId")
  @method("get")
  async getTemplateByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("Template", ctx);
  }
}
