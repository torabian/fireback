import { type Context, type DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { type IResponse } from "../../definitions/JSONStyle";
import { type IResponseList } from "../../sdk/core/http-tools";
import { EmailSenderDto } from "../../sdk/messaging/EmailSenderDto";
import { mdb } from "../database/databases";
import { QueryToId } from "../database/memory-db";

export class EmailSenderMockServer {
  @uriMatch("email-senders")
  @method("get")
  async getEmailSenders(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<EmailSenderDto>>> {
    return {
      data: {
        items: mdb.emailSender.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: mdb.emailSender.total(),
      },
    };
  }

  @uriMatch("email-sender/:uniqueId")
  @method("get")
  async getEmailSenderByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderDto>>> {
    return {
      data: mdb.emailSender.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("email-sender")
  @method("patch")
  async patchEmailSenderByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderDto>>> {
    return {
      data: mdb.emailSender.patchOne(ctx.body),
    };
  }

  @uriMatch("email-sender")
  @method("post")
  async postRole(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderDto>>> {
    const entity = mdb.emailSender.create(
      ctx.body as Partial<EmailSenderDto>
    );

    return {
      data: entity,
    };
  }

  @uriMatch("email-sender")
  @method("delete")
  async deleteRole(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderDto>>> {
    mdb.emailSender.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }
}
