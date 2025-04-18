import { Context, DeepPartial, method, uriMatch } from "../../hooks/mock-tools";
import { IResponse } from "../../definitions/JSONStyle";
import { IResponseList } from "../../sdk/core/http-tools";
import { EmailProviderEntity } from "../../sdk/modules/abac/EmailProviderEntity";
import { mdb } from "../database/databases";
import { QueryToId } from "../database/memory-db";

export class EmailProviderMockServer {
  @uriMatch("email-providers")
  @method("get")
  async getEmailProviders(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<EmailProviderEntity>>> {
    return {
      data: {
        items: mdb.emailProvider.items(ctx),
        itemsPerPage: ctx.itemsPerPage,
        totalItems: mdb.emailProvider.total(),
      },
    };
  }

  @uriMatch("email-provider/:uniqueId")
  @method("get")
  async getEmailProviderByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailProviderEntity>>> {
    return {
      data: mdb.emailProvider.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("email-provider")
  @method("patch")
  async patchEmailProviderByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailProviderEntity>>> {
    return {
      data: mdb.emailProvider.patchOne(ctx.body),
    };
  }

  @uriMatch("email-provider")
  @method("post")
  async postRole(
    ctx: Context
  ): Promise<IResponse<DeepPartial<EmailProviderEntity>>> {
    const entity = mdb.emailProvider.create(
      ctx.body as Partial<EmailProviderEntity>
    );

    return {
      data: entity,
    };
  }

  @uriMatch("email-provider")
  @method("delete")
  async deleteRole(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<EmailProviderEntity>>> {
    mdb.emailProvider.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }
}
