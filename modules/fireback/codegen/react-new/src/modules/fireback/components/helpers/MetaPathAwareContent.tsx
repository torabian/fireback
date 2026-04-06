import { BUILD_VARIABLES } from "../../hooks/build-variables";

/**
 * Transforms the fbtusid___ and other kind of meta path into real accessible path
 */
export function replacePossibleMetaPaths(data?: string): string {
  let value = (data || "").replaceAll(
    /fbtusid_____(.*)_____/g,
    BUILD_VARIABLES.REMOTE_SERVICE + "files/$1"
  );

  value = (value || "").replaceAll(
    /directasset_____(.*)_____/g,
    BUILD_VARIABLES.REMOTE_SERVICE + "$1"
  );

  return value;
}

export function MetaPathAwareContent({ data }: { data?: string }) {
  const value = replacePossibleMetaPaths(data);
  return <span dangerouslySetInnerHTML={{ __html: value }}></span>;
}

export function getFileUrlFromTusId(tusId: string) {
  return BUILD_VARIABLES.REMOTE_SERVICE + "files/" + tusId;
}
