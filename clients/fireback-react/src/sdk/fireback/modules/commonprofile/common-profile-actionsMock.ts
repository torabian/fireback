// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: commonprofile
 */

import * as workspaces from "../workspaces";

import * as commonprofile from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class CommonProfileActions {
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

  static isCommonProfileEntityEqual(
    a: commonprofile.CommonProfileEntity,
    b: commonprofile.CommonProfileEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getCommonProfileEntityPrimaryKey(
    a: commonprofile.CommonProfileEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): CommonProfileActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): CommonProfileActions {
    return new CommonProfileActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): CommonProfileActions {
    return new CommonProfileActions(fn);
  }

  uniqueId(id: string): CommonProfileActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): CommonProfileActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): CommonProfileActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): CommonProfileActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): CommonProfileActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): CommonProfileActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): CommonProfileActions {
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

  async getCommonProfiles(): Promise<
    IResponseList<commonprofile.CommonProfileEntity>
  > {
    return this.apiFn(
      "GET",
      `commonProfiles?action=CommonProfileActionQuery&${this.paramsAsString}`
    );
  }

  async getCommonProfilesExport(): Promise<
    IResponseList<commonprofile.CommonProfileEntity>
  > {
    return this.apiFn(
      "GET",
      `commonProfiles/export?action=CommonProfileActionExport&${this.paramsAsString}`
    );
  }

  async getCommonProfileByUniqueId(
    uniqueId: string
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> {
    return this.apiFn(
      "GET",
      `commonProfile/:uniqueId?action=CommonProfileActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postCommonProfile(
    entity: commonprofile.CommonProfileEntity
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> {
    return this.apiFn(
      "POST",
      `commonProfile?action=CommonProfileActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchCommonProfile(
    entity: commonprofile.CommonProfileEntity
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> {
    return this.apiFn(
      "PATCH",
      `commonProfile?action=CommonProfileActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchCommonProfileDistinct(
    entity: commonprofile.CommonProfileEntity
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> {
    return this.apiFn(
      "PATCH",
      `commonProfileDistinct?action=CommonProfileDistinctActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async getCommonProfileDistinct(): Promise<
    IResponse<commonprofile.CommonProfileEntity>
  > {
    return this.apiFn(
      "GET",
      `commonProfileDistinct?action=CommonProfileDistinctActionGetOne&${this.paramsAsString}`
    );
  }

  async patchCommonProfiles(
    entity: core.BulkRecordRequest<commonprofile.CommonProfileEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[commonprofile.CommonProfileEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `commonProfiles?action=CommonProfileActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteCommonProfile(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `commonProfile?action=CommonProfileActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
