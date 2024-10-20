import { ErrorsView } from "@/modules/fireback/components/error-view/ErrorView";
import { FormButton } from "@/modules/fireback/components/forms/form-button/FormButton";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { PageSection } from "@/modules/fireback/components/page-section/PageSection";
import { KeyValue } from "@/modules/fireback/definitions/definitions";
import { AppConfigContext } from "@/modules/fireback/hooks/appConfigTools";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { enTranslations } from "../../translations/en";
import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { en, strings } from "./strings/translations";
import { useS } from "@/modules/fireback/hooks/useS";

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
  const s = useS(strings);

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
      <p>{s.accessibility.description}</p>
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
