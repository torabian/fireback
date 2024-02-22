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

export class TableViewSizingActions {
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

  static isTableViewSizingEntityEqual(
    a: workspaces.TableViewSizingEntity,
    b: workspaces.TableViewSizingEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getTableViewSizingEntityPrimaryKey(
    a: workspaces.TableViewSizingEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): TableViewSizingActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): TableViewSizingActions {
    return new TableViewSizingActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): TableViewSizingActions {
    return new TableViewSizingActions(fn);
  }

  uniqueId(id: string): TableViewSizingActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): TableViewSizingActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): TableViewSizingActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): TableViewSizingActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): TableViewSizingActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): TableViewSizingActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): TableViewSizingActions {
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

  async getTableViewSizings(): Promise<
    IResponseList<workspaces.TableViewSizingEntity>
  > {
    return this.apiFn(
      "GET",
      `tableViewSizings?action=TableViewSizingActionQuery&${this.paramsAsString}`
    );
  }

  async getTableViewSizingsExport(): Promise<
    IResponseList<workspaces.TableViewSizingEntity>
  > {
    return this.apiFn(
      "GET",
      `tableViewSizings/export?action=TableViewSizingActionExport&${this.paramsAsString}`
    );
  }

  async getTableViewSizingByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.TableViewSizingEntity>> {
    return this.apiFn(
      "GET",
      `tableViewSizing/:uniqueId?action=TableViewSizingActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postTableViewSizing(
    entity: workspaces.TableViewSizingEntity
  ): Promise<IResponse<workspaces.TableViewSizingEntity>> {
    return this.apiFn(
      "POST",
      `tableViewSizing?action=TableViewSizingActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchTableViewSizing(
    entity: workspaces.TableViewSizingEntity
  ): Promise<IResponse<workspaces.TableViewSizingEntity>> {
    return this.apiFn(
      "PATCH",
      `tableViewSizing?action=TableViewSizingActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchTableViewSizings(
    entity: core.BulkRecordRequest<workspaces.TableViewSizingEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[workspaces.TableViewSizingEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `tableViewSizings?action=TableViewSizingActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteTableViewSizing(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `tableViewSizing?action=TableViewSizingActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
