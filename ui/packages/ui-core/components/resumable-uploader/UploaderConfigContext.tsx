import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import type { ReactNode } from "react";
import { Upload } from "tus-js-client";
import { validateFileAgainstRules } from "./fileValidation";
import { defaultUploadTranslations, mergeTranslations } from "./translations";
import type { UploaderConfig, UploadItem, UploadTranslations } from "./types";

interface UploaderContextValue {
  config: UploaderConfig;
  translations: UploadTranslations;
  items: UploadItem[];
  isOnline: boolean;
  addFiles: (files: File[], ownerId: string) => void;
  pauseItem: (id: string) => void;
  resumeItem: (id: string) => void;
  retryItem: (id: string) => void;
  cancelItem: (id: string) => void;
  removeItem: (id: string) => void;
}

const UploaderConfigContext = createContext<UploaderContextValue | null>(null);

function makeId(): string {
  if (typeof crypto !== "undefined" && "randomUUID" in crypto) {
    return crypto.randomUUID();
  }
  return `${Date.now()}-${Math.random().toString(16).slice(2)}`;
}

function isPreviewable(mime: string): boolean {
  return (
    mime.startsWith("image/") ||
    mime.startsWith("video/") ||
    mime.startsWith("audio/") ||
    mime === "application/pdf"
  );
}

async function resolveHeaders(
  config: UploaderConfig,
): Promise<Record<string, string>> {
  if (!config.headers) return {};
  if (typeof config.headers === "function") {
    return (await config.headers()) || {};
  }
  return config.headers;
}

