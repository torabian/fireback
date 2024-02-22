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

interface AccessibilityConfig {
  preferredHand: string;
}

const AccessibilityConfigFields = {
  preferredHand: "preferredHand",
};

const updateSettings = (
  values: Partial<AccessibilityConfig>,
  d: FormikHelpers<Partial<AccessibilityConfig>>
) => {
  if (values.preferredHand) {
    localStorage.setItem("app_preferredHand_address", values.preferredHand);
  }
};

const availableRichAccessibilitys = (t: typeof enTranslations): KeyValue[] => [
  {
    label: t.accesibility.leftHand,
    value: "left",
  },
  {
    label: t.accesibility.rightHand,
    value: "right",
  },
];

export function AccessiblitySettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const t = useT();
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<AccessibilityConfig>>({});

  const onSubmit = (
    values: Partial<AccessibilityConfig>,
    d: FormikHelpers<Partial<AccessibilityConfig>>
  ) => {
    if (!values.preferredHand) {
      return;
    }

    patchConfig({ preferredHand: values.preferredHand });
    updateSettings(values, d);
  };

  useEffect(() => {
    formik.current?.setValues({ preferredHand: config.preferredHand });
  }, [config.remote]);

  return (
    <PageSection title={t.generalSettings.accessibility.title}>
      <p>{t.generalSettings.accessibility.description}</p>
      <Formik
        innerRef={(r) => {
          if (r) formik.current = r;
        }}
        initialValues={{}}
        onSubmit={onSubmit}
      >
        {(form: FormikProps<Partial<AccessibilityConfig>>) => (
          <form
            className="richtext-editor-config-form"
            onSubmit={(e) => e.preventDefault()}
          >
            <ErrorsView errors={form.errors} />
            <FormSelect
              value={form.values.preferredHand}
              onChange={(value) => {
                form.setFieldValue(
                  AccessibilityConfigFields.preferredHand,
                  value
                );
              }}
              errorMessage={form.errors.preferredHand}
              options={availableRichAccessibilitys(t)}
              label={t.settings.preferredHand}
              hint={t.settings.preferredHandHint}
            />

            <FormButton
              disabled={
                form.values.preferredHand === "" ||
                form.values.preferredHand === config.preferredHand
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
