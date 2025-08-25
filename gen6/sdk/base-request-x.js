export var HttpMethod;
(function (HttpMethod) {
    HttpMethod["GET"] = "GET";
    HttpMethod["POST"] = "POST";
    HttpMethod["PUT"] = "PUT";
    HttpMethod["PATCH"] = "PATCH";
    HttpMethod["DELETE"] = "DELETE";
})(HttpMethod || (HttpMethod = {}));
// Without decorators
export class JsonMessage {
}
JsonMessage.__jsonParsable = true;
// Decorator
export function JsonParsable(constructor) {
    // add a flag to the constructor
    constructor.__jsonParsable = true;
    return constructor;
}
// helper to check
export function isJsonParsable(cls) {
    return !!cls.__jsonParsable;
}
/**
 * @description Base class for network-based requests.
 * Provides common properties and methods for URL handling,
 * typed responses/messages, and cleanup.
 */
export class BaseRequest {
    constructor(MessageClass, ErrorClass) {
        this.MessageClass = MessageClass;
        this.ErrorClass = ErrorClass;
    }
    /**
     * @description Sets or replaces the URL
     * @param url - The endpoint or server URL
     */
    setUrl(url) {
        this.url = url;
        return this;
    }
    /**
     * @description Returns the current URL
     */
    getUrl() {
        return this.url;
    }
    /**
     * @description Sets a controller for aborting or cleanup
     * @param controller - AbortController or custom controller
     */
    setController(controller) {
        this.controller = controller;
        return this;
    }
    /**
     * @description Cleans up resources (connections, streams, etc.)
     * Override in subclasses if needed
     */
    cleanup() {
        if (this.controller) {
            this.controller.abort();
        }
    }
    deserialize(cls, response, raw) {
        if (!cls)
            return raw;
        const instance = new cls();
        // JsonParsable default logic
        if (isJsonParsable(cls)) {
            Object.assign(instance, typeof raw === "string" ? JSON.parse(raw) : raw);
            return instance;
        }
        // static unserialize
        if (typeof cls.unserialize === "function") {
            const result = cls.unserialize(raw, response);
            if (result !== undefined) {
                Object.assign(instance, result);
            }
        }
        return instance;
    }
    processMessage(response, raw) {
        if (!this.MessageClass)
            return raw;
        return this.deserialize(this.MessageClass, response, raw);
    }
    processError(response, err) {
        if (!this.ErrorClass)
            return err;
        return this.deserialize(this.ErrorClass, response, err);
    }
}
