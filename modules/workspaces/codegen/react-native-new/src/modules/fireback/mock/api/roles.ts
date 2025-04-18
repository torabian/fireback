import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse, IResponseList } from "../../sdk/core/http-tools";
import { RoleEntity } from "../../sdk/modules/abac/RoleEntity";
import { MockRoles } from "./../database/role.db";
import { QueryToId } from "../database/memory-db";

export class RoleMockServer {
  @uriMatch("roles")
  @method("get")
  async getRoles(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<RoleEntity>>> {
    return {
      data: {
        items: MockRoles.items(),
      },
    };
  }

  @uriMatch("role/:uniqueId")
  @method("get")
  async getRoleByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<RoleEntity>>> {
    return {
      data: MockRoles.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("role")
  @method("patch")
  async patchRoleByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<RoleEntity>>> {
    return {
      data: MockRoles.patchOne(ctx.body),
    };
  }

  @uriMatch("role")
  @method("delete")
  async deleteRole(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<RoleEntity>>> {
    MockRoles.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("role")
  @method("post")
  async postRole(ctx: Context): Promise<IResponse<DeepPartial<RoleEntity>>> {
    const entity = MockRoles.create(ctx.body as Partial<RoleEntity>);

    return {
      data: entity,
    };
  }
}
