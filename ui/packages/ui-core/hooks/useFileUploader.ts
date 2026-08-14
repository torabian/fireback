import { useState } from "react";
import { Upload } from "tus-js-client";
import { useAuthentication } from "@fireback/auth-client";

export interface ActiveUpload {
  uploadId: string;
  bytesSent: number;
  bytesTotal: number;
  filename?: string;
}

/**
 * Uploads files via tus, tagging requests with the current session's
 * token/workspace. Moved here verbatim (behavior-wise) from the old
 * `sdk/core/react-tools.tsx`'s `useFileUploader`, just re-sourced from
 * `useAuthentication()` instead of `RemoteQueryContext`.
 *
 * `activeUploads` is local to each `useFileUploader()` call rather than
 * shared app-wide state - nothing outside the old `RemoteQueryContext`
 * itself ever read the shared list, so there's no cross-component progress
 * display depending on it.
 */
export function useFileUploader() {
  const { token, selectedWorkspace } = useAuthentication();
  const [activeUploads, setActiveUploads] = useState<ActiveUpload[]>([]);

  const upload = (files: File[]): Promise<string>[] => {
    const result = files.map((file) => {
      return new Promise((resolve: (t: string) => void) => {
        const upload = new Upload(file, {
          endpoint: "http://localhost:51230/files/",
          onBeforeRequest(req: any) {
            req.setHeader("authorization", token);
            req.setHeader("workspace-id", selectedWorkspace?.workspaceId);
          },
          headers: {
            // authorization: authorization,
          },
          metadata: {
            filename: file.name,
            path: "/database/users",
            filetype: file.type,
          },
          onSuccess() {
            const uploadId = upload.url?.match(/([a-z0-9]){10,}/gi);
            resolve(`${uploadId}`);
          },

          onProgress(bytesSent, bytesTotal) {
            const uploadId = upload.url?.match(/([a-z0-9]){10,}/gi)?.toString();
            let updated = false;

            setActiveUploads((items) =>
              items?.map((item) => {
                if (item.uploadId === uploadId) {
                  updated = true;
                  return {
                    uploadId,
                    bytesSent,
                    filename: file.name,
                    bytesTotal,
                  };
                }

                return item;
              }),
            );

            if (!updated && uploadId) {
              setActiveUploads((activeUploads) => [
                ...activeUploads,
                { uploadId, bytesSent, bytesTotal, filename: file.name },
              ]);
            }
          },
        });

        upload.start();
      });
    });

    return result;
  };

  return { upload, activeUploads };
}
