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
export class WorkspaceConfigMockProvider {
  @uriMatch("workspace-configs")
  @method("get")
  async getWorkspaceConfigs(ctx: Context): Promise<IResponse<DeepPartial<any>>> {
    return getJson("WorkspaceConfig", ctx);
  }
  @uriMatch("workspace-config/:uniqueId")
  @method("get")
  async getWorkspaceConfigByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("WorkspaceConfig", ctx);
  }
}
