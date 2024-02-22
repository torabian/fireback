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

export class AppMenuActions {
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

  static isAppMenuEntityEqual(
    a: workspaces.AppMenuEntity,
    b: workspaces.AppMenuEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getAppMenuEntityPrimaryKey(a: workspaces.AppMenuEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): AppMenuActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): AppMenuActions {
    return new AppMenuActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): AppMenuActions {
    return new AppMenuActions(fn);
  }

  uniqueId(id: string): AppMenuActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): AppMenuActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): AppMenuActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): AppMenuActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): AppMenuActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): AppMenuActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): AppMenuActions {
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

  async getCteAppMenus(): Promise<IResponseList<workspaces.AppMenuEntity>> {
    return this.apiFn(
      "GET",
      `cteAppMenus?action=AppMenuActionCteQuery&${this.paramsAsString}`
    );
  }

  async getAppMenus(): Promise<IResponseList<workspaces.AppMenuEntity>> {
    return this.apiFn(
      "GET",
      `appMenus?action=AppMenuActionQuery&${this.paramsAsString}`
    );
  }

  async getAppMenusExport(): Promise<IResponseList<workspaces.AppMenuEntity>> {
    return this.apiFn(
      "GET",
      `appMenus/export?action=AppMenuActionExport&${this.paramsAsString}`
    );
  }

  async getAppMenuByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.AppMenuEntity>> {
    return this.apiFn(
      "GET",
      `appMenu/:uniqueId?action=AppMenuActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postAppMenu(
    entity: workspaces.AppMenuEntity
  ): Promise<IResponse<workspaces.AppMenuEntity>> {
    return this.apiFn(
      "POST",
      `appMenu?action=AppMenuActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchAppMenu(
    entity: workspaces.AppMenuEntity
  ): Promise<IResponse<workspaces.AppMenuEntity>> {
    return this.apiFn(
      "PATCH",
      `appMenu?action=AppMenuActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchAppMenus(
    entity: core.BulkRecordRequest<workspaces.AppMenuEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.AppMenuEntity]>> {
    return this.apiFn(
      "PATCH",
      `appMenus?action=AppMenuActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteAppMenu(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `appMenu?action=AppMenuActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
