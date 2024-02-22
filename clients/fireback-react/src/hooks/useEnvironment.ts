import { replacePossibleMetaPaths } from "@/components/helpers/MetaPathAwareContent";
import { FileEntity } from "src/sdk/fireback";

export enum Compiler {
  Nextjs = "nextjs",
  CRA = "cra",
  ReactNative = "reactnative",
  Unknown = "unknown",
}

export function useCompiler(): { compiler: Compiler } {
  if (process.env.RUNNING_ON_NEXT) {
    return { compiler: Compiler.Nextjs };
  }

  return { compiler: Compiler.Unknown };
}

export function useRemoteInformation() {
  const directPath = (d?: FileEntity) => {
    if (!d?.diskPath && d?.uniqueId) {
      return replacePossibleMetaPaths(d.uniqueId);
    }

    if (process.env.REACT_APP_SKIP_INLINE_FILES === "true") {
      return d?.diskPath;
    }

    return `${process.env.REACT_APP_REMOTE_FILE_SERVER}files-inline/${d?.diskPath}`;
  };

  const downloadPath = (d?: FileEntity) => {
    if (!d?.diskPath && d?.uniqueId) {
      return replacePossibleMetaPaths(d.uniqueId);
    }

    return `${process.env.REACT_APP_REMOTE_FILE_SERVER}files/${d?.diskPath}`;
  };

  return { directPath, downloadPath };
}
