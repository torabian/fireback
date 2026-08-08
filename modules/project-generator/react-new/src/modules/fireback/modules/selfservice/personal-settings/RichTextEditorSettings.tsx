import { ErrorsView } from "../../../../fireback-ui/components/error-view/ErrorView";
import { FormButton } from "../../../../fireback-ui/components/forms/form-button/FormButton";
import { FormSelect } from "../../../../fireback-ui/components/forms/form-select/FormSelect";
import { PageSection } from "../../../../fireback-ui/components/page-section/PageSection";
import { type KeyValue } from "../../../definitions/definitions";
import { AppConfigContext } from "../../../../fireback-ui/hooks/appConfigTools";
import { createQuerySource } from "../../../../fireback-ui/hooks/useAsQuery";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { enTranslations } from "../../../translations/en";
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
  d: FormikHelpers<Partial<TextEditorConfig>>
) => {
  if (values.textEditorModule) {
    localStorage.setItem(
      "app_textEditorModule_address",
      values.textEditorModule
    );
  }
};

const availableRichTextEditors = (t: typeof enTranslations): KeyValue[] => [
  {
    label: t.simpleTextEditor,
    value: "bare",
  },
  {
    label: t.tinymceeditor,
    value: "tinymce",
  },
];

export function RichTextEditorSettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const t = useT();
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<TextEditorConfig>>({});

  const onSubmit = (
    values: Partial<TextEditorConfig>,
    d: FormikHelpers<Partial<TextEditorConfig>>
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

  const editors = availableRichTextEditors(t);
  const editorsQuerySource = createQuerySource(editors);

  return (
    <PageSection title={t.generalSettings.richTextEditor.title}>
      <p>{t.generalSettings.richTextEditor.description}</p>
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
              label={t.settings.textEditorModule}
              hint={t.settings.textEditorModuleHint}
            />

            <FormButton
              disabled={
                form.values.textEditorModule === "" ||
                form.values.textEditorModule === config.textEditorModule
              }
              label={t.settings.apply}
              onClick={() => form.submitForm()}
            />
          </form>
        )}
      </Formik>
    </PageSection>
  );
}
