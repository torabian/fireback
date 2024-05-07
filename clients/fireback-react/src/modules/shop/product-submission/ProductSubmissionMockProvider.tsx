import {
  Context,
  DeepPartial,
  emptyList,
  getJson,
  method,
  uriMatch,
  getItemUid,
} from "@/fireback/hooks/mock-tools";
import { IResponse } from "@/sdk/fireback/core/http-tools";
export class ProductSubmissionMockProvider {
  @uriMatch("productsubmissions")
  @method("get")
  async getProductSubmissions(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getJson("ProductSubmission", ctx);
  }
  @uriMatch("<no value>/:uniqueId")
  @method("get")
  async getProductSubmissionByUniqueId(
    ctx: Context
  ): Promise<IResponse<DeepPartial<any>>> {
    return getItemUid("ProductSubmission", ctx);
  }
}
