import type { UploadTranslations } from "../types";

export const en: UploadTranslations = {
  attachFile: "Attach file",
  dropHere: "Drop files here, or click to browse",
  browse: "Browse",
  queued: "Queued",
  uploading: "Uploading",
  paused: "Paused",
  offlinePaused: "Paused (offline)",
  completed: "Completed",
  failed: "Failed",
  canceled: "Canceled",
  pause: "Pause",
  resume: "Resume",
  retry: "Retry",
  remove: "Remove",
  clear: "Clear",
  replace: "Replace",
  currentFile: "Current file",
  offlineNotice:
    "You are offline. Uploads are paused and will resume automatically once your connection is back.",
  onlineResuming: "Connection restored, resuming uploads...",
  fileTooLarge: (maxMB: number) => `File too large. Max allowed size is ${maxMB}MB.`,
  fileTypeNotAllowed: "File type not allowed.",
  maxSizeLabel: (size: string) => `Max file size: ${size}`,
};
