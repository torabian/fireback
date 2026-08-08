import { useEffect, useState } from "react";
import "./resumable-uploader.css";

const TEXT_PREVIEW_MAX_BYTES = 200 * 1024;

function fileExtensionLabel(filename: string): string {
  const idx = filename.lastIndexOf(".");
  return idx >= 0 ? filename.slice(idx + 1).toUpperCase() : "FILE";
}

export function FilePreview({
  file,
  previewUrl,
}: {
  file: File;
  previewUrl: string | null;
}) {
  const mime = file.type || "";
  const isTextLike = mime.startsWith("text/") || mime === "application/json";
  const [textContent, setTextContent] = useState<string | null>(null);

  useEffect(() => {
    if (!isTextLike || file.size > TEXT_PREVIEW_MAX_BYTES) {
      setTextContent(null);
      return;
    }
    let cancelled = false;
    file.text().then((content) => {
      if (!cancelled) setTextContent(content);
    });
    return () => {
      cancelled = true;
    };
  }, [file, isTextLike]);

  if (mime.startsWith("image/") && previewUrl) {
    return <img src={previewUrl} alt={file.name} className="ru-preview-image" />;
  }

  if (mime.startsWith("video/") && previewUrl) {
    return <video src={previewUrl} controls className="ru-preview-video" />;
  }

  if (mime.startsWith("audio/") && previewUrl) {
    return <audio src={previewUrl} controls className="ru-preview-audio" />;
  }

  if (mime === "application/pdf" && previewUrl) {
    return <iframe src={previewUrl} title={file.name} className="ru-preview-pdf" />;
  }

  if (textContent !== null) {
    return <pre className="ru-preview-text">{textContent}</pre>;
  }

  return <div className="ru-preview-generic">{fileExtensionLabel(file.name)}</div>;
}
