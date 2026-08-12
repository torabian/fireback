/**
 * Used in fetch context, to detect if the response is not a buffer
 * or arraybufer, then create class instance of it.
 * In such cases when response is blob, its better to pass the original classes to caller.
 * @param obj
 * @returns
 */
export const isPlausibleObject = (obj: any) => {
  const isBuffer =
    typeof globalThis.Buffer !== "undefined" &&
    typeof globalThis.Buffer.isBuffer === "function" &&
    globalThis.Buffer.isBuffer(obj);

  const isBlob =
    typeof globalThis.Blob !== "undefined" && obj instanceof globalThis.Blob;

  return (
    obj &&
    typeof obj === "object" &&
    !Array.isArray(obj) &&
    !isBuffer &&
    !(obj instanceof ArrayBuffer) &&
    !isBlob
  );
};
