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

export class PriceTagActions {
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

  static isPriceTagEntityEqual(
    a: currency.PriceTagEntity,
    b: currency.PriceTagEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getPriceTagEntityPrimaryKey(a: currency.PriceTagEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): PriceTagActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): PriceTagActions {
    return new PriceTagActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): PriceTagActions {
    return new PriceTagActions(fn);
  }

  uniqueId(id: string): PriceTagActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): PriceTagActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): PriceTagActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): PriceTagActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): PriceTagActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): PriceTagActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): PriceTagActions {
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

  async getPriceTags(): Promise<IResponseList<currency.PriceTagEntity>> {
    return this.apiFn(
      "GET",
      `priceTags?action=PriceTagActionQuery&${this.paramsAsString}`
    );
  }

  async getPriceTagsExport(): Promise<IResponseList<currency.PriceTagEntity>> {
    return this.apiFn(
      "GET",
      `priceTags/export?action=PriceTagActionExport&${this.paramsAsString}`
    );
  }

  async getPriceTagByUniqueId(
    uniqueId: string
  ): Promise<IResponse<currency.PriceTagEntity>> {
    return this.apiFn(
      "GET",
      `priceTag/:uniqueId?action=PriceTagActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postPriceTag(
    entity: currency.PriceTagEntity
  ): Promise<IResponse<currency.PriceTagEntity>> {
    return this.apiFn(
      "POST",
      `priceTag?action=PriceTagActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPriceTag(
    entity: currency.PriceTagEntity
  ): Promise<IResponse<currency.PriceTagEntity>> {
    return this.apiFn(
      "PATCH",
      `priceTag?action=PriceTagActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPriceTags(
    entity: core.BulkRecordRequest<currency.PriceTagEntity>
  ): Promise<IResponse<core.BulkRecordRequest[currency.PriceTagEntity]>> {
    return this.apiFn(
      "PATCH",
      `priceTags?action=PriceTagActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deletePriceTag(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `priceTag?action=PriceTagActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
