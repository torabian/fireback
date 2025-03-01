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
import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
export class WorkspaceConfigMockProvider {
  @uriMatch("workspace-config")
  @method("get")
  async getWorkspaceConfig(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceConfigEntity>>> {
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
  ): Promise<IResponse<DeepPartial<WorkspaceConfigEntity>>> {
    return {
      data: ctx.body,
    };
  }
}
