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

export class UserRoleWorkspaceActions {
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

  static isUserRoleWorkspaceEntityEqual(
    a: workspaces.UserRoleWorkspaceEntity,
    b: workspaces.UserRoleWorkspaceEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getUserRoleWorkspaceEntityPrimaryKey(
    a: workspaces.UserRoleWorkspaceEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): UserRoleWorkspaceActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): UserRoleWorkspaceActions {
    return new UserRoleWorkspaceActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): UserRoleWorkspaceActions {
    return new UserRoleWorkspaceActions(fn);
  }

  uniqueId(id: string): UserRoleWorkspaceActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): UserRoleWorkspaceActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): UserRoleWorkspaceActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): UserRoleWorkspaceActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): UserRoleWorkspaceActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): UserRoleWorkspaceActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): UserRoleWorkspaceActions {
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

  async getUserRoleWorkspaces(): Promise<
    IResponseList<workspaces.UserRoleWorkspaceEntity>
  > {
    return this.apiFn("GET", `userRoleWorkspaces?${this.paramsAsString}`);
  }
}
