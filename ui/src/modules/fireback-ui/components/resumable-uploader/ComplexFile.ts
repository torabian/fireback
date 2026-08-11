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
  if (!url) return null;
  try {
    const { pathname } = new URL(url);
    return pathname.split("/").filter(Boolean).pop() ?? null;
  } catch {
    // Bug fix: this used to return null here, silently discarding
    // whatever was passed in. FileUploader's own onChange hands callers
    // back a ComplexFile whose .id is already just the bare id (this
    // function's happy-path result, not the full upload URL) - so the
    // obvious, expected round trip of storing that id and later passing
    // it straight back in as `value` (exactly what the README's own
    // "Previewing an existing value" example shows: `value={record.
    // avatarUploadId}`) hit this catch branch, since a bare id like
    // "a1b2c3d4e5" isn't a parseable absolute URL. That turned .id into
    // null for every reopened form, breaking the thumbnail and the
    // clear/replace controls. Treat "not a URL" as "already the id"
    // instead of discarding it.
    return url;
  }
}
