import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { ErrorsView } from "../../../../fireback-ui/components/error-view/ErrorView";
import { FormButton } from "../../../../fireback-ui/components/forms/form-button/FormButton";

import { PageSection } from "../../../../fireback-ui/components/page-section/PageSection";
import { type KeyValue } from "../../../definitions/definitions";
import { AppConfigContext } from "../../../../fireback-ui/hooks/appConfigTools";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { FormSelect } from "../../../../fireback-ui/components/forms/form-select/FormSelect";
import { createQuerySource } from "../../../../fireback-ui/hooks/useAsQuery";
import { interfaceLanguages } from "./Langugages";
import { strings } from "./strings/translations";

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

  const s = useS(strings);
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

  const languages = interfaceLanguages(s);
  const languagesQuerySource = createQuerySource(languages);

  return (
    <PageSection title={s.interfaceLang.title}>
      <p>{s.interfaceLang.description}</p>
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
              label={s.interfaceLang.label}
              hint={s.interfaceLang.hint}
            />

            <FormButton
              disabled={
                form.values.interfaceLanguage === "" ||
                form.values.interfaceLanguage === config.interfaceLanguage
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
