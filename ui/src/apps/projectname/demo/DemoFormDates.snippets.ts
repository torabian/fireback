export const snippets = {
  "Example1": `const Example1 = () => {
  class FormDataSample {
    date: string;

    static Fields = {
      date: "date",
    };
  }

  return (
    <div>
      <h2>Form Date demo</h2>
      <p>
        In many examples you want to select only a date string, nothing more.
        This input does that clearly.
      </p>
      <Formik
        initialValues={{ date: "2020-10-10" } as FormDataSample}
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormDate
              value={form.values.date}
              label="When did you born?"
              onChange={(value) =>
                form.setFieldValue(FormDataSample.Fields.date, value)
              }
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example2": `const Example2 = () => {
  class FormDataSample {
    time: string;

    static Fields = {
      time: "time",
    };
  }

  return (
    <div>
      <h2>Form Date demo</h2>
      <p>
        Sometimes we just need to store a time, without anything else. 5
        characters 00:00
      </p>
      <Formik
        initialValues={{ time: "22:10" } as FormDataSample}
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormTime
              value={form.values.time}
              label="At which hour did you born?"
              onChange={(value) =>
                form.setFieldValue(FormDataSample.Fields.time, value)
              }
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example3": `const Example3 = () => {
  class FormDataSample {
    datetime: string;

    static Fields = {
      datetime: "datetime",
    };
  }

  return (
    <div>
      <h2>Form DateTime demo</h2>
      <p>
        In some cases, you want to store the datetime values with timezone in
        the database. this the component to use.
      </p>
      <Formik
        initialValues={
          { datetime: "2025-05-02T10:06:00.000Z" } as FormDataSample
        }
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormDateTime
              value={form.values.datetime}
              label="When did you born?"
              onChange={(value) =>
                form.setFieldValue(FormDataSample.Fields.datetime, value)
              }
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example5": `const Example5 = () => {
  class FormDataSample {
    daterange: {
      startDate?: Date | null;
      endDate?: Date | null;
    };

    static Fields = {
      daterange$: "daterange",
      daterange: {
        startDate: "startDate",
        endDate: "endDate",
      },
    };
  }

  return (
    <div>
      <h2>Form DateRange demo</h2>
      <p>
        Choosing a date range also is an important thing in many applications,
        without timestamp.
      </p>
      <Formik
        initialValues={
          {
            daterange: {
              endDate: new Date(),
              startDate: new Date(),
            },
          } as FormDataSample
        }
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormDateRange
              value={form.values.daterange}
              label="How many days take to eggs become chicken?"
              onChange={(value) =>
                form.setFieldValue(FormDataSample.Fields.daterange$, value)
              }
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example4": `const Example4 = () => {
  class FormDataSample {
    daterange: {
      startDate?: Date | null;
      endDte?: Date | null;
    };

    static Fields = {
      daterange$: "daterange",
      daterange: {
        startDate: "startDate",
        endDate: "endDate",
      },
    };
  }

  return (
    <div>
      <h2>Form DateTimeRange demo</h2>
      <p>
        Choosing a date range also is an important thing in many applications, a
        localised timezone.
      </p>
      <Formik
        initialValues={
          {
            daterange: {
              endDate: new Date(),
              startDate: new Date(),
            },
          } as FormDataSample
        }
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormDateTimeRange
              value={form.values.daterange}
              label="Exactly what time egg came and gone??"
              onChange={(value) =>
                form.setFieldValue(FormDataSample.Fields.daterange$, value)
              }
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`
};
