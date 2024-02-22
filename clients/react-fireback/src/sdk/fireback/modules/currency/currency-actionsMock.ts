// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: currency
 */

import * as workspaces from "../workspaces";

import * as currency from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class CurrencyActions {
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

  static isCurrencyEntityEqual(
    a: currency.CurrencyEntity,
    b: currency.CurrencyEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getCurrencyEntityPrimaryKey(a: currency.CurrencyEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): CurrencyActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): CurrencyActions {
    return new CurrencyActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): CurrencyActions {
    return new CurrencyActions(fn);
  }

  uniqueId(id: string): CurrencyActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): CurrencyActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): CurrencyActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): CurrencyActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): CurrencyActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): CurrencyActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): CurrencyActions {
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

  async getCurrencys(): Promise<IResponseList<currency.CurrencyEntity>> {
    return this.apiFn(
      "GET",
      `currencys?action=CurrencyActionQuery&${this.paramsAsString}`
    );
  }

  async getCurrencysExport(): Promise<IResponseList<currency.CurrencyEntity>> {
    return this.apiFn(
      "GET",
      `currencys/export?action=CurrencyActionExport&${this.paramsAsString}`
    );
  }

  async getCurrencyByUniqueId(
    uniqueId: string
  ): Promise<IResponse<currency.CurrencyEntity>> {
    return this.apiFn(
      "GET",
      `currency/:uniqueId?action=CurrencyActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postCurrency(
    entity: currency.CurrencyEntity
  ): Promise<IResponse<currency.CurrencyEntity>> {
    return this.apiFn(
      "POST",
      `currency?action=CurrencyActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchCurrency(
    entity: currency.CurrencyEntity
  ): Promise<IResponse<currency.CurrencyEntity>> {
    return this.apiFn(
      "PATCH",
      `currency?action=CurrencyActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchCurrencys(
    entity: core.BulkRecordRequest<currency.CurrencyEntity>
  ): Promise<IResponse<core.BulkRecordRequest[currency.CurrencyEntity]>> {
    return this.apiFn(
      "PATCH",
      `currencys?action=CurrencyActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteCurrency(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `currency?action=CurrencyActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
