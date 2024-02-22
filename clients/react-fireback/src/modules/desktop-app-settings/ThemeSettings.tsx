import { ErrorsView } from "@/components/error-view/ErrorView";
import { FormButton } from "@/components/forms/form-button/FormButton";
import { FormSelect } from "@/components/forms/form-select/FormSelect";
import { PageSection } from "@/components/page-section/PageSection";
import { KeyValue } from "@/definitions/definitions";
import { AppConfigContext } from "@/hooks/appConfigTools";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect } from "react";

interface ThemeConfig {
  theme: string;
}

const ThemeConfigFields = {
  theme: "theme",
};

const updateSettings = (
  values: Partial<ThemeConfig>,
  d: FormikHelpers<Partial<ThemeConfig>>
) => {
  if (values.theme) {
    localStorage.setItem("ui_theme", values.theme);
    const b: any = document.getElementsByTagName("body")[0].classList;

    for (const klass of b.value.split(" ")) {
      if (klass.endsWith("-theme")) {
        b.remove(klass);
      }
    }
    values.theme.split(" ").forEach((item) => {
      b.add(item);
    });
  }
};

const availableRichThemes: KeyValue[] = [
  // {
  //   label: "Minimal",
  //   value: "minimal",
  // },
  {
    label: "MacOSX Automatic",
    value: "mac-theme",
  },
  {
    label: "MacOSX Light",
    value: "mac-theme light-theme",
  },
  {
    label: "MacOSX Dark",
    value: "mac-theme dark-theme",
  },
  // {
  //   label: "Windows",
  //   value: "windows",
  // },
  // {
  //   label: "IPhone",
  //   value: "ios",
  // },
  // {
  //   label: "Android",
  //   value: "android",
  // },
];

export function ThemeSettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const t = useT();
  const { formik } = useCommonEntityManager<Partial<ThemeConfig>>({});

  const onSubmit = (
    values: Partial<ThemeConfig>,
    d: FormikHelpers<Partial<ThemeConfig>>
  ) => {
    if (!values.theme) {
      return;
    }

    patchConfig({ theme: values.theme });
    updateSettings(values, d);
  };

  useEffect(() => {
    formik.current?.setValues({ theme: config.theme });
  }, [config.remote]);

  return (
    <PageSection title={t.generalSettings.theme.title}>
      <p>{t.generalSettings.theme.description}</p>
      <Formik
        innerRef={(r) => {
          if (r) formik.current = r;
        }}
        initialValues={{}}
        onSubmit={onSubmit}
      >
        {(form: FormikProps<Partial<ThemeConfig>>) => (
          <form
            className="richtext-editor-config-form"
            onSubmit={(e) => e.preventDefault()}
          >
            <ErrorsView errors={form.errors} />
            <FormSelect
              value={form.values.theme}
              onChange={(value) => {
                form.setFieldValue(ThemeConfigFields.theme, value, false);
              }}
              errorMessage={form.errors.theme}
              options={availableRichThemes}
              label={t.settings.theme}
              hint={t.settings.themeHint}
            />

            <FormButton
              disabled={
                form.values.theme === "" || form.values.theme === config.theme
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
