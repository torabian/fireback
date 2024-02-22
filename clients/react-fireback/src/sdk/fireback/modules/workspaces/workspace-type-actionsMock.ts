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

export class WorkspaceTypeActions {
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

  static isWorkspaceTypeEntityEqual(
    a: workspaces.WorkspaceTypeEntity,
    b: workspaces.WorkspaceTypeEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getWorkspaceTypeEntityPrimaryKey(
    a: workspaces.WorkspaceTypeEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): WorkspaceTypeActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): WorkspaceTypeActions {
    return new WorkspaceTypeActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): WorkspaceTypeActions {
    return new WorkspaceTypeActions(fn);
  }

  uniqueId(id: string): WorkspaceTypeActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): WorkspaceTypeActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): WorkspaceTypeActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): WorkspaceTypeActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): WorkspaceTypeActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): WorkspaceTypeActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): WorkspaceTypeActions {
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

  async getWorkspaceTypes(): Promise<
    IResponseList<workspaces.WorkspaceTypeEntity>
  > {
    return this.apiFn(
      "GET",
      `workspaceTypes?action=WorkspaceTypeActionQuery&${this.paramsAsString}`
    );
  }

  async getWorkspaceTypesExport(): Promise<
    IResponseList<workspaces.WorkspaceTypeEntity>
  > {
    return this.apiFn(
      "GET",
      `workspaceTypes/export?action=WorkspaceTypeActionExport&${this.paramsAsString}`
    );
  }

  async getWorkspaceTypeByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> {
    return this.apiFn(
      "GET",
      `workspaceType/:uniqueId?action=WorkspaceTypeActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postWorkspaceType(
    entity: workspaces.WorkspaceTypeEntity
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> {
    return this.apiFn(
      "POST",
      `workspaceType?action=WorkspaceTypeActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWorkspaceType(
    entity: workspaces.WorkspaceTypeEntity
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> {
    return this.apiFn(
      "PATCH",
      `workspaceType?action=WorkspaceTypeActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWorkspaceTypeDistinct(
    entity: workspaces.WorkspaceTypeEntity
  ): Promise<IResponse<workspaces.WorkspaceTypeEntity>> {
    return this.apiFn(
      "PATCH",
      `workspaceTypeDistinct?action=WorkspaceTypeDistinctActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async getWorkspaceTypeDistinct(): Promise<
    IResponse<workspaces.WorkspaceTypeEntity>
  > {
    return this.apiFn(
      "GET",
      `workspaceTypeDistinct?action=WorkspaceTypeDistinctActionGetOne&${this.paramsAsString}`
    );
  }

  async patchWorkspaceTypes(
    entity: core.BulkRecordRequest<workspaces.WorkspaceTypeEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[workspaces.WorkspaceTypeEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `workspaceTypes?action=WorkspaceTypeActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteWorkspaceType(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `workspaceType?action=WorkspaceTypeActionRemove&${this.paramsAsString}`,
      entity
    );
  }

  async getPublicWorkspaceTypes(): Promise<
    IResponseList<workspaces.WorkspaceTypeEntity>
  > {
    return this.apiFn(
      "GET",
      `publicWorkspaceTypes?action=WorkspaceTypeActionPublicQuery&${this.paramsAsString}`
    );
  }
}
