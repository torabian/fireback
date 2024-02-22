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

export class CapabilityActions {
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

  static isCapabilityEntityEqual(
    a: workspaces.CapabilityEntity,
    b: workspaces.CapabilityEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getCapabilityEntityPrimaryKey(a: workspaces.CapabilityEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): CapabilityActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): CapabilityActions {
    return new CapabilityActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): CapabilityActions {
    return new CapabilityActions(fn);
  }

  uniqueId(id: string): CapabilityActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): CapabilityActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): CapabilityActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): CapabilityActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): CapabilityActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): CapabilityActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): CapabilityActions {
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

  async getCapabilitys(): Promise<IResponseList<workspaces.CapabilityEntity>> {
    return this.apiFn(
      "GET",
      `capabilitys?action=CapabilityActionQuery&${this.paramsAsString}`
    );
  }

  async getCapabilitysExport(): Promise<
    IResponseList<workspaces.CapabilityEntity>
  > {
    return this.apiFn(
      "GET",
      `capabilitys/export?action=CapabilityActionExport&${this.paramsAsString}`
    );
  }

  async getCapabilityByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.CapabilityEntity>> {
    return this.apiFn(
      "GET",
      `capability/:uniqueId?action=CapabilityActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postCapability(
    entity: workspaces.CapabilityEntity
  ): Promise<IResponse<workspaces.CapabilityEntity>> {
    return this.apiFn(
      "POST",
      `capability?action=CapabilityActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchCapability(
    entity: workspaces.CapabilityEntity
  ): Promise<IResponse<workspaces.CapabilityEntity>> {
    return this.apiFn(
      "PATCH",
      `capability?action=CapabilityActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchCapabilitys(
    entity: core.BulkRecordRequest<workspaces.CapabilityEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.CapabilityEntity]>> {
    return this.apiFn(
      "PATCH",
      `capabilitys?action=CapabilityActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteCapability(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `capability?action=CapabilityActionRemove&${this.paramsAsString}`,
      entity
    );
  }

  async getCapabilitiesTree(): Promise<
    IResponse<workspaces.CapabilitiesResult>
  > {
    return this.apiFn("GET", `capabilitiesTree?${this.paramsAsString}`);
  }
}
