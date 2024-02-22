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

export class ProductPlanActions {
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

  static isProductPlanEntityEqual(
    a: licenses.ProductPlanEntity,
    b: licenses.ProductPlanEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getProductPlanEntityPrimaryKey(a: licenses.ProductPlanEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): ProductPlanActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): ProductPlanActions {
    return new ProductPlanActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): ProductPlanActions {
    return new ProductPlanActions(fn);
  }

  uniqueId(id: string): ProductPlanActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): ProductPlanActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): ProductPlanActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): ProductPlanActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): ProductPlanActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): ProductPlanActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): ProductPlanActions {
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

  async getProductPlans(): Promise<IResponseList<licenses.ProductPlanEntity>> {
    return this.apiFn(
      "GET",
      `productPlans?action=ProductPlanActionQuery&${this.paramsAsString}`
    );
  }

  async getProductPlansExport(): Promise<
    IResponseList<licenses.ProductPlanEntity>
  > {
    return this.apiFn(
      "GET",
      `productPlans/export?action=ProductPlanActionExport&${this.paramsAsString}`
    );
  }

  async getProductPlanByUniqueId(
    uniqueId: string
  ): Promise<IResponse<licenses.ProductPlanEntity>> {
    return this.apiFn(
      "GET",
      `productPlan/:uniqueId?action=ProductPlanActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postProductPlan(
    entity: licenses.ProductPlanEntity
  ): Promise<IResponse<licenses.ProductPlanEntity>> {
    return this.apiFn(
      "POST",
      `productPlan?action=ProductPlanActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchProductPlan(
    entity: licenses.ProductPlanEntity
  ): Promise<IResponse<licenses.ProductPlanEntity>> {
    return this.apiFn(
      "PATCH",
      `productPlan?action=ProductPlanActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchProductPlans(
    entity: core.BulkRecordRequest<licenses.ProductPlanEntity>
  ): Promise<IResponse<core.BulkRecordRequest[licenses.ProductPlanEntity]>> {
    return this.apiFn(
      "PATCH",
      `productPlans?action=ProductPlanActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteProductPlan(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `productPlan?action=ProductPlanActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
