import { type Context, type DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { type IResponseList } from "../../sdk/core/http-tools";
import { UserDto } from "../../sdk/abac/UserDto";
import { AppMenuEntities } from "../database/app-menu";

export class SidebarMockServer {
  @uriMatch("cte-app-menus")
  @method("get")
  async getAppMenu(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserDto>>> {
    return {
      data: {
        items: AppMenuEntities as any,
      },
    };
  }
}
