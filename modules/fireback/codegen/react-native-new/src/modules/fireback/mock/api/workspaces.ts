import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse, IResponseList } from "../../sdk/core/http-tools";
import { WorkspaceEntity } from "../../sdk/modules/abac/WorkspaceEntity";

import { QueryToId } from "../database/memory-db";
import { mdb } from "../database/databases";

export class WorkspaceMockServer {
  @uriMatch("workspaces")
  @method("get")
  async getWorkspaces(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceEntity>>> {
    return {
      data: {
        items: mdb.fireback.items(),
      },
    };
  }
  @uriMatch("cte-workspaces")
  @method("get")
  async getWorkspacesCte(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceEntity>>> {
    return {
      data: {
        items: mdb.fireback.items(),
      },
    };
  }

  @uriMatch("workspace/:uniqueId")
  @method("get")
  async getWorkspaceByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceEntity>>> {
    return {
      data: mdb.fireback.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("workspace")
  @method("patch")
  async patchWorkspaceByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceEntity>>> {
    return {
      data: mdb.fireback.patchOne(ctx.body),
    };
  }

  @uriMatch("workspace")
  @method("delete")
  async deleteWorkspace(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceEntity>>> {
    mdb.fireback.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("workspace")
  @method("post")
  async postWorkspace(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceEntity>>> {
    const entity = mdb.fireback.create(ctx.body as Partial<WorkspaceEntity>);

    return {
      data: entity,
    };
  }
}
