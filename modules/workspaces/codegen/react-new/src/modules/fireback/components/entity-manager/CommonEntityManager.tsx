import { httpErrorHanlder } from "@/modules/fireback/hooks/api";
import { Toast } from "@/modules/fireback/hooks/toast";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { IResponse } from "../../sdk/core/http-tools";
import { Formik, FormikHelpers, FormikProps } from "formik";
import { useEffect } from "react";
import { KeyboardAction } from "../../definitions/definitions";
import { useBackButton, useCommonCrudActions } from "../action-menu/ActionMenu";
import { QueryErrorView } from "../error-view/QueryError";
import { usePageTitle } from "../page-title/PageTitle";

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
  onCancel?: () => void;
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
  Form?: any;
}

export const CommonEntityManager = ({
  data,
  Form,
  getSingleHook,
  postHook,
  onCancel,
  onFinishUriResolver,
  patchHook,
  onCreateTitle,
  onEditTitle,
  setInnerRef,
  forceEdit,
  customClass,
  beforeSubmit,
  onSuccessPatchOrPost,
}: CommonEntityManagerProps<any>) => {
  const { router, isEditing, locale, formik, t } = useCommonEntityManager<
    Partial<any>
  >({
    data,
  });

  useBackButton(onCancel, KeyboardAction.CommonBack);

  usePageTitle((isEditing || forceEdit ? onEditTitle : onCreateTitle) || "");

  const { query: getQuery } = getSingleHook;

  useEffect(() => {
    if (getQuery.data?.data) {
      formik.current?.setValues({
        ...getQuery.data.data,
      });
    }
  }, [getQuery.data]);

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
      isEditing || forceEdit
        ? patchHook?.submit(values, d)
        : postHook?.submit(values, d);

    op.then((response: any) => {
      if (response.data?.uniqueId) {
        if (onSuccessPatchOrPost) {
          onSuccessPatchOrPost(response);
        } else if (onFinishUriResolver) {
          router.goBackOrDefault(onFinishUriResolver(response, locale));
        } else {
          Toast("Done", { type: "success" });
        }
      }
    }).catch((err) => httpErrorHanlder(err, t));
  };

  useCommonCrudActions({
    // onCancel: onCancel,
    onSave() {
      formik.current?.submitForm();
    },
  });

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
          className={
            customClass == undefined
              ? "headless-form-entity-manager"
              : customClass
          }
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
          <button type="submit" className="d-none" />
        </form>
      )}
    </Formik>
  );
};
