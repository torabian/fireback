import { FormDateRange } from "@/modules/fireback/components/forms/form-date-range/FormDateRange";
import { FormDate } from "@/modules/fireback/components/forms/form-date/FormDate";
import { FormDateTime } from "@/modules/fireback/components/forms/form-datetime/FormDateTime";
import { Formik, FormikProps } from "formik";

export function DemoFormDates() {
  return (
    <div>
      <h2>FormDate* component</h2>
      <p>
        Selecting date, time, datetime, daterange is an important aspect of many
        different apps and softwares. Fireback react comes with a different set
        of such components.
      </p>

      <div className="mt-5 mb-5">
        <FormDateExample />
      </div>
      <div className="mt-5 mb-5">
        <FormDateTimeExample />
      </div>
      <div className="mt-5 mb-5">
        <FormDateRangeExample />
      </div>
    </div>
  );
}

function FormDateExample() {
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
}

function FormDateTimeExample() {
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
}

function FormDateRangeExample() {
  class FormDataSample {
    daterange: {
      start?: Date | null;
      end?: Date | null;
    };

    static Fields = {
      daterange$: "daterange",
      daterange: {
        start: "start",
        end: "end",
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
              end: new Date(),
              start: new Date(),
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
}
