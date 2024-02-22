// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: workspaces
 */

import * as workspaces from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class PassportMethodActions {
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

  static isPassportMethodEntityEqual(
    a: workspaces.PassportMethodEntity,
    b: workspaces.PassportMethodEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getPassportMethodEntityPrimaryKey(
    a: workspaces.PassportMethodEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): PassportMethodActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): PassportMethodActions {
    return new PassportMethodActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): PassportMethodActions {
    return new PassportMethodActions(fn);
  }

  uniqueId(id: string): PassportMethodActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): PassportMethodActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): PassportMethodActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): PassportMethodActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): PassportMethodActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): PassportMethodActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): PassportMethodActions {
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

  async getPassportMethods(): Promise<
    IResponseList<workspaces.PassportMethodEntity>
  > {
    return this.apiFn(
      "GET",
      `passportMethods?action=PassportMethodActionQuery&${this.paramsAsString}`
    );
  }

  async getPassportMethodsExport(): Promise<
    IResponseList<workspaces.PassportMethodEntity>
  > {
    return this.apiFn(
      "GET",
      `passportMethods/export?action=PassportMethodActionExport&${this.paramsAsString}`
    );
  }

  async getPassportMethodByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.PassportMethodEntity>> {
    return this.apiFn(
      "GET",
      `passportMethod/:uniqueId?action=PassportMethodActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postPassportMethod(
    entity: workspaces.PassportMethodEntity
  ): Promise<IResponse<workspaces.PassportMethodEntity>> {
    return this.apiFn(
      "POST",
      `passportMethod?action=PassportMethodActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPassportMethod(
    entity: workspaces.PassportMethodEntity
  ): Promise<IResponse<workspaces.PassportMethodEntity>> {
    return this.apiFn(
      "PATCH",
      `passportMethod?action=PassportMethodActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPassportMethods(
    entity: core.BulkRecordRequest<workspaces.PassportMethodEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[workspaces.PassportMethodEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `passportMethods?action=PassportMethodActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deletePassportMethod(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `passportMethod?action=PassportMethodActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
