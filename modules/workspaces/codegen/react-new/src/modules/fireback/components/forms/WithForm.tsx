import { Formik, FormikHelpers, FormikProps } from "formik";
import { useRef } from "react";

export const WithForm = ({
  Form,
  setFormRef,
  onSubmit,
  initialValues,
}: {
  Form: any;
  onSubmit?: (
    values?: Partial<any>,
    formikHelpers?: FormikHelpers<Partial<any>>
  ) => void | Promise<any>;
  setFormRef?: (ref: FormikProps<Partial<any>>) => void;
  initialValues?: Partial<any>;
}) => {
  const formik = useRef<FormikProps<Partial<any>> | null>();

  return (
    <Formik
      innerRef={(p) => {
        if (p) formik.current = p;
        if (setFormRef) {
          setFormRef(p);
        }
      }}
      initialValues={initialValues || {}}
      onSubmit={onSubmit}
    >
      {(formik: FormikProps<Partial<any>>) => {
        return (
          <form
            onSubmit={(e) => {
              e.preventDefault();
              formik.submitForm();
            }}
          >
            <Form form={formik} />
          </form>
        );
      }}
    </Formik>
  );
};
