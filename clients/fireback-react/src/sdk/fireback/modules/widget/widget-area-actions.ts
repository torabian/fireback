// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: widget
 */

import * as workspaces from "../workspaces";

import * as widget from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class WidgetAreaActions {
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

  static isWidgetAreaEntityEqual(
    a: widget.WidgetAreaEntity,
    b: widget.WidgetAreaEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getWidgetAreaEntityPrimaryKey(a: widget.WidgetAreaEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): WidgetAreaActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): WidgetAreaActions {
    return new WidgetAreaActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): WidgetAreaActions {
    return new WidgetAreaActions(fn);
  }

  uniqueId(id: string): WidgetAreaActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): WidgetAreaActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): WidgetAreaActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): WidgetAreaActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): WidgetAreaActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): WidgetAreaActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): WidgetAreaActions {
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

  async getWidgetAreas(): Promise<IResponseList<widget.WidgetAreaEntity>> {
    return this.apiFn(
      "GET",
      `widgetAreas?action=WidgetAreaActionQuery&${this.paramsAsString}`
    );
  }

  async getWidgetAreasExport(): Promise<
    IResponseList<widget.WidgetAreaEntity>
  > {
    return this.apiFn(
      "GET",
      `widgetAreas/export?action=WidgetAreaActionExport&${this.paramsAsString}`
    );
  }

  async getWidgetAreaByUniqueId(
    uniqueId: string
  ): Promise<IResponse<widget.WidgetAreaEntity>> {
    return this.apiFn(
      "GET",
      `widgetArea/:uniqueId?action=WidgetAreaActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postWidgetArea(
    entity: widget.WidgetAreaEntity
  ): Promise<IResponse<widget.WidgetAreaEntity>> {
    return this.apiFn(
      "POST",
      `widgetArea?action=WidgetAreaActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWidgetArea(
    entity: widget.WidgetAreaEntity
  ): Promise<IResponse<widget.WidgetAreaEntity>> {
    return this.apiFn(
      "PATCH",
      `widgetArea?action=WidgetAreaActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWidgetAreas(
    entity: core.BulkRecordRequest<widget.WidgetAreaEntity>
  ): Promise<IResponse<core.BulkRecordRequest[widget.WidgetAreaEntity]>> {
    return this.apiFn(
      "PATCH",
      `widgetAreas?action=WidgetAreaActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteWidgetArea(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `widgetArea?action=WidgetAreaActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
