import { ErrorsView } from "@/components/error-view/ErrorView";
import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormSelect } from "@/components/forms/form-select/FormSelect";
import { PageSection } from "@/components/page-section/PageSection";
import { KeyValue } from "@/definitions/definitions";
import { AppConfigContext } from "@/hooks/appConfigTools";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { enTranslations } from "@/translations/en";
import { Formik, FormikHelpers, FormikProps } from "formik";
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
              value={form.values.textEditorModule}
              onChange={(value) => {
                form.setFieldValue(
                  TextEditorConfigFields.textEditorModule,
                  value
                );
              }}
              errorMessage={form.errors.textEditorModule}
              options={availableRichTextEditors(t)}
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
