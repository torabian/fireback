import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { ErrorsView } from "../../../components/error-view/ErrorView";
import { FormButton } from "../../../components/forms/form-button/FormButton";

import { PageSection } from "../../../components/page-section/PageSection";
import { KeyValue } from "../../../definitions/definitions";
import { AppConfigContext } from "../../../hooks/appConfigTools";
import { useCommonEntityManager } from "../../../hooks/useCommonEntityManager";
import { useT } from "../../../hooks/useT";
import { enTranslations } from "../../../translations/en";
import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { createQuerySource } from "../../../hooks/useAsQuery";
import { interfaceLanguages } from "./Langugages";

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

  const languages = interfaceLanguages(t);
  const languagesQuerySource = createQuerySource(languages);

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
              keyExtractor={(item) => item.value}
              formEffect={{
                form,
                field: "interfaceLanguage",
                beforeSet(item) {
                  return item.value;
                },
              }}
              errorMessage={form.errors.interfaceLanguage}
              querySource={languagesQuerySource}
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
