import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { ExecutionContext, createParamDecorator } from "@nestjs/common";
import { TypedRequestInit, URLSearchParamsX, fetchx } from "./sdk";
/**
 * Action to communicate with the action checkClassicPassport
 */

/**
 * FetchCheckClassicPassportAction
 */

export class FetchCheckClassicPassportAction {
  static URL = "/workspace/passport/check";
  static Method = "post";

  static Axios = (
    clientInstance: AxiosInstance,
    config: AxiosRequestConfig<unknown>
  ) =>
    clientInstance
      .request<unknown, AxiosResponse<unknown>, unknown>(config)

      .then((res) => {
        return {
          ...res,

          // if there is a output class, create instance out of it.
          data: new CheckClassicPassportRes(res.data),
        };
      });

  static Fetch = (
    init?: TypedRequestInit<unknown, unknown> | undefined,
    qs?: CheckClassicPassportQueryParams,
    overrideUrl?: string
  ) =>
    fetchx<unknown, unknown, unknown>(
      new URL(
        (overrideUrl ?? FetchCheckClassicPassportAction.URL) +
          "?" +
          qs?.toString()
      ),
      init
    )
      .then((res) => res.json())

      .then((data) => new CheckClassicPassportRes(data));
}

/**
 * @decription The base class definition for checkClassicPassportReq
 **/

export class CheckClassicPassportReq {
  /**
   * @type {string}
   * @description
   **/
  value;
  /**
   * @returns {string}
   * @description
   **/
  getValue() {
    return this[`value`];
  }
  /**
   * @param {string}
   * @description
   **/
  setValue(value) {
    this[`value`] = value;
    return this;
  }

  /**
   * @type {string}
   * @description This can be the value of recaptcha2, recaptch3, or generate security image or voice for verification. Will be used based on the configuration.
   **/
  securityToken;
  /**
   * @returns {string}
   * @description This can be the value of recaptcha2, recaptch3, or generate security image or voice for verification. Will be used based on the configuration.
   **/
  getSecurityToken() {
    return this[`securityToken`];
  }
  /**
   * @param {string}
   * @description This can be the value of recaptcha2, recaptch3, or generate security image or voice for verification. Will be used based on the configuration.
   **/
  setSecurityToken(value) {
    this[`securityToken`] = value;
    return this;
  }

  /** a placeholder for WebRequestX auto patching the json content to the object **/
  static __jsonParsable;

  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@CheckClassicPassportReq.Nest() headers: CheckClassicPassportReq): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    // @ts-ignore
    const request = ctx.switchToHttp().getRequest();
    // @ts-ignore
    return new CheckClassicPassportReq(request.body);
  });
}

/**
 * @decription The base class definition for checkClassicPassportRes
 **/

export class CheckClassicPassportRes {
  /**
   * @type {string[]}
   * @description The next possible action which is suggested.
   **/
  next;
  /**
   * @returns {string[]}
   * @description The next possible action which is suggested.
   **/
  getNext() {
    return this[`next`];
  }
  /**
   * @param {string[]}
   * @description The next possible action which is suggested.
   **/
  setNext(value) {
    this[`next`] = value;
    return this;
  }

  /**
   * @type {string[]}
   * @description Extra information that can be useful actually when doing onboarding. Make sure sensetive information doesn't go out.
   **/
  flags;
  /**
   * @returns {string[]}
   * @description Extra information that can be useful actually when doing onboarding. Make sure sensetive information doesn't go out.
   **/
  getFlags() {
    return this[`flags`];
  }
  /**
   * @param {string[]}
   * @description Extra information that can be useful actually when doing onboarding. Make sure sensetive information doesn't go out.
   **/
  setFlags(value) {
    this[`flags`] = value;
    return this;
  }

  /**
   * @type {CheckClassicPassportRes.OtpInfo}
   * @description If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   **/
  otpInfo;
  /**
   * @returns {CheckClassicPassportRes.OtpInfo}
   * @description If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   **/
  getOtpInfo() {
    return this[`otpInfo`];
  }
  /**
   * @param {CheckClassicPassportRes.OtpInfo}
   * @description If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
   **/
  setOtpInfo(value) {
    this[`otpInfo`] = value;
    return this;
  }

