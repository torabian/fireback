import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse, IResponseList } from "../../sdk/core/http-tools";
import { UserEntity } from "../../sdk/modules/workspaces/UserEntity";
import { MockUsers } from "./../database/user.db";
import { QueryToId } from "../database/memory-db";

export class UserMockServer {
  @uriMatch("users")
  @method("get")
  async getUsers(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return {
      data: {
        items: MockUsers.items(),
      },
    };
  }

  @uriMatch("user")
  @method("delete")
  async deleteUser(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    MockUsers.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("user/:uniqueId")
  @method("get")
  async getUserByUniqueId(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return {
      data: MockUsers.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("user")
  @method("patch")
  async patchUserByUniqueId(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return {
      data: MockUsers.patchOne(ctx.body),
    };
  }

  @uriMatch("user")
  @method("post")
  async postUser(ctx: Context): Promise<IResponse<DeepPartial<UserEntity>>> {
    const entity = MockUsers.create(ctx.body as Partial<UserEntity>);

    return {
      data: entity,
    };
  }
}
