import { useT } from "../../../hooks/useT";
import { useFileUploader } from "../../../modules/drive/DriveTools";

import { debounce } from "lodash";
import { useRef } from "react";
import { useFileListener } from "../../window-drop/WindowDrop";
import { useRemoteInformation } from "../../../hooks/useEnvironment";
import { FileEntity } from "@/modules/fireback/sdk/modules/workspaces/FileEntity";

export interface ImageAsset {
  width?: number | undefined;
  height?: number | undefined;
  attachment?: FileEntity | undefined;
  attachmentId?: string | undefined;
}

export interface ImageAssetPlaceholder {
  width: number;
  height: number;
  title?: string;
  description?: string;
}

interface FormImageAssetProps {
  onChange?: (value: ImageAsset[]) => void;
  value?: ImageAsset[];
  label?: string;
  hint?: string;
  placeHolders: ImageAssetPlaceholder[];
}

export const FormImageAsset = ({
  onChange,
  value,
  label,
  placeHolders,
}: FormImageAssetProps) => {
  const { directPath } = useRemoteInformation();
  const readonly = !!onChange;
  const { upload } = useFileUploader();
  const data = useRef<FileEntity[]>([]);
  const t = useT();
  // Use debounced onChange
  const onChangeDebounced = debounce(
    (items: FileEntity[]) => {
      // onChange && onChange(items);
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

  const itemExtractor = (
    items: ImageAsset[],
    width: number,
    height: number
  ) => {
    return (items || []).find(
      (item) => item.width === width && item.height === height
    );
  };

  return (
    <div>
      {label && <label>{label}</label>}

      <div className="image-assets">
        {placeHolders.map((x) => {
          const { width, height } = calculateProportionalSize(
            x.width,
            x.height
          );

          const item = itemExtractor(value, x.width, x.height);

          if (item && item.attachment) {
            return (
              <div
                className="image-asset-item"
                style={{
                  fontSize: "11px",
                  display: "flex",
                  width: `${width}px`,
                  height: `${height}px`,
                  flexDirection: "column",
                }}
                key={`${x.width}_${x.height}`}
              >
                {x.width}x{x.height}
                <img
                  style={{ width: "100%" }}
                  src={directPath(item.attachment as any)}
                />
              </div>
            );
          }

          return (
            <div
              className="image-asset-item"
              style={{
                width: `${width}px`,
                height: `${height}px`,
              }}
              key={`${x.width}_${x.height}`}
            >
              {x.width}x{x.height}
            </div>
          );
        })}
      </div>
    </div>
  );
};

const MAX_SIZE = 180;

const calculateProportionalSize = (width, height) => {
  const scaleFactor = MAX_SIZE / Math.max(width, height);
  return {
    width: Math.round(width * scaleFactor),
    height: Math.round(height * scaleFactor),
  };
};
