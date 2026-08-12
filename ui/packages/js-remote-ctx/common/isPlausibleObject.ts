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
