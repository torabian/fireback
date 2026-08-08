import {
  type Context,
  type DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "../../../hooks/mock-tools";
import { type IResponse } from "../../../sdk/core/http-tools";
import { WorkspaceConfigDto } from "../../../sdk/abac/WorkspaceConfigDto";
export class WorkspaceConfigMockProvider {
  @uriMatch("workspace-config")
  @method("get")
  async getWorkspaceConfig(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceConfigDto>>> {
    return {
      data: {
        enableOtp: true,
        forcePasswordOnPhone: true,
      },
    };
  }

  @uriMatch("workspace-wconfig/distiwnct")
  @method("patch")
  async setWorkspaceConfig(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceConfigDto>>> {
    return {
      data: ctx.body,
    };
  }
}
