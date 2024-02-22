import { Editor } from "@tinymce/tinymce-react";

import classNames from "classnames";
import React, {
  useCallback,
  useContext,
  useEffect,
  useRef,
  useState,
} from "react";

import {
  BaseFormElement,
  BaseFormElementProps,
} from "../base-form-element/BaseFormElement";
import { AppConfigContext } from "@/hooks/appConfigTools";
import { LineLoader } from "@/components/line-loader/LineLoader";
import { useFileUploader } from "@/modules/drive/DriveTools";
import { Blob } from "buffer";
import { useRemoteInformation } from "@/hooks/useEnvironment";
import { addKatexSupport } from "./katex";
import { useT } from "@/hooks/useT";

export interface FormRichTextProps extends BaseFormElementProps {
  placeholder?: string;
  label?: string;
  disabled?: boolean;
  onChange?: (value: string) => void;
  secureTextEntry?: boolean;
  Icon?: any;
  errorMessage?: string;
  autoFocus?: boolean;
  validMessage?: string;
  onReady?: () => void;
  value?: any | null;
  className?: string;
  dir?: string;
  focused?: boolean;
  forceBasic?: boolean;
  forceRich?: boolean;
  height?: number | string;
  getInputRef?: (ref: any) => void;
}

export enum EditorTypes {
  TinyMCE = "tinymce",
  TextArea = "textarea",
}

export const FormRichText = (props: FormRichTextProps) => {
  const { config } = useContext(AppConfigContext);
  const t = useT();
  const {
    placeholder,
    label,
    getInputRef,
    secureTextEntry,
    Icon,
    onChange,
    value,
    height,
    disabled,
    forceBasic,
    forceRich,
    focused: f = false,
    autoFocus,
    ...restProps
  } = props;

  const [focused, setFocused] = useState(false);
  const ref = useRef<Editor | null>();
  const isTinyMceLoaded = useRef<boolean>(false);
  const [editorType, setEditorType] = useState(EditorTypes.TinyMCE);

  const { upload } = useFileUploader();
  const { directPath } = useRemoteInformation();

  useEffect(() => {
    if (config.textEditorModule !== "tinymce") {
      props.onReady && props.onReady();
    } else {
      const timeout = setTimeout(() => {
        if (isTinyMceLoaded.current === false) {
          setEditorType(EditorTypes.TextArea);
          props.onReady && props.onReady();
        }
      }, 5000);

      return () => {
        clearTimeout(timeout);
      };
    }
  }, []);

  const uploadUploadHandler: any = async (
    blobInfo: { blob: Blob },
    progress: () => void
  ) => {
    const resp = await upload(
      [new File([(blobInfo.blob as any)()], "filename")],
      true
    )[0];
    return directPath({ diskPath: resp as any } as any);
  };

  const isDark =
    window.matchMedia("(prefers-color-scheme: dark)").matches ||
    document.getElementsByTagName("body")[0].classList.contains("dark-theme");

  return (
    <BaseFormElement focused={focused} {...props}>
      {(config.textEditorModule === "tinymce" && !forceBasic) || forceRich ? (
        <Editor
          onInit={(evt, editor) => {
            (ref as any).current = editor;

            setTimeout(() => {
              editor.setContent(value || "", { format: "raw" });
            }, 0);

            props.onReady && props.onReady();
          }}
          // value={`<span class="katex-inline-editor"><span class="katex"><math xmlns="http://www.w3.org/1998/Math/MathML"><semantics><mrow><mi>f</mi><mo stretchy="false">(</mo><mi>x</mi><mo stretchy="false">)</mo><mo>=</mo><mstyle scriptlevel="0" displaystyle="true"><msubsup><mo>∫</mo><mrow><mo>−</mo><mi mathvariant="normal">∞</mi></mrow><mi mathvariant="normal">∞</mi></msubsup><mover accent="true"><mi>f</mi><mo>^</mo></mover><mo stretchy="false">(</mo><mi>ξ</mi><mo stretchy="false">)</mo><mtext> </mtext><msup><mi>e</mi><mrow><mn>2</mn><mi>π</mi><mi>i</mi><mi>ξ</mi><mi>x</mi></mrow></msup><mtext> </mtext><mi>d</mi><mi>ξ</mi><mo>+</mo><munder><mrow><mi>lim</mi><mo>⁡</mo></mrow><mrow><mi>h</mi><mo>→</mo><mn>0</mn></mrow></munder><mfrac><mrow><mi>f</mi><mo stretchy="false">(</mo><mi>x</mi><mo>+</mo><mi>h</mi><mo stretchy="false">)</mo><mo>−</mo><mi>f</mi><mo stretchy="false">(</mo><mi>x</mi><mo stretchy="false">)</mo></mrow><mi>h</mi></mfrac></mstyle></mrow><annotation encoding="application/x-tex">f(x) = \\displaystyle\\int_{-\\infty}^\\infty \\hat f(\\xi)\\,e^{2 \\pi i \\xi x} \\,d\\xi + \\lim_{h \\rightarrow 0 } \\frac{f(x+h)-f(x)}{h}</annotation></semantics></math></span></span><p><br data-mce-bogus="1"></p>`}
          onEditorChange={(e, editor) => {
            onChange && onChange(editor.getContent({ format: "raw" }));
          }}
          onScriptsLoad={() => addKatexSupport(t)}
          onLoadContent={() => {
            isTinyMceLoaded.current = true;
          }}
          apiKey="4dh1g4gxp1gbmxi3hnkro4wf9lfgmqr86khygey2bwb7ps74"
          onBlur={() => setFocused(false)}
          tinymceScriptSrc={
            (process.env.REACT_APP_PUBLIC_URL || "") +
            "plugins/js/tinymce/tinymce.min.js"
          }
          onFocus={() => setFocused(true)}
          init={{
            menubar: false,
            height: height || 400,
            images_upload_handler: uploadUploadHandler,
            language: "fa",
            skin: isDark ? "oxide-dark" : "oxide",
            content_css: isDark ? "dark" : "default",
            plugins: [
              "example",
              "image",
              "directionality",
              "image",
              // "advlist directionality autolink autosave link image lists charmap print preview hr anchor pagebreak",
              // "searchreplace wordcount visualblocks visualchars code fullscreen insertdatetime media nonbreaking",
              // "table contextmenu textcolor paste textcolor",
              // " autolink lists advlist link image charmap print preview anchor",
              // "searchreplace visualblocks code fullscreen",
              // "insertdatetime media table paste code help wordcount",
            ],
            toolbar:
              "undo redo | formatselect | example | image | rtl ltr | link | bullist numlist " +
              "bold italic backcolor h2 h3 | alignleft aligncenter " +
              "alignright alignjustify | bullist numlist outdent indent | " +
              "removeformat | help",
            content_style: "body {font-size:18px }",
          }}
        />
      ) : (
        <textarea
          {...restProps}
          value={value}
          style={{ minHeight: "140px" }}
          autoFocus={autoFocus}
          className={classNames(
            "form-control",
            props.errorMessage && "is-invalid",
            props.validMessage && "is-valid"
          )}
          onChange={(e) => onChange && onChange(e.target.value)}
          onBlur={() => setFocused(false)}
          onFocus={() => setFocused(true)}
        />
      )}
    </BaseFormElement>
  );
};
