import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { IResponse } from "@/sdk/fireback/core/http-tools";
import { Formik, FormikHelpers, FormikProps } from "formik";

import { useEffect } from "react";
import { toast } from "react-toastify";
import { QueryErrorView } from "../error-view/QueryError";

export interface HeadlessEntityManagerProps<T> {
  data?: T | null;
  Form: any;
  getSingleHook?: any;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  postHook?: any;
  forceEdit?: boolean;
  patchHook?: any;
  beforeSubmit?: (data: T) => T;
  onSuccessPatchOrPost?: (response: IResponse<any>) => void;
  onFinishUriResolver?: (response: IResponse<any>, locale: string) => string;
}

export interface DtoEntity<T, V = null> {
  data?: Partial<T> | null;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  enabledFields?: Partial<V>;
  onSuccess?: (response: IResponse<T>) => void;
  showSubmit?: boolean;
}

export const HeadlessEntityManager = ({
  data,
  Form,
  getSingleHook,
  postHook,
  onFinishUriResolver,
  patchHook,
  setInnerRef,
  forceEdit,
  beforeSubmit,
  onSuccessPatchOrPost,
}: HeadlessEntityManagerProps<any>) => {
  const { router, isEditing, locale, formik } = useCommonEntityManager<
    Partial<any>
  >({
    data,
  });

  // const { query: getQuery } = getSingleHook;

  useEffect(() => {
    if (getSingleHook?.query.data?.data) {
      formik.current?.setValues({
        ...getSingleHook.query.data.data,
      });
    }
  }, [getSingleHook?.query.data]);

  useEffect(() => {
    formik.current?.setSubmitting(
      postHook?.mutation.isLoading || patchHook?.mutation.isLoading
    );
  }, [postHook?.isLoading, patchHook?.isLoading]);

  const onSubmit = (values: Partial<any>, d: FormikHelpers<Partial<any>>) => {
    if (beforeSubmit) {
      values = beforeSubmit(values);
    }

    const op =
      (isEditing || forceEdit) && patchHook
        ? patchHook?.submit(values, d)
        : postHook?.submit(values, d);

    op.then((response: any) => {
      if (response.data?.uniqueId) {
        if (onSuccessPatchOrPost) {
          onSuccessPatchOrPost(response);
        } else if (onFinishUriResolver) {
          router.goBackOrDefault(onFinishUriResolver(response, locale));
        } else {
          toast("Done", { type: "success" });
        }
      }
    });
  };

  return (
    <Formik
      innerRef={(r) => {
        if (r) {
          formik.current = r;
          setInnerRef && setInnerRef(r);
        }
      }}
      initialValues={{}}
      onSubmit={onSubmit}
    >
      {(form: FormikProps<Partial<any>>) => (
        <form
          onSubmit={(e) => {
            e.preventDefault();
            form.submitForm();
          }}
          className="headless-form-entity-manager"
        >
          {/* <pre>{JSON.stringify(form.values, null, 2)}</pre> */}
          {/* <ErrorsView errors={form.errors} /> */}
          <div style={{ marginBottom: "15px" }}>
            <QueryErrorView
              query={
                postHook?.mutation?.isError
                  ? postHook.mutation
                  : patchHook?.mutation?.isError
                  ? patchHook.mutation
                  : getSingleHook?.query?.isError
                  ? getSingleHook.query
                  : null
              }
            />
          </div>
          <Form isEditing={isEditing} form={form} />
          {/* <ProductPlanForm form={form} /> */}
          <button className="btn btn-primary" type="submit">
            Submit
          </button>
        </form>
      )}
    </Formik>
  );
};
