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

interface InterfaceSettingsInformation {
  interfaceLanguage: string;
}

const InterfaceSettingsInformationFields = {
  interfaceLanguage: "interfaceLanguage",
};

const updateSettings = (
  values: Partial<InterfaceSettingsInformation>,
  d: FormikHelpers<Partial<InterfaceSettingsInformation>>
) => {
  if (values.interfaceLanguage) {
    localStorage.setItem(
      "app_interfaceLanguage_address",
      values.interfaceLanguage
    );
  }
};

const interfaceLanguages = (t: typeof enTranslations): KeyValue[] => [
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("en")
  {
    label: t.locale.englishWorldwide,
    value: "en",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("fa")
  {
    label: t.locale.persianIran,
    value: "fa",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("ru")
  {
    label: "Russian (Русский)",
    value: "ru",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("pl")
  {
    label: t.locale.polishPoland,
    value: "pl",
  },
  /// #endif
  /// #if REACT_APP_SUPPORTED_LANGUAGES.includes("ua")
  {
    label: "Ukrainain (українська)",
    value: "ua",
  },
  /// #endif
];

export function InterfaceSettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const t = useT();
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<InterfaceSettingsInformation>>({});

  const onSubmit = (
    values: Partial<InterfaceSettingsInformation>,
    d: FormikHelpers<Partial<InterfaceSettingsInformation>>
  ) => {
    if (!values.interfaceLanguage) {
      return;
    }

    patchConfig({ interfaceLanguage: values.interfaceLanguage });
    updateSettings(values, d);

    router.push(`/${values.interfaceLanguage}/settings`);
  };

  useEffect(() => {
    formik.current?.setValues({ interfaceLanguage: config.interfaceLanguage });
  }, [config.remote]);

  return (
    <PageSection title={t.generalSettings.interfaceLang.title}>
      <p>{t.generalSettings.interfaceLang.description}</p>
      <Formik
        innerRef={(r) => {
          if (r) formik.current = r;
        }}
        initialValues={{}}
        onSubmit={onSubmit}
      >
        {(form: FormikProps<Partial<InterfaceSettingsInformation>>) => (
          <form
            className="remote-service-form"
            onSubmit={(e) => e.preventDefault()}
          >
            <ErrorsView errors={form.errors} />
            <FormSelect
              value={form.values.interfaceLanguage}
              onChange={(value) => {
                form.setFieldValue(
                  InterfaceSettingsInformationFields.interfaceLanguage,
                  value
                );
              }}
              errorMessage={form.errors.interfaceLanguage}
              options={interfaceLanguages(t)}
              label={t.settings.interfaceLanguage}
              hint={t.settings.interfaceLanguageHint}
            />

            <FormButton
              disabled={
                form.values.interfaceLanguage === "" ||
                form.values.interfaceLanguage === config.interfaceLanguage
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
