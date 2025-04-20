// sometimes we want to allow users to upload small images, for example
// icons for a menu. Use this one, which only interacts svg, png files
// and saves them in database as string.
// make sure database is keeping type text instead of string

import { useT } from "../../../hooks/useT";
import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";

interface FormInlineImageProps extends BaseFormElementProps {
  onChange?: (value: string) => void;
  value?: string | any;
  label?: string;
  hint?: string;
  errorMessage?: string;
  maxFileSizeBytes?: number;
}

export const FormInlineImage = ({
  onChange,
  value,
  label,
  maxFileSizeBytes = 5 * 1024,
  ...props
}: FormInlineImageProps) => {
  const readonly = !!onChange;
  const t = useT();

  const uploadFn = (file: File) => {
    if (!file) return;

    // Check file type
    if (file.type !== "image/png" && file.type !== "image/svg+xml") {
      alert("Only PNG or SVG files are allowed.");
      return;
    }

    // Check file size
    if (file.size > maxFileSizeBytes) {
      alert("File size exceeds 5KB.");
      return;
    }

    // Read the file as a base64 string
    const reader = new FileReader();
    reader.onload = () => {
      const fileContent = reader.result as string;

      if (onChange) onChange(fileContent);
    };
    reader.readAsDataURL(file);
  };

  const onUploadDialog = () => {
    var input = document.createElement("input");
    input.type = "file";
    input.multiple = true;

    input.onchange = (e: any) => {
      const file = e.target.files?.[0];
      uploadFn(file);
    };

    input.click();
  };

  return (
    <BaseFormElement {...props}>
      {label && <label></label>}
      {value && (
        <div>
          <img
            src={value}
            alt="Selected"
            style={{ maxWidth: "200px", maxHeight: "200px" }}
          />
        </div>
      )}
      {readonly !== false && (
        <button
          className="btn btn-primary"
          type="button"
          onClick={onUploadDialog}
        >
          {t.drive.attachFile}
        </button>
      )}
    </BaseFormElement>
  );
};
