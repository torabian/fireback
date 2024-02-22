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

export class UserActions {
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

  static isUserEntityEqual(
    a: workspaces.UserEntity,
    b: workspaces.UserEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getUserEntityPrimaryKey(a: workspaces.UserEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): UserActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): UserActions {
    return new UserActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): UserActions {
    return new UserActions(fn);
  }

  uniqueId(id: string): UserActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): UserActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): UserActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): UserActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): UserActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): UserActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): UserActions {
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

  async getUsers(): Promise<IResponseList<workspaces.UserEntity>> {
    return this.apiFn(
      "GET",
      `users?action=UserActionQuery&${this.paramsAsString}`
    );
  }

  async getUsersExport(): Promise<IResponseList<workspaces.UserEntity>> {
    return this.apiFn(
      "GET",
      `users/export?action=UserActionExport&${this.paramsAsString}`
    );
  }

  async getUserByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.UserEntity>> {
    return this.apiFn(
      "GET",
      `user/:uniqueId?action=UserActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postUser(
    entity: workspaces.UserEntity
  ): Promise<IResponse<workspaces.UserEntity>> {
    return this.apiFn(
      "POST",
      `user?action=UserActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchUser(
    entity: workspaces.UserEntity
  ): Promise<IResponse<workspaces.UserEntity>> {
    return this.apiFn(
      "PATCH",
      `user?action=UserActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchUsers(
    entity: core.BulkRecordRequest<workspaces.UserEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.UserEntity]>> {
    return this.apiFn(
      "PATCH",
      `users?action=UserActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteUser(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `user?action=UserActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
