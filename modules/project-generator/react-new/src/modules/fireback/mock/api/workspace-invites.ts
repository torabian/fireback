import { type Context, type DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { type IResponse } from "../../definitions/JSONStyle";
import { type IResponseList } from "../../sdk/core/http-tools";
import { WorkspaceInviteDto } from "../../sdk/abac/WorkspaceInviteDto";
import { mdb } from "../database/databases";
import { QueryToId } from "../database/memory-db";

export class WorkspaceInviteMockServer {
  @uriMatch("workspace-invites")
  @method("get")
  async getWorkspaceInvites(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceInviteDto>>> {
    return {
      data: {
        items: mdb.workspaceInvite.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: mdb.workspaceInvite.total(),
      },
    };
  }

  @uriMatch("workspace-invite/:uniqueId")
  @method("get")
  async getWorkspaceInviteByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceInviteDto>>> {
    return {
      data: mdb.workspaceInvite.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("workspace-invite")
  @method("patch")
  async patchWorkspaceInviteByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceInviteDto>>> {
    return {
      data: mdb.workspaceInvite.patchOne(ctx.body),
    };
  }

  @uriMatch("workspace/invite")
  @method("post")
  async postWorkspaceInvite(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceInviteDto>>> {
    const entity = mdb.workspaceInvite.create(
      ctx.body as Partial<WorkspaceInviteDto>
    );

    return {
      data: entity,
    };
  }

  @uriMatch("workspace-invite")
  @method("delete")
  async deleteWorkspaceInvite(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceInviteDto>>> {
    mdb.workspaceInvite.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }
}
