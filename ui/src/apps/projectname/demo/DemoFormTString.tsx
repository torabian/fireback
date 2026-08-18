import { FormTString } from "@fireback/ui-core/components/forms/form-tstring/FormTString";
import { Formik, type FormikProps } from "formik";
import { useState } from "react";
import { CodeViewer } from "./CodeViewer";
import { snippets } from "./DemoFormTString.snippets";

export const DemoFormTString = () => {
  return (
    <div>
      <h2>FormTString component</h2>
      <p>
        A TString (modules/fireback/complexes/TString.go) is a locale -&gt; text map,
        e.g. <code>{`{"en": "Home", "fa": "خانه"}`}</code> - one value per language
        instead of one flat string. FormTString edits it the same way
        DataGridList's own TString column filter does: a modal with one text input
        per language, opened from a single field. When the field holds text in more
        than one language, it crossfades between them every 2 seconds so there's a
        hint something's there without opening the modal.
      </p>

      <div className="mt-5 mb-5">
        <Example1 />
        <CodeViewer codeString={snippets.Example1} />
      </div>
      <div className="mt-5 mb-5">
        <Example2 />
        <CodeViewer codeString={snippets.Example2} />
      </div>
    </div>
  );
};

const Example1 = () => {
  const [value, setValue] = useState<Record<string, string>>({
    en: "Home",
    fa: "خانه",
  });

  return (
    <div>
      <h2>Plain state</h2>
      <p>
        The simplest case: the value lives in a react state, updated whenever the
        modal is saved.
      </p>
      <pre>Value: {JSON.stringify(value, null, 2)}</pre>
      <FormTString label="Page title" value={value} onChange={(v) => setValue(v)} />
    </div>
  );
};

const Example2 = () => {
  class FormDataSample {
    label: Record<string, string>;

    static Fields = {
      label: "label",
    };
  }

  return (
    <div>
      <h2>Inside a Formik form</h2>
      <p>
        Same component, wired to a Formik field like every other Form* component in
        this app - starts with only one language filled in, to show the
        no-animation single-value case too.
      </p>
      <Formik
        initialValues={{ label: { en: "Untitled" } } as FormDataSample}
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormTString
              value={form.values.label}
              label="Category label"
              onChange={(value) =>
                form.setFieldValue(FormDataSample.Fields.label, value)
              }
            />
          </div>
        )}
      </Formik>
    </div>
  );
};
