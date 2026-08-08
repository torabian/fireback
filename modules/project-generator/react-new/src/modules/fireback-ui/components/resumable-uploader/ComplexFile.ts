/** Plain-object shape a ComplexFile can be built from or serialized back into. */
export interface ComplexFileData {
  id: string;
  filesize?: number;
  thumbnail?: string;
  filename?: string;
  mimeType?: string;
  [key: string]: unknown;
}

export type ComplexFileInput = string | ComplexFileData | ComplexFile;

/**
 * Represents the value a file field holds: either just an id (e.g. a plain
 * tus upload id already stored on a record) or a richer object carrying
 * whatever metadata the backend also returned (filesize, thumbnail, ...).
 * FileUploader constructs one of these from a completed upload and hands it
 * to onChange, so the form always deals with one consistent shape whether
 * the value came from a fresh upload or from re-opening an existing record.
 */
export class ComplexFile implements ComplexFileData {
  id: string;
  filesize?: number;
  thumbnail?: string;
  filename?: string;
  mimeType?: string;
  [key: string]: unknown;

  constructor(value: ComplexFileInput) {
    if (typeof value === "string") {
      this.id = extractHash(value);
      return;
    }
    Object.assign(this, { ...value, id: extractHash(value.id) });
  }

  toJSON(): ComplexFileData {
    return { ...this };
  }

  toString(): string {
    return extractHash(this.id);
  }
}

function extractHash(url: string): string | null {
  try {
    const { pathname } = new URL(url);
    return pathname.split("/").filter(Boolean).pop() ?? null;
  } catch {
    return null;
  }
}
