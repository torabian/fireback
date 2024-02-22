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

export class RoleActions {
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

  static isRoleEntityEqual(
    a: workspaces.RoleEntity,
    b: workspaces.RoleEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getRoleEntityPrimaryKey(a: workspaces.RoleEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): RoleActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): RoleActions {
    return new RoleActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): RoleActions {
    return new RoleActions(fn);
  }

  uniqueId(id: string): RoleActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): RoleActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): RoleActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): RoleActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): RoleActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): RoleActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): RoleActions {
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

  async getRoles(): Promise<IResponseList<workspaces.RoleEntity>> {
    return this.apiFn(
      "GET",
      `roles?action=RoleActionQuery&${this.paramsAsString}`
    );
  }

  async getRolesExport(): Promise<IResponseList<workspaces.RoleEntity>> {
    return this.apiFn(
      "GET",
      `roles/export?action=RoleActionExport&${this.paramsAsString}`
    );
  }

  async getRoleByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.RoleEntity>> {
    return this.apiFn(
      "GET",
      `role/:uniqueId?action=RoleActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postRole(
    entity: workspaces.RoleEntity
  ): Promise<IResponse<workspaces.RoleEntity>> {
    return this.apiFn(
      "POST",
      `role?action=RoleActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchRole(
    entity: workspaces.RoleEntity
  ): Promise<IResponse<workspaces.RoleEntity>> {
    return this.apiFn(
      "PATCH",
      `role?action=RoleActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchRoles(
    entity: core.BulkRecordRequest<workspaces.RoleEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.RoleEntity]>> {
    return this.apiFn(
      "PATCH",
      `roles?action=RoleActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteRole(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `role?action=RoleActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
