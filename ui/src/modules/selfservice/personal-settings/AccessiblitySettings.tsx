import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { ErrorsView } from "../../fireback-ui/components/error-view/ErrorView";
import { FormButton } from "../../fireback-ui/components/forms/form-button/FormButton";
import { PageSection } from "../../fireback-ui/components/page-section/PageSection";
import { type KeyValue } from "../../fireback-ui/types/KeyValue";
import { AppConfigContext } from "../../fireback-ui/hooks/appConfigTools";
import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { createQuerySource } from "../../fireback-ui/hooks/useAsQuery";

interface AccessibilityConfig {
  preferredHand: string;
}

const AccessibilityConfigFields = {
  preferredHand: "preferredHand",
};

const updateSettings = (
  values: Partial<AccessibilityConfig>,
  d: FormikHelpers<Partial<AccessibilityConfig>>,
) => {
  if (values.preferredHand) {
    localStorage.setItem("app_preferredHand_address", values.preferredHand);
  }
};

const availableRichAccessibilitys = (s: typeof strings): KeyValue[] => [
  {
    label: s.accessibility.leftHand,
    value: "left",
  },
  {
    label: s.accessibility.rightHand,
    value: "right",
  },
];

export function AccessiblitySettings({}: {}) {
  const { config, patchConfig } = useContext(AppConfigContext);

  const s = useS(strings);

  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<AccessibilityConfig>>({});

  const onSubmit = (
    values: Partial<AccessibilityConfig>,
    d: FormikHelpers<Partial<AccessibilityConfig>>,
  ) => {
    if (!values.preferredHand) {
      return;
    }

    patchConfig({ preferredHand: values.preferredHand });
    updateSettings(values, d);
  };

  const availbleAccessbilitySource = createQuerySource(
    availableRichAccessibilitys(s),
  );

  useEffect(() => {
    formik.current?.setValues({ preferredHand: config.preferredHand });
  }, [config.remote]);

  return (
    <PageSection title={s.accessibility.title}>
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
              formEffect={{
                form,
                field: "preferredHand",
                beforeSet(item) {
                  return item.value;
                },
              }}
              keyExtractor={(item) => item.value}
              errorMessage={form.errors.preferredHand}
              querySource={availbleAccessbilitySource}
              label={s.accessibility.preferredHand}
              hint={s.accessibility.preferredHandHint}
            />

            <FormButton
              disabled={
                form.values.preferredHand === "" ||
                form.values.preferredHand === config.preferredHand
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
