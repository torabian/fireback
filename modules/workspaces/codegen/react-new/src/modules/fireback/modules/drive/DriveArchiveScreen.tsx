import { useT } from "@/modules/fireback/hooks/useT";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { useFileListener } from "@/modules/fireback/components/window-drop/WindowDrop";
import { DriveList } from "./DriveList";
import { useFileUploader } from "./DriveTools";
import { useQueryClient } from "react-query";

export const DriveArchiveScreen = () => {
  const t = useT();
  const { upload } = useFileUploader();
  const queryClient = useQueryClient();

  const handleUpload = (files: File[]) => {
    Promise.all(upload(files))
      .then((result) => {
        queryClient.invalidateQueries("*drive.FileEntity");
      })
      .catch((err) => {
        alert(err);
      });
  };

  useFileListener({
    label: "Add files or documents to drive",
    extentions: ["*"],
    onCaptureFile(files) {
      handleUpload(files);
    },
  });

  const onUploadDialog = () => {
    var input = document.createElement("input");
    input.type = "file";

    input.onchange = (e: any) => {
      handleUpload(Array.from(e.target.files));
    };

    input.click();
  };

  return (
    <CommonArchiveManager
      pageTitle={t.drive.driveTitle}
      newEntityHandler={() => {
        onUploadDialog();
      }}
    >
      <DriveList />
    </CommonArchiveManager>
  );
};
