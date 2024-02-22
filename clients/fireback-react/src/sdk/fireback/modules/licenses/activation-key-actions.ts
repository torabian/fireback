// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: licenses
 */

import * as workspaces from "../workspaces";

import * as licenses from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class ActivationKeyActions {
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

  static isActivationKeyEntityEqual(
    a: licenses.ActivationKeyEntity,
    b: licenses.ActivationKeyEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getActivationKeyEntityPrimaryKey(
    a: licenses.ActivationKeyEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): ActivationKeyActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): ActivationKeyActions {
    return new ActivationKeyActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): ActivationKeyActions {
    return new ActivationKeyActions(fn);
  }

  uniqueId(id: string): ActivationKeyActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): ActivationKeyActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): ActivationKeyActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): ActivationKeyActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): ActivationKeyActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): ActivationKeyActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): ActivationKeyActions {
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

  async getActivationKeys(): Promise<
    IResponseList<licenses.ActivationKeyEntity>
  > {
    return this.apiFn(
      "GET",
      `activationKeys?action=ActivationKeyActionQuery&${this.paramsAsString}`
    );
  }

  async getActivationKeysExport(): Promise<
    IResponseList<licenses.ActivationKeyEntity>
  > {
    return this.apiFn(
      "GET",
      `activationKeys/export?action=ActivationKeyActionExport&${this.paramsAsString}`
    );
  }

  async getActivationKeyByUniqueId(
    uniqueId: string
  ): Promise<IResponse<licenses.ActivationKeyEntity>> {
    return this.apiFn(
      "GET",
      `activationKey/:uniqueId?action=ActivationKeyActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postActivationKey(
    entity: licenses.ActivationKeyEntity
  ): Promise<IResponse<licenses.ActivationKeyEntity>> {
    return this.apiFn(
      "POST",
      `activationKey?action=ActivationKeyActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchActivationKey(
    entity: licenses.ActivationKeyEntity
  ): Promise<IResponse<licenses.ActivationKeyEntity>> {
    return this.apiFn(
      "PATCH",
      `activationKey?action=ActivationKeyActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchActivationKeys(
    entity: core.BulkRecordRequest<licenses.ActivationKeyEntity>
  ): Promise<IResponse<core.BulkRecordRequest[licenses.ActivationKeyEntity]>> {
    return this.apiFn(
      "PATCH",
      `activationKeys?action=ActivationKeyActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteActivationKey(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `activationKey?action=ActivationKeyActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
