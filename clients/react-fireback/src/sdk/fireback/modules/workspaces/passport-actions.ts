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

export class PassportActions {
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

  static isPassportEntityEqual(
    a: workspaces.PassportEntity,
    b: workspaces.PassportEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getPassportEntityPrimaryKey(a: workspaces.PassportEntity): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): PassportActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): PassportActions {
    return new PassportActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): PassportActions {
    return new PassportActions(fn);
  }

  uniqueId(id: string): PassportActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): PassportActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): PassportActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): PassportActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): PassportActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): PassportActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): PassportActions {
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

  async getPassports(): Promise<IResponseList<workspaces.PassportEntity>> {
    return this.apiFn(
      "GET",
      `passports?action=PassportActionQuery&${this.paramsAsString}`
    );
  }

  async getPassportsExport(): Promise<
    IResponseList<workspaces.PassportEntity>
  > {
    return this.apiFn(
      "GET",
      `passports/export?action=PassportActionExport&${this.paramsAsString}`
    );
  }

  async getPassportByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.PassportEntity>> {
    return this.apiFn(
      "GET",
      `passport/:uniqueId?action=PassportActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postPassport(
    entity: workspaces.PassportEntity
  ): Promise<IResponse<workspaces.PassportEntity>> {
    return this.apiFn(
      "POST",
      `passport?action=PassportActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPassport(
    entity: workspaces.PassportEntity
  ): Promise<IResponse<workspaces.PassportEntity>> {
    return this.apiFn(
      "PATCH",
      `passport?action=PassportActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchPassports(
    entity: core.BulkRecordRequest<workspaces.PassportEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.PassportEntity]>> {
    return this.apiFn(
      "PATCH",
      `passports?action=PassportActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deletePassport(entity: core.DeleteRequest): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `passport?action=PassportActionRemove&${this.paramsAsString}`,
      entity
    );
  }

  async postPassportSignupEmail(
    entity: workspaces.EmailAccountSignupDto
  ): Promise<IResponse<workspaces.UserSessionDto>> {
    return this.apiFn(
      "POST",
      `passport/signup/email?${this.paramsAsString}`,
      entity
    );
  }

  async postPassportSigninEmail(
    entity: workspaces.EmailAccountSigninDto
  ): Promise<IResponse<workspaces.UserSessionDto>> {
    return this.apiFn(
      "POST",
      `passport/signin/email?${this.paramsAsString}`,
      entity
    );
  }

  async postPassportAuthorizeOs(
    entity: workspaces.EmailAccountSigninDto
  ): Promise<IResponse<workspaces.UserSessionDto>> {
    return this.apiFn(
      "POST",
      `passport/authorizeOs?action=PassportActionAuthorizeOs&${this.paramsAsString}`,
      entity
    );
  }

  async postPassportRequestResetMailPassword(
    entity: workspaces.OtpAuthenticateDto
  ): Promise<IResponse<workspaces.EmailOtpResponse>> {
    return this.apiFn(
      "POST",
      `passport/request-reset-mail-password?${this.paramsAsString}`,
      entity
    );
  }
}
