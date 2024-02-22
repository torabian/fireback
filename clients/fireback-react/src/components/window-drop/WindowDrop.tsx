import { getOS } from "@/hooks/useHtmlClass";
import { useT } from "@/hooks/useT";
import React, { useContext, useEffect, useState } from "react";
import Dropzone from "react-dropzone";

export interface FileListenerRef {
  id: string;
  condition: FileListernerCondition;
}

export function useFileListener(con: FileListernerCondition) {
  const fl = useContext(FileListenerContext);
  useEffect(() => {
    const r = fl.listenFiles(con);
    return () => fl.removeSubscription(r);
  }, []);
}

export interface FileListernerCondition {
  extentions: [string];
  label: string;
  onCaptureFile: (files: File[]) => void;
  enabled?: boolean;
}

export interface IFileListenerContext {
  listenFiles: (x: FileListernerCondition) => string;
  removeSubscription: (refId: string) => void;
  refs: Array<FileListenerRef>;
}

export const FileListenerContext = React.createContext<IFileListenerContext>({
  listenFiles() {
    return "";
  },
  removeSubscription() {},
  refs: [],
});

export function FileListenerProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [modalRefs, setModalRefs] = useState<Array<FileListenerRef>>([]);
  const listenFiles = (condition: FileListernerCondition) => {
    let r = (Math.random() + 1).toString(36).substring(2);

    const newRef: FileListenerRef = { id: r, condition };
    setModalRefs((r) => [...r, newRef]);

    return r;
  };

  const removeSubscription = (refId: string) => {
    setModalRefs((r) => r.filter((t) => t.id !== refId));
  };

  return (
    <FileListenerContext.Provider
      value={{
        refs: modalRefs,
        removeSubscription,
        listenFiles,
      }}
    >
      {children}
    </FileListenerContext.Provider>
  );
}

export function WindowDrop({ children }: { children: React.ReactNode }) {
  const [filesToSettle, setFilesToSettle$] = useState<File[]>([]);
  const { refs } = useContext(FileListenerContext);
  const t = useT();

  const setFilesToSettle = (files: File[]) => {
    setFilesToSettle$(files);

    for (const ref of refs) {
      if (ref.condition?.onCaptureFile) {
        ref.condition?.onCaptureFile(files);
      }
    }
  };

  if (getOS() === "ios") {
    return <>{children}</>;
  }

  return (
    <>
      <Dropzone
        onDrop={(acceptedFiles) => {
          setFilesToSettle(acceptedFiles);
        }}
        noClick
        noKeyboard
      >
        {({ getRootProps, getInputProps, isDragActive }) => (
          <div {...getRootProps()} style={{}}>
            {isDragActive && (
              <div
                style={{ flexDirection: "column" }}
                className="file-dropping-indicator"
              >
                <span className="dropin-files-hint">
                  {t.dropNFiles.replace("{n}", `${refs.length}`)}
                </span>
                {refs.map((r) => (
                  <span key={r.id}>{r.condition.label}</span>
                ))}
              </div>
            )}

            <>{children}</>
          </div>
        )}
      </Dropzone>
    </>
  );
}
