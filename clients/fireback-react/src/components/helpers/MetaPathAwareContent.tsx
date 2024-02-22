/**
 * Transforms the fbtusid___ and other kind of meta path into real accessible path
 */
export function replacePossibleMetaPaths(data?: string): string {
  let value = (data || "").replaceAll(
    /fbtusid_____(.*)_____/g,
    process.env.REACT_APP_REMOTE_FILE_SERVER + "files/$1"
  );

  value = (value || "").replaceAll(
    /directasset_____(.*)_____/g,
    process.env.REACT_APP_PUBLIC_URL + "$1"
  );

  return value;
}

export function MetaPathAwareContent({ data }: { data?: string }) {
  const value = replacePossibleMetaPaths(data);
  return <span dangerouslySetInnerHTML={{ __html: value }}></span>;
}

export function getFileUrlFromTusId(tusId: string) {
  return process.env.REACT_APP_REMOTE_FILE_SERVER + "files/" + tusId;
}
