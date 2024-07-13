import {
  Context,
  DeepPartial,
  emptyList,
  getItemUid,
  getJson,
  getJsonList,
  getJsonRaw,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse, IResponseList } from "../sdk/core/http-tools";
import { KeyboardShortcutEntity } from "../sdk/modules/keyboardActions/KeyboardShortcutEntity";
import { PublicJoinKeyEntity } from "../sdk/modules/workspaces/PublicJoinKeyEntity";
import { RoleEntity } from "../sdk/modules/workspaces/RoleEntity";
import { UserEntity } from "../sdk/modules/workspaces/UserEntity";
import { UserSessionDto } from "../sdk/modules/workspaces/UserSessionDto";
import { WorkspaceInviteEntity } from "../sdk/modules/workspaces/WorkspaceInviteEntity";

export class AbacModuleMockProvider {
  @uriMatch("passport/signin/email")
  @method("post")
  async postUserSignin(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return getJsonRaw("UserSessionDto", ctx);
  }

  @uriMatch("passport/authorizeOs")
  @method("post")
  async postOsAuthorize(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return this.postUserSignin(ctx);
  }

  @uriMatch("passport/signup/email")
  @method("post")
  async postUserSignup(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return this.postUserSignin(ctx);
  }

  @uriMatch("userRoleWorkspaces")
  @method("get")
  async getUserRoleWorkspaces(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return getJson("UserRoleWorkspaces", ctx);
  }

  @uriMatch("users")
  @method("get")
  async getUsers(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return emptyList;
  }
  @uriMatch("workspace-invites")
  @method("get")
  async getWorkspaceInvites(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return emptyList;
  }

  @uriMatch("cte-app-menus")
  @method("get")
  async getAppMenu(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return getJsonList("AppMenu", ctx);
  }

  @uriMatch("workspace/publicjoinkeys")
  @method("get")
  async getPublicJoinKeys(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<PublicJoinKeyEntity>>> {
    return emptyList;
  }

  @uriMatch("workspace/invites")
  @method("get")
  async getInvites(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceInviteEntity>>> {
    return emptyList;
  }

  @uriMatch("workspace/roles")
  @method("get")
  async getWorkspaceRoles(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<RoleEntity>>> {
    return emptyList;
  }

  @uriMatch("workspace-types")
  @method("get")
  async getWorkspaceTypes(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return getJson("WorkspaceType", ctx);
  }
  @uriMatch("public-workspace-types")
  @method("get")
  async getPublicWorkspaceTypes(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return getJson("WorkspaceType", ctx);
  }

  @uriMatch("workspaceType/:uniqueId")
  @method("get")
  async getWorkspaceType(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return getItemUid("WorkspaceType", ctx);
  }

  @uriMatch("email-senders")
  @method("get")
  async getEmailSenders(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return emptyList;
  }

  @uriMatch("emailProviders")
  @method("get")
  async getEmailProviders(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return emptyList;
  }

  @uriMatch("workspaces")
  @method("get")
  async getWorkspaces(ctx: Context): Promise<IResponseList<DeepPartial<any>>> {
    return emptyList;
  }

  @uriMatch("drive")
  @method("get")
  async getDrive(ctx: Context): Promise<IResponseList<DeepPartial<any>>> {
    return emptyList;
  }

  @uriMatch("cteWorkspaces")
  @method("get")
  async getcteWorkspaces(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return emptyList;
  }

  @uriMatch("roles")
  @method("get")
  async getroles(ctx: Context): Promise<IResponseList<DeepPartial<any>>> {
    return emptyList;
  }

  @uriMatch("keyboardShortcuts")
  @method("get")
  async getKeyboardShortcuts(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<KeyboardShortcutEntity>>> {
    return getJson("KeyboardShortcut", ctx);
  }
}
