import { useFileListener } from "../../window-drop/WindowDrop";
import { useFileUploader } from "@/modules/fireback/modules/drive/DriveTools";
import { useRemoteInformation } from "@/modules/fireback/hooks/useEnvironment";
import { useT } from "@/modules/fireback/hooks/useT";
import { FileEntity } from "../../../sdk/modules/workspaces/FileEntity";
import { debounce } from "lodash";
import { useRef } from "react";

interface FormUploaderProps {
  onChange?: (value: FileEntity[]) => void;
  value?: FileEntity[] | any;
}

function AttachmentViewer({ attachments }: { attachments: FileEntity[] }) {
  const { directPath, downloadPath } = useRemoteInformation();

  return (
    <div className="file-viewer-files">
      {(attachments || []).map((attachment) => {
        return (
          <div className="file-viewer-file" key={attachment.uniqueId}>
            <span className="file-viewer-type">{attachment.type}</span>
            <span className="file-viewer-size">{attachment.size}</span>
            <span className="file-viewer-name">{attachment.name}</span>
            <div>
              <a
                target="_blank"
                rel="noreferrer"
                referrerPolicy="no-referrer"
                href={directPath(attachment) || ""}
                className="btn"
              >
                View
              </a>
              <a href={downloadPath(attachment)} className="btn">
                Download
              </a>
            </div>
          </div>
        );
      })}
    </div>
  );
}

export const FormUploader = ({ onChange, value }: FormUploaderProps) => {
  const readonly = !!onChange;
  const { upload } = useFileUploader();
  const data = useRef<FileEntity[]>([]);
  const t = useT();
  // Use debounced onChange
  const onChangeDebounced = debounce(
    (items: FileEntity[]) => {
      onChange && onChange(items);
    },
    250,
    { maxWait: 1000 }
  );

  const uploadFn = (files: File[]) => {
    const items = upload(files);
    items.forEach((item) => {
      item.then((x) => {
        data.current.push({ uniqueId: x } as any);
        onChangeDebounced(data.current);
      });
    });

    return items;
  };

  useFileListener({
    enabled: !readonly,
    label: "Attach documents about the payment",
    extentions: ["*"],
    onCaptureFile(files) {
      alert("hi");
      Promise.all(uploadFn(files)).then((result) => {});
    },
  });

  const onUploadDialog = () => {
    var input = document.createElement("input");
    input.type = "file";
    input.multiple = true;

    input.onchange = (e: any) => {
      Promise.all(uploadFn(Array.from(e.target.files))).then((result) => {});
    };

    input.click();
  };

  return (
    <div>
      {readonly !== false && (
        <button
          className="btn btn-primary"
          type="button"
          onClick={onUploadDialog}
        >
          {t.drive.attachFile}
        </button>
      )}
      <AttachmentViewer attachments={value || []} />
    </div>
  );
};
