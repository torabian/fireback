import { withPrefix } from "./sdk/common/withPrefix";
/**
 * The base class definition for responseDto
 **/
export class ResponseDto<T> {
  /**
   * Version of the API used for this response.
   * @type {string}
   **/
  #apiVersion?: string | null = undefined;
  /**
   * Version of the API used for this response.
   * @returns {string}
   **/
  get apiVersion() {
    return this.#apiVersion;
  }
  /**
   * Version of the API used for this response.
   * @type {string}
   **/
  set apiVersion(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#apiVersion = correctType ? value : String(value);
  }
  setApiVersion(value: string | null | undefined) {
    this.apiVersion = value;
    return this;
  }
  /**
   * Context string provided by the client or system for request tracking.
   * @type {string}
   **/
  #context?: string | null = undefined;
  /**
   * Context string provided by the client or system for request tracking.
   * @returns {string}
   **/
  get context() {
    return this.#context;
  }
  /**
   * Context string provided by the client or system for request tracking.
   * @type {string}
   **/
  set context(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#context = correctType ? value : String(value);
  }
  setContext(value: string | null | undefined) {
    this.context = value;
    return this;
  }
  /**
   * Unique identifier assigned to the request/response.
   * @type {string}
   **/
  #id?: string | null = undefined;
  /**
   * Unique identifier assigned to the request/response.
   * @returns {string}
   **/
  get id() {
    return this.#id;
  }
  /**
   * Unique identifier assigned to the request/response.
   * @type {string}
   **/
  set id(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#id = correctType ? value : String(value);
  }
  setId(value: string | null | undefined) {
    this.id = value;
    return this;
  }
  /**
   * Name of the API method invoked.
   * @type {string}
   **/
  #method?: string | null = undefined;
  /**
   * Name of the API method invoked.
   * @returns {string}
   **/
  get method() {
    return this.#method;
  }
  /**
   * Name of the API method invoked.
   * @type {string}
   **/
  set method(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#method = correctType ? value : String(value);
  }
  setMethod(value: string | null | undefined) {
    this.method = value;
    return this;
  }
  /**
   * Parameters sent with the request.
   * @type {any}
   **/
  #params: any = null;
  /**
   * Parameters sent with the request.
   * @returns {any}
   **/
  get params() {
    return this.#params;
  }
  /**
   * Parameters sent with the request.
   * @type {any}
   **/
  set params(value: any) {
    this.#params = value;
  }
  setParams(value: any) {
    this.params = value;
    return this;
  }
  /**
   * Main data payload of the response.
   * @type {ResponseDto.Data<T>}
   **/
  #data!: InstanceType<typeof ResponseDto.Data<T>>;
  /**
   * Main data payload of the response.
   * @returns {ResponseDto.Data<T>}
   **/
  get data() {
    return this.#data;
  }
  /**
   * Main data payload of the response.
   * @type {ResponseDto.Data}
   **/
  set data(value: InstanceType<typeof ResponseDto.Data<T>>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof ResponseDto.Data) {
      this.#data = value;
    } else {
      this.#data = new ResponseDto.Data(value);
    }
  }
  setData(value: InstanceType<typeof ResponseDto.Data<T>>) {
    this.data = value;
    return this;
  }
  /**
   * Error details, if the request failed.
   * @type {ResponseDto.Error}
   **/
  #error!: InstanceType<typeof ResponseDto.Error>;
  /**
   * Error details, if the request failed.
   * @returns {ResponseDto.Error}
   **/
  get error() {
    return this.#error;
  }
  /**
   * Error details, if the request failed.
   * @type {ResponseDto.Error}
   **/
  set error(value: InstanceType<typeof ResponseDto.Error>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof ResponseDto.Error) {
      this.#error = value;
    } else {
      this.#error = new ResponseDto.Error(value);
    }
  }
  setError(value: InstanceType<typeof ResponseDto.Error>) {
    this.error = value;
    return this;
  }
  /**
   * The base class definition for data
   **/
  static Data = class Data<T> {
    /**
     * Single item returned by the API.
     * @type {any}
     **/
    #item: T = null;
    /**
     * Single item returned by the API.
     * @returns {T}
     **/
    get item(): T {
      return this.#item;
    }
    /**
     * Single item returned by the API.
     * @type {any}
     **/
    set item(value: any) {
      this.#item = value;
    }
    setItem(value: any) {
      this.item = value;
      return this;
    }
    /**
     * List of items returned by the API.
     * @type {any}
     **/
    #items: any = null;
    /**
     * List of items returned by the API.
     * @returns {any}
     **/
    get items() {
      return this.#items;
    }
    /**
     * List of items returned by the API.
     * @type {any}
     **/
    set items(value: any) {
      this.#items = value;
    }
    setItems(value: any) {
      this.items = value;
      return this;
    }
    /**
     * Link to edit this resource.
     * @type {string}
     **/
    #editLink?: string | null = undefined;
    /**
     * Link to edit this resource.
     * @returns {string}
     **/
    get editLink() {
      return this.#editLink;
    }
    /**
     * Link to edit this resource.
     * @type {string}
     **/
    set editLink(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#editLink = correctType ? value : String(value);
    }
    setEditLink(value: string | null | undefined) {
      this.editLink = value;
      return this;
    }
    /**
     * Link to retrieve this resource.
     * @type {string}
     **/
    #selfLink?: string | null = undefined;
    /**
     * Link to retrieve this resource.
     * @returns {string}
     **/
    get selfLink() {
      return this.#selfLink;
    }
    /**
     * Link to retrieve this resource.
     * @type {string}
     **/
    set selfLink(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#selfLink = correctType ? value : String(value);
    }
    setSelfLink(value: string | null | undefined) {
      this.selfLink = value;
      return this;
    }
    /**
     * Resource type (kind) identifier.
     * @type {string}
     **/
    #kind?: string | null = undefined;
    /**
     * Resource type (kind) identifier.
     * @returns {string}
     **/
    get kind() {
      return this.#kind;
    }
    /**
     * Resource type (kind) identifier.
     * @type {string}
     **/
    set kind(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#kind = correctType ? value : String(value);
    }
    setKind(value: string | null | undefined) {
      this.kind = value;
      return this;
    }
    /**
     * Selector specifying which fields are included in a partial response.
     * @type {string}
     **/
    #fields?: string | null = undefined;
    /**
     * Selector specifying which fields are included in a partial response.
     * @returns {string}
     **/
    get fields() {
      return this.#fields;
    }
    /**
     * Selector specifying which fields are included in a partial response.
     * @type {string}
     **/
    set fields(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#fields = correctType ? value : String(value);
    }
    setFields(value: string | null | undefined) {
      this.fields = value;
      return this;
    }
    /**
     * ETag of the resource, used for caching/version control.
     * @type {string}
     **/
    #etag?: string | null = undefined;
    /**
     * ETag of the resource, used for caching/version control.
     * @returns {string}
     **/
    get etag() {
      return this.#etag;
    }
    /**
     * ETag of the resource, used for caching/version control.
     * @type {string}
     **/
    set etag(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#etag = correctType ? value : String(value);
    }
    setEtag(value: string | null | undefined) {
      this.etag = value;
      return this;
    }
    /**
     * Cursor for paginated data fetching.
     * @type {string}
     **/
    #cursor?: string | null = undefined;
    /**
     * Cursor for paginated data fetching.
     * @returns {string}
     **/
    get cursor() {
      return this.#cursor;
    }
    /**
     * Cursor for paginated data fetching.
     * @type {string}
     **/
    set cursor(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#cursor = correctType ? value : String(value);
    }
    setCursor(value: string | null | undefined) {
      this.cursor = value;
      return this;
    }
    /**
     * Unique identifier of the resource.
     * @type {string}
     **/
    #id?: string | null = undefined;
    /**
     * Unique identifier of the resource.
     * @returns {string}
     **/
    get id() {
      return this.#id;
    }
    /**
     * Unique identifier of the resource.
     * @type {string}
     **/
    set id(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#id = correctType ? value : String(value);
    }
    setId(value: string | null | undefined) {
      this.id = value;
      return this;
    }
    /**
     * Language code of the response data.
     * @type {string}
     **/
    #lang?: string | null = undefined;
    /**
     * Language code of the response data.
     * @returns {string}
     **/
    get lang() {
      return this.#lang;
    }
    /**
     * Language code of the response data.
     * @type {string}
     **/
    set lang(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#lang = correctType ? value : String(value);
    }
    setLang(value: string | null | undefined) {
      this.lang = value;
      return this;
    }
    /**
     * Last modification time of the resource.
     * @type {string}
     **/
    #updated?: string | null = undefined;
    /**
     * Last modification time of the resource.
     * @returns {string}
     **/
    get updated() {
      return this.#updated;
    }
    /**
     * Last modification time of the resource.
     * @type {string}
     **/
    set updated(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#updated = correctType ? value : String(value);
    }
    setUpdated(value: string | null | undefined) {
      this.updated = value;
      return this;
    }
    /**
     * Number of items in the current response page.
     * @type {number}
     **/
    #currentItemCount?: number | null = undefined;
    /**
     * Number of items in the current response page.
     * @returns {number}
     **/
    get currentItemCount() {
      return this.#currentItemCount;
    }
    /**
     * Number of items in the current response page.
     * @type {number}
     **/
    set currentItemCount(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#currentItemCount = parsedValue;
      }
    }
    setCurrentItemCount(value: number | null | undefined) {
      this.currentItemCount = value;
      return this;
    }
    /**
     * Maximum number of items per page.
     * @type {number}
     **/
    #itemsPerPage?: number | null = undefined;
    /**
     * Maximum number of items per page.
     * @returns {number}
     **/
    get itemsPerPage() {
      return this.#itemsPerPage;
    }
    /**
     * Maximum number of items per page.
     * @type {number}
     **/
    set itemsPerPage(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#itemsPerPage = parsedValue;
      }
    }
    setItemsPerPage(value: number | null | undefined) {
      this.itemsPerPage = value;
      return this;
    }
    /**
     * Index of the first item in the current page.
     * @type {number}
     **/
    #startIndex?: number | null = undefined;
    /**
     * Index of the first item in the current page.
     * @returns {number}
     **/
    get startIndex() {
      return this.#startIndex;
    }
    /**
     * Index of the first item in the current page.
     * @type {number}
     **/
    set startIndex(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#startIndex = parsedValue;
      }
    }
    setStartIndex(value: number | null | undefined) {
      this.startIndex = value;
      return this;
    }
    /**
     * Total number of items available.
     * @type {number}
     **/
    #totalItems?: number | null = undefined;
    /**
     * Total number of items available.
     * @returns {number}
     **/
    get totalItems() {
      return this.#totalItems;
    }
    /**
     * Total number of items available.
     * @type {number}
     **/
    set totalItems(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#totalItems = parsedValue;
      }
    }
    setTotalItems(value: number | null | undefined) {
      this.totalItems = value;
      return this;
    }
    /**
     * Number of items available for this user/query.
     * @type {number}
     **/
    #totalAvailableItems?: number | null = undefined;
    /**
     * Number of items available for this user/query.
     * @returns {number}
     **/
    get totalAvailableItems() {
      return this.#totalAvailableItems;
    }
    /**
     * Number of items available for this user/query.
     * @type {number}
     **/
    set totalAvailableItems(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#totalAvailableItems = parsedValue;
      }
    }
    setTotalAvailableItems(value: number | null | undefined) {
      this.totalAvailableItems = value;
      return this;
    }
    /**
     * Current page index in the pagination.
     * @type {number}
     **/
    #pageIndex?: number | null = undefined;
    /**
     * Current page index in the pagination.
     * @returns {number}
     **/
    get pageIndex() {
      return this.#pageIndex;
    }
    /**
     * Current page index in the pagination.
     * @type {number}
     **/
    set pageIndex(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#pageIndex = parsedValue;
      }
    }
    setPageIndex(value: number | null | undefined) {
      this.pageIndex = value;
      return this;
    }
    /**
     * Total number of pages in the pagination.
     * @type {number}
     **/
    #totalPages?: number | null = undefined;
    /**
     * Total number of pages in the pagination.
     * @returns {number}
     **/
    get totalPages() {
      return this.#totalPages;
    }
    /**
     * Total number of pages in the pagination.
     * @type {number}
     **/
    set totalPages(value: number | null | undefined) {
      const correctType =
        typeof value === "number" || value === undefined || value === null;
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#totalPages = parsedValue;
      }
    }
    setTotalPages(value: number | null | undefined) {
      this.totalPages = value;
      return this;
    }
    constructor(data: unknown = undefined) {
      if (data === null || data === undefined) {
        return;
      }
      if (typeof data === "string") {
        this.applyFromObject(JSON.parse(data));
      } else if (this.#isJsonAppliable(data)) {
        this.applyFromObject(data);
      } else {
        throw new Error(
          "Instance cannot be created on an unknown value, check the content being passed. got: " +
            typeof data,
        );
      }
    }
    #isJsonAppliable(obj: unknown) {
      const g = globalThis as unknown as { Buffer: any; Blob: any };
      const isBuffer =
        typeof g.Buffer !== "undefined" &&
        typeof g.Buffer.isBuffer === "function" &&
        g.Buffer.isBuffer(obj);
      const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
      return (
        obj &&
        typeof obj === "object" &&
        !Array.isArray(obj) &&
        !isBuffer &&
        !(obj instanceof ArrayBuffer) &&
        !isBlob
      );
    }
    /**
     * casts the fields of a javascript object into the class properties one by one
     **/
    applyFromObject(data = {}) {
      const d = data as Partial<Data<T>>;
      if (d.item !== undefined) {
        this.item = d.item;
      }
      if (d.items !== undefined) {
        this.items = d.items;
      }
      if (d.editLink !== undefined) {
        this.editLink = d.editLink;
      }
      if (d.selfLink !== undefined) {
        this.selfLink = d.selfLink;
      }
      if (d.kind !== undefined) {
        this.kind = d.kind;
      }
      if (d.fields !== undefined) {
        this.fields = d.fields;
      }
      if (d.etag !== undefined) {
        this.etag = d.etag;
      }
      if (d.cursor !== undefined) {
        this.cursor = d.cursor;
      }
      if (d.id !== undefined) {
        this.id = d.id;
      }
      if (d.lang !== undefined) {
        this.lang = d.lang;
      }
      if (d.updated !== undefined) {
        this.updated = d.updated;
      }
      if (d.currentItemCount !== undefined) {
        this.currentItemCount = d.currentItemCount;
      }
      if (d.itemsPerPage !== undefined) {
        this.itemsPerPage = d.itemsPerPage;
      }
      if (d.startIndex !== undefined) {
        this.startIndex = d.startIndex;
      }
      if (d.totalItems !== undefined) {
        this.totalItems = d.totalItems;
      }
      if (d.totalAvailableItems !== undefined) {
        this.totalAvailableItems = d.totalAvailableItems;
      }
      if (d.pageIndex !== undefined) {
        this.pageIndex = d.pageIndex;
      }
      if (d.totalPages !== undefined) {
        this.totalPages = d.totalPages;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        item: this.#item,
        items: this.#items,
        editLink: this.#editLink,
        selfLink: this.#selfLink,
        kind: this.#kind,
        fields: this.#fields,
        etag: this.#etag,
        cursor: this.#cursor,
        id: this.#id,
        lang: this.#lang,
        updated: this.#updated,
        currentItemCount: this.#currentItemCount,
        itemsPerPage: this.#itemsPerPage,
        startIndex: this.#startIndex,
        totalItems: this.#totalItems,
        totalAvailableItems: this.#totalAvailableItems,
        pageIndex: this.#pageIndex,
        totalPages: this.#totalPages,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item: "item",
        items: "items",
        editLink: "editLink",
        selfLink: "selfLink",
        kind: "kind",
        fields: "fields",
        etag: "etag",
        cursor: "cursor",
        id: "id",
        lang: "lang",
        updated: "updated",
        currentItemCount: "currentItemCount",
        itemsPerPage: "itemsPerPage",
        startIndex: "startIndex",
        totalItems: "totalItems",
        totalAvailableItems: "totalAvailableItems",
        pageIndex: "pageIndex",
        totalPages: "totalPages",
      };
    }
    /**
     * Creates an instance of ResponseDto.Data, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from<T = unknown>(possibleDtoObject: ResponseDtoType.DataType<T>) {
      return new ResponseDto.Data(possibleDtoObject);
    }
    /**
     * Creates an instance of ResponseDto.Data, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with<T = unknown>(
      partialDtoObject: PartialDeep<ResponseDtoType.DataType<T>>,
    ) {
      return new ResponseDto.Data(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<ResponseDtoType.DataType<unknown>>,
    ): InstanceType<typeof ResponseDto.Data> {
      return new ResponseDto.Data({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof ResponseDto.Data> {
      return new ResponseDto.Data(this.toJSON());
    }
  };
  /**
   * The base class definition for error
   **/
  static Error = class Error {
    /**
     * Numeric error code representing the failure.
     * @type {number}
     **/
    #code: number = 0;
    /**
     * Numeric error code representing the failure.
     * @returns {number}
     **/
    get code() {
      return this.#code;
    }
    /**
     * Numeric error code representing the failure.
     * @type {number}
     **/
    set code(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#code = parsedValue;
      }
    }
    setCode(value: number) {
      this.code = value;
      return this;
    }
    /**
     * Human-readable explanation of the error.
     * @type {string}
     **/
    #message: string = "";
    /**
     * Human-readable explanation of the error.
     * @returns {string}
     **/
    get message() {
      return this.#message;
    }
    /**
     * Human-readable explanation of the error.
     * @type {string}
     **/
    set message(value: string) {
      this.#message = String(value);
    }
    setMessage(value: string) {
      this.message = value;
      return this;
    }
    /**
     * Localized/translated version of the error message.
     * @type {string}
     **/
    #messageTranslated: string = "";
    /**
     * Localized/translated version of the error message.
     * @returns {string}
     **/
    get messageTranslated() {
      return this.#messageTranslated;
    }
    /**
     * Localized/translated version of the error message.
     * @type {string}
     **/
    set messageTranslated(value: string) {
      this.#messageTranslated = String(value);
    }
    setMessageTranslated(value: string) {
      this.messageTranslated = value;
      return this;
    }
    /**
     * Detailed list of error objects.
     * @type {ResponseDto.Error.Errors}
     **/
    #errors: InstanceType<typeof ResponseDto.Error.Errors>[] = [];
    /**
     * Detailed list of error objects.
     * @returns {ResponseDto.Error.Errors}
     **/
    get errors() {
      return this.#errors;
    }
    /**
     * Detailed list of error objects.
     * @type {ResponseDto.Error.Errors}
     **/
    set errors(value: InstanceType<typeof ResponseDto.Error.Errors>[]) {
      // For arrays, you only can pass arrays to the object
      if (!Array.isArray(value)) {
        return;
      }
      if (value.length > 0 && value[0] instanceof ResponseDto.Error.Errors) {
        this.#errors = value;
      } else {
        this.#errors = value.map((item) => new ResponseDto.Error.Errors(item));
      }
    }
    setErrors(value: InstanceType<typeof ResponseDto.Error.Errors>[]) {
      this.errors = value;
      return this;
    }
    /**
     * The base class definition for errors
     **/
    static Errors = class Errors {
      /**
       * Logical grouping of the error (e.g., global, usageLimits).
       * @type {string}
       **/
      #domain?: string | null = undefined;
      /**
       * Logical grouping of the error (e.g., global, usageLimits).
       * @returns {string}
       **/
      get domain() {
        return this.#domain;
      }
      /**
       * Logical grouping of the error (e.g., global, usageLimits).
       * @type {string}
       **/
      set domain(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#domain = correctType ? value : String(value);
      }
      setDomain(value: string | null | undefined) {
        this.domain = value;
        return this;
      }
      /**
       * Reason identifier for the error.
       * @type {string}
       **/
      #reason?: string | null = undefined;
      /**
       * Reason identifier for the error.
       * @returns {string}
       **/
      get reason() {
        return this.#reason;
      }
      /**
       * Reason identifier for the error.
       * @type {string}
       **/
      set reason(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#reason = correctType ? value : String(value);
      }
      setReason(value: string | null | undefined) {
        this.reason = value;
        return this;
      }
      /**
       * Human-readable explanation of the sub-error.
       * @type {string}
       **/
      #message?: string | null = undefined;
      /**
       * Human-readable explanation of the sub-error.
       * @returns {string}
       **/
      get message() {
        return this.#message;
      }
      /**
       * Human-readable explanation of the sub-error.
       * @type {string}
       **/
      set message(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#message = correctType ? value : String(value);
      }
      setMessage(value: string | null | undefined) {
        this.message = value;
        return this;
      }
      /**
       * Localized/translated version of the sub-error message.
       * @type {string}
       **/
      #messageTranslated?: string | null = undefined;
      /**
       * Localized/translated version of the sub-error message.
       * @returns {string}
       **/
      get messageTranslated() {
        return this.#messageTranslated;
      }
      /**
       * Localized/translated version of the sub-error message.
       * @type {string}
       **/
      set messageTranslated(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#messageTranslated = correctType ? value : String(value);
      }
      setMessageTranslated(value: string | null | undefined) {
        this.messageTranslated = value;
        return this;
      }
      /**
       * Field or parameter in which the error occurred.
       * @type {string}
       **/
      #location?: string | null = undefined;
      /**
       * Field or parameter in which the error occurred.
       * @returns {string}
       **/
      get location() {
        return this.#location;
      }
      /**
       * Field or parameter in which the error occurred.
       * @type {string}
       **/
      set location(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#location = correctType ? value : String(value);
      }
      setLocation(value: string | null | undefined) {
        this.location = value;
        return this;
      }
      /**
       * Type of location (e.g., parameter, header).
       * @type {string}
       **/
      #locationType?: string | null = undefined;
      /**
       * Type of location (e.g., parameter, header).
       * @returns {string}
       **/
      get locationType() {
        return this.#locationType;
      }
      /**
       * Type of location (e.g., parameter, header).
       * @type {string}
       **/
      set locationType(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#locationType = correctType ? value : String(value);
      }
      setLocationType(value: string | null | undefined) {
        this.locationType = value;
        return this;
      }
      /**
       * URL linking to additional documentation about the error.
       * @type {string}
       **/
      #extendedHelp?: string | null = undefined;
      /**
       * URL linking to additional documentation about the error.
       * @returns {string}
       **/
      get extendedHelp() {
        return this.#extendedHelp;
      }
      /**
       * URL linking to additional documentation about the error.
       * @type {string}
       **/
      set extendedHelp(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#extendedHelp = correctType ? value : String(value);
      }
      setExtendedHelp(value: string | null | undefined) {
        this.extendedHelp = value;
        return this;
      }
      /**
       * URL to submit a report for this error.
       * @type {string}
       **/
      #sendReport?: string | null = undefined;
      /**
       * URL to submit a report for this error.
       * @returns {string}
       **/
      get sendReport() {
        return this.#sendReport;
      }
      /**
       * URL to submit a report for this error.
       * @type {string}
       **/
      set sendReport(value: string | null | undefined) {
        const correctType =
          typeof value === "string" || value === undefined || value === null;
        this.#sendReport = correctType ? value : String(value);
      }
      setSendReport(value: string | null | undefined) {
        this.sendReport = value;
        return this;
      }
      constructor(data: unknown = undefined) {
        if (data === null || data === undefined) {
          return;
        }
        if (typeof data === "string") {
          this.applyFromObject(JSON.parse(data));
        } else if (this.#isJsonAppliable(data)) {
          this.applyFromObject(data);
        } else {
          throw new Error(
            "Instance cannot be created on an unknown value, check the content being passed. got: " +
              typeof data,
          );
        }
      }
      #isJsonAppliable(obj: unknown) {
        const g = globalThis as unknown as { Buffer: any; Blob: any };
        const isBuffer =
          typeof g.Buffer !== "undefined" &&
          typeof g.Buffer.isBuffer === "function" &&
          g.Buffer.isBuffer(obj);
        const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
        return (
          obj &&
          typeof obj === "object" &&
          !Array.isArray(obj) &&
          !isBuffer &&
          !(obj instanceof ArrayBuffer) &&
          !isBlob
        );
      }
      /**
       * casts the fields of a javascript object into the class properties one by one
       **/
      applyFromObject(data = {}) {
        const d = data as Partial<Errors>;
        if (d.domain !== undefined) {
          this.domain = d.domain;
        }
        if (d.reason !== undefined) {
          this.reason = d.reason;
        }
        if (d.message !== undefined) {
          this.message = d.message;
        }
        if (d.messageTranslated !== undefined) {
          this.messageTranslated = d.messageTranslated;
        }
        if (d.location !== undefined) {
          this.location = d.location;
        }
        if (d.locationType !== undefined) {
          this.locationType = d.locationType;
        }
        if (d.extendedHelp !== undefined) {
          this.extendedHelp = d.extendedHelp;
        }
        if (d.sendReport !== undefined) {
          this.sendReport = d.sendReport;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          domain: this.#domain,
          reason: this.#reason,
          message: this.#message,
          messageTranslated: this.#messageTranslated,
          location: this.#location,
          locationType: this.#locationType,
          extendedHelp: this.#extendedHelp,
          sendReport: this.#sendReport,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          domain: "domain",
          reason: "reason",
          message: "message",
          messageTranslated: "messageTranslated",
          location: "location",
          locationType: "locationType",
          extendedHelp: "extendedHelp",
          sendReport: "sendReport",
        };
      }
      /**
       * Creates an instance of ResponseDto.Error.Errors, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(possibleDtoObject: ResponseDtoType.ErrorType.ErrorsType) {
        return new ResponseDto.Error.Errors(possibleDtoObject);
      }
      /**
       * Creates an instance of ResponseDto.Error.Errors, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<ResponseDtoType.ErrorType.ErrorsType>,
      ) {
        return new ResponseDto.Error.Errors(partialDtoObject);
      }
      copyWith(
        partial: PartialDeep<ResponseDtoType.ErrorType.ErrorsType>,
      ): InstanceType<typeof ResponseDto.Error.Errors> {
        return new ResponseDto.Error.Errors({ ...this.toJSON(), ...partial });
      }
      clone(): InstanceType<typeof ResponseDto.Error.Errors> {
        return new ResponseDto.Error.Errors(this.toJSON());
      }
    };
    constructor(data: unknown = undefined) {
      if (data === null || data === undefined) {
        return;
      }
      if (typeof data === "string") {
        this.applyFromObject(JSON.parse(data));
      } else if (this.#isJsonAppliable(data)) {
        this.applyFromObject(data);
      } else {
        throw new Error(
          "Instance cannot be created on an unknown value, check the content being passed. got: " +
            typeof data,
        );
      }
    }
    #isJsonAppliable(obj: unknown) {
      const g = globalThis as unknown as { Buffer: any; Blob: any };
      const isBuffer =
        typeof g.Buffer !== "undefined" &&
        typeof g.Buffer.isBuffer === "function" &&
        g.Buffer.isBuffer(obj);
      const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
      return (
        obj &&
        typeof obj === "object" &&
        !Array.isArray(obj) &&
        !isBuffer &&
        !(obj instanceof ArrayBuffer) &&
        !isBlob
      );
    }
    /**
     * casts the fields of a javascript object into the class properties one by one
     **/
    applyFromObject(data = {}) {
      const d = data as Partial<Error>;
      if (d.code !== undefined) {
        this.code = d.code;
      }
      if (d.message !== undefined) {
        this.message = d.message;
      }
      if (d.messageTranslated !== undefined) {
        this.messageTranslated = d.messageTranslated;
      }
      if (d.errors !== undefined) {
        this.errors = d.errors;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        code: this.#code,
        message: this.#message,
        messageTranslated: this.#messageTranslated,
        errors: this.#errors,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        code: "code",
        message: "message",
        messageTranslated: "messageTranslated",
        errors$: "errors",
        get errors() {
          return withPrefix(
            "error.errors[:i]",
            ResponseDto.Error.Errors.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of ResponseDto.Error, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: ResponseDtoType.ErrorType) {
      return new ResponseDto.Error(possibleDtoObject);
    }
    /**
     * Creates an instance of ResponseDto.Error, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<ResponseDtoType.ErrorType>) {
      return new ResponseDto.Error(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<ResponseDtoType.ErrorType>,
    ): InstanceType<typeof ResponseDto.Error> {
      return new ResponseDto.Error({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof ResponseDto.Error> {
      return new ResponseDto.Error(this.toJSON());
    }
  };
  constructor(data: unknown = undefined) {
    if (data === null || data === undefined) {
      this.#lateInitFields();
      return;
    }
    if (typeof data === "string") {
      this.applyFromObject(JSON.parse(data));
    } else if (this.#isJsonAppliable(data)) {
      this.applyFromObject(data);
    } else {
      throw new Error(
        "Instance cannot be created on an unknown value, check the content being passed. got: " +
          typeof data,
      );
    }
  }
  #isJsonAppliable(obj: unknown) {
    const g = globalThis as unknown as { Buffer: any; Blob: any };
    const isBuffer =
      typeof g.Buffer !== "undefined" &&
      typeof g.Buffer.isBuffer === "function" &&
      g.Buffer.isBuffer(obj);
    const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
    return (
      obj &&
      typeof obj === "object" &&
      !Array.isArray(obj) &&
      !isBuffer &&
      !(obj instanceof ArrayBuffer) &&
      !isBlob
    );
  }
  /**
   * casts the fields of a javascript object into the class properties one by one
   **/
  applyFromObject(data = {}) {
    const d = data as Partial<ResponseDto<T>>;
    if (d.apiVersion !== undefined) {
      this.apiVersion = d.apiVersion;
    }
    if (d.context !== undefined) {
      this.context = d.context;
    }
    if (d.id !== undefined) {
      this.id = d.id;
    }
    if (d.method !== undefined) {
      this.method = d.method;
    }
    if (d.params !== undefined) {
      this.params = d.params;
    }
    if (d.data !== undefined) {
      this.data = d.data;
    }
    if (d.error !== undefined) {
      this.error = d.error;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<ResponseDto<T>>;
    if (!(d.data instanceof ResponseDto.Data)) {
      this.data = new ResponseDto.Data(d.data || {});
    }
    if (!(d.error instanceof ResponseDto.Error)) {
      this.error = new ResponseDto.Error(d.error || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      apiVersion: this.#apiVersion,
      context: this.#context,
      id: this.#id,
      method: this.#method,
      params: this.#params,
      data: this.#data,
      error: this.#error,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      apiVersion: "apiVersion",
      context: "context",
      id: "id",
      method: "method",
      params: "params",
      data$: "data",
      get data() {
        return withPrefix("data", ResponseDto.Data.Fields);
      },
      error$: "error",
      get error() {
        return withPrefix("error", ResponseDto.Error.Fields);
      },
    };
  }
  /**
   * Creates an instance of ResponseDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from<T = unknown>(possibleDtoObject: ResponseDtoType<T>) {
    return new ResponseDto(possibleDtoObject);
  }
  /**
   * Creates an instance of ResponseDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with<T = unknown>(partialDtoObject: PartialDeep<ResponseDtoType<T>>) {
    return new ResponseDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ResponseDtoType<T>>,
  ): InstanceType<typeof ResponseDto> {
    return new ResponseDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ResponseDto> {
    return new ResponseDto(this.toJSON());
  }
}
export abstract class ResponseDtoFactory {
  abstract create(data: unknown): ResponseDto<unknown>;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for responseDto
 **/
export type ResponseDtoType<T> = {
  /**
   * Version of the API used for this response.
   * @type {string}
   **/
  apiVersion?: string;
  /**
   * Context string provided by the client or system for request tracking.
   * @type {string}
   **/
  context?: string;
  /**
   * Unique identifier assigned to the request/response.
   * @type {string}
   **/
  id?: string;
  /**
   * Name of the API method invoked.
   * @type {string}
   **/
  method?: string;
  /**
   * Parameters sent with the request.
   * @type {any}
   **/
  params: any;
  /**
   * Main data payload of the response.
   * @type {ResponseDtoType.DataType}
   **/
  data: ResponseDtoType.DataType<T>;
  /**
   * Error details, if the request failed.
   * @type {ResponseDtoType.ErrorType}
   **/
  error: ResponseDtoType.ErrorType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ResponseDtoType {
  /**
   * The base type definition for dataType
   **/
  export type DataType<T> = {
    /**
     * Single item returned by the API.
     * @type {any}
     **/
    item2: T;
    /**
     * List of items returned by the API.
     * @type {any}
     **/
    items: any;
    /**
     * Link to edit this resource.
     * @type {string}
     **/
    editLink?: string;
    /**
     * Link to retrieve this resource.
     * @type {string}
     **/
    selfLink?: string;
    /**
     * Resource type (kind) identifier.
     * @type {string}
     **/
    kind?: string;
    /**
     * Selector specifying which fields are included in a partial response.
     * @type {string}
     **/
    fields?: string;
    /**
     * ETag of the resource, used for caching/version control.
     * @type {string}
     **/
    etag?: string;
    /**
     * Cursor for paginated data fetching.
     * @type {string}
     **/
    cursor?: string;
    /**
     * Unique identifier of the resource.
     * @type {string}
     **/
    id?: string;
    /**
     * Language code of the response data.
     * @type {string}
     **/
    lang?: string;
    /**
     * Last modification time of the resource.
     * @type {string}
     **/
    updated?: string;
    /**
     * Number of items in the current response page.
     * @type {number}
     **/
    currentItemCount?: number;
    /**
     * Maximum number of items per page.
     * @type {number}
     **/
    itemsPerPage?: number;
    /**
     * Index of the first item in the current page.
     * @type {number}
     **/
    startIndex?: number;
    /**
     * Total number of items available.
     * @type {number}
     **/
    totalItems?: number;
    /**
     * Number of items available for this user/query.
     * @type {number}
     **/
    totalAvailableItems?: number;
    /**
     * Current page index in the pagination.
     * @type {number}
     **/
    pageIndex?: number;
    /**
     * Total number of pages in the pagination.
     * @type {number}
     **/
    totalPages?: number;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace DataType {}
  /**
   * The base type definition for errorType
   **/
  export type ErrorType = {
    /**
     * Numeric error code representing the failure.
     * @type {number}
     **/
    code: number;
    /**
     * Human-readable explanation of the error.
     * @type {string}
     **/
    message: string;
    /**
     * Localized/translated version of the error message.
     * @type {string}
     **/
    messageTranslated: string;
    /**
     * Detailed list of error objects.
     * @type {ResponseDtoType.ErrorType.ErrorsType[]}
     **/
    errors: ResponseDtoType.ErrorType.ErrorsType[];
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ErrorType {
    /**
     * The base type definition for errorsType
     **/
    export type ErrorsType = {
      /**
       * Logical grouping of the error (e.g., global, usageLimits).
       * @type {string}
       **/
      domain?: string;
      /**
       * Reason identifier for the error.
       * @type {string}
       **/
      reason?: string;
      /**
       * Human-readable explanation of the sub-error.
       * @type {string}
       **/
      message?: string;
      /**
       * Localized/translated version of the sub-error message.
       * @type {string}
       **/
      messageTranslated?: string;
      /**
       * Field or parameter in which the error occurred.
       * @type {string}
       **/
      location?: string;
      /**
       * Type of location (e.g., parameter, header).
       * @type {string}
       **/
      locationType?: string;
      /**
       * URL linking to additional documentation about the error.
       * @type {string}
       **/
      extendedHelp?: string;
      /**
       * URL to submit a report for this error.
       * @type {string}
       **/
      sendReport?: string;
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace ErrorsType {}
  }
}
