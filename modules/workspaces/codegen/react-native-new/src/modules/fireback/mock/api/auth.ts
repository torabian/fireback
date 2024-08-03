import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse } from "../../definitions/JSONStyle";
import { UserSessionDto } from "../../sdk/modules/workspaces/UserSessionDto";

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
}
