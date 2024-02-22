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
} from "@/hooks/mock-tools";
import {
  IResponse,
  IResponseList,
  KeyboardShortcutEntity,
  PublicJoinKeyEntity,
  RoleEntity,
  UserEntity,
  UserSessionDto,
  WorkspaceInviteEntity,
} from "src/sdk/fireback";

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
  @uriMatch("workspaceInvites")
  @method("get")
  async getWorkspaceInvites(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<UserEntity>>> {
    return emptyList;
  }

  @uriMatch("cteAppMenus")
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

  @uriMatch("workspaceTypes")
  @method("get")
  async getWorkspaceTypes(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<any>>> {
    return getJson("WorkspaceType", ctx);
  }
  @uriMatch("publicWorkspaceTypes")
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

  @uriMatch("emailSenders")
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
