export const snippets = {
  "Example1": `const Example1 = () => {
  const users = useMemo(() => generateUsers(100000), []);
  const querySource = createQuerySource(users);
  const [selectedValue, setValue] = usePresistentState<{
    name: string;
    id: number;
  }>("samplefromstaticjson", users[0]);

  return (
    <div>
      <h2>Selecting from static array</h2>
      <p>
        In many cases, you already have an array your app hard coded, then you
        want to allow user to select from them, and you store them into a form
        or a react state. In this example we create large list of users, and
        preselect the first one.
      </p>

      <pre>Value: {JSON.stringify(selectedValue, null, 2)}</pre>
      <FormSelect
        value={selectedValue}
        label="User"
        keyExtractor={(value) => value.id}
        fnLabelFormat={(value) => value.name}
        querySource={querySource}
        onChange={(value) => {
          setValue(value);
        }}
      />

      <div>Code:</div>
    </div>
  );
}`,
  "Example2": `const Example2 = () => {
  const users = useMemo(() => generateUsers(10_000), []);
  const querySource = createQuerySource(users);
  const [value, setValue] = useState<{ name: string; id: number }[]>([
    users[0],
    users[1],
    users[2],
  ]);

  return (
    <div>
      <h2>Selecting multiple from static array</h2>
      <p>
        In this example, we use a large list of users array from a static json,
        and then user can make multiple selection, and we keep that into a react
        state.
      </p>

      <pre>Value: {JSON.stringify(value, null, 2)}</pre>
      <FormSelectMultiple
        value={value}
        label="Multiple users"
        keyExtractor={(value) => value.id as any}
        fnLabelFormat={(value) => value.name}
        querySource={querySource}
        onChange={(value) => setValue(value)}
      />
    </div>
  );
}`,
  "Example3": `const Example3 = () => {
  const [value, setValue] = useState<RoleEntity[]>();

  return (
    <div>
      <h2>Select multiple entities from Fireback generated code</h2>
      <p>
        As all of the entities generated via Fireback are searchable through the
        generated sdk, by using react-query, in this example we are selecting a
        role and storing it into a react state. There are samples to store that
        on formik form using formEffect later in this document.
      </p>
      <pre>Value: {JSON.stringify(value, null, 2)}</pre>
      <FormSelectMultiple
        value={value}
        label="Multiple users"
        keyExtractor={(value) => value.uniqueId}
        fnLabelFormat={(value) => value.name}
        querySource={useGetRoles}
        onChange={(value) => setValue(value)}
      />
    </div>
  );
}`,
  "Example4": `const Example4 = () => {
  const [value, setValue] = usePresistentState("Example4", undefined);

  return (
    <div>
      <h2>Select single entity (role) from backend</h2>
      <p>
        In this scenario we allow user to select a single entity and assign it
        to the react usestate.
      </p>
      <pre>Value: {JSON.stringify(value, null, 2)}</pre>
      <FormSelect
        value={value}
        label="Select single role"
        keyExtractor={(value) => value.uniqueId}
        fnLabelFormat={(value) => value.name}
        querySource={useGetRoles}
        onChange={(value) => setValue(value)}
      />
    </div>
  );
}`,
  "Example5": `const Example5 = () => {
  class FormDataSample {
    user: {
      role?: RoleEntity;

      // This is how fireback works actually, to choose an entity you need to select it with
      // the unique id of the record (not the primary key), and the object will be filled for you
      // upon query by gorm
      roleId?: string;
    };

    static Fields = {
      user$: "user",
      user: {
        role: "user.role",
        roleId: "user.roleId",
      },
    };
  }

  return (
    <div>
      <h2>Selecting role with formEffect property</h2>
      <p>
        A lot of time we are working with formik forms. In order to avoid value,
        onChange settings for each field, FormSelect and FormMultipleSelect
        allow for <strong>formEffect</strong>
        property, which would automatically operate on the form values and
        modify them.
      </p>
      <Formik
        initialValues={{ user: {} } as FormDataSample}
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormSelect
              value={form.values.user.role}
              label="Select single role"
              keyExtractor={(value) => value.uniqueId}
              fnLabelFormat={(value) => value.name}
              querySource={useGetRoles}
              formEffect={{ field: FormDataSample.Fields.user.role, form }}
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example6": `const Example6 = () => {
  class FormDataSample {
    user: {
      roles?: RoleEntity[];
    };

    static Fields = {
      user$: "user",
      user: {
        roles: "user.roles",
      },
    };
  }

  return (
    <div>
      <h2>Selecting multiple role with formEffect</h2>
      <p>
        In this example, we allow a user to fill an array in the formik form, by
        selecting multiple roles and assign them to the user.
      </p>
      <Formik
        initialValues={{ user: {} } as FormDataSample}
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormSelectMultiple
              value={form.values.user.roles}
              label="Select multiple roles"
              keyExtractor={(value) => value.uniqueId}
              fnLabelFormat={(value) => value.name}
              querySource={useGetRoles}
              formEffect={{ field: FormDataSample.Fields.user.roles, form }}
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example9": `const Example9 = () => {
  const [selectedValue, setValue] = usePresistentState<number>(
    "samplePrimitivenumeric",
    3
  );

  const querySource = createQuerySource([
    { sisters: 1 },
    { sisters: 2 },
    { sisters: 3 },
  ]);

  return (
    <div>
      <h2>Selecting and changing only pure primitives</h2>
      <p>
        There are reasons that you want to set a primitive such as string or
        number when working with input select. In fact, by default a lot of
        components out there in react community let you do this, and you need to
        build FormSelect and FormMultipleSelect yourself.
      </p>

      <pre>Value: {JSON.stringify(selectedValue, null, 2)}</pre>
      <FormSelect
        value={selectedValue}
        label="Select a number"
        onChange={(value) => setValue(value.sisters)}
        keyExtractor={(value) => value.sisters}
        fnLabelFormat={(value) => value.sisters + " Sisters"}
        querySource={querySource}
      />
    </div>
  );
}`,
  "Example8": `const Example8 = () => {
  class FormDataSample {
    user: {
      sisters?: number;
    };

    static Fields = {
      user$: "user",
      user: {
        sisters: "user.sisters",
      },
    };
  }

  const querySource = createQuerySource([
    { sisters: 1 },
    { sisters: 2 },
    { sisters: 3 },
  ]);

  return (
    <div>
      <h2>Selecting primitives with form effect</h2>
      <p>
        Direct change, and read primitives such as string and number are
        available also as formeffect, just take a deeper look on the{" "}
        <strong>beforeSet</strong> function in this case. You need to take out
        the value you want in this callback.
      </p>
      <Formik
        initialValues={{ user: { sisters: 2 } } as FormDataSample}
        onSubmit={(data) => {
          alert(JSON.stringify(data, null, 2));
        }}
      >
        {(form: FormikProps<Partial<FormDataSample>>) => (
          <div>
            <pre>Form: {JSON.stringify(form.values, null, 2)}</pre>
            <FormSelect
              value={form.values.user.sisters}
              label="Select how many sisters user has"
              keyExtractor={(value) => value.sisters}
              fnLabelFormat={(value) => value.sisters + " sisters!"}
              querySource={querySource}
              formEffect={{
                field: FormDataSample.Fields.user.sisters,
                form,
                beforeSet(item) {
                  return item.sisters;
                },
              }}
            />
          </div>
        )}
      </Formik>
    </div>
  );
}`,
  "Example7": `const Example7 = () => {
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
}`
};
