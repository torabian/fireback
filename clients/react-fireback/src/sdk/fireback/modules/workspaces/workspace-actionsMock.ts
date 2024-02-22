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

export class WorkspaceActions {
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

  static isWorkspaceEntityEqual(
    a: workspaces.WorkspaceEntity,
    b: workspaces.WorkspaceEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getWorkspaceEntityPrimaryKey(a: workspaces.WorkspaceEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): WorkspaceActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): WorkspaceActions {
    return new WorkspaceActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): WorkspaceActions {
    return new WorkspaceActions(fn);
  }

  uniqueId(id: string): WorkspaceActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): WorkspaceActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): WorkspaceActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): WorkspaceActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): WorkspaceActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): WorkspaceActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): WorkspaceActions {
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

  async getCteWorkspaces(): Promise<IResponseList<workspaces.WorkspaceEntity>> {
    return this.apiFn(
      "GET",
      `cteWorkspaces?action=WorkspaceActionCteQuery&${this.paramsAsString}`
    );
  }

  async getWorkspaces(): Promise<IResponseList<workspaces.WorkspaceEntity>> {
    return this.apiFn(
      "GET",
      `workspaces?action=WorkspaceActionQuery&${this.paramsAsString}`
    );
  }

  async getWorkspacesExport(): Promise<
    IResponseList<workspaces.WorkspaceEntity>
  > {
    return this.apiFn(
      "GET",
      `workspaces/export?action=WorkspaceActionExport&${this.paramsAsString}`
    );
  }

  async getWorkspaceByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.WorkspaceEntity>> {
    return this.apiFn(
      "GET",
      `workspace/:uniqueId?action=WorkspaceActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postWorkspace(
    entity: workspaces.WorkspaceEntity
  ): Promise<IResponse<workspaces.WorkspaceEntity>> {
    return this.apiFn(
      "POST",
      `workspace?action=WorkspaceActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWorkspace(
    entity: workspaces.WorkspaceEntity
  ): Promise<IResponse<workspaces.WorkspaceEntity>> {
    return this.apiFn(
      "PATCH",
      `workspace?action=WorkspaceActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWorkspaces(
    entity: core.BulkRecordRequest<workspaces.WorkspaceEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.WorkspaceEntity]>> {
    return this.apiFn(
      "PATCH",
      `workspaces?action=WorkspaceActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteWorkspace(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `workspace?action=WorkspaceActionRemove&${this.paramsAsString}`,
      entity
    );
  }

  async reactiveReactiveSearch(): Promise<
    IResponse<core.ReactiveSearchResultDto>
  > {
    return this.apiFn("REACTIVE", `reactiveSearch?${this.paramsAsString}`);
  }
}
