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

export class PublicJoinKeyActions {
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

  static isPublicJoinKeyEntityEqual(
    a: workspaces.PublicJoinKeyEntity,
    b: workspaces.PublicJoinKeyEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getPublicJoinKeyEntityPrimaryKey(
    a: workspaces.PublicJoinKeyEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): PublicJoinKeyActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): PublicJoinKeyActions {
    return new PublicJoinKeyActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): PublicJoinKeyActions {
    return new PublicJoinKeyActions(fn);
  }

  uniqueId(id: string): PublicJoinKeyActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): PublicJoinKeyActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): PublicJoinKeyActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): PublicJoinKeyActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): PublicJoinKeyActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): PublicJoinKeyActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): PublicJoinKeyActions {
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

  async getPublicJoinKeys(): Promise<
    IResponseList<workspaces.PublicJoinKeyEntity>
  > {
    return this.apiFn(
      "GET",
      `publicJoinKeys?action=PublicJoinKeyActionQuery&${this.paramsAsString}`
    );
  }

  async getPublicJoinKeysExport(): Promise<
    IResponseList<workspaces.PublicJoinKeyEntity>
  > {
    return this.apiFn(
      "GET",
      `publicJoinKeys/export?action=PublicJoinKeyActionExport&${this.paramsAsString}`
    );
  }

  async getPublicJoinKeyByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.PublicJoinKeyEntity>> {
    return this.apiFn(
      "GET",
      `publicJoinKey/:uniqueId?action=PublicJoinKeyActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postPublicJoinKey(
    entity: workspaces.PublicJoinKeyEntity
  ): Promise<IResponse<workspaces.PublicJoinKeyEntity>> {
    return this.apiFn(
      "POST",
      `publicJoinKey?action=PublicJoinKeyActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPublicJoinKey(
    entity: workspaces.PublicJoinKeyEntity
  ): Promise<IResponse<workspaces.PublicJoinKeyEntity>> {
    return this.apiFn(
      "PATCH",
      `publicJoinKey?action=PublicJoinKeyActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPublicJoinKeys(
    entity: core.BulkRecordRequest<workspaces.PublicJoinKeyEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[workspaces.PublicJoinKeyEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `publicJoinKeys?action=PublicJoinKeyActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deletePublicJoinKey(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `publicJoinKey?action=PublicJoinKeyActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
