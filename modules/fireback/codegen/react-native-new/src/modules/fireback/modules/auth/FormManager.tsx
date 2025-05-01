import Button from '../../components/button/Button';
import {IResponse} from '@/modules/fireback/sdk/core/http-tools';
import {Formik, FormikHelpers, FormikProps} from 'formik';
import {useEffect, useRef} from 'react';
import {Alert} from 'react-native';

export interface CommonEntityManagerProps<T> {
  data?: T | null;
  Form?: any;
  getSingleHook?: any;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  postHook?: any;
  forceEdit?: boolean;
  patchHook?: any;
  onEditTitle?: string;
  customClass?: string;
  onCreateTitle?: string;
  onSubmit: (
    values: {},
    formikHelpers: FormikHelpers<{}>,
  ) => void | Promise<any>;

  onCancel?: () => void;
  beforeSubmit?: (data: T) => T;
  onSuccess?: (response: IResponse<any>) => void;
  onFinishUriResolver?: (response: IResponse<any>, locale: string) => string;
}

export interface DtoEntity<T, V = null> {
  data?: Partial<T> | null;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  enabledFields?: Partial<V>;
  onSuccess?: (response: IResponse<T>) => void;
  showSubmit?: boolean;
  Form?: any;
}

export const FormManager = ({
  data,
  Form,
  postHook,
  onSubmit,
  patchHook,
  setInnerRef,
}: CommonEntityManagerProps<any>) => {
  const form = useRef<FormikProps<any>>();
  const isEditing = true;

  useEffect(() => {
    form.current?.setSubmitting(
      postHook?.mutation.isLoading || patchHook?.mutation.isLoading,
    );
  }, [postHook?.isLoading, patchHook?.isLoading]);

  return (
    <Formik
      innerRef={r => {
        if (r) {
          form.current = r;
          setInnerRef && setInnerRef(r);
        }
      }}
      initialValues={{}}
      onSubmit={onSubmit}>
      {(form: FormikProps<Partial<any>>) => (
        <>
          <Form isEditing={isEditing} form={form} />
          <Button title="Continue" onPress={() => form.submitForm()} />
        </>
      )}
    </Formik>
  );
};
