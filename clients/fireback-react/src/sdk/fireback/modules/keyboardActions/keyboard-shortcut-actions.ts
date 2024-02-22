// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: keyboardActions
 */

import * as workspaces from "../workspaces";

import * as keyboardActions from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class KeyboardShortcutActions {
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

  static isKeyboardShortcutEntityEqual(
    a: keyboardActions.KeyboardShortcutEntity,
    b: keyboardActions.KeyboardShortcutEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getKeyboardShortcutEntityPrimaryKey(
    a: keyboardActions.KeyboardShortcutEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): KeyboardShortcutActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): KeyboardShortcutActions {
    return new KeyboardShortcutActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): KeyboardShortcutActions {
    return new KeyboardShortcutActions(fn);
  }

  uniqueId(id: string): KeyboardShortcutActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): KeyboardShortcutActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): KeyboardShortcutActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): KeyboardShortcutActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): KeyboardShortcutActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): KeyboardShortcutActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): KeyboardShortcutActions {
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

  async getKeyboardShortcuts(): Promise<
    IResponseList<keyboardActions.KeyboardShortcutEntity>
  > {
    return this.apiFn(
      "GET",
      `keyboardShortcuts?action=KeyboardShortcutActionQuery&${this.paramsAsString}`
    );
  }

  async getKeyboardShortcutsExport(): Promise<
    IResponseList<keyboardActions.KeyboardShortcutEntity>
  > {
    return this.apiFn(
      "GET",
      `keyboardShortcuts/export?action=KeyboardShortcutActionExport&${this.paramsAsString}`
    );
  }

  async getKeyboardShortcutByUniqueId(
    uniqueId: string
  ): Promise<IResponse<keyboardActions.KeyboardShortcutEntity>> {
    return this.apiFn(
      "GET",
      `keyboardShortcut/:uniqueId?action=KeyboardShortcutActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postKeyboardShortcut(
    entity: keyboardActions.KeyboardShortcutEntity
  ): Promise<IResponse<keyboardActions.KeyboardShortcutEntity>> {
    return this.apiFn(
      "POST",
      `keyboardShortcut?action=KeyboardShortcutActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchKeyboardShortcut(
    entity: keyboardActions.KeyboardShortcutEntity
  ): Promise<IResponse<keyboardActions.KeyboardShortcutEntity>> {
    return this.apiFn(
      "PATCH",
      `keyboardShortcut?action=KeyboardShortcutActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchKeyboardShortcuts(
    entity: core.BulkRecordRequest<keyboardActions.KeyboardShortcutEntity>
  ): Promise<
    IResponse<core.BulkRecordRequest[keyboardActions.KeyboardShortcutEntity]>
  > {
    return this.apiFn(
      "PATCH",
      `keyboardShortcuts?action=KeyboardShortcutActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteKeyboardShortcut(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `keyboardShortcut?action=KeyboardShortcutActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
