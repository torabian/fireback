import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { MArray } from "@fireback/js-remote-ctx/common/operators";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
import { withPrefix } from "@fireback/js-remote-ctx/common/withPrefix";
/**
 * Action to communicate with the action InternalStatsSnapshot
 */
export type InternalStatsSnapshotActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type InternalStatsSnapshotActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<InternalStatsSnapshotActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  InternalStatsSnapshotActionOptions &
  Partial<{
    creatorFn: (item: unknown) => InternalStatsSnapshotActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useInternalStatsSnapshotActionQuery = (
  options: InternalStatsSnapshotActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return InternalStatsSnapshotAction.Fetch(
      {
        headers: options?.headers,
      },
      {
        creatorFn: options?.creatorFn,
        qs: options?.qs,
        ctx,
        onMessage: options?.onMessage,
        overrideUrl: options?.overrideUrl,
      },
    ).then((x) => {
      x.done.then(() => {
        setCompleteState(true);
      });
      setResponse(x.response);
      return x.response.result;
    });
  };
  const result = useQuery({
    queryKey: [InternalStatsSnapshotAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type InternalStatsSnapshotActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  InternalStatsSnapshotActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => InternalStatsSnapshotActionRes;
  }>;
export const useInternalStatsSnapshotAction = (
  options?: InternalStatsSnapshotActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return InternalStatsSnapshotAction.Fetch(
      {
        body,
        headers: options?.headers,
      },
      {
        creatorFn: options?.creatorFn,
        qs: options?.qs,
        ctx,
        onMessage: options?.onMessage,
        overrideUrl: options?.overrideUrl,
      },
    ).then((x) => {
      x.done.then(() => {
        setCompleteState(true);
      });
      setResponse(x.response);
      return x.response.result;
    });
  };
  const result = useMutation({
    mutationFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
/**
 * InternalStatsSnapshotAction
 */
export class InternalStatsSnapshotAction {
  //
  static URL = "/internal-stats/snapshot";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(InternalStatsSnapshotAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<InternalStatsSnapshotActionRes>, unknown, unknown>(
      overrideUrl ?? InternalStatsSnapshotAction.NewUrl(qs),
      {
        method: InternalStatsSnapshotAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => InternalStatsSnapshotActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new InternalStatsSnapshotActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new InternalStatsSnapshotActionRes(item));
    const res = await InternalStatsSnapshotAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<InternalStatsSnapshotActionRes>();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "InternalStatsSnapshot",
    cliName: "snapshot",
    url: "/internal-stats/snapshot",
    method: "get",
    description:
      "One point-in-time snapshot of every measured stat, as JSON. Root workspace token required by default - see InternalStatsModuleConfig.Authorize.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "generatedAt",
          description: "RFC3339 timestamp this snapshot was collected at.",
          type: "string",
        },
        {
          name: "hostname",
          type: "string",
        },
        {
          name: "items",
          description:
            "Every measured stat, in a stable display order (grouped by category).",
          type: "array",
          fields: [
            {
              name: "key",
              description:
                "Stable machine-readable identifier, e.g. cpu.usedPercent.",
              type: "string",
            },
            {
              name: "label",
              description: "Human-readable label for display, e.g. CPU Usage.",
              type: "string",
            },
            {
              name: "category",
              description:
                "Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.",
              type: "string",
            },
            {
              name: "value",
              description: "Pre-formatted display value, e.g. 62.3% or 8.0 GB.",
              type: "string",
            },
            {
              name: "rawValue",
              description:
                "The same value as a plain number, for programmatic use (0 when not numeric).",
              type: "float64",
            },
            {
              name: "unit",
              description:
                "Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.",
              type: "string",
            },
            {
              name: "severity",
              description:
                "One of ok, warn, critical, info - a coarse threshold-based read on this stat.",
              type: "string",
            },
          ],
        },
      ],
    },
  };
}
/**
 * The base class definition for internalStatsSnapshotActionRes
 **/
export class InternalStatsSnapshotActionRes {
  /**
   * RFC3339 timestamp this snapshot was collected at.
   * @type {string}
   **/
  #generatedAt: string = "";
  /**
   * RFC3339 timestamp this snapshot was collected at.
   * @returns {string}
   **/
  get generatedAt() {
    return this.#generatedAt;
  }
  /**
   * RFC3339 timestamp this snapshot was collected at.
   * @type {string}
   **/
  set generatedAt(value: string) {
    this.#generatedAt = String(value);
  }
  setGeneratedAt(value: string) {
    this.generatedAt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #hostname: string = "";
  /**
   *
   * @returns {string}
   **/
  get hostname() {
    return this.#hostname;
  }
  /**
   *
   * @type {string}
   **/
  set hostname(value: string) {
    this.#hostname = String(value);
  }
  setHostname(value: string) {
    this.hostname = value;
    return this;
  }
  /**
   * Every measured stat, in a stable display order (grouped by category).
   * @type {InternalStatsSnapshotActionRes.Items}
   **/
  #items: MArray<InstanceType<typeof InternalStatsSnapshotActionRes.Items>> =
    MArray.of([]);
  /**
   * Every measured stat, in a stable display order (grouped by category).
   * @returns {InternalStatsSnapshotActionRes.Items}
   **/
  get items() {
    return this.#items;
  }
  /**
   * Every measured stat, in a stable display order (grouped by category).
   * @type {InternalStatsSnapshotActionRes.Items}
   **/
  set items(
    value:
      | MArray<InstanceType<typeof InternalStatsSnapshotActionRes.Items>>
      | InstanceType<typeof InternalStatsSnapshotActionRes.Items>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof InternalStatsSnapshotActionRes.Items
      ) {
        this.#items = MArray.of(value);
      } else {
        this.#items = MArray.of(
          value.map((item) => new InternalStatsSnapshotActionRes.Items(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#items = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast<unknown>(value);
    if (ok) {
      this.#items = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to items, because it needs MArray instance or an Array.",
    );
  }
  setItems(
    value:
      | MArray<InstanceType<typeof InternalStatsSnapshotActionRes.Items>>
      | InstanceType<typeof InternalStatsSnapshotActionRes.Items>[],
  ) {
    this.items = value;
    return this;
  }
  /**
   * The base class definition for items
   **/
  static Items = class Items {
    /**
     * Stable machine-readable identifier, e.g. cpu.usedPercent.
     * @type {string}
     **/
    #key: string = "";
    /**
     * Stable machine-readable identifier, e.g. cpu.usedPercent.
     * @returns {string}
     **/
    get key() {
      return this.#key;
    }
    /**
     * Stable machine-readable identifier, e.g. cpu.usedPercent.
     * @type {string}
     **/
    set key(value: string) {
      this.#key = String(value);
    }
    setKey(value: string) {
      this.key = value;
      return this;
    }
    /**
     * Human-readable label for display, e.g. CPU Usage.
     * @type {string}
     **/
    #label: string = "";
    /**
     * Human-readable label for display, e.g. CPU Usage.
     * @returns {string}
     **/
    get label() {
      return this.#label;
    }
    /**
     * Human-readable label for display, e.g. CPU Usage.
     * @type {string}
     **/
    set label(value: string) {
      this.#label = String(value);
    }
    setLabel(value: string) {
      this.label = value;
      return this;
    }
    /**
     * Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.
     * @type {string}
     **/
    #category: string = "";
    /**
     * Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.
     * @returns {string}
     **/
    get category() {
      return this.#category;
    }
    /**
     * Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.
     * @type {string}
     **/
    set category(value: string) {
      this.#category = String(value);
    }
    setCategory(value: string) {
      this.category = value;
      return this;
    }
    /**
     * Pre-formatted display value, e.g. 62.3% or 8.0 GB.
     * @type {string}
     **/
    #value: string = "";
    /**
     * Pre-formatted display value, e.g. 62.3% or 8.0 GB.
     * @returns {string}
     **/
    get value() {
      return this.#value;
    }
    /**
     * Pre-formatted display value, e.g. 62.3% or 8.0 GB.
     * @type {string}
     **/
    set value(value: string) {
      this.#value = String(value);
    }
    setValue(value: string) {
      this.value = value;
      return this;
    }
    /**
     * The same value as a plain number, for programmatic use (0 when not numeric).
     * @type {number}
     **/
    #rawValue: number = 0.0;
    /**
     * The same value as a plain number, for programmatic use (0 when not numeric).
     * @returns {number}
     **/
    get rawValue() {
      return this.#rawValue;
    }
    /**
     * The same value as a plain number, for programmatic use (0 when not numeric).
     * @type {number}
     **/
    set rawValue(value: number) {
      this.#rawValue = value;
    }
    setRawValue(value: number) {
      this.rawValue = value;
      return this;
    }
    /**
     * Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.
     * @type {string}
     **/
    #unit: string = "";
    /**
     * Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.
     * @returns {string}
     **/
    get unit() {
      return this.#unit;
    }
    /**
     * Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.
     * @type {string}
     **/
    set unit(value: string) {
      this.#unit = String(value);
    }
    setUnit(value: string) {
      this.unit = value;
      return this;
    }
    /**
     * One of ok, warn, critical, info - a coarse threshold-based read on this stat.
     * @type {string}
     **/
    #severity: string = "";
    /**
     * One of ok, warn, critical, info - a coarse threshold-based read on this stat.
     * @returns {string}
     **/
    get severity() {
      return this.#severity;
    }
    /**
     * One of ok, warn, critical, info - a coarse threshold-based read on this stat.
     * @type {string}
     **/
    set severity(value: string) {
      this.#severity = String(value);
    }
    setSeverity(value: string) {
      this.severity = value;
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
      const d = data as Partial<Items>;
      if (d.key !== undefined) {
        this.key = d.key;
      }
      if (d.label !== undefined) {
        this.label = d.label;
      }
      if (d.category !== undefined) {
        this.category = d.category;
      }
      if (d.value !== undefined) {
        this.value = d.value;
      }
      if (d.rawValue !== undefined) {
        this.rawValue = d.rawValue;
      }
      if (d.unit !== undefined) {
        this.unit = d.unit;
      }
      if (d.severity !== undefined) {
        this.severity = d.severity;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        key: this.#key,
        label: this.#label,
        category: this.#category,
        value: this.#value,
        rawValue: this.#rawValue,
        unit: this.#unit,
        severity: this.#severity,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        key: "key",
        label: "label",
        category: "category",
        value: "value",
        rawValue: "rawValue",
        unit: "unit",
        severity: "severity",
      };
    }
    /**
     * Creates an instance of InternalStatsSnapshotActionRes.Items, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: InternalStatsSnapshotActionResType.ItemsType,
    ) {
      return new InternalStatsSnapshotActionRes.Items(possibleDtoObject);
    }
    /**
     * Creates an instance of InternalStatsSnapshotActionRes.Items, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<InternalStatsSnapshotActionResType.ItemsType>,
    ) {
      return new InternalStatsSnapshotActionRes.Items(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<InternalStatsSnapshotActionResType.ItemsType>,
    ): InstanceType<typeof InternalStatsSnapshotActionRes.Items> {
      return new InternalStatsSnapshotActionRes.Items({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof InternalStatsSnapshotActionRes.Items> {
      return new InternalStatsSnapshotActionRes.Items(this.toJSON());
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
    const d = data as Partial<InternalStatsSnapshotActionRes>;
    if (d.generatedAt !== undefined) {
      this.generatedAt = d.generatedAt;
    }
    if (d.hostname !== undefined) {
      this.hostname = d.hostname;
    }
    if (d.items !== undefined) {
      this.items = d.items;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      generatedAt: this.#generatedAt,
      hostname: this.#hostname,
      items: this.#items,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      generatedAt: "generatedAt",
      hostname: "hostname",
      items$: "items",
      get items() {
        return withPrefix(
          "items[:i]",
          InternalStatsSnapshotActionRes.Items.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of InternalStatsSnapshotActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: InternalStatsSnapshotActionResType) {
    return new InternalStatsSnapshotActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of InternalStatsSnapshotActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<InternalStatsSnapshotActionResType>,
  ) {
    return new InternalStatsSnapshotActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<InternalStatsSnapshotActionResType>,
  ): InstanceType<typeof InternalStatsSnapshotActionRes> {
    return new InternalStatsSnapshotActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof InternalStatsSnapshotActionRes> {
    return new InternalStatsSnapshotActionRes(this.toJSON());
  }
}
export abstract class InternalStatsSnapshotActionResFactory {
  abstract create(data: unknown): InternalStatsSnapshotActionRes;
}
/**
 * The base type definition for internalStatsSnapshotActionRes
 **/
export type InternalStatsSnapshotActionResType = {
  /**
   * RFC3339 timestamp this snapshot was collected at.
   * @type {string}
   **/
  generatedAt: string;
  /**
   *
   * @type {string}
   **/
  hostname: string;
  /**
   * Every measured stat, in a stable display order (grouped by category).
   * @type {InternalStatsSnapshotActionResType.ItemsType[]}
   **/
  items: InternalStatsSnapshotActionResType.ItemsType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace InternalStatsSnapshotActionResType {
  /**
   * The base type definition for itemsType
   **/
  export type ItemsType = {
    /**
     * Stable machine-readable identifier, e.g. cpu.usedPercent.
     * @type {string}
     **/
    key: string;
    /**
     * Human-readable label for display, e.g. CPU Usage.
     * @type {string}
     **/
    label: string;
    /**
     * Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.
     * @type {string}
     **/
    category: string;
    /**
     * Pre-formatted display value, e.g. 62.3% or 8.0 GB.
     * @type {string}
     **/
    value: string;
    /**
     * The same value as a plain number, for programmatic use (0 when not numeric).
     * @type {number}
     **/
    rawValue: number;
    /**
     * Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.
     * @type {string}
     **/
    unit: string;
    /**
     * One of ok, warn, critical, info - a coarse threshold-based read on this stat.
     * @type {string}
     **/
    severity: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace ItemsType {}
}
