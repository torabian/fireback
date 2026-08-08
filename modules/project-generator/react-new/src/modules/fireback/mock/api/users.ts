import { type Context, type DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { type IResponse, type IResponseList } from "../../sdk/core/http-tools";
import { UserDto } from "../../sdk/abac/UserDto";
import { MockUsers } from "./../database/user.db";
import { QueryToId } from "../database/memory-db";

export class UserMockServer {
  @uriMatch("users")
  @method("get")
  async getUsers(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserDto>>> {
    return {
      data: {
        items: MockUsers.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: MockUsers.total(),
      },
    };
  }

  @uriMatch("user")
  @method("delete")
  async deleteUser(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserDto>>> {
    MockUsers.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }

  @uriMatch("user/:uniqueId")
  @method("get")
  async getUserByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserDto>>> {
    return {
      data: MockUsers.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("user")
  @method("patch")
  async patchUserByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserDto>>> {
    return {
      data: MockUsers.patchOne(ctx.body),
    };
  }

  @uriMatch("user")
  @method("post")
  async postUser(ctx: Context): Promise<IResponse<DeepPartial<UserDto>>> {
    const entity = MockUsers.create(ctx.body as Partial<UserDto>);

    return {
      data: entity,
    };
  }
}
