import { IResponse, IResponseList } from "../../definitions/JSONStyle";
import {
  Context,
  DeepPartial,
  emptyList,
  method,
  uriMatch,
} from "../../hooks/mock-tools";
import {
  CheckClassicPassportActionResDto,
  CheckPassportMethodsActionResDto,
  ClassicSignupActionResDto,
  ConfirmClassicPassportTotpActionResDto,
} from "../../sdk/modules/abac/AbacActionsDto";
import { UserSessionDto } from "../../sdk/modules/abac/UserSessionDto";
import { WorkspaceInviteEntity } from "../../sdk/modules/abac/WorkspaceInviteEntity";

import { WorkspaceTypeEntity } from "../../sdk/modules/abac/WorkspaceTypeEntity";

const commonSession: IResponse<DeepPartial<UserSessionDto>> = {
  data: {
    user: {
      firstName: "Ali",
      lastName: "Torabi",
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

  @uriMatch("users/invitations")
  @method("get")
  async getUserInvites(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceInviteEntity>>> {
    return emptyList;
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
        phone: true,
        recaptcha2ClientKey: "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI",
      },
    };
  }

  @uriMatch("workspace/passport/check")
  @method("post")
  async postWorkspacePassportCheck(
    ctx: Context
  ): Promise<IResponse<DeepPartial<CheckClassicPassportActionResDto>>> {
    const isEmail = ctx?.body?.value.includes("@");

    if (isEmail) {
      return {
        data: {
          next: ["otp", "create-with-password"],
          flags: ["enable-totp", "force-totp"],
          otpInfo: null,
        },
      };
    } else {
      return {
        data: {
          next: ["otp"],
          flags: ["enable-totp", "force-totp"],
          otpInfo: null,
        },
      };
    }
  }

  @uriMatch("passports/signup/classic")
  @method("post")
  async postPassportSignupClassic(
    ctx: Context
  ): Promise<IResponse<DeepPartial<ClassicSignupActionResDto>>> {
    const isEmail = ctx?.body?.value.includes("@");

    return {
      data: {
        session: null,
        sessionId: null,
        totpUrl:
          "otpauth://totp/Fireback:ali@ali.com?algorithm=SHA1\u0026digits=6\u0026issuer=Fireback\u0026period=30\u0026secret=R2AQ4NPS7FKECL3ZVTF3JMTLBYGDAAVU",
        continueToTotp: true,
        forcedTotp: true,
      },
    };
  }

  @uriMatch("passport/totp/confirm")
  @method("post")
  async postConfirm(
    ctx: Context
  ): Promise<IResponse<DeepPartial<ConfirmClassicPassportTotpActionResDto>>> {
    return {
      data: {
        session: commonSession.data,
      },
    };
  }

  @uriMatch("workspace/passport/otp")
  @method("post")
  async postOtp(
    ctx: Context
  ): Promise<IResponse<DeepPartial<ConfirmClassicPassportTotpActionResDto>>> {
    return {
      data: {
        session: commonSession.data,
      },
    };
  }

  @uriMatch("workspace/public/types")
  @method("get")
  async getWorkspaceTypes(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<WorkspaceTypeEntity>>> {
    return {
      data: {
        items: [
          {
            description: null,
            slug: "customer",
            title: "customer",
            uniqueId: "nG012z7VNyYKMJPqWjV04",
          },
        ],
        itemsPerPage: 20,
        startIndex: 0,
        totalItems: 2,
      },
    };
  }
}
