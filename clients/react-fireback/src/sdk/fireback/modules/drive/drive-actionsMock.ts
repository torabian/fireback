// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: drive
 */

import * as workspaces from "../workspaces";

import * as drive from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class DriveActions {
  private _itemsPerPage?: number = undefined;
  private _startIndex?: number = undefined;
  private _sort?: number = undefined;
  private _query?: string = undefined;
  private _jsonQuery?: any = undefined;
  private _withPreloads?: string = undefined;
  private _uniqueId?: string = undefined;
  private _deep?: boolean = undefined;

  constructor(private apiFn: ExecApi) {
    this.apiFn = apiFn;
  }

  static isFileEntityEqual(a: drive.FileEntity, b: drive.FileEntity): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getFileEntityPrimaryKey(a: drive.FileEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): DriveActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): DriveActions {
    return new DriveActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): DriveActions {
    return new DriveActions(fn);
  }

  uniqueId(id: string): DriveActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): DriveActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): DriveActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): DriveActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): DriveActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): DriveActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): DriveActions {
    this._itemsPerPage = limit;
    return this;
  }

  get paramsAsString(): string {
    const q: any = {
      startIndex: this._startIndex,
      itemsPerPage: this._itemsPerPage,
      query: this._query,
      deep: this._deep,
      jsonQuery: JSON.stringify(this._jsonQuery),
      withPreloads: this._withPreloads,
      uniqueId: this._uniqueId,
      sort: this._sort,
    };

    const searchParams = new URLSearchParams();
    Object.keys(q).forEach((key) => {
      if (q[key]) {
        searchParams.append(key, q[key]);
      }
    });

    return searchParams.toString();
  }

  async getDrive(): Promise<IResponseList<drive.FileEntity>> {
    return this.apiFn("GET", `drive?${this.paramsAsString}`);
  }

  async getDriveByUniqueId(
    uniqueId: string
  ): Promise<IResponse<drive.FileEntity>> {
    return this.apiFn(
      "GET",
      `drive/:uniqueId?${this.paramsAsString}`.replace(":uniqueId", uniqueId)
    );
  }
}
