import {
  FormSelectMultiple,
  FormSelect,
} from "@/modules/fireback/components/forms/form-select/FormSelect";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";
import usePresistentState from "@/modules/fireback/hooks/usePresistentState";
import { RoleEntity } from "@/modules/fireback/sdk/modules/workspaces/RoleEntity";
import { useGetRoles } from "@/modules/fireback/sdk/modules/workspaces/useGetRoles";
import { Formik, FormikProps } from "formik";
import { useMemo, useState } from "react";

export function DemoFormSelect() {
  return (
    <div>
      <h2>FormSelect</h2>
      <p>
        Selecting items are one of the most important aspect of any application.
        You want always give the user the option to select, search, deselect
        items and assign that selection in some part of an DTO or entity.
      </p>
      <div className="mt-5 mb-5">
        <SampleFromStaticJson />
      </div>
      <div className="mt-5 mb-5">
        <SampleMultipleSelect />
      </div>
      <div className="mt-5 mb-5">
        <SampleBackendRoles />
      </div>
      <div className="mt-5 mb-5">
        <SampleBackendRole />
      </div>
      <div className="mt-5 mb-5">
        <AffectingPrimitveOnForm />
      </div>
      <div className="mt-5 mb-5">
        <AffectingFormArray />
      </div>
      <div className="mt-5 mb-5">
        <SelectingPrimitives />
      </div>
      <div className="mt-5 mb-5">
        <SelectingPrimitivesOnFormEffect />
      </div>
    </div>
  );
}

const firstNames = `
    Ali Reza Negar Sina Parisa Mehdi Hamed Kiana Bahram Nima Farzad Samira 
    Shahram Yasmin Dariush Elham Kamran Roya Shirin Behnaz Omid Nasrin Saeed 
    Shahab Zohreh Babak Ladan Fariba Mohsen Mojgan Amir Hossein Farhad Leila 
    Arash Mahsa Behrad Taraneh Keyvan Setareh Vahid Soraya Peyman Neda Soheil 
    Forough Parsa Sara Kourosh Fereshteh Niloofar Mehrazin Matin Armin Samin 
    Pouya Anahita Shapour Laleh Dariya Navid Elnaz Siamak Shadi Behzad Rozita 
    Hassan Tarannom Baharak Pejman Mansour Parsa Mobin Yasna Yashar Mahdieh
    `.split(/\s+/);

const lastNames = `
    Torabi Moghaddam Khosravi Jafari Gholami Ahmadi Shams Karimi Hashemi 
    Zand Rajabi Shariatmadari Tavakoli Hedayati Amini Behnam Farhadi Yazdani 
    Mirzaei Eskandari Shafiei Motamedi Monfared Eslami Rashidi Daneshgar Kianian 
    Nazari Alavi Bahrami Kordestani Noori Sharifi Abbasi Asgari Hemmati Shirazi 
    Keshavarz Rezazadeh Kaviani Namdar Baniameri Kamali Moradi Azimi Sotoudeh 
    Amiri Nikpour Fakhimi Karamat Taheri Javid Salimi Saidi Yousefi Rostami 
    Najafi Ranjbar Darvishi Fallahian Ghanbari Panahi Hosseinzadeh Fattahi Rahbar 
    Sousa Oliveira Gomez Rodriguez`.split(/\s+/);

function generateUsers(count: number) {
  return Array.from({ length: count }, (_, id) => ({
    name: `${firstNames[Math.floor(Math.random() * firstNames.length)]} ${
      lastNames[Math.floor(Math.random() * lastNames.length)]
    }`,
    id: id + 1,
  }));
}

function SampleFromStaticJson() {
  // First, cast your array to a query source so it can be queried, paginated and searched
  // We create a large list to show that it works with as many items as possible.
  const users = useMemo(() => generateUsers(100000), []);
  const querySource = createQuerySource(users);
  const [selectedValue, setValue] = usePresistentState<{
    name: string;
    id: number;
  }>("samplefromstaticjson", users[0]);

  return (
    <div>
      <h2>Selecting user from static array @</h2>
      <p>
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
    </div>
  );
}

function SampleMultipleSelect() {
  // First, cast your array to a query source so it can be queried, paginated and searched
  // We create a large list to show that it works with as many items as possible.
  const users = useMemo(() => generateUsers(10_000), []);
  const querySource = createQuerySource(users);
  const [value, setValue] = useState<{ name: string; id: number }[]>([
    users[0],
    users[1],
    users[2],
  ]);

  return (
    <div>
      <h2>Selecting multiple from array</h2>
      <p>
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
}

function SampleBackendRoles() {
  const [value, setValue] = useState<RoleEntity[]>();

  return (
    <div>
      <h2>Selecting role</h2>
      <p>
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
}

function SampleBackendRole() {
  const [value, setValue] = usePresistentState("samplebackendrole", undefined);

  return (
    <div>
      <h2>Selecting role</h2>
      <p>
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
}

function AffectingPrimitveOnForm() {
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
      <h2>Selecting role with affecting form</h2>
      <p>
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
}

function AffectingFormArray() {
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
      <h2>Selecting multiple role with affecting form</h2>
      <p>
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
}

function SelectingPrimitives() {
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
        In following example, you can select between a set of users from a json
        object, allowing you to search for them by name:
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
}

function SelectingPrimitivesOnFormEffect() {
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
      <p>If you want to change primites directly into a form.</p>
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
}
