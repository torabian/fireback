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

export class WidgetActions {
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

  static isWidgetEntityEqual(
    a: widget.WidgetEntity,
    b: widget.WidgetEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getWidgetEntityPrimaryKey(a: widget.WidgetEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): WidgetActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): WidgetActions {
    return new WidgetActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): WidgetActions {
    return new WidgetActions(fn);
  }

  uniqueId(id: string): WidgetActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): WidgetActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): WidgetActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): WidgetActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): WidgetActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): WidgetActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): WidgetActions {
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

  async getWidgets(): Promise<IResponseList<widget.WidgetEntity>> {
    return this.apiFn(
      "GET",
      `widgets?action=WidgetActionQuery&${this.paramsAsString}`
    );
  }

  async getWidgetsExport(): Promise<IResponseList<widget.WidgetEntity>> {
    return this.apiFn(
      "GET",
      `widgets/export?action=WidgetActionExport&${this.paramsAsString}`
    );
  }

  async getWidgetByUniqueId(
    uniqueId: string
  ): Promise<IResponse<widget.WidgetEntity>> {
    return this.apiFn(
      "GET",
      `widget/:uniqueId?action=WidgetActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postWidget(
    entity: widget.WidgetEntity
  ): Promise<IResponse<widget.WidgetEntity>> {
    return this.apiFn(
      "POST",
      `widget?action=WidgetActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWidget(
    entity: widget.WidgetEntity
  ): Promise<IResponse<widget.WidgetEntity>> {
    return this.apiFn(
      "PATCH",
      `widget?action=WidgetActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchWidgets(
    entity: core.BulkRecordRequest<widget.WidgetEntity>
  ): Promise<IResponse<core.BulkRecordRequest[widget.WidgetEntity]>> {
    return this.apiFn(
      "PATCH",
      `widgets?action=WidgetActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteWidget(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `widget?action=WidgetActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
