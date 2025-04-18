import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse } from "../../sdk/core/http-tools";
import { WorkspaceConfigEntity } from "../../sdk/modules/abac/WorkspaceConfigEntity";

export class WorkspaceConfigMockServer {
  @uriMatch("workspace/config")
  @method("get")
  async getWorkspaceConfig(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceConfigEntity>>> {
    return {
      data: {},
    };
  }

  @uriMatch("workspace")
  @method("patch")
  async patchWorkspace(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceConfigEntity>>> {
    return {
      data: {},
    };
  }
}
