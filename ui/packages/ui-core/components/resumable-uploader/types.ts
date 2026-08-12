export type UploadStatus =
  | "queued"
  | "uploading"
  | "paused"
  | "offline-paused"
  | "completed"
  | "error"
  | "canceled";

export interface FileValidationRule {
  mimeStartsWith?: string; // e.g. 'image/', 'application/pdf'
  extension?: string; // e.g. '.jpg', '.docx'
  maxSize?: number; // bytes
}

export interface UploadItem {
  id: string;
  /**
   * Identifies which <FileUploader> instance added this item, so multiple
   * uploaders sharing one <UploaderConfigProvider> (e.g. several fields on
   * the same form) each only see and react to their own uploads instead of
   * the entire shared queue.
   */
  ownerId: string;
  file: File;
  fileId?: string;
  status: UploadStatus;
  bytesUploaded: number;
  bytesTotal: number;
  uploadUrl: string | null;
  error: string | null;
  previewUrl: string | null;
}

export type UploaderLocale = "en" | "fa";

export type HeaderProvider =
  | Record<string, string>
  | (() => Record<string, string> | Promise<Record<string, string>>);

export interface UploadTranslations {
  attachFile: string;
  dropHere: string;
  browse: string;
  queued: string;
  uploading: string;
  paused: string;
  offlinePaused: string;
  completed: string;
  failed: string;
  canceled: string;
  pause: string;
  resume: string;
  retry: string;
  remove: string;
  clear: string;
  replace: string;
  currentFile: string;
  offlineNotice: string;
  onlineResuming: string;
  fileTooLarge: (maxMB: number) => string;
  fileTypeNotAllowed: string;
  maxSizeLabel: (size: string) => string;
}

export interface UploaderConfig {
  /** tus server endpoint, e.g. "http://localhost:4500/storage/files" */
  endpoint: string;
  /** static headers, or a (sync/async) function evaluated before every request - useful for tokens that can expire mid-upload */
  headers?: HeaderProvider;
  metadata?: Record<string, string>;
  chunkSize?: number;
  retryDelays?: number[];
  validateFile?: FileValidationRule[];
  /** Selects the built-in translation set ("en" or "fa"). Ignored for keys overridden in `translations`. */
  locale?: UploaderLocale;
  translations?: Partial<UploadTranslations>;
  /**
   * Given the value previously stored for this field (the tus upload id/URL
   * returned from a past upload), returns a URL the backend serves a
   * thumbnail/preview image from. Used to render a preview when a form is
   * reopened later with an existing value but no local File object.
   */
  getThumbnailUrl?: (value: string) => string;
}
