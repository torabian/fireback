import { debounce } from "lodash";
import { useEffect, useRef, useState } from "react";
import { useT } from "../../../hooks/useT";
import { useFileUploader } from "../../../modules/manage/drive/DriveTools";
import { useFileListener } from "../../window-drop/WindowDrop";

export type XFile =
  | {
      base64?: string;
      blob?: string;
    }
  | string;

interface FileValidationRule {
  mimeStartsWith?: string; // e.g. 'image/', 'application/pdf'
  extension?: string; // e.g. '.jpg', '.docx'
  maxSize?: number; // bytes
}

function buildAcceptString(rules: FileValidationRule[]): string {
  const parts = rules.map((r) => {
    if (r.extension) return r.extension;
    if (r.mimeStartsWith?.endsWith("/")) return r.mimeStartsWith + "*";
    return r.mimeStartsWith || "";
  });

  return parts.filter(Boolean).join(",");
}

function validateFileAgainstRules(
  file: File,
  rules: FileValidationRule[]
): string | null {
  for (const rule of rules) {
    const matchesType =
      !rule.mimeStartsWith || file.type.startsWith(rule.mimeStartsWith);
    if (matchesType) {
      if (file.size > rule.maxSize) {
        return `File too large. Max allowed size is ${Math.round(
          rule.maxSize / 1024 / 1024
        )}MB.`;
      }
      return null; // valid
    }
  }
  return "File type not allowed.";
}

interface FormXFileProps {
  onChange?: (file: XFile) => void;
  value?: XFile | null;
  label?: string;
  hint?: string;
  validateFile?: FileValidationRule[];
}

const MAX_INLINE_SIZE = 5 * 1024 * 1024;

function isBase64DataUrl(value: string): boolean {
  return value.startsWith("data:");
}

function encodeBase64(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = reject;
    reader.readAsDataURL(file);
  });
}

function detectInitialValueType(val: any): "base64" | "tus" | "unknown" {
  if (typeof val === "string" && isBase64DataUrl(val)) return "base64";
  if (typeof val === "object" && val?.tusId) return "tus";
  return "unknown";
}

function getFileTypeFromBase64(base64) {
  const match = base64.match(/^data:(.+?);base64,/);
  return match ? match[1] : null;
}

export const FormXFile = ({
  onChange,
  value,
  label,
  validateFile,
}: FormXFileProps) => {
  const t = useT();
  const { uploadSingle } = useFileUploader(); // assumes a custom tus uploader
  const [uploadError, setUploadError] = useState();
  const inputRef = useRef<HTMLInputElement | null>(null);
  const [file, setFile] = useState<File | null>(null);
  let [previewUrl, setPreviewUrl] = useState<string | null>(null);
  const [textContent, setTextContent] = useState<string | null>(null);

  const handleFiles = async (files: File[]) => {
    const file = files[0];
    if (!file) return;

    if (validateFile?.length) {
      const error = validateFileAgainstRules(file, validateFile);
      if (error) {
        alert(error);
        return;
      }
    }

    const mime = file.type;
    setFile(file);

    // Text files
    if (mime.startsWith("text/") || mime === "application/json") {
      console.log("asd");
      const reader = new FileReader();
      reader.onload = () => {
        setTextContent(reader.result as string);
      };
      reader.readAsText(file);
    } else {
      const url = URL.createObjectURL(file);
      setPreviewUrl(url);
    }

    if (file.size > MAX_INLINE_SIZE) {
      try {
        const tusResult = await uploadSingle(file as any); // should return { tusId, filename, mime }
        console.log("Upload completed", tusResult);
        onChange?.(tusResult as any);
      } catch (err) {
        setUploadError(err);
        console.log(4, err);
      }
    } else {
      const base64 = await encodeBase64(file);
      onChange?.(base64 as any);
    }
  };

  useEffect(() => {
    return () => {
      if (previewUrl) URL.revokeObjectURL(previewUrl);
    };
  }, [previewUrl]);

  const openFileDialog = () => {
    if (!inputRef.current) {
      inputRef.current = document.createElement("input");
      inputRef.current.type = "file";

      if (validateFile?.length) {
        inputRef.current.accept = buildAcceptString(validateFile);
      }

      inputRef.current.onchange = (e: any) => {
        const files = Array.from(e.target.files || []);
        handleFiles(files as any);
      };
    }
    inputRef.current.click();
  };

  useFileListener({
    enabled: !!onChange,
    label: "Drop a file here",
    extentions: ["*"],
    onCaptureFile(files) {
      handleFiles(files);
    },
  });

  let fileType = file?.type || "";

  if (!previewUrl && typeof value === "string") {
    previewUrl = value;
  } else if (!previewUrl && typeof value === "object") {
    previewUrl = value.blob;
    fileType = getFileTypeFromBase64(previewUrl) || "";
  }

  return (
    <div className="space-y-2">
      {label && <label className="block font-semibold">{label}</label>}
      {onChange && (
        <button
          type="button"
          className="btn btn-primary"
          onClick={openFileDialog}
        >
          {t.drive.attachFile}
        </button>
      )}
      {uploadError ? <pre>Error: {(uploadError as any).toString()}</pre> : null}
      <div className="space-y-2">
        {file && (
          <div className="mt-4">
            <strong>{file.name}</strong> ({Math.round(file.size / 1024)} KB)
          </div>
        )}

        {fileType.startsWith("image/") && previewUrl && (
          <img
            src={previewUrl}
            alt="Preview"
            className="max-w-full max-h-80"
            style={{ maxWidth: "50%" }}
          />
        )}

        {fileType.startsWith("video/") && previewUrl && (
          <video src={previewUrl} controls className="max-w-full max-h-96" />
        )}

        {fileType.startsWith("audio/") && previewUrl && (
          <audio src={previewUrl} controls className="w-full" />
        )}

        {fileType === "application/pdf" && previewUrl && (
          <iframe
            src={previewUrl}
            className="w-full h-96 border rounded"
            title="PDF Preview"
          />
        )}

        {textContent && (
          <pre className="bg-gray-100 p-2 rounded max-h-96 overflow-auto text-sm">
            {textContent}
          </pre>
        )}
      </div>
    </div>
  );
};
