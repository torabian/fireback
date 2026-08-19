import { Formik, type FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { ErrorsView } from "@fireback/ui-core/components/error-view/ErrorView";
import { FormButton } from "@fireback/ui-core/components/forms/form-button/FormButton";

import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { type KeyValue } from "@fireback/ui-core/types/KeyValue";
import { AppConfigContext } from "@fireback/ui-core/hooks/appConfigTools";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { setLocale } from "@fireback/ui-core/hooks/localeStore";
import { useS } from "@fireback/ui-core/hooks/useS";
import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { createQuerySource } from "@fireback/ui-core/hooks/useAsQuery";
import { interfaceLanguages } from "./Langugages";
import { strings } from "./strings/translations";

interface InterfaceSettingsInformation {
  interfaceLanguage: string;
}

const InterfaceSettingsInformationFields = {
  interfaceLanguage: "interfaceLanguage",
};

export function InterfaceSettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const s = useS(strings);
  const { uniqueId, queryClient, isEditing, formik } = useCommonEntityManager<
    Partial<InterfaceSettingsInformation>
  >({});

  const onSubmit = (values: Partial<InterfaceSettingsInformation>) => {
    if (!values.interfaceLanguage) {
      return;
    }

    // Bug fix: this used to router.push(`/${values.interfaceLanguage}/settings`)
    // - back when locale lived in the URL, that navigation was what actually
    // made useLocale() pick up the new value. Locale now comes from
    // localeStore (see its own doc comment) - setLocale() alone re-renders
    // every useLocale()/usePureLocale() instance in the app immediately, no
    // navigation needed (we're already on /settings, and staying there).
    patchConfig({ interfaceLanguage: values.interfaceLanguage });
    setLocale(values.interfaceLanguage);
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
