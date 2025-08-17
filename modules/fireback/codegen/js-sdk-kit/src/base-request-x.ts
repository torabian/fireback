export type WebResult<T, E> =
  | { kind: "success"; value: T; raw: Response }
  | { kind: "error"; error: E };

export enum HttpMethod {
  GET = "GET",
  POST = "POST",
  PUT = "PUT",
  PATCH = "PATCH",
  DELETE = "DELETE",
}

// Without decorators
export class JsonMessage {
  static __jsonParsable = true;
}

// Decorator
export function JsonParsable<T extends { new (...args: any[]): {} }>(
  constructor: T
) {
  // add a flag to the constructor
  (constructor as any).__jsonParsable = true;
  return constructor;
}

// helper to check
export function isJsonParsable(cls: any) {
  return !!cls.__jsonParsable;
}

/**
 * @description Base class for network-based requests.
 * Provides common properties and methods for URL handling,
 * typed responses/messages, and cleanup.
 */
export abstract class BaseRequest<TMessage = unknown, TError = unknown> {
  /** The URL or endpoint for the request/connection */
  protected url: string;

  /** Optional typed class for response/message deserialization */
  protected MessageClass?: { new (raw?: string): TMessage };

  /** Optional typed class for error deserialization */
  protected ErrorClass?: { new (): TError };

  /** Optional abort or cleanup controller */
  protected controller?: AbortController;

  constructor(
    MessageClass?: { new (): TMessage },
    ErrorClass?: { new (): TError }
  ) {
    this.MessageClass = MessageClass;
    this.ErrorClass = ErrorClass;
  }

  /**
   * @description Sets or replaces the URL
   * @param url - The endpoint or server URL
   */
  setUrl(url: string) {
    this.url = url;
    return this;
  }

  /**
   * @description Returns the current URL
   */
  getUrl(): string {
    return this.url;
  }

  /**
   * @description Sets a controller for aborting or cleanup
   * @param controller - AbortController or custom controller
   */
  setController(controller: AbortController) {
    this.controller = controller;
    return this;
  }

  /**
   * @description Cleans up resources (connections, streams, etc.)
   * Override in subclasses if needed
   */
  cleanup(): void {
    if (this.controller) {
      this.controller.abort();
    }
  }

  protected deserialize<T>(cls: new () => T, response: Response, raw: any): T {
    if (!cls) return raw;

    const instance = new cls();

    // JsonParsable default logic
    if (isJsonParsable(cls)) {
      Object.assign(instance, typeof raw === "string" ? JSON.parse(raw) : raw);
      return instance;
    }

    // static unserialize
    if (typeof (cls as any).unserialize === "function") {
      const result = (cls as any).unserialize(raw, response);
      if (result !== undefined) {
        Object.assign(instance, result);
      }
    }

    return instance;
  }

  protected processMessage(response: Response, raw: any): TMessage {
    if (!this.MessageClass) return raw;
    return this.deserialize(this.MessageClass, response, raw);
  }

  protected processError(response: Response, err: any): TError {
    if (!this.ErrorClass) return err;
    return this.deserialize(this.ErrorClass, response, err);
  }
}
