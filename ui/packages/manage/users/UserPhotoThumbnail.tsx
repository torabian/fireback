import type { CSSProperties } from "react";
import { useAuthentication } from "@fireback/auth-client";
import { AuthenticatedThumbnail } from "@fireback/ui-core/components/resumable-uploader";
import { storageDownloadUrl } from "@fireback/ui-core/hooks/useStorageUploaderConfig";

/**
 * Renders a user's `photo` field (a bare storage-module upload id, set via
 * UserEditForm's <FileUploader>) as an actual image - anywhere in the app
 * that just needs to display it, no upload widget involved (the users list
 * column, the single-user view, ...).
 *
 * Goes through AuthenticatedThumbnail rather than a plain `<img src=...>`:
 * an authenticated upload is owned (its `user_id` is set), and the storage
 * module's download route 403s a plain `<img>`'s inevitably-anonymous
 * request for an owned file - see modules/storage/README.md §6's "Caveat:
 * embedding an owned file in <img src=...>".
 */
export function UserPhotoThumbnail({
  photo,
  className,
  style,
}: {
  photo?: string | null;
  className?: string;
  style?: CSSProperties;
}) {
  const { token, selectedWorkspace } = useAuthentication();
  const src = storageDownloadUrl(photo);

  if (!src) {
    return null;
  }

  return (
    <AuthenticatedThumbnail
      src={src}
      alt=""
      className={className}
      style={style}
      headers={() => ({
        authorization: token || "",
        "workspace-id": selectedWorkspace?.workspaceId || "",
      })}
    />
  );
}
