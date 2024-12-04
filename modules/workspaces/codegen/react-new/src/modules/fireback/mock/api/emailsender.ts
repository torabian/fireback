import { Context, DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { IResponse } from "../../definitions/JSONStyle";
import { IResponseList } from "../../sdk/core/http-tools";
import { EmailSenderEntity } from "../../sdk/modules/workspaces/EmailSenderEntity";
import { mdb } from "../database/databases";
import { QueryToId } from "../database/memory-db";

export class EmailSenderMockServer {
  @uriMatch("email-senders")
  @method("get")
  async getEmailSenders(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<EmailSenderEntity>>> {
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
  ): Promise<IResponse<DeepPartial<EmailSenderEntity>>> {
    return {
      data: mdb.emailSender.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("email-sender")
  @method("patch")
  async patchEmailSenderByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderEntity>>> {
    return {
      data: mdb.emailSender.patchOne(ctx.body),
    };
  }

  @uriMatch("email-sender")
  @method("post")
  async postRole(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderEntity>>> {
    const entity = mdb.emailSender.create(
      ctx.body as Partial<EmailSenderEntity>
    );

    return {
      data: entity,
    };
  }

  @uriMatch("email-sender")
  @method("delete")
  async deleteRole(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailSenderEntity>>> {
    mdb.emailSender.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }
}
