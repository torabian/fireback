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

export class EmailProviderActions {
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

  static isEmailProviderEntityEqual(
    a: workspaces.EmailProviderEntity,
    b: workspaces.EmailProviderEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getEmailProviderEntityPrimaryKey(
    a: workspaces.EmailProviderEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): EmailProviderActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): EmailProviderActions {
    return new EmailProviderActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): EmailProviderActions {
    return new EmailProviderActions(fn);
  }

  uniqueId(id: string): EmailProviderActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): EmailProviderActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): EmailProviderActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): EmailProviderActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): EmailProviderActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): EmailProviderActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): EmailProviderActions {
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

  async getEmailProviders(): Promise<
    IResponseList<workspaces.EmailProviderEntity>
  > {
    return this.apiFn(
      "GET",
      `emailProviders?action=EmailProviderActionQuery&${this.paramsAsString}`
    );
  }

  async getEmailProvidersExport(): Promise<
    IResponseList<workspaces.EmailProviderEntity>
  > {
    return this.apiFn(
      "GET",
      `emailProviders/export?action=EmailProviderActionExport&${this.paramsAsString}`
    );
  }

  async getEmailProviderByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.EmailProviderEntity>> {
    return this.apiFn(
      "GET",
      `emailProvider/:uniqueId?action=EmailProviderActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postEmailProvider(
    entity: workspaces.EmailProviderEntity
  ): Promise<IResponse<workspaces.EmailProviderEntity>> {
    return this.apiFn(
      "POST",
      `emailProvider?action=EmailProviderActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchEmailProvider(
    entity: workspaces.EmailProviderEntity
  ): Promise<IResponse<workspaces.EmailProviderEntity>> {
    return this.apiFn(
      "PATCH",
      `emailProvider?action=EmailProviderActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchEmailProviders(
    entity: core.BulkRecordRequest<workspaces.EmailProviderEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[workspaces.EmailProviderEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `emailProviders?action=EmailProviderActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteEmailProvider(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `emailProvider?action=EmailProviderActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
