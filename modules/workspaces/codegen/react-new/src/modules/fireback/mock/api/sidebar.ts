import { Context, DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { IResponseList } from "../../sdk/core/http-tools";
import { UserEntity } from "../../sdk/modules/workspaces/UserEntity";
import { AppMenuEntities } from "../database/app-menu";

export class SidebarMockServer {
  @uriMatch("cte-app-menus")
  @method("get")
  async getAppMenu(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return {
      data: {
        items: AppMenuEntities as any,
      },
    };
  }
}