export function UploaderConfigProvider({
  config,
  children,
}: {
  config: UploaderConfig;
  children: ReactNode;
}) {
  const translations = useMemo(
    () => mergeTranslations(config.translations, config.locale),
    [config.translations, config.locale],
  );

  const [items, setItems] = useState<UploadItem[]>([]);
  const [isOnline, setIsOnline] = useState(
    typeof navigator === "undefined" ? true : navigator.onLine,
  );

  // Keep a mutable mirror of `items` so the online/offline handlers (attached
  // once) always see the latest list without having to re-subscribe.
  const itemsRef = useRef<UploadItem[]>(items);
  useEffect(() => {
    itemsRef.current = items;
  }, [items]);

  // tus Upload instances live outside React state: they are stateful objects
  // with their own internal offset tracking, which is what lets us call
  // .abort() / .start() again on the very same instance to pause/resume.
  const uploadsRef = useRef<Map<string, Upload>>(new Map());
  const manuallyPausedRef = useRef<Set<string>>(new Set());
  const configRef = useRef(config);
  configRef.current = config;

  const updateItem = useCallback((id: string, patch: Partial<UploadItem>) => {
    setItems((prev) =>
      prev.map((it) => (it.id === id ? { ...it, ...patch } : it)),
    );
  }, []);

  const startUpload = useCallback(
    async (item: UploadItem) => {
      const cfg = configRef.current;

      const upload = new Upload(item.file, {
        endpoint: cfg.endpoint,
        // tus-js-client merges options via `{ ...defaultOptions, ...options }`,
        // so an explicit `chunkSize: undefined` key here would override its
        // own Infinity default instead of falling back to it, then get cast
        // via `Number(undefined)` into NaN - which silently breaks the
        // browser file source's slice() bounds (every chunk reads back empty
        // without ever reporting "done", so the upload spins forever sending
        // 0-byte PATCH requests with no error). Must always resolve to a
        // real, finite number.
        chunkSize: cfg.chunkSize ?? Number.POSITIVE_INFINITY,
        retryDelays: cfg.retryDelays ?? [0, 1000, 3000, 5000, 10000],
        metadata: {
          filename: item.file.name,
          filetype: item.file.type || "application/octet-stream",
          ...cfg.metadata,
        },
        // Bug fix: this used to *also* pass a `headers` option here (a
        // snapshot resolved once up front), on top of onBeforeRequest below
        // re-resolving and re-applying the same headers again for every
        // request tus-js-client makes - including this same upload's very
        // first creation POST, which goes through both addRequiredHeaders
        // (the static `headers` option) and onBeforeRequest on the same
        // underlying XMLHttpRequest. XMLHttpRequest.setRequestHeader()
        // *appends*, comma-joined, when called twice for the same header
        // name rather than overwriting - so e.g. "authorization" arrived at
        // the server as "<token>, <token>", which the backend's token
        // lookup obviously can't match to a real session, 401ing every
        // upload for any config that set an authorization header (i.e. any
        // authenticated use of this module at all). onBeforeRequest alone
        // already covers every request, so the static option was not just
        // redundant but actively wrong - removed instead of deduplicated.
        onBeforeRequest: async (req) => {
          const freshHeaders = await resolveHeaders(configRef.current);
          for (const [key, value] of Object.entries(freshHeaders)) {
            req.setHeader(key, value);
          }
        },
        onProgress: (bytesSent, bytesTotal) => {
          updateItem(item.id, { bytesUploaded: bytesSent, bytesTotal });
        },
        onSuccess: () => {
          uploadsRef.current.delete(item.id);
          updateItem(item.id, {
            status: "completed",
            uploadUrl: upload.url,
            fileId: extractHash(upload.url),
          });
        },
        onError: (error) => {
          // A disconnect surfaces here as a request error. If we are already
          // offline, treat it as a pause rather than a failure - the
          // online-listener below will restart this same instance later.
          const offline = typeof navigator !== "undefined" && !navigator.onLine;
          if (offline) {
            updateItem(item.id, { status: "offline-paused" });
            return;
          }
          updateItem(item.id, {
            status: "error",
            error: error.message ?? String(error),
          });
        },
      });

      uploadsRef.current.set(item.id, upload);
      updateItem(item.id, { status: "uploading", error: null });

      try {
        const previousUploads = await upload.findPreviousUploads();
        if (previousUploads.length > 0) {
          upload.resumeFromPreviousUpload(previousUploads[0]);
        }
      } catch {
        // URL storage lookup failing (e.g. localStorage unavailable) just
        // means we start a fresh upload instead of resuming one.
      }

      upload.start();
    },
    [updateItem],
  );

  const addFiles = useCallback(
    (files: File[], ownerId: string) => {
      for (const file of files) {
        const failure = validateFileAgainstRules(
          file,
          configRef.current.validateFile,
        );
        if (failure) {
          const message =
            failure.code === "size"
              ? translations.fileTooLarge(
                  Math.round(failure.maxSize / 1024 / 1024),
                )
              : translations.fileTypeNotAllowed;
          setItems((prev) => [
            ...prev,
            {
              id: makeId(),
              ownerId,
              file,
              status: "error",
              bytesUploaded: 0,
              bytesTotal: file.size,
              uploadUrl: null,
              error: message,
              previewUrl: null,
            },
          ]);
          continue;
        }

        const item: UploadItem = {
          id: makeId(),
          ownerId,
          file,
          status: "queued",
          bytesUploaded: 0,
          bytesTotal: file.size,
          uploadUrl: null,
          error: null,
          previewUrl: isPreviewable(file.type)
            ? URL.createObjectURL(file)
            : null,
        };

        setItems((prev) => [...prev, item]);
        void startUpload(item);
      }
    },
    [startUpload, translations],
  );

  const pauseItem = useCallback(
    (id: string) => {
      manuallyPausedRef.current.add(id);
      uploadsRef.current.get(id)?.abort();
      updateItem(id, { status: "paused" });
    },
    [updateItem],
  );

  const resumeItem = useCallback(
    (id: string) => {
      manuallyPausedRef.current.delete(id);
      const upload = uploadsRef.current.get(id);
      if (upload) {
        upload.start();
        updateItem(id, { status: "uploading" });
      }
    },
    [updateItem],
  );

  const retryItem = useCallback(
    (id: string) => {
      manuallyPausedRef.current.delete(id);
      const item = itemsRef.current.find((it) => it.id === id);
      if (item) void startUpload(item);
    },
    [startUpload],
  );

  const cancelItem = useCallback(
    (id: string) => {
      uploadsRef.current.get(id)?.abort();
      uploadsRef.current.delete(id);
      manuallyPausedRef.current.delete(id);
      updateItem(id, { status: "canceled" });
    },
    [updateItem],
  );

  const removeItem = useCallback((id: string) => {
    const item = itemsRef.current.find((it) => it.id === id);
    if (item?.previewUrl) URL.revokeObjectURL(item.previewUrl);
    uploadsRef.current.get(id)?.abort();
    uploadsRef.current.delete(id);
    manuallyPausedRef.current.delete(id);
    setItems((prev) => prev.filter((it) => it.id !== id));
  }, []);

  // Connectivity handling: pause in-flight uploads the moment the browser
  // reports "offline" (rather than waiting for requests to time out), and
  // transparently resume every auto-paused upload once back "online" - this
  // is what lets an upload keep going across a network drop without the
  // user having to do anything.
  useEffect(() => {
    const handleOffline = () => {
      setIsOnline(false);
      for (const item of itemsRef.current) {
        if (item.status === "uploading") {
          uploadsRef.current.get(item.id)?.abort();
          updateItem(item.id, { status: "offline-paused" });
        }
      }
    };

    const handleOnline = () => {
      setIsOnline(true);
      for (const item of itemsRef.current) {
        if (
          item.status === "offline-paused" &&
          !manuallyPausedRef.current.has(item.id)
        ) {
          const upload = uploadsRef.current.get(item.id);
          if (upload) {
            upload.start();
            updateItem(item.id, { status: "uploading" });
          } else {
            void startUpload(item);
          }
        }
      }
    };

    window.addEventListener("offline", handleOffline);
    window.addEventListener("online", handleOnline);
    return () => {
      window.removeEventListener("offline", handleOffline);
      window.removeEventListener("online", handleOnline);
    };
  }, [startUpload, updateItem]);

  const value: UploaderContextValue = {
    config,
    translations,
    items,
    isOnline,
    addFiles,
    pauseItem,
    resumeItem,
    retryItem,
    cancelItem,
    removeItem,
  };

  return (
    <UploaderConfigContext.Provider value={value}>
      {children}
    </UploaderConfigContext.Provider>
  );
}

export function useUploaderConfig(): UploaderContextValue {
  const ctx = useContext(UploaderConfigContext);
  if (!ctx) {
    throw new Error(
      "useUploaderConfig must be used within a <UploaderConfigProvider>",
    );
  }
  return ctx;
}

export { defaultUploadTranslations };

function extractHash(url: string): string | null {
  try {
    const { pathname } = new URL(url);
    return pathname.split("/").filter(Boolean).pop() ?? null;
  } catch {
    return null;
  }
}
