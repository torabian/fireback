import {
  Context,
  DeepPartial,
  method,
  uriMatch,
} from "@/modules/fireback/hooks/mock-tools";
import { IResponseList } from "../../sdk/core/http-tools";
import { FileEntity } from "../../sdk/modules/abac/FileEntity";
import { MockFiles } from "../database/file.db";

export class DriveMockServer {
  @uriMatch("files")
  @method("get")
  async getFiles(
    ctx: Context
  ): Promise<IResponseList<DeepPartial<FileEntity>>> {
    return {
      data: {
        items: MockFiles.items(),
      },
    };
  }
}
