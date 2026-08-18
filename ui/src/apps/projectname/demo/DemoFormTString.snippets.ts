export const snippets = {
  "Example1": `const Example1 = () => {
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
}`,
  "Example2": `const Example2 = () => {
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
}`,
};
