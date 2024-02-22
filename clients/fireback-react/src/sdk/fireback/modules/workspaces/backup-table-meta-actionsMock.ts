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

export class BackupTableMetaActions {
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

  static isBackupTableMetaEntityEqual(
    a: workspaces.BackupTableMetaEntity,
    b: workspaces.BackupTableMetaEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getBackupTableMetaEntityPrimaryKey(
    a: workspaces.BackupTableMetaEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): BackupTableMetaActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): BackupTableMetaActions {
    return new BackupTableMetaActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): BackupTableMetaActions {
    return new BackupTableMetaActions(fn);
  }

  uniqueId(id: string): BackupTableMetaActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): BackupTableMetaActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): BackupTableMetaActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): BackupTableMetaActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): BackupTableMetaActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): BackupTableMetaActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): BackupTableMetaActions {
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

  async getBackupTableMetas(): Promise<
    IResponseList<workspaces.BackupTableMetaEntity>
  > {
    return this.apiFn(
      "GET",
      `backupTableMetas?action=BackupTableMetaActionQuery&${this.paramsAsString}`
    );
  }

  async getBackupTableMetasExport(): Promise<
    IResponseList<workspaces.BackupTableMetaEntity>
  > {
    return this.apiFn(
      "GET",
      `backupTableMetas/export?action=BackupTableMetaActionExport&${this.paramsAsString}`
    );
  }

  async getBackupTableMetaByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.BackupTableMetaEntity>> {
    return this.apiFn(
      "GET",
      `backupTableMeta/:uniqueId?action=BackupTableMetaActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postBackupTableMeta(
    entity: workspaces.BackupTableMetaEntity
  ): Promise<IResponse<workspaces.BackupTableMetaEntity>> {
    return this.apiFn(
      "POST",
      `backupTableMeta?action=BackupTableMetaActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchBackupTableMeta(
    entity: workspaces.BackupTableMetaEntity
  ): Promise<IResponse<workspaces.BackupTableMetaEntity>> {
    return this.apiFn(
      "PATCH",
      `backupTableMeta?action=BackupTableMetaActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchBackupTableMetas(
    entity: core.BulkRecordRequest<workspaces.BackupTableMetaEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[workspaces.BackupTableMetaEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `backupTableMetas?action=BackupTableMetaActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteBackupTableMeta(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `backupTableMeta?action=BackupTableMetaActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
