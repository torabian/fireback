import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";

export function UploaderStatsCard() {
  const { activeUploads, setActiveUploads } = useContext(RemoteQueryContext);

  const discardActiveUploads = () => {
    setActiveUploads([]);
  };

  if (activeUploads.length === 0) {
    return null;
  }

  return (
    <div className="active-upload-box">
      <div className="upload-header">
        <span>{activeUploads.length} Uploads</span>
        <span className="action-section">
          <button onClick={discardActiveUploads}>
            <img src="/common/close.svg" />
          </button>
        </span>
      </div>
      {activeUploads.map((item) => (
        <div key={item.uploadId} className="upload-file-item">
          <span>{item.filename}</span>
          <span>{Math.ceil((item.bytesSent / item.bytesTotal) * 100)}%</span>
        </div>
      ))}
    </div>
  );
}
