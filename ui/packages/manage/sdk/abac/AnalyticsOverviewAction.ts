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
 * Action to communicate with the action AnalyticsOverview
 */
export type AnalyticsOverviewActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AnalyticsOverviewActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<AnalyticsOverviewActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  AnalyticsOverviewActionOptions &
  Partial<{
    creatorFn: (item: unknown) => AnalyticsOverviewActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useAnalyticsOverviewActionQuery = (
  options: AnalyticsOverviewActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return AnalyticsOverviewAction.Fetch(
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
    queryKey: [AnalyticsOverviewAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type AnalyticsOverviewActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AnalyticsOverviewActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AnalyticsOverviewActionRes;
  }>;
export const useAnalyticsOverviewAction = (
  options?: AnalyticsOverviewActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return AnalyticsOverviewAction.Fetch(
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
 * AnalyticsOverviewAction
 */
export class AnalyticsOverviewAction {
  //
  static URL = "/analytics/overview";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AnalyticsOverviewAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<AnalyticsOverviewActionRes>, unknown, unknown>(
      overrideUrl ?? AnalyticsOverviewAction.NewUrl(qs),
      {
        method: AnalyticsOverviewAction.Method,
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
      creatorFn?: ((item: unknown) => AnalyticsOverviewActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new AnalyticsOverviewActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AnalyticsOverviewActionRes(item));
    const res = await AnalyticsOverviewAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<AnalyticsOverviewActionRes>();
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
    name: "AnalyticsOverview",
    cliName: "overview",
    url: "/analytics/overview",
    method: "get",
    description:
      "Root-only. Aggregate stats across users/workspaces/passports/roles/sessions - a fixed set of headline numbers (items, same key/label/category/value/rawValue/ unit/severity shape InternalStatsSnapshot uses, so the UI can reuse the same stat-tile renderer) plus a handful of named chart datasets (series - monthly trends and categorical breakdowns) for the chart section below them. See AnalyticsActionImplementation.go for exactly what's measured and why.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "generatedAt",
          description: "RFC3339 timestamp this snapshot was computed at.",
          type: "string",
        },
        {
          name: "items",
          description:
            "Headline stats, in a stable display order (grouped by category).",
          type: "array",
          fields: [
            {
              name: "key",
              description:
                "Stable machine-readable identifier, e.g. users.total.",
              type: "string",
            },
            {
              name: "label",
              description:
                "Human-readable label for display, e.g. Total Users.",
              type: "string",
            },
            {
              name: "category",
              description:
                "Display grouping, e.g. Users, Workspaces, Access, Engagement.",
              type: "string",
            },
            {
              name: "value",
              description: "Pre-formatted display value, e.g. 128 or 42.3%.",
              type: "string",
            },
            {
              name: "rawValue",
              description:
                "The same value as a plain number, for programmatic use.",
              type: "float64",
            },
            {
              name: "unit",
              description:
                "Unit of rawValue, e.g. count, percent. Empty for non-numeric stats.",
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
        {
          name: "series",
          description:
            "Named chart datasets (monthly trends, categorical breakdowns) for the chart section.",
          type: "array",
          fields: [
            {
              name: "key",
              description:
                "Stable machine-readable identifier, e.g. users.monthlySignups.",
              type: "string",
            },
            {
              name: "label",
              description: "Human-readable chart title.",
              type: "string",
            },
            {
              name: "category",
              description:
                "Display grouping, matching items' own category values where related.",
              type: "string",
            },
            {
              name: "chartType",
              description:
                "One of 'line', 'bar' - a hint for which mark the UI should use.",
              type: "string",
            },
            {
              name: "unit",
              description: "Unit of each point's value, e.g. count, percent.",
              type: "string",
            },
            {
              name: "points",
              description:
                "Ordered data points - for time series, oldest first.",
              type: "array",
              fields: [
                {
                  name: "label",
                  description:
                    'X-axis label, e.g. "2026-03" for a month bucket, or a category name.',
                  type: "string",
                },
                {
                  name: "value",
                  type: "float64",
                },
              ],
            },
          ],
        },
      ],
    },
  };
}
/**
 * The base class definition for analyticsOverviewActionRes
 **/
export class AnalyticsOverviewActionRes {
  /**
   * RFC3339 timestamp this snapshot was computed at.
   * @type {string}
   **/
  #generatedAt: string = "";
  /**
   * RFC3339 timestamp this snapshot was computed at.
   * @returns {string}
   **/
  get generatedAt() {
    return this.#generatedAt;
  }
  /**
   * RFC3339 timestamp this snapshot was computed at.
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
   * Headline stats, in a stable display order (grouped by category).
   * @type {AnalyticsOverviewActionRes.Items}
   **/
  #items: MArray<InstanceType<typeof AnalyticsOverviewActionRes.Items>> =
    MArray.of([]);
  /**
   * Headline stats, in a stable display order (grouped by category).
   * @returns {AnalyticsOverviewActionRes.Items}
   **/
  get items() {
    return this.#items;
  }
  /**
   * Headline stats, in a stable display order (grouped by category).
   * @type {AnalyticsOverviewActionRes.Items}
   **/
  set items(
    value:
      | MArray<InstanceType<typeof AnalyticsOverviewActionRes.Items>>
      | InstanceType<typeof AnalyticsOverviewActionRes.Items>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof AnalyticsOverviewActionRes.Items
      ) {
        this.#items = MArray.of(value);
      } else {
        this.#items = MArray.of(
          value.map((item) => new AnalyticsOverviewActionRes.Items(item)),
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
      | MArray<InstanceType<typeof AnalyticsOverviewActionRes.Items>>
      | InstanceType<typeof AnalyticsOverviewActionRes.Items>[],
  ) {
    this.items = value;
    return this;
  }
  /**
   * Named chart datasets (monthly trends, categorical breakdowns) for the chart section.
   * @type {AnalyticsOverviewActionRes.Series}
   **/
  #series: MArray<InstanceType<typeof AnalyticsOverviewActionRes.Series>> =
    MArray.of([]);
  /**
   * Named chart datasets (monthly trends, categorical breakdowns) for the chart section.
   * @returns {AnalyticsOverviewActionRes.Series}
   **/
  get series() {
    return this.#series;
  }
  /**
   * Named chart datasets (monthly trends, categorical breakdowns) for the chart section.
   * @type {AnalyticsOverviewActionRes.Series}
   **/
  set series(
    value:
      | MArray<InstanceType<typeof AnalyticsOverviewActionRes.Series>>
      | InstanceType<typeof AnalyticsOverviewActionRes.Series>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof AnalyticsOverviewActionRes.Series
      ) {
        this.#series = MArray.of(value);
      } else {
        this.#series = MArray.of(
          value.map((item) => new AnalyticsOverviewActionRes.Series(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#series = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast<unknown>(value);
    if (ok) {
      this.#series = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to series, because it needs MArray instance or an Array.",
    );
  }
  setSeries(
    value:
      | MArray<InstanceType<typeof AnalyticsOverviewActionRes.Series>>
      | InstanceType<typeof AnalyticsOverviewActionRes.Series>[],
  ) {
    this.series = value;
    return this;
  }
  /**
   * The base class definition for items
   **/
  static Items = class Items {
    /**
     * Stable machine-readable identifier, e.g. users.total.
     * @type {string}
     **/
    #key: string = "";
    /**
     * Stable machine-readable identifier, e.g. users.total.
     * @returns {string}
     **/
    get key() {
      return this.#key;
    }
    /**
     * Stable machine-readable identifier, e.g. users.total.
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
     * Human-readable label for display, e.g. Total Users.
     * @type {string}
     **/
    #label: string = "";
    /**
     * Human-readable label for display, e.g. Total Users.
     * @returns {string}
     **/
    get label() {
      return this.#label;
    }
    /**
     * Human-readable label for display, e.g. Total Users.
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
     * Display grouping, e.g. Users, Workspaces, Access, Engagement.
     * @type {string}
     **/
    #category: string = "";
    /**
     * Display grouping, e.g. Users, Workspaces, Access, Engagement.
     * @returns {string}
     **/
    get category() {
      return this.#category;
    }
    /**
     * Display grouping, e.g. Users, Workspaces, Access, Engagement.
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
     * Pre-formatted display value, e.g. 128 or 42.3%.
     * @type {string}
     **/
    #value: string = "";
    /**
     * Pre-formatted display value, e.g. 128 or 42.3%.
     * @returns {string}
     **/
    get value() {
      return this.#value;
    }
    /**
     * Pre-formatted display value, e.g. 128 or 42.3%.
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
     * The same value as a plain number, for programmatic use.
     * @type {number}
     **/
    #rawValue: number = 0.0;
    /**
     * The same value as a plain number, for programmatic use.
     * @returns {number}
     **/
    get rawValue() {
      return this.#rawValue;
    }
    /**
     * The same value as a plain number, for programmatic use.
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
     * Unit of rawValue, e.g. count, percent. Empty for non-numeric stats.
     * @type {string}
     **/
    #unit: string = "";
    /**
     * Unit of rawValue, e.g. count, percent. Empty for non-numeric stats.
     * @returns {string}
     **/
    get unit() {
      return this.#unit;
    }
    /**
     * Unit of rawValue, e.g. count, percent. Empty for non-numeric stats.
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
     * Creates an instance of AnalyticsOverviewActionRes.Items, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: AnalyticsOverviewActionResType.ItemsType) {
      return new AnalyticsOverviewActionRes.Items(possibleDtoObject);
    }
    /**
     * Creates an instance of AnalyticsOverviewActionRes.Items, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<AnalyticsOverviewActionResType.ItemsType>,
    ) {
      return new AnalyticsOverviewActionRes.Items(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<AnalyticsOverviewActionResType.ItemsType>,
    ): InstanceType<typeof AnalyticsOverviewActionRes.Items> {
      return new AnalyticsOverviewActionRes.Items({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof AnalyticsOverviewActionRes.Items> {
      return new AnalyticsOverviewActionRes.Items(this.toJSON());
    }
  };
  /**
   * The base class definition for series
   **/
  static Series = class Series {
    /**
     * Stable machine-readable identifier, e.g. users.monthlySignups.
     * @type {string}
     **/
    #key: string = "";
    /**
     * Stable machine-readable identifier, e.g. users.monthlySignups.
     * @returns {string}
     **/
    get key() {
      return this.#key;
    }
    /**
     * Stable machine-readable identifier, e.g. users.monthlySignups.
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
     * Human-readable chart title.
     * @type {string}
     **/
    #label: string = "";
    /**
     * Human-readable chart title.
     * @returns {string}
     **/
    get label() {
      return this.#label;
    }
    /**
     * Human-readable chart title.
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
     * Display grouping, matching items' own category values where related.
     * @type {string}
     **/
    #category: string = "";
    /**
     * Display grouping, matching items' own category values where related.
     * @returns {string}
     **/
    get category() {
      return this.#category;
    }
    /**
     * Display grouping, matching items' own category values where related.
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
     * One of 'line', 'bar' - a hint for which mark the UI should use.
     * @type {string}
     **/
    #chartType: string = "";
    /**
     * One of 'line', 'bar' - a hint for which mark the UI should use.
     * @returns {string}
     **/
    get chartType() {
      return this.#chartType;
    }
    /**
     * One of 'line', 'bar' - a hint for which mark the UI should use.
     * @type {string}
     **/
    set chartType(value: string) {
      this.#chartType = String(value);
    }
    setChartType(value: string) {
      this.chartType = value;
      return this;
    }
    /**
     * Unit of each point's value, e.g. count, percent.
     * @type {string}
     **/
    #unit: string = "";
    /**
     * Unit of each point's value, e.g. count, percent.
     * @returns {string}
     **/
    get unit() {
      return this.#unit;
    }
    /**
     * Unit of each point's value, e.g. count, percent.
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
     * Ordered data points - for time series, oldest first.
     * @type {AnalyticsOverviewActionRes.Series.Points}
     **/
    #points: MArray<
      InstanceType<typeof AnalyticsOverviewActionRes.Series.Points>
    > = MArray.of([]);
    /**
     * Ordered data points - for time series, oldest first.
     * @returns {AnalyticsOverviewActionRes.Series.Points}
     **/
    get points() {
      return this.#points;
    }
    /**
     * Ordered data points - for time series, oldest first.
     * @type {AnalyticsOverviewActionRes.Series.Points}
     **/
    set points(
      value:
        | MArray<InstanceType<typeof AnalyticsOverviewActionRes.Series.Points>>
        | InstanceType<typeof AnalyticsOverviewActionRes.Series.Points>[],
    ) {
      // When the passed value is already an array, we check if we need to
      // cast the inner items into class instance.
      if (Array.isArray(value)) {
        if (
          value.length > 0 &&
          value[0] instanceof AnalyticsOverviewActionRes.Series.Points
        ) {
          this.#points = MArray.of(value);
        } else {
          this.#points = MArray.of(
            value.map(
              (item) => new AnalyticsOverviewActionRes.Series.Points(item),
            ),
          );
        }
        return;
      }
      // If the instance is already an MArray, we assume it's all good.
      if (value instanceof MArray) {
        this.#points = value;
        return;
      }
      // If the value is not array, and is not a MArray, we need to be consider,
      // it might be eligible to be casted into MArray.
      const { ok, value: mcastValue } = MArray.cast<unknown>(value);
      if (ok) {
        this.#points = mcastValue as any;
        return;
      }
      console.warn(
        "Cannot assing value to points, because it needs MArray instance or an Array.",
      );
    }
    setPoints(
      value:
        | MArray<InstanceType<typeof AnalyticsOverviewActionRes.Series.Points>>
        | InstanceType<typeof AnalyticsOverviewActionRes.Series.Points>[],
    ) {
      this.points = value;
      return this;
    }
    /**
     * The base class definition for points
     **/
    static Points = class Points {
      /**
       * X-axis label, e.g. "2026-03" for a month bucket, or a category name.
       * @type {string}
       **/
      #label: string = "";
      /**
       * X-axis label, e.g. "2026-03" for a month bucket, or a category name.
       * @returns {string}
       **/
      get label() {
        return this.#label;
      }
      /**
       * X-axis label, e.g. "2026-03" for a month bucket, or a category name.
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
       *
       * @type {number}
       **/
      #value: number = 0.0;
      /**
       *
       * @returns {number}
       **/
      get value() {
        return this.#value;
      }
      /**
       *
       * @type {number}
       **/
      set value(value: number) {
        this.#value = value;
      }
      setValue(value: number) {
        this.value = value;
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
        const d = data as Partial<Points>;
        if (d.label !== undefined) {
          this.label = d.label;
        }
        if (d.value !== undefined) {
          this.value = d.value;
        }
      }
      /**
       *	Special toJSON override, since the field are private,
       *	Json stringify won't see them unless we mention it explicitly.
       **/
      toJSON() {
        return {
          label: this.#label,
          value: this.#value,
        };
      }
      toString() {
        return JSON.stringify(this);
      }
      static get Fields() {
        return {
          label: "label",
          value: "value",
        };
      }
      /**
       * Creates an instance of AnalyticsOverviewActionRes.Series.Points, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(
        possibleDtoObject: AnalyticsOverviewActionResType.SeriesType.PointsType,
      ) {
        return new AnalyticsOverviewActionRes.Series.Points(possibleDtoObject);
      }
      /**
       * Creates an instance of AnalyticsOverviewActionRes.Series.Points, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<AnalyticsOverviewActionResType.SeriesType.PointsType>,
      ) {
        return new AnalyticsOverviewActionRes.Series.Points(partialDtoObject);
      }
      copyWith(
        partial: PartialDeep<AnalyticsOverviewActionResType.SeriesType.PointsType>,
      ): InstanceType<typeof AnalyticsOverviewActionRes.Series.Points> {
        return new AnalyticsOverviewActionRes.Series.Points({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<typeof AnalyticsOverviewActionRes.Series.Points> {
        return new AnalyticsOverviewActionRes.Series.Points(this.toJSON());
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
      const d = data as Partial<Series>;
      if (d.key !== undefined) {
        this.key = d.key;
      }
      if (d.label !== undefined) {
        this.label = d.label;
      }
      if (d.category !== undefined) {
        this.category = d.category;
      }
      if (d.chartType !== undefined) {
        this.chartType = d.chartType;
      }
      if (d.unit !== undefined) {
        this.unit = d.unit;
      }
      if (d.points !== undefined) {
        this.points = d.points;
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
        chartType: this.#chartType,
        unit: this.#unit,
        points: this.#points,
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
        chartType: "chartType",
        unit: "unit",
        points$: "points",
        get points() {
          return withPrefix(
            "series.points[:i]",
            AnalyticsOverviewActionRes.Series.Points.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of AnalyticsOverviewActionRes.Series, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: AnalyticsOverviewActionResType.SeriesType) {
      return new AnalyticsOverviewActionRes.Series(possibleDtoObject);
    }
    /**
     * Creates an instance of AnalyticsOverviewActionRes.Series, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<AnalyticsOverviewActionResType.SeriesType>,
    ) {
      return new AnalyticsOverviewActionRes.Series(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<AnalyticsOverviewActionResType.SeriesType>,
    ): InstanceType<typeof AnalyticsOverviewActionRes.Series> {
      return new AnalyticsOverviewActionRes.Series({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof AnalyticsOverviewActionRes.Series> {
      return new AnalyticsOverviewActionRes.Series(this.toJSON());
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
    const d = data as Partial<AnalyticsOverviewActionRes>;
    if (d.generatedAt !== undefined) {
      this.generatedAt = d.generatedAt;
    }
    if (d.items !== undefined) {
      this.items = d.items;
    }
    if (d.series !== undefined) {
      this.series = d.series;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      generatedAt: this.#generatedAt,
      items: this.#items,
      series: this.#series,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      generatedAt: "generatedAt",
      items$: "items",
      get items() {
        return withPrefix("items[:i]", AnalyticsOverviewActionRes.Items.Fields);
      },
      series$: "series",
      get series() {
        return withPrefix(
          "series[:i]",
          AnalyticsOverviewActionRes.Series.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of AnalyticsOverviewActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AnalyticsOverviewActionResType) {
    return new AnalyticsOverviewActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of AnalyticsOverviewActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AnalyticsOverviewActionResType>) {
    return new AnalyticsOverviewActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AnalyticsOverviewActionResType>,
  ): InstanceType<typeof AnalyticsOverviewActionRes> {
    return new AnalyticsOverviewActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AnalyticsOverviewActionRes> {
    return new AnalyticsOverviewActionRes(this.toJSON());
  }
}
export abstract class AnalyticsOverviewActionResFactory {
  abstract create(data: unknown): AnalyticsOverviewActionRes;
}
/**
 * The base type definition for analyticsOverviewActionRes
 **/
export type AnalyticsOverviewActionResType = {
  /**
   * RFC3339 timestamp this snapshot was computed at.
   * @type {string}
   **/
  generatedAt: string;
  /**
   * Headline stats, in a stable display order (grouped by category).
   * @type {AnalyticsOverviewActionResType.ItemsType[]}
   **/
  items: AnalyticsOverviewActionResType.ItemsType[];
  /**
   * Named chart datasets (monthly trends, categorical breakdowns) for the chart section.
   * @type {AnalyticsOverviewActionResType.SeriesType[]}
   **/
  series: AnalyticsOverviewActionResType.SeriesType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AnalyticsOverviewActionResType {
  /**
   * The base type definition for itemsType
   **/
  export type ItemsType = {
    /**
     * Stable machine-readable identifier, e.g. users.total.
     * @type {string}
     **/
    key: string;
    /**
     * Human-readable label for display, e.g. Total Users.
     * @type {string}
     **/
    label: string;
    /**
     * Display grouping, e.g. Users, Workspaces, Access, Engagement.
     * @type {string}
     **/
    category: string;
    /**
     * Pre-formatted display value, e.g. 128 or 42.3%.
     * @type {string}
     **/
    value: string;
    /**
     * The same value as a plain number, for programmatic use.
     * @type {number}
     **/
    rawValue: number;
    /**
     * Unit of rawValue, e.g. count, percent. Empty for non-numeric stats.
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
  /**
   * The base type definition for seriesType
   **/
  export type SeriesType = {
    /**
     * Stable machine-readable identifier, e.g. users.monthlySignups.
     * @type {string}
     **/
    key: string;
    /**
     * Human-readable chart title.
     * @type {string}
     **/
    label: string;
    /**
     * Display grouping, matching items' own category values where related.
     * @type {string}
     **/
    category: string;
    /**
     * One of 'line', 'bar' - a hint for which mark the UI should use.
     * @type {string}
     **/
    chartType: string;
    /**
     * Unit of each point's value, e.g. count, percent.
     * @type {string}
     **/
    unit: string;
    /**
     * Ordered data points - for time series, oldest first.
     * @type {AnalyticsOverviewActionResType.SeriesType.PointsType[]}
     **/
    points: AnalyticsOverviewActionResType.SeriesType.PointsType[];
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace SeriesType {
    /**
     * The base type definition for pointsType
     **/
    export type PointsType = {
      /**
       * X-axis label, e.g. "2026-03" for a month bucket, or a category name.
       * @type {string}
       **/
      label: string;
      /**
       *
       * @type {number}
       **/
      value: number;
    };
    // eslint-disable-next-line @typescript-eslint/no-namespace
    export namespace PointsType {}
  }
}
