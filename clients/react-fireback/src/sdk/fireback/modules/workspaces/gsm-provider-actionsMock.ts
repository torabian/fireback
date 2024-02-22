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

export class GsmProviderActions {
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

  static isGsmProviderEntityEqual(
    a: workspaces.GsmProviderEntity,
    b: workspaces.GsmProviderEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getGsmProviderEntityPrimaryKey(
    a: workspaces.GsmProviderEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): GsmProviderActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): GsmProviderActions {
    return new GsmProviderActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): GsmProviderActions {
    return new GsmProviderActions(fn);
  }

  uniqueId(id: string): GsmProviderActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): GsmProviderActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): GsmProviderActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): GsmProviderActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): GsmProviderActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): GsmProviderActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): GsmProviderActions {
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

  async getGsmProviders(): Promise<
    IResponseList<workspaces.GsmProviderEntity>
  > {
    return this.apiFn(
      "GET",
      `gsmProviders?action=GsmProviderActionQuery&${this.paramsAsString}`
    );
  }

  async getGsmProvidersExport(): Promise<
    IResponseList<workspaces.GsmProviderEntity>
  > {
    return this.apiFn(
      "GET",
      `gsmProviders/export?action=GsmProviderActionExport&${this.paramsAsString}`
    );
  }

  async getGsmProviderByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.GsmProviderEntity>> {
    return this.apiFn(
      "GET",
      `gsmProvider/:uniqueId?action=GsmProviderActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postGsmProvider(
    entity: workspaces.GsmProviderEntity
  ): Promise<IResponse<workspaces.GsmProviderEntity>> {
    return this.apiFn(
      "POST",
      `gsmProvider?action=GsmProviderActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchGsmProvider(
    entity: workspaces.GsmProviderEntity
  ): Promise<IResponse<workspaces.GsmProviderEntity>> {
    return this.apiFn(
      "PATCH",
      `gsmProvider?action=GsmProviderActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchGsmProviders(
    entity: core.BulkRecordRequest<workspaces.GsmProviderEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.GsmProviderEntity]>> {
    return this.apiFn(
      "PATCH",
      `gsmProviders?action=GsmProviderActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteGsmProvider(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `gsmProvider?action=GsmProviderActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
