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

export class LicenseActions {
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

  static isLicenseEntityEqual(
    a: licenses.LicenseEntity,
    b: licenses.LicenseEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getLicenseEntityPrimaryKey(a: licenses.LicenseEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): LicenseActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): LicenseActions {
    return new LicenseActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): LicenseActions {
    return new LicenseActions(fn);
  }

  uniqueId(id: string): LicenseActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): LicenseActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): LicenseActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): LicenseActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): LicenseActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): LicenseActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): LicenseActions {
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

  async getLicenses(): Promise<IResponseList<licenses.LicenseEntity>> {
    return this.apiFn(
      "GET",
      `licenses?action=LicenseActionQuery&${this.paramsAsString}`
    );
  }

  async getLicensesExport(): Promise<IResponseList<licenses.LicenseEntity>> {
    return this.apiFn(
      "GET",
      `licenses/export?action=LicenseActionExport&${this.paramsAsString}`
    );
  }

  async getLicenseByUniqueId(
    uniqueId: string
  ): Promise<IResponse<licenses.LicenseEntity>> {
    return this.apiFn(
      "GET",
      `license/:uniqueId?action=LicenseActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postLicense(
    entity: licenses.LicenseEntity
  ): Promise<IResponse<licenses.LicenseEntity>> {
    return this.apiFn(
      "POST",
      `license?action=LicenseActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchLicense(
    entity: licenses.LicenseEntity
  ): Promise<IResponse<licenses.LicenseEntity>> {
    return this.apiFn(
      "PATCH",
      `license?action=LicenseActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchLicenses(
    entity: core.BulkRecordRequest<licenses.LicenseEntity>
  ): Promise<IResponse<core.BulkRecordRequest[licenses.LicenseEntity]>> {
    return this.apiFn(
      "PATCH",
      `licenses?action=LicenseActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteLicense(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `license?action=LicenseActionRemove&${this.paramsAsString}`,
      entity
    );
  }

  async postLicenseFromPlanByUniqueId(
    uniqueId: string,

    entity: licenses.LicenseFromPlanIdDto
  ): Promise<IResponse<licenses.LicenseEntity>> {
    return this.apiFn(
      "POST",
      `license/from-plan/:uniqueId?${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      ),
      entity
    );
  }
}
