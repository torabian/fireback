import { Context, DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { IResponse } from "../../definitions/JSONStyle";
import { UserSessionDto } from "../../sdk/modules/workspaces/UserSessionDto";
import {
  CheckClassicPassportActionResDto,
  CheckPassportMethodsActionResDto,
} from "../../sdk/modules/workspaces/WorkspacesActionsDto";

const commonSession: IResponse<DeepPartial<UserSessionDto>> = {
  data: {
    user: {
      person: {
        firstName: "Ali",
        lastName: "Torabi",
      },
    },
    exchangeKey: "key1",
    token: "token",
  },
};

export class AuthMockServer {
  @uriMatch("passport/authorizeOs")
  @method("post")
  async passportAuthroizeOs(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return commonSession;
  }

  @uriMatch("passports/signin/classic")
  @method("post")
  async postSigninClassic(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return commonSession;
  }

  @uriMatch("passport/request-reset-mail-password")
  @method("post")
  async postRequestResetMail(
    ctx: Context
  ): Promise<IResponse<DeepPartial<UserSessionDto>>> {
    return commonSession;
  }

  @uriMatch("passports/available-methods")
  @method("get")
  async getAvailableMethods(
    ctx: Context
  ): Promise<IResponse<DeepPartial<CheckPassportMethodsActionResDto>>> {
    return {
      data: {
        email: true,
        enabledRecaptcha2: false,
        google: null,
        phone: false,
        recaptcha2ClientKey: undefined,
      },
    };
  }

  @uriMatch("workspace/passport/check")
  @method("post")
  async postWorkspacePassportCheck(
    ctx: Context
  ): Promise<IResponse<DeepPartial<CheckClassicPassportActionResDto>>> {
    return { data: { continueWithPassword: true } };
  }
}
