import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponse } from "../../definitions/JSONStyle";
import { IResponseList } from "../../sdk/core/http-tools";
import { PublicJoinKeyEntity } from "../../sdk/modules/abac/PublicJoinKeyEntity";
import { mdb } from "../database/databases";
import { QueryToId } from "../database/memory-db";

export class PublicJoinKeyMockServer {
  @uriMatch("public-join-keys")
  @method("get")
  async getPublicJoinKeys(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<PublicJoinKeyEntity>>> {
    return {
      data: {
        items: mdb.publicJoinKey.items(),
      },
    };
  }

  @uriMatch("public-join-key/:uniqueId")
  @method("get")
  async getPublicJoinKeyByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<PublicJoinKeyEntity>>> {
    return {
      data: mdb.publicJoinKey.getOne(ctx.paramValues[0]),
    };
  }

  @uriMatch("public-join-key")
  @method("patch")
  async patchPublicJoinKeyByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<PublicJoinKeyEntity>>> {
    return {
      data: mdb.publicJoinKey.patchOne(ctx.body),
    };
  }

  @uriMatch("public-join-key")
  @method("post")
  async postPublicJoinKey(
    ctx: Context
  ): Promise<IResponse<DeepPartial<PublicJoinKeyEntity>>> {
    const entity = mdb.publicJoinKey.create(
      ctx.body as Partial<PublicJoinKeyEntity>
    );

    return {
      data: entity,
    };
  }

  @uriMatch("public-join-key")
  @method("delete")
  async deletePublicJoinKey(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<PublicJoinKeyEntity>>> {
    mdb.publicJoinKey.deletes(QueryToId(ctx.body.query));

    return {
      data: {},
    };
  }
}
