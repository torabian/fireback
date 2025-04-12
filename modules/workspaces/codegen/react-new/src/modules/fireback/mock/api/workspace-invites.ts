import { Context, DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { IResponse } from "../../definitions/JSONStyle";
import { IResponseList } from "../../sdk/core/http-tools";
import { WorkspaceInviteEntity } from "../../sdk/modules/abac/WorkspaceInviteEntity";
import { mdb } from "../database/databases";
import { QueryToId } from "../database/memory-db";

export class WorkspaceInviteMockServer {
  @uriMatch("workspace-invites")
  @method("get")
  async getWorkspaceInvites(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceInviteEntity>>> {
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
  ): Promise<IResponse<DeepPartial<WorkspaceInviteEntity>>> {
    return {
      data: mdb.workspaceInvite.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("workspace-invite")
  @method("patch")
  async patchWorkspaceInviteByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceInviteEntity>>> {
    return {
      data: mdb.workspaceInvite.patchOne(ctx.body),
    };
  }

  @uriMatch("workspace/invite")
  @method("post")
  async postWorkspaceInvite(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceInviteEntity>>> {
    const entity = mdb.workspaceInvite.create(
      ctx.body as Partial<WorkspaceInviteEntity>
    );

    return {
      data: entity,
    };
  }

  @uriMatch("workspace-invite")
  @method("delete")
  async deleteWorkspaceInvite(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceInviteEntity>>> {
    mdb.workspaceInvite.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }
}
