import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse, IResponseList } from "../../sdk/core/http-tools";
import { WorkspaceTypeEntity } from "../../sdk/modules/workspaces/WorkspaceTypeEntity";
import { MockWorkspaceType } from "./../database/workspace-type.db";
import { QueryToId } from "../database/memory-db";

export class WorkspaceTypeMockServer {
  @uriMatch("workspace-types")
  @method("get")
  async getWorkspaceTypes(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceTypeEntity>>> {
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
  ): Promise<IResponse<DeepPartial<WorkspaceTypeEntity>>> {
    return {
      data: MockWorkspaceType.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("workspace-type")
  @method("patch")
  async patchWorkspaceTypeByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceTypeEntity>>> {
    return {
      data: MockWorkspaceType.patchOne(ctx.body),
    };
  }

  @uriMatch("workspace-type")
  @method("delete")
  async deleteWorkspaceType(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceTypeEntity>>> {
    MockWorkspaceType.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("workspace-type")
  @method("post")
  async postWorkspaceType(
    ctx: Context
  ): Promise<IResponse<DeepPartial<WorkspaceTypeEntity>>> {
    const entity = MockWorkspaceType.create(
      ctx.body as Partial<WorkspaceTypeEntity>
    );

    return {
      data: entity,
    };
  }
}
