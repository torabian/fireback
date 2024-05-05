import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { Upload } from "tus-js-client";

export interface ActiveUpload {
  uploadId: string;
  bytesSent: number;
  bytesTotal: number;
  filename?: string;
}

function appendTheProgress(current: ActiveUpload[], newData: ActiveUpload) {
  const next: ActiveUpload[] = [];

  let updated = false;
  for (let item of current) {
    if (item.uploadId === newData.uploadId) {
      updated = true;
      next.push(newData);
    } else {
      next.push(item);
    }
  }

  if (updated === false) {
    next.push(newData);
  }

  return next;
}

export function useFileUploader() {
  const { session, selectedUrw, activeUploads, setActiveUploads } =
    useContext(RemoteQueryContext);

  const uploadBlob = (blob: Blob, filename: string): Promise<string>[] => {
    return upload([new File([blob], filename)]);
  };

  const upload = (
    files: File[],
    silent: boolean = false
  ): Promise<string>[] => {
    console.log("start", files);
    const result = files.map((file) => {
      return new Promise(
        (resolve: (t: string) => void, reject: (err: any) => void) => {
          const upload = new Upload(file, {
            endpoint: process.env.REACT_APP_REMOTE_FILE_SERVER + "files/",
            onBeforeRequest(req: any) {
              req.setHeader("authorization", session.token);
              req.setHeader("workspace-id", selectedUrw?.workspaceId);
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
              console.log("Success", upload);
            },
            onError(error) {
              alert(error);
              reject(error);
            },

            onProgress(bytesSent, bytesTotal) {
              const uploadId = upload.url
                ?.match(/([a-z0-9]){10,}/gi)
                ?.toString();
              if (uploadId) {
                const item: ActiveUpload = {
                  uploadId,
                  bytesSent,
                  filename: file.name,
                  bytesTotal,
                };

                if (silent !== true) {
                  setActiveUploads((activeUploads) =>
                    appendTheProgress(activeUploads, item)
                  );
                }
              }
            },
          });

          upload.start();
        }
      );
    });

    console.log("items", result);
    return result;
  };

  return { upload, activeUploads, uploadBlob };
}
