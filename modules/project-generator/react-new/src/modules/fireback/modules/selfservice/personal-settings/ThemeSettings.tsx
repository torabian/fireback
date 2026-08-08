import { ErrorsView } from "../../../../fireback-ui/components/error-view/ErrorView";
import { FormButton } from "../../../../fireback-ui/components/forms/form-button/FormButton";
import { FormSelect } from "../../../../fireback-ui/components/forms/form-select/FormSelect";

import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { PageSection } from "../../../../fireback-ui/components/page-section/PageSection";

import { AppConfigContext } from "../../../../fireback-ui/hooks/appConfigTools";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { createQuerySource } from "../../../../fireback-ui/hooks/useAsQuery";
import { strings } from "./strings/translations";

export interface StringKeyValue {
  label?: string;
  value?: string;
}

interface ThemeConfig {
  theme: string;
}

const ThemeConfigFields = {
  theme: "theme",
};

const updateSettings = (
  values: Partial<ThemeConfig>,
  d: FormikHelpers<Partial<ThemeConfig>>,
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

const availableRichThemes = (s: typeof strings): StringKeyValue[] => [
  {
    label: s.richThemes.macAutomatic,
    value: "mac-theme",
  },
  {
    label: s.richThemes.macLight,
    value: "mac-theme light-theme",
  },
  {
    label: s.richThemes.macDark,
    value: "mac-theme dark-theme",
  },
];

export function ThemeSettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const s = useS(strings);
  const { formik } = useCommonEntityManager<Partial<ThemeConfig>>({});

  const onSubmit = (
    values: Partial<ThemeConfig>,
    d: FormikHelpers<Partial<ThemeConfig>>,
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

  const themeQuerySource = createQuerySource(availableRichThemes(s));

  return (
    <PageSection title={s.theme.title}>
      <p>{s.theme.description}</p>
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
              keyExtractor={(t) => t.value}
              formEffect={{
                form,
                field: "theme",
                beforeSet(item) {
                  return item.value;
                },
              }}
              errorMessage={form.errors.theme}
              querySource={themeQuerySource}
              label={s.theme.label}
              hint={s.theme.hint}
            />
            <FormButton
              disabled={
                form.values.theme === "" || form.values.theme === config.theme
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
