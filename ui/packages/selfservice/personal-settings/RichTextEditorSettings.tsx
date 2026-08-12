import { ErrorsView } from "@fireback/ui-core/components/error-view/ErrorView";
import { FormButton } from "@fireback/ui-core/components/forms/form-button/FormButton";
import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { type KeyValue } from "@fireback/ui-core/types/KeyValue";
import { AppConfigContext } from "@fireback/ui-core/hooks/appConfigTools";
import { createQuerySource } from "@fireback/ui-core/hooks/useAsQuery";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useContext, useEffect } from "react";

interface TextEditorConfig {
  textEditorModule: string;
}

const TextEditorConfigFields = {
  textEditorModule: "textEditorModule",
};

const updateSettings = (
  values: Partial<TextEditorConfig>,
  d: FormikHelpers<Partial<TextEditorConfig>>,
) => {
  if (values.textEditorModule) {
    localStorage.setItem(
      "app_textEditorModule_address",
      values.textEditorModule,
    );
  }
};

const availableRichTextEditors = (s: typeof strings): KeyValue[] => [
  {
    label: s.simpleTextEditor,
    value: "bare",
  },
  {
    label: s.tinymceeditor,
    value: "tinymce",
  },
];

export function RichTextEditorSettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const s = useS(strings);
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<TextEditorConfig>>({});

  const onSubmit = (
    values: Partial<TextEditorConfig>,
    d: FormikHelpers<Partial<TextEditorConfig>>,
  ) => {
    if (!values.textEditorModule) {
      return;
    }

    patchConfig({ textEditorModule: values.textEditorModule });
    updateSettings(values, d);
  };

  useEffect(() => {
    formik.current?.setValues({ textEditorModule: config.textEditorModule });
  }, [config.remote]);

  const editors = availableRichTextEditors(s);
  const editorsQuerySource = createQuerySource(editors);

  return (
    <PageSection title={s.richTextEditor.title}>
      <p>{s.richTextEditor.description}</p>
      <Formik
        innerRef={(r) => {
          if (r) formik.current = r;
        }}
        initialValues={{}}
        onSubmit={onSubmit}
      >
        {(form: FormikProps<Partial<TextEditorConfig>>) => (
          <form
            className="richtext-editor-config-form"
            onSubmit={(e) => e.preventDefault()}
          >
            <ErrorsView errors={form.errors} />
            <FormSelect
              formEffect={{
                form,
                field: "textEditorModule",
                beforeSet(item) {
                  return item.value;
                },
              }}
              keyExtractor={(v) => v.value}
              querySource={editorsQuerySource}
              errorMessage={form.errors.textEditorModule}
              label={s.richTextEditor.label}
              hint={s.richTextEditor.hint}
            />

            <FormButton
              disabled={
                form.values.textEditorModule === "" ||
                form.values.textEditorModule === config.textEditorModule
              }
              label={s.apply}
              onClick={() => form.submitForm()}
            />
          </form>
        )}
      </Formik>
    </PageSection>
  );
}
