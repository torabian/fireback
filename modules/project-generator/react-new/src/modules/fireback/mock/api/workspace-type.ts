import { type Context, type DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { type IResponse, type IResponseList } from "../../sdk/core/http-tools";
import { WorkspaceTypeDto } from "../../sdk/abac/WorkspaceTypeDto";
import { MockWorkspaceType } from "./../database/workspace-type.db";
import { QueryToId } from "../database/memory-db";

export class WorkspaceTypeMockServer {
  @uriMatch("workspace-types")
  @method("get")
  async getWorkspaceTypes(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceTypeDto>>> {
    return {
      data: {
        items: MockWorkspaceType.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: MockWorkspaceType.total(),
      },
    };
  }

  @uriMatch("workspace-type/:uniqueId")
  @method("get")
  async getWorkspaceTypeByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceTypeDto>>> {
    return {
      data: MockWorkspaceType.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("workspace-type")
  @method("patch")
  async patchWorkspaceTypeByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceTypeDto>>> {
    return {
      data: MockWorkspaceType.patchOne(ctx.body),
    };
  }

  @uriMatch("workspace-type")
  @method("delete")
  async deleteWorkspaceType(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceTypeDto>>> {
    MockWorkspaceType.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("workspace-type")
  @method("post")
  async postWorkspaceType(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceTypeDto>>> {
    const entity = MockWorkspaceType.create(
      ctx.body as Partial<WorkspaceTypeDto>
    );

    return {
      data: entity,
    };
  }
}
