import { GResponse } from "../../sdk/envelopes/index";
import { buildUrl } from "../../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * Action to communicate with the action CheckClassicPassport
 */
export type CheckClassicPassportActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CheckClassicPassportActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CheckClassicPassportActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CheckClassicPassportActionRes;
  }>;
export const useCheckClassicPassportAction = (
  options?: CheckClassicPassportActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CheckClassicPassportActionReq) => {
    setCompleteState(false);
    return CheckClassicPassportAction.Fetch(
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
 * CheckClassicPassportAction
 */
export class CheckClassicPassportAction {
  //
  static URL = "/workspace/passport/check";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CheckClassicPassportAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<CheckClassicPassportActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<CheckClassicPassportActionRes>,
      CheckClassicPassportActionReq,
      unknown
    >(
      overrideUrl ?? CheckClassicPassportAction.NewUrl(qs),
      {
        method: CheckClassicPassportAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CheckClassicPassportActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => CheckClassicPassportActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CheckClassicPassportActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new CheckClassicPassportActionRes(item));
    const res = await CheckClassicPassportAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CheckClassicPassportActionRes>();
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
    name: "CheckClassicPassport",
    cliName: "ccp",
    url: "/workspace/passport/check",
    method: "post",
    description:
      "Checks if a classic passport (email, phone) exists or not, used in multi step authentication",
    in: {
      fields: [
        {
          name: "value",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "securityToken",
          description:
            "This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.",
          type: "string",
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "next",
          description: "The next possible action which is suggested.",
          type: "slice",
          primitive: "string",
        },
        {
          name: "flags",
          description:
            "Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.",
          type: "slice",
          primitive: "string",
        },
        {
          name: "otpInfo",
          description:
            "If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.",
          type: "object?",
          fields: [
            {
              name: "suspendUntil",
              type: "int64",
            },
            {
              name: "validUntil",
              type: "int64",
            },
            {
              name: "blockedUntil",
              type: "int64",
            },
            {
              name: "secondsToUnblock",
              description:
                "The amount of time left to unblock for next request",
              type: "int64",
            },
          ],
        },
      ],
    },
  };
}
/**
 * The base class definition for checkClassicPassportActionReq
 **/
export class CheckClassicPassportActionReq {
  /**
   *
   * @type {string}
   **/
  #value: string = "";
  /**
   *
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   *
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
   * This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.
   * @type {string}
   **/
  #securityToken: string = "";
  /**
   * This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.
   * @returns {string}
   **/
  get securityToken() {
    return this.#securityToken;
  }
  /**
   * This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.
   * @type {string}
   **/
  set securityToken(value: string) {
    this.#securityToken = String(value);
  }
  setSecurityToken(value: string) {
    this.securityToken = value;
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
    const d = data as Partial<CheckClassicPassportActionReq>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.securityToken !== undefined) {
      this.securityToken = d.securityToken;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      securityToken: this.#securityToken,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      securityToken: "securityToken",
    };
  }
  /**
   * Creates an instance of CheckClassicPassportActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CheckClassicPassportActionReqType) {
    return new CheckClassicPassportActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CheckClassicPassportActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CheckClassicPassportActionReqType>,
  ) {
    return new CheckClassicPassportActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CheckClassicPassportActionReqType>,
  ): InstanceType<typeof CheckClassicPassportActionReq> {
    return new CheckClassicPassportActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CheckClassicPassportActionReq> {
    return new CheckClassicPassportActionReq(this.toJSON());
  }
}
export abstract class CheckClassicPassportActionReqFactory {
  abstract create(data: unknown): CheckClassicPassportActionReq;
}
/**
 * The base type definition for checkClassicPassportActionReq
 **/
export type CheckClassicPassportActionReqType = {
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   * This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.
   * @type {string}
   **/
  securityToken: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CheckClassicPassportActionReqType {}
/**
 * The base class definition for checkClassicPassportActionRes
 **/
export class CheckClassicPassportActionRes {
  /**
   * The next possible action which is suggested.
   * @type {string[]}
   **/
  #next: string[] = [];
  /**
   * The next possible action which is suggested.
   * @returns {string[]}
   **/
  get next() {
    return this.#next;
  }
  /**
   * The next possible action which is suggested.
   * @type {string[]}
   **/
  set next(value: string[]) {
    this.#next = value;
  }
  setNext(value: string[]) {
    this.next = value;
    return this;
  }
  /**
   * Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.
   * @type {string[]}
   **/
  #flags: string[] = [];
  /**
   * Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.
   * @returns {string[]}
   **/
  get flags() {
    return this.#flags;
  }
  /**
   * Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.
   * @type {string[]}
   **/
  set flags(value: string[]) {
    this.#flags = value;
  }
  setFlags(value: string[]) {
    this.flags = value;
    return this;
  }
  /**
   * If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   * @type {CheckClassicPassportActionRes.OtpInfo}
   **/
  #otpInfo?:
    | InstanceType<typeof CheckClassicPassportActionRes.OtpInfo>
    | null
    | undefined
    | null = undefined;
  /**
   * If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   * @returns {CheckClassicPassportActionRes.OtpInfo}
   **/
  get otpInfo() {
    return this.#otpInfo;
  }
  /**
   * If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   * @type {CheckClassicPassportActionRes.OtpInfo}
   **/
  set otpInfo(
    value:
      | InstanceType<typeof CheckClassicPassportActionRes.OtpInfo>
      | null
      | undefined
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof CheckClassicPassportActionRes.OtpInfo) {
      this.#otpInfo = value;
    } else {
      this.#otpInfo = new CheckClassicPassportActionRes.OtpInfo(value);
    }
  }
  setOtpInfo(
    value:
      | InstanceType<typeof CheckClassicPassportActionRes.OtpInfo>
      | null
      | undefined
      | null
      | undefined,
  ) {
    this.otpInfo = value;
    return this;
  }
  /**
   * The base class definition for otpInfo
   **/
  static OtpInfo = class OtpInfo {
    /**
     *
     * @type {number}
     **/
    #suspendUntil: number = 0;
    /**
     *
     * @returns {number}
     **/
    get suspendUntil() {
      return this.#suspendUntil;
    }
    /**
     *
     * @type {number}
     **/
    set suspendUntil(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#suspendUntil = parsedValue;
      }
    }
    setSuspendUntil(value: number) {
      this.suspendUntil = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #validUntil: number = 0;
    /**
     *
     * @returns {number}
     **/
    get validUntil() {
      return this.#validUntil;
    }
    /**
     *
     * @type {number}
     **/
    set validUntil(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#validUntil = parsedValue;
      }
    }
    setValidUntil(value: number) {
      this.validUntil = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #blockedUntil: number = 0;
    /**
     *
     * @returns {number}
     **/
    get blockedUntil() {
      return this.#blockedUntil;
    }
    /**
     *
     * @type {number}
     **/
    set blockedUntil(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#blockedUntil = parsedValue;
      }
    }
    setBlockedUntil(value: number) {
      this.blockedUntil = value;
      return this;
    }
    /**
     * The amount of time left to unblock for next request
     * @type {number}
     **/
    #secondsToUnblock: number = 0;
    /**
     * The amount of time left to unblock for next request
     * @returns {number}
     **/
    get secondsToUnblock() {
      return this.#secondsToUnblock;
    }
    /**
     * The amount of time left to unblock for next request
     * @type {number}
     **/
    set secondsToUnblock(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#secondsToUnblock = parsedValue;
      }
    }
    setSecondsToUnblock(value: number) {
      this.secondsToUnblock = value;
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
      const d = data as Partial<OtpInfo>;
      if (d.suspendUntil !== undefined) {
        this.suspendUntil = d.suspendUntil;
      }
      if (d.validUntil !== undefined) {
        this.validUntil = d.validUntil;
      }
      if (d.blockedUntil !== undefined) {
        this.blockedUntil = d.blockedUntil;
      }
      if (d.secondsToUnblock !== undefined) {
        this.secondsToUnblock = d.secondsToUnblock;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        suspendUntil: this.#suspendUntil,
        validUntil: this.#validUntil,
        blockedUntil: this.#blockedUntil,
        secondsToUnblock: this.#secondsToUnblock,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        suspendUntil: "suspendUntil",
        validUntil: "validUntil",
        blockedUntil: "blockedUntil",
        secondsToUnblock: "secondsToUnblock",
      };
    }
    /**
     * Creates an instance of CheckClassicPassportActionRes.OtpInfo, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: CheckClassicPassportActionResType.OtpInfoType,
    ) {
      return new CheckClassicPassportActionRes.OtpInfo(possibleDtoObject);
    }
    /**
     * Creates an instance of CheckClassicPassportActionRes.OtpInfo, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<CheckClassicPassportActionResType.OtpInfoType>,
    ) {
      return new CheckClassicPassportActionRes.OtpInfo(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<CheckClassicPassportActionResType.OtpInfoType>,
    ): InstanceType<typeof CheckClassicPassportActionRes.OtpInfo> {
      return new CheckClassicPassportActionRes.OtpInfo({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof CheckClassicPassportActionRes.OtpInfo> {
      return new CheckClassicPassportActionRes.OtpInfo(this.toJSON());
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
    const d = data as Partial<CheckClassicPassportActionRes>;
    if (d.next !== undefined) {
      this.next = d.next;
    }
    if (d.flags !== undefined) {
      this.flags = d.flags;
    }
    if (d.otpInfo !== undefined) {
      this.otpInfo = d.otpInfo;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      next: this.#next,
      flags: this.#flags,
      otpInfo: this.#otpInfo,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      next$: "next",
      get next() {
        return "next[:i]";
      },
      flags$: "flags",
      get flags() {
        return "flags[:i]";
      },
      otpInfo$: "otpInfo",
      get otpInfo() {
        return withPrefix(
          "otpInfo",
          CheckClassicPassportActionRes.OtpInfo.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of CheckClassicPassportActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CheckClassicPassportActionResType) {
    return new CheckClassicPassportActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CheckClassicPassportActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CheckClassicPassportActionResType>,
  ) {
    return new CheckClassicPassportActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CheckClassicPassportActionResType>,
  ): InstanceType<typeof CheckClassicPassportActionRes> {
    return new CheckClassicPassportActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CheckClassicPassportActionRes> {
    return new CheckClassicPassportActionRes(this.toJSON());
  }
}
export abstract class CheckClassicPassportActionResFactory {
  abstract create(data: unknown): CheckClassicPassportActionRes;
}
/**
 * The base type definition for checkClassicPassportActionRes
 **/
export type CheckClassicPassportActionResType = {
  /**
   * The next possible action which is suggested.
   * @type {string[]}
   **/
  next: string[];
  /**
   * Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.
   * @type {string[]}
   **/
  flags: string[];
  /**
   * If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   * @type {CheckClassicPassportActionResType.OtpInfoType}
   **/
  otpInfo?: CheckClassicPassportActionResType.OtpInfoType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CheckClassicPassportActionResType {
  /**
   * The base type definition for otpInfoType
   **/
  export type OtpInfoType = {
    /**
     *
     * @type {number}
     **/
    suspendUntil: number;
    /**
     *
     * @type {number}
     **/
    validUntil: number;
    /**
     *
     * @type {number}
     **/
    blockedUntil: number;
    /**
     * The amount of time left to unblock for next request
     * @type {number}
     **/
    secondsToUnblock: number;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace OtpInfoType {}
}
