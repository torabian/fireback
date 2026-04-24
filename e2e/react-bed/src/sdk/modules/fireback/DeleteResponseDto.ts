// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * The base class definition for deleteResponseDto
 **/
export class DeleteResponseDto {
  /**
   *
   * @type {DeleteResponseDto.Data}
   **/
  #data!: InstanceType<typeof DeleteResponseDto.Data>;
  /**
   *
   * @returns {DeleteResponseDto.Data}
   **/
  get data() {
    return this.#data;
  }
  /**
   *
   * @type {DeleteResponseDto.Data}
   **/
  set data(value: InstanceType<typeof DeleteResponseDto.Data>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof DeleteResponseDto.Data) {
      this.#data = value;
    } else {
      this.#data = new DeleteResponseDto.Data(value);
    }
  }
  setData(value: InstanceType<typeof DeleteResponseDto.Data>) {
    this.data = value;
    return this;
  }
  /**
   * The base class definition for data
   **/
  static Data = class Data {
    /**
     *
     * @type {DeleteResponseDto.Data.Item}
     **/
    #item!: InstanceType<typeof DeleteResponseDto.Data.Item>;
    /**
     *
     * @returns {DeleteResponseDto.Data.Item}
     **/
    get item() {
      return this.#item;
    }
    /**
     *
     * @type {DeleteResponseDto.Data.Item}
     **/
    set item(value: InstanceType<typeof DeleteResponseDto.Data.Item>) {
      // For objects, the sub type needs to always be instance of the sub class.
      if (value instanceof DeleteResponseDto.Data.Item) {
        this.#item = value;
      } else {
        this.#item = new DeleteResponseDto.Data.Item(value);
      }
    }
    setItem(value: InstanceType<typeof DeleteResponseDto.Data.Item>) {
      this.item = value;
      return this;
    }
    /**
     * The base class definition for item
     **/
    static Item = class Item {
      /**
       * If the deletion executed immediately.
       * @type {boolean}
       **/
      #executed!: boolean;
      /**
       * If the deletion executed immediately.
       * @returns {boolean}
       **/
      get executed() {
        return this.#executed;
      }
      /**
       * If the deletion executed immediately.
       * @type {boolean}
       **/
      set executed(value: boolean) {
        this.#executed = Boolean(value);
      }
      setExecuted(value: boolean) {
        this.executed = value;
        return this;
      }
      /**
       * The query selector which would be used to delete the content.
       * @type {number}
       **/
      #rowsAffected: number = 0;
      /**
       * The query selector which would be used to delete the content.
       * @returns {number}
       **/
      get rowsAffected() {
        return this.#rowsAffected;
      }
      /**
       * The query selector which would be used to delete the content.
       * @type {number}
       **/
      set rowsAffected(value: number) {
        const correctType = typeof value === "number";
        const parsedValue = correctType ? value : Number(value);
        if (!Number.isNaN(parsedValue)) {
          this.#rowsAffected = parsedValue;
        }
      }
      setRowsAffected(value: number) {
        this.rowsAffected = value;
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
        const d = data as Partial<Item>;
        if (d.executed !== undefined) {
          this.executed = d.executed;
        }
        if (d.rowsAffected !== undefined) {
          this.rowsAffected = d.rowsAffected;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          executed: this.#executed,
          rowsAffected: this.#rowsAffected,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          executed: "executed",
          rowsAffected: "rowsAffected",
        };
      }
      /**
       * Creates an instance of DeleteResponseDto.Data.Item, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(possibleDtoObject: DeleteResponseDtoType.DataType.ItemType) {
        return new DeleteResponseDto.Data.Item(possibleDtoObject);
      }
      /**
       * Creates an instance of DeleteResponseDto.Data.Item, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<DeleteResponseDtoType.DataType.ItemType>,
      ) {
        return new DeleteResponseDto.Data.Item(partialDtoObject);
      }
      copyWith(
        partial: PartialDeep<DeleteResponseDtoType.DataType.ItemType>,
      ): InstanceType<typeof DeleteResponseDto.Data.Item> {
        return new DeleteResponseDto.Data.Item({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<typeof DeleteResponseDto.Data.Item> {
        return new DeleteResponseDto.Data.Item(this.toJSON());
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
      const d = data as Partial<Data>;
      if (d.item !== undefined) {
        this.item = d.item;
      }
      this.#lateInitFields(data);
    }
    /**
     * These are the class instances, which need to be initialised, regardless of the constructor incoming data
     **/
    #lateInitFields(data = {}) {
      const d = data as Partial<Data>;
      if (!(d.item instanceof DeleteResponseDto.Data.Item)) {
        this.item = new DeleteResponseDto.Data.Item(d.item || {});
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        item: this.#item,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        item$: "item",
        get item() {
          return withPrefix("data.item", DeleteResponseDto.Data.Item.Fields);
        },
      };
    }
    /**
     * Creates an instance of DeleteResponseDto.Data, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: DeleteResponseDtoType.DataType) {
      return new DeleteResponseDto.Data(possibleDtoObject);
    }
    /**
     * Creates an instance of DeleteResponseDto.Data, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<DeleteResponseDtoType.DataType>) {
      return new DeleteResponseDto.Data(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<DeleteResponseDtoType.DataType>,
    ): InstanceType<typeof DeleteResponseDto.Data> {
      return new DeleteResponseDto.Data({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof DeleteResponseDto.Data> {
      return new DeleteResponseDto.Data(this.toJSON());
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
    const d = data as Partial<DeleteResponseDto>;
    if (d.data !== undefined) {
      this.data = d.data;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<DeleteResponseDto>;
    if (!(d.data instanceof DeleteResponseDto.Data)) {
      this.data = new DeleteResponseDto.Data(d.data || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      data: this.#data,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      data$: "data",
      get data() {
        return withPrefix("data", DeleteResponseDto.Data.Fields);
      },
    };
  }
  /**
   * Creates an instance of DeleteResponseDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: DeleteResponseDtoType) {
    return new DeleteResponseDto(possibleDtoObject);
  }
  /**
   * Creates an instance of DeleteResponseDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<DeleteResponseDtoType>) {
    return new DeleteResponseDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<DeleteResponseDtoType>,
  ): InstanceType<typeof DeleteResponseDto> {
    return new DeleteResponseDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof DeleteResponseDto> {
    return new DeleteResponseDto(this.toJSON());
  }
}
export abstract class DeleteResponseDtoFactory {
  abstract create(data: unknown): DeleteResponseDto;
}
/**
 * The base type definition for deleteResponseDto
 **/
export type DeleteResponseDtoType = {
  /**
   *
   * @type {DeleteResponseDtoType.DataType}
   **/
  data: DeleteResponseDtoType.DataType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DeleteResponseDtoType {
  /**
   * The base type definition for dataType
   **/
  export type DataType = {
    /**
     *
     * @type {DeleteResponseDtoType.DataType.ItemType}
     **/
    item: DeleteResponseDtoType.DataType.ItemType;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace DataType {
    /**
     * The base type definition for itemType
     **/
    export type ItemType = {
      /**
       * If the deletion executed immediately.
       * @type {boolean}
       **/
      executed: boolean;
      /**
       * The query selector which would be used to delete the content.
       * @type {number}
       **/
      rowsAffected: number;
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace ItemType {}
  }
}
