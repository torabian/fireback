import { type Context, type DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { type IResponse, type IResponseList } from "../../sdk/core/http-tools";
import { WorkspaceDto } from "../../sdk/abac/WorkspaceDto";

import { QueryToId } from "../database/memory-db";
import { mdb } from "../database/databases";

export class WorkspaceMockServer {
  @uriMatch("workspaces")
  @method("get")
  async getWorkspaces(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceDto>>> {
    return {
      data: {
        items: mdb.workspaces.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: mdb.workspaces.total(),
      },
    };
  }
  @uriMatch("cte-workspaces")
  @method("get")
  async getWorkspacesCte(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceDto>>> {
    return {
      data: {
        items: mdb.workspaces.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: mdb.workspaces.total(),
      },
    };
  }

  @uriMatch("workspace/:uniqueId")
  @method("get")
  async getWorkspaceByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceDto>>> {
    return {
      data: mdb.workspaces.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("workspace")
  @method("patch")
  async patchWorkspaceByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceDto>>> {
    return {
      data: mdb.workspaces.patchOne(ctx.body),
    };
  }

  @uriMatch("workspace")
  @method("delete")
  async deleteWorkspace(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceDto>>> {
    mdb.workspaces.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("workspace")
  @method("post")
  async postWorkspace(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceDto>>> {
    const entity = mdb.workspaces.create(ctx.body as Partial<WorkspaceDto>);

    return {
      data: entity,
    };
  }
}