  /**
   * @decription The base class definition for otpInfo
   **/

  static OtpInfo = class OtpInfo {
    /**
     * @type {number}
     * @description
     **/
    suspendUntil;
    /**
     * @returns {number}
     * @description
     **/
    getSuspendUntil() {
      return this[`suspendUntil`];
    }
    /**
     * @param {number}
     * @description
     **/
    setSuspendUntil(value) {
      this[`suspendUntil`] = value;
      return this;
    }

    /**
     * @type {number}
     * @description
     **/
    validUntil;
    /**
     * @returns {number}
     * @description
     **/
    getValidUntil() {
      return this[`validUntil`];
    }
    /**
     * @param {number}
     * @description
     **/
    setValidUntil(value) {
      this[`validUntil`] = value;
      return this;
    }

    /**
     * @type {number}
     * @description
     **/
    blockedUntil;
    /**
     * @returns {number}
     * @description
     **/
    getBlockedUntil() {
      return this[`blockedUntil`];
    }
    /**
     * @param {number}
     * @description
     **/
    setBlockedUntil(value) {
      this[`blockedUntil`] = value;
      return this;
    }

    /**
     * @type {number}
     * @description The amount of time left to unblock for next request
     **/
    secondsToUnblock;
    /**
     * @returns {number}
     * @description The amount of time left to unblock for next request
     **/
    getSecondsToUnblock() {
      return this[`secondsToUnblock`];
    }
    /**
     * @param {number}
     * @description The amount of time left to unblock for next request
     **/
    setSecondsToUnblock(value) {
      this[`secondsToUnblock`] = value;
      return this;
    }

    /** a placeholder for WebRequestX auto patching the json content to the object **/
    static __jsonParsable;
  };

  /** a placeholder for WebRequestX auto patching the json content to the object **/
  static __jsonParsable;

  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@CheckClassicPassportRes.Nest() headers: CheckClassicPassportRes): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    // @ts-ignore
    const request = ctx.switchToHttp().getRequest();
    // @ts-ignore
    return new CheckClassicPassportRes(request.body);
  });
}

/**
 * CheckClassicPassportHeaders class
 * Auto-generated from Module3Action
 */
export class CheckClassicPassportHeaders extends Headers {
  // the getters generated by us would be casting types before returning.
  // you still can use .get function to get the string value.
  #getTyped(key, type) {
    const val = this.get(key);
    if (val == null) return null;

    const t = type.toLowerCase();

    if (t.includes("number")) return Number(val);
    if (t.includes("bool")) return val === "true";
    return val; // string or any other fallback
  }

  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }

  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@CheckClassicPassportHeaders.Nest() headers: CheckClassicPassportHeaders): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    // @ts-ignore
    const request = ctx.switchToHttp().getRequest();
    // @ts-ignore
    return new CheckClassicPassportHeaders(Object.entries(request.headers));
  });
}

/**
 * CheckClassicPassportQueryParams class
 * Auto-generated from Module3Action
 */
export class CheckClassicPassportQueryParams extends URLSearchParamsX {
  // the getters generated by us would be casting types before returning.
  // you still can use .get function to get the string value.
  #getTyped(key, type) {
    const val = this.get(key);
    if (val == null) return null;

    const t = type.toLowerCase();

    if (t.includes("number")) return Number(val);
    if (t.includes("bool")) return val === "true";
    return val; // string or any other fallback
  }

  /**
   * Nest.js decorator for controller query. Instead of using @Query() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@CheckClassicPassportQueryParams.Nest() query: CheckClassicPassportQueryParams): string {
   *  return JSON.stringify(query.getMyfield());
   * }
   */
  static Nest = createParamDecorator((_data, ctx) => {
    // @ts-ignore
    const request = ctx.switchToHttp().getRequest();
    // @ts-ignore
    return new CheckClassicPassportQueryParams(request.query);
  });
}
